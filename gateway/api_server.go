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

package gateway

import (
	"net"
	"sync"
	"time"

	"github.com/coreos/go-semver/semver"
	"github.com/pkg/errors"
	"github.com/zetamesh/zetamesh/message"
	"github.com/zetamesh/zetamesh/version"
	"go.uber.org/zap"
)

type (
	// PeerInfo represents the peer of the Zetamesh system.
	PeerInfo struct {
		VirtAddress   string    `json:"virt_address"`
		UDPAddress    string    `json:"udp_address"`
		LastHeartbeat time.Time `json:"-"`
	}

	// server represents the HTTP server which serves for current
	// gateway.
	server struct {
		notifier *notifier // Notifier is used to notify the peers of current tunnel
		key      string    // The key of gateway
		peers    sync.Map  // All peers connected to the gateway
	}
)

// newServer returns a new gateway server instance and the gateway server is
// used to handle the HTTP request and store the peer information.
func newServer(notifier *notifier, key string) *server {
	return &server{
		notifier: notifier,
		key:      key,
	}
}

// OpenTunnel handles the `OpenTunnelRequest` POST request. It will validate the
// client information of `Version/Key` and notify the two endpoint if the peer
// validation successfully.
func (s *server) OpenTunnel(req *message.OpenTunnelRequest) (*message.OpenTunnelResponse, error) {
	ver, err := semver.NewVersion(req.Version)
	if err != nil {
		return nil, withcode(err, message.StatusCode_InvalidVersion)
	}

	if ver.Major < version.MajorVersion {
		err := errors.Errorf("client version %s doesn't match the server version %s", req.Version, version.NewVersion().String())
		return nil, withcode(err, message.StatusCode_VersionTooOld)
	}

	// TODO: check encryption

	src := s.Peer(req.Source)
	if src == nil {
		return nil, errors.Errorf("source peer '%s' not found in cache", req.Source)
	}
	dst := s.Peer(req.Destination)
	if dst == nil {
		return nil, errors.Errorf("destination peer '%s' not found in cache", req.Destination)
	}
	s.notifier.OpenTunnel(src, dst)

	return &message.OpenTunnelResponse{}, nil
}

// Heartbeat handles the peer heartbeat packet and update the peer information
// to the latest to keep it up to date.
func (s *server) Heartbeat(remote *net.UDPAddr, heartbeat *message.CtrlHeartbeat) {
	val, found := s.peers.Load(heartbeat.VirtAddress)
	if found {
		peer := val.(*PeerInfo)
		dest := remote.String()
		peer.LastHeartbeat = time.Now()
		if peer.UDPAddress != dest {
			peer.UDPAddress = dest
		}
		return
	}

	zap.L().Info("New peer added", zap.String("peer", heartbeat.VirtAddress), zap.Stringer("remote", remote))

	peer := &PeerInfo{
		VirtAddress:   heartbeat.VirtAddress,
		UDPAddress:    remote.String(),
		LastHeartbeat: time.Now(),
	}
	s.peers.Store(heartbeat.VirtAddress, peer)
}

// Peer returns the peers and nil will be returned if the peer corresponding
// to the virtual address is not found.
func (s *server) Peer(virtAddr string) *PeerInfo {
	val, found := s.peers.Load(virtAddr)
	if !found {
		return nil
	}
	return val.(*PeerInfo)
}
