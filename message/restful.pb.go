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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: restful.proto

package message

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// OpenTunnelRequest represents the request when trying to open
// a tunnel
type OpenTunnelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Algorithm   string `protobuf:"bytes,2,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Nonce       string `protobuf:"bytes,3,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Cipher      string `protobuf:"bytes,4,opt,name=cipher,proto3" json:"cipher,omitempty"`
	Source      string `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	Destination string `protobuf:"bytes,6,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *OpenTunnelRequest) Reset() {
	*x = OpenTunnelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restful_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenTunnelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenTunnelRequest) ProtoMessage() {}

func (x *OpenTunnelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restful_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenTunnelRequest.ProtoReflect.Descriptor instead.
func (*OpenTunnelRequest) Descriptor() ([]byte, []int) {
	return file_restful_proto_rawDescGZIP(), []int{0}
}

func (x *OpenTunnelRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *OpenTunnelRequest) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *OpenTunnelRequest) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *OpenTunnelRequest) GetCipher() string {
	if x != nil {
		return x.Cipher
	}
	return ""
}

func (x *OpenTunnelRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *OpenTunnelRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

// OpenTunnelResponse represent the response of trying to open
// a tunnel
type OpenTunnelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encrypt string `protobuf:"bytes,1,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
}

func (x *OpenTunnelResponse) Reset() {
	*x = OpenTunnelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restful_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenTunnelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenTunnelResponse) ProtoMessage() {}

func (x *OpenTunnelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restful_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenTunnelResponse.ProtoReflect.Descriptor instead.
func (*OpenTunnelResponse) Descriptor() ([]byte, []int) {
	return file_restful_proto_rawDescGZIP(), []int{1}
}

func (x *OpenTunnelResponse) GetEncrypt() string {
	if x != nil {
		return x.Encrypt
	}
	return ""
}

var File_restful_proto protoreflect.FileDescriptor

var file_restful_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x66, 0x75, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb3, 0x01, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x12, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restful_proto_rawDescOnce sync.Once
	file_restful_proto_rawDescData = file_restful_proto_rawDesc
)

func file_restful_proto_rawDescGZIP() []byte {
	file_restful_proto_rawDescOnce.Do(func() {
		file_restful_proto_rawDescData = protoimpl.X.CompressGZIP(file_restful_proto_rawDescData)
	})
	return file_restful_proto_rawDescData
}

var file_restful_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_restful_proto_goTypes = []interface{}{
	(*OpenTunnelRequest)(nil),  // 0: OpenTunnelRequest
	(*OpenTunnelResponse)(nil), // 1: OpenTunnelResponse
}
var file_restful_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_restful_proto_init() }
func file_restful_proto_init() {
	if File_restful_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restful_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenTunnelRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_restful_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenTunnelResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_restful_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_restful_proto_goTypes,
		DependencyIndexes: file_restful_proto_depIdxs,
		MessageInfos:      file_restful_proto_msgTypes,
	}.Build()
	File_restful_proto = out.File
	file_restful_proto_rawDesc = nil
	file_restful_proto_goTypes = nil
	file_restful_proto_depIdxs = nil
}
