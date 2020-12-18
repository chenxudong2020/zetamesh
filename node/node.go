// Copyright 2020 ZetaMesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package node

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/libp2p/go-reuseport"
	"github.com/pkg/errors"
	"github.com/zetamesh/zetamesh/codec"
	"github.com/zetamesh/zetamesh/constant"
	"github.com/zetamesh/zetamesh/message"
	"github.com/zetamesh/zetamesh/node/tun"
	"go.uber.org/zap"
)

var subnetMask = []byte{0xff, 0xff, 0, 0}

// Options represents the CLI arguments of the Zetamesh peer node
type Options struct {
	Gateway string
	Key     string
	Address string
	TLS     bool
}

// Node represents a local peer node of ZetaMesh
type Node struct {
	opt       Options
	apiClient *client
	dialer    *net.Dialer
	gateway   *net.UDPConn
	pipeline  chan []byte

	mask        net.IP   // Only packet sent to the same subnet will be handled
	pending     sync.Map // virtAddr -> time.Time
	connections sync.Map // virtAddr -> connection
}

// New returns a new instance of local peer node
func New(opt Options) *Node {
	return &Node{
		opt:       opt,
		apiClient: newClient(opt.Gateway, opt.Key, opt.TLS),
		pipeline:  make(chan []byte, 512),
	}
}

// Serve starts the local peer and connect to the matcher
func (n *Node) Serve() error {
	// Random select a free port to serve the current node
	port, err := func() (int, error) {
		conn, err := net.ListenUDP("udp", &net.UDPAddr{})
		if err != nil {
			return -1, err
		}
		defer conn.Close()
		return conn.LocalAddr().(*net.UDPAddr).Port, nil
	}()
	if err != nil {
		return errors.WithMessage(err, "no free port available")
	}

	n.dialer = &net.Dialer{
		LocalAddr: &net.UDPAddr{IP: net.IPv4zero, Port: port},
		Control:   reuseport.Control,
	}

	// Initialize the subnet mask
	n.mask = net.ParseIP(n.opt.Address).Mask([]byte{0xff, 0xff, 0, 0})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Dial to the gateway and the connection is used to keep heartbeat with gateway
	conn, err := n.dialer.DialContext(ctx, "udp", n.opt.Gateway)
	if err != nil {
		return errors.WithStack(err)
	}
	n.gateway = conn.(*net.UDPConn)

	zap.L().Info("Setup local address successfully", zap.Stringer("local", conn.LocalAddr()))

	// Setup virtual network interface tunnel
	dev, err := tun.NewTUN(n.opt.Address)
	if err != nil {
		return err
	}
	defer dev.Close()

	zap.L().Info("Setup virtual network successfully", zap.String("interface", dev.Name()))

	// Begin virtual network interface traffic handling
	go n.serveDev(ctx, dev)

	// Begin forward heartbeat message eventually
	go n.heartbeat(ctx)

	// Begin schedule all UDP messages
	return n.schedule(ctx)
}

// Stop stops the local peer and disconnect to the matcher
func (n *Node) Stop() {
	if n.gateway != nil {
		_ = n.gateway.Close()
	}
}

func (n *Node) serveDev(ctx context.Context, dev tun.Device) {
	read := func() {
		buffer := make([]byte, constant.MaxBufferSize)
		for {
			select {
			case <-ctx.Done():
				zap.L().Info("Serve virtual device read cancelled", zap.Error(ctx.Err()))
				return

			default:
				c, err := dev.Read(buffer)
				if err != nil {
					continue
				}
				packet := gopacket.NewPacket(buffer[:c], layers.LayerTypeIPv4, gopacket.DecodeOptions{Lazy: true, NoCopy: true})
				ipv4, ok := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
				if !ok {
					continue
				}

				// Skip the packet because it has different subnet
				if !ipv4.DstIP.Mask(subnetMask).Equal(n.mask) {
					continue
				}

				// Write pipeline back if the destination is the current virtual address
				destination := ipv4.DstIP.String()
				if destination == n.opt.Address {
					dataCopy := make([]byte, c)
					copy(dataCopy, buffer[:c])
					n.pipeline <- dataCopy
					continue
				}

				n.forward(destination, buffer[:c])
			}
		}
	}

	write := func() {
		for {
			select {
			case <-ctx.Done():
				zap.L().Info("Serve virtual device write cancelled", zap.Error(ctx.Err()))
				return

			case data := <-n.pipeline:
				_, err := dev.Write(data)
				if err != nil {
					zap.L().Error("Write data into virtual device failed", zap.Error(err))
				}
			}
		}
	}

	go read()
	go write()
}

func (n *Node) forward(virtAddress string, data []byte) {
	zap.L().Debug("Send packet", zap.String("peer", virtAddress))

	// Open a new tunnel if cannot find the connection between the peers
	conn, found := n.connections.Load(virtAddress)
	if found {
		conn := conn.(*connection)
		if conn.state == StateEstablished {
			// We must make a copy of the data
			dataCopy := make([]byte, len(data))
			copy(dataCopy, data)

			select {
			case conn.pipeline <- codec.EncodeRaw(dataCopy):
			default:
				zap.L().Warn("Drop data due to channel full", zap.Reflect("peer", virtAddress))
			}
		} else {
			_, _ = n.gateway.Write(codec.Encode(message.PacketType_Relay, &message.CtrlRelay{
				VirtAddress: virtAddress,
				Data:        data,
			}))
			zap.L().Debug("Relay data due to connection not ready", zap.Stringer("state", conn.state), zap.Int("length", len(data)))
		}

		return
	}

	// The connection is trying to establish
	pending, found := n.pending.Load(virtAddress)
	if found && pending.(time.Time).Add(time.Second).After(time.Now()) {
		return
	}

	zap.L().Info("Try to establish connection", zap.String("peer", virtAddress))

	n.pending.Store(virtAddress, time.Now())
	go func() {
		defer n.pending.Delete(virtAddress)
		err := n.apiClient.OpenTunnel(n.opt.Address, virtAddress)
		if err != nil {
			zap.L().Error("Try to establish connection failed", zap.Error(err), zap.String("peer", virtAddress))
		}
	}()
}

// heartbeat keeps alive with the gateway and forward UDP heartbeat to
// the gateway every `HeartbeatInterval` seconds
func (n *Node) heartbeat(ctx context.Context) {
	data := codec.Encode(message.PacketType_Heartbeat, &message.CtrlHeartbeat{
		VirtAddress: n.opt.Address,
	})

	timer := time.After(0)
	for {
		select {
		case <-ctx.Done():
			zap.L().Info("UDP heartbeat cancelled", zap.Error(ctx.Err()))
			return

		case <-timer:
			timer = time.After(time.Second * constant.HeartbeatInterval)
			_, err := n.gateway.Write(data)
			if err != nil {
				zap.L().Error("Send heartbeat failed", zap.Error(err))
				continue
			}
		}
	}
}
