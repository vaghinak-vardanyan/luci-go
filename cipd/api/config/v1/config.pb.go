// Copyright 2019 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.26.1
// source: go.chromium.org/luci/cipd/api/config/v1/config.proto

package api

import (
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

// Defines a client whose requests should be monitored.
type ClientMonitoringConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an IP whitelist in the auth service. If a request is received from
	// an IP matching this whitelist, it will be reported.
	IpWhitelist string `protobuf:"bytes,1,opt,name=ip_whitelist,json=ipWhitelist,proto3" json:"ip_whitelist,omitempty"`
	// Monitoring label to apply when reporting metrics for this client.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *ClientMonitoringConfig) Reset() {
	*x = ClientMonitoringConfig{}
	mi := &file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientMonitoringConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMonitoringConfig) ProtoMessage() {}

func (x *ClientMonitoringConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMonitoringConfig.ProtoReflect.Descriptor instead.
func (*ClientMonitoringConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *ClientMonitoringConfig) GetIpWhitelist() string {
	if x != nil {
		return x.IpWhitelist
	}
	return ""
}

func (x *ClientMonitoringConfig) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// A schema for the monitoring.cfg config file.
//
// It defines a list of clients whose requests should be monitored.
type ClientMonitoringWhitelist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of configurations for clients to monitor. When a request is
	// received, the list is traversed in order and the first match is the
	// monitoring config to use. If none of the configs match the request is
	// unmonitored.
	ClientMonitoringConfig []*ClientMonitoringConfig `protobuf:"bytes,1,rep,name=client_monitoring_config,json=clientMonitoringConfig,proto3" json:"client_monitoring_config,omitempty"`
}

func (x *ClientMonitoringWhitelist) Reset() {
	*x = ClientMonitoringWhitelist{}
	mi := &file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientMonitoringWhitelist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMonitoringWhitelist) ProtoMessage() {}

func (x *ClientMonitoringWhitelist) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMonitoringWhitelist.ProtoReflect.Descriptor instead.
func (*ClientMonitoringWhitelist) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *ClientMonitoringWhitelist) GetClientMonitoringConfig() []*ClientMonitoringConfig {
	if x != nil {
		return x.ClientMonitoringConfig
	}
	return nil
}

// A schema for the bootstrap.cfg config file.
//
// It defines a list of packages that contain executables that should be
// accessible via direct download URLs.
type BootstrapConfigFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of all known bootstrap packages, will be scanned in order.
	BootstrapConfig []*BootstrapConfig `protobuf:"bytes,1,rep,name=bootstrap_config,json=bootstrapConfig,proto3" json:"bootstrap_config,omitempty"`
}

func (x *BootstrapConfigFile) Reset() {
	*x = BootstrapConfigFile{}
	mi := &file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BootstrapConfigFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapConfigFile) ProtoMessage() {}

func (x *BootstrapConfigFile) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapConfigFile.ProtoReflect.Descriptor instead.
func (*BootstrapConfigFile) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescGZIP(), []int{2}
}

func (x *BootstrapConfigFile) GetBootstrapConfig() []*BootstrapConfig {
	if x != nil {
		return x.BootstrapConfig
	}
	return nil
}

// BootstrapConfig defines a set of bootstrap packages under a single prefix.
//
// Each package should contain exactly one file (presumable an executable). It
// will be extracted and put into the storage, to allow the CIPD backend to
// generate direct download URLs to it. This is useful to allow clients to
// directly download such binaries.
type BootstrapConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The package prefix of matching packages e.g. "infra/tools/my-tool".
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *BootstrapConfig) Reset() {
	*x = BootstrapConfig{}
	mi := &file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BootstrapConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapConfig) ProtoMessage() {}

func (x *BootstrapConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapConfig.ProtoReflect.Descriptor instead.
func (*BootstrapConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescGZIP(), []int{3}
}

func (x *BootstrapConfig) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

var File_go_chromium_org_luci_cipd_api_config_v1_config_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x51, 0x0a, 0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x70, 0x5f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x7a, 0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x5d, 0x0a, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x16, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x5e, 0x0a, 0x13, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x62, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x69, 0x70, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x29, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x69, 0x70, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescData = file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDesc
)

func file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDescData
}

var file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_cipd_api_config_v1_config_proto_goTypes = []any{
	(*ClientMonitoringConfig)(nil),    // 0: cipd.config.ClientMonitoringConfig
	(*ClientMonitoringWhitelist)(nil), // 1: cipd.config.ClientMonitoringWhitelist
	(*BootstrapConfigFile)(nil),       // 2: cipd.config.BootstrapConfigFile
	(*BootstrapConfig)(nil),           // 3: cipd.config.BootstrapConfig
}
var file_go_chromium_org_luci_cipd_api_config_v1_config_proto_depIdxs = []int32{
	0, // 0: cipd.config.ClientMonitoringWhitelist.client_monitoring_config:type_name -> cipd.config.ClientMonitoringConfig
	3, // 1: cipd.config.BootstrapConfigFile.bootstrap_config:type_name -> cipd.config.BootstrapConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cipd_api_config_v1_config_proto_init() }
func file_go_chromium_org_luci_cipd_api_config_v1_config_proto_init() {
	if File_go_chromium_org_luci_cipd_api_config_v1_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cipd_api_config_v1_config_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cipd_api_config_v1_config_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cipd_api_config_v1_config_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cipd_api_config_v1_config_proto = out.File
	file_go_chromium_org_luci_cipd_api_config_v1_config_proto_rawDesc = nil
	file_go_chromium_org_luci_cipd_api_config_v1_config_proto_goTypes = nil
	file_go_chromium_org_luci_cipd_api_config_v1_config_proto_depIdxs = nil
}
