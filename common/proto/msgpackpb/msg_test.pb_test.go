// Copyright 2022 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: go.chromium.org/luci/common/proto/msgpackpb/msg_test.proto

package msgpackpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VALUE int32

const (
	VALUE_ZERO VALUE = 0
	VALUE_ONE  VALUE = 1
	VALUE_TWO  VALUE = 2
)

// Enum value maps for VALUE.
var (
	VALUE_name = map[int32]string{
		0: "ZERO",
		1: "ONE",
		2: "TWO",
	}
	VALUE_value = map[string]int32{
		"ZERO": 0,
		"ONE":  1,
		"TWO":  2,
	}
)

func (x VALUE) Enum() *VALUE {
	p := new(VALUE)
	*p = x
	return p
}

func (x VALUE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VALUE) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_enumTypes[0].Descriptor()
}

func (VALUE) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_enumTypes[0]
}

func (x VALUE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VALUE.Descriptor instead.
func (VALUE) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDescGZIP(), []int{0}
}

type TestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Boolval        bool                    `protobuf:"varint,2,opt,name=boolval,proto3" json:"boolval,omitempty"`
	Intval         int64                   `protobuf:"varint,3,opt,name=intval,proto3" json:"intval,omitempty"`
	Uintval        uint64                  `protobuf:"varint,4,opt,name=uintval,proto3" json:"uintval,omitempty"`
	ShortIntval    int32                   `protobuf:"varint,5,opt,name=short_intval,json=shortIntval,proto3" json:"short_intval,omitempty"`
	ShortUintval   uint32                  `protobuf:"varint,6,opt,name=short_uintval,json=shortUintval,proto3" json:"short_uintval,omitempty"`
	Strval         string                  `protobuf:"bytes,7,opt,name=strval,proto3" json:"strval,omitempty"`
	Floatval       float64                 `protobuf:"fixed64,8,opt,name=floatval,proto3" json:"floatval,omitempty"`
	ShortFloatval  float32                 `protobuf:"fixed32,9,opt,name=short_floatval,json=shortFloatval,proto3" json:"short_floatval,omitempty"`
	Value          VALUE                   `protobuf:"varint,10,opt,name=value,proto3,enum=go.chromium.org.luci.common.proto.msgpackpb.VALUE" json:"value,omitempty"`
	Mapfield       map[string]*TestMessage `protobuf:"bytes,11,rep,name=mapfield,proto3" json:"mapfield,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Duration       *durationpb.Duration    `protobuf:"bytes,12,opt,name=duration,proto3" json:"duration,omitempty"`
	Strings        []string                `protobuf:"bytes,13,rep,name=strings,proto3" json:"strings,omitempty"`
	SingleRecurse  *TestMessage            `protobuf:"bytes,14,opt,name=single_recurse,json=singleRecurse,proto3" json:"single_recurse,omitempty"`
	MultiRecursion []*TestMessage          `protobuf:"bytes,15,rep,name=multi_recursion,json=multiRecursion,proto3" json:"multi_recursion,omitempty"`
	// Types that are assignable to Choice:
	//
	//	*TestMessage_Intchoice
	//	*TestMessage_Strchoice
	Choice isTestMessage_Choice `protobuf_oneof:"choice"`
}

func (x *TestMessage) Reset() {
	*x = TestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage) ProtoMessage() {}

func (x *TestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage.ProtoReflect.Descriptor instead.
func (*TestMessage) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestMessage) GetBoolval() bool {
	if x != nil {
		return x.Boolval
	}
	return false
}

func (x *TestMessage) GetIntval() int64 {
	if x != nil {
		return x.Intval
	}
	return 0
}

func (x *TestMessage) GetUintval() uint64 {
	if x != nil {
		return x.Uintval
	}
	return 0
}

func (x *TestMessage) GetShortIntval() int32 {
	if x != nil {
		return x.ShortIntval
	}
	return 0
}

func (x *TestMessage) GetShortUintval() uint32 {
	if x != nil {
		return x.ShortUintval
	}
	return 0
}

func (x *TestMessage) GetStrval() string {
	if x != nil {
		return x.Strval
	}
	return ""
}

func (x *TestMessage) GetFloatval() float64 {
	if x != nil {
		return x.Floatval
	}
	return 0
}

func (x *TestMessage) GetShortFloatval() float32 {
	if x != nil {
		return x.ShortFloatval
	}
	return 0
}

func (x *TestMessage) GetValue() VALUE {
	if x != nil {
		return x.Value
	}
	return VALUE_ZERO
}

func (x *TestMessage) GetMapfield() map[string]*TestMessage {
	if x != nil {
		return x.Mapfield
	}
	return nil
}

func (x *TestMessage) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *TestMessage) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

func (x *TestMessage) GetSingleRecurse() *TestMessage {
	if x != nil {
		return x.SingleRecurse
	}
	return nil
}

func (x *TestMessage) GetMultiRecursion() []*TestMessage {
	if x != nil {
		return x.MultiRecursion
	}
	return nil
}

func (m *TestMessage) GetChoice() isTestMessage_Choice {
	if m != nil {
		return m.Choice
	}
	return nil
}

func (x *TestMessage) GetIntchoice() int32 {
	if x, ok := x.GetChoice().(*TestMessage_Intchoice); ok {
		return x.Intchoice
	}
	return 0
}

func (x *TestMessage) GetStrchoice() string {
	if x, ok := x.GetChoice().(*TestMessage_Strchoice); ok {
		return x.Strchoice
	}
	return ""
}

type isTestMessage_Choice interface {
	isTestMessage_Choice()
}

type TestMessage_Intchoice struct {
	Intchoice int32 `protobuf:"varint,16,opt,name=intchoice,proto3,oneof"`
}

type TestMessage_Strchoice struct {
	Strchoice string `protobuf:"bytes,17,opt,name=strchoice,proto3,oneof"`
}

func (*TestMessage_Intchoice) isTestMessage_Choice() {}

func (*TestMessage_Strchoice) isTestMessage_Choice() {}

var File_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b, 0x70, 0x62, 0x2f, 0x6d, 0x73,
	0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x75,
	0x63, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b, 0x70, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x07, 0x0a, 0x0b, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6f, 0x6f,
	0x6c, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c,
	0x76, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x75,
	0x69, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x75, 0x69,
	0x6e, 0x74, 0x76, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x69,
	0x6e, 0x74, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x49, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x69, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x76, 0x61,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x76, 0x61,
	0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x76, 0x61, 0x6c, 0x12, 0x48, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x73, 0x67, 0x70,
	0x61, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x62, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b,
	0x70, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x61,
	0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5f, 0x0a, 0x0e, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x5f, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0f, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x5f, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x75, 0x63, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b, 0x70, 0x62, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x09, 0x69, 0x6e, 0x74, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x09, 0x73,
	0x74, 0x72, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x73, 0x74, 0x72, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x75, 0x0a, 0x0d, 0x4d,
	0x61, 0x70, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4e,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x6c, 0x75, 0x63, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x2a, 0x23, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x12, 0x08, 0x0a, 0x04, 0x5a,
	0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x6f, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67,
	0x70, 0x61, 0x63, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDescData = file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDesc
)

func file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDescData)
	})
	return file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDescData
}

var file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_goTypes = []interface{}{
	(VALUE)(0),                  // 0: go.chromium.org.luci.common.proto.msgpackpb.VALUE
	(*TestMessage)(nil),         // 1: go.chromium.org.luci.common.proto.msgpackpb.TestMessage
	nil,                         // 2: go.chromium.org.luci.common.proto.msgpackpb.TestMessage.MapfieldEntry
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_depIdxs = []int32{
	0, // 0: go.chromium.org.luci.common.proto.msgpackpb.TestMessage.value:type_name -> go.chromium.org.luci.common.proto.msgpackpb.VALUE
	2, // 1: go.chromium.org.luci.common.proto.msgpackpb.TestMessage.mapfield:type_name -> go.chromium.org.luci.common.proto.msgpackpb.TestMessage.MapfieldEntry
	3, // 2: go.chromium.org.luci.common.proto.msgpackpb.TestMessage.duration:type_name -> google.protobuf.Duration
	1, // 3: go.chromium.org.luci.common.proto.msgpackpb.TestMessage.single_recurse:type_name -> go.chromium.org.luci.common.proto.msgpackpb.TestMessage
	1, // 4: go.chromium.org.luci.common.proto.msgpackpb.TestMessage.multi_recursion:type_name -> go.chromium.org.luci.common.proto.msgpackpb.TestMessage
	1, // 5: go.chromium.org.luci.common.proto.msgpackpb.TestMessage.MapfieldEntry.value:type_name -> go.chromium.org.luci.common.proto.msgpackpb.TestMessage
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_init() }
func file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_init() {
	if File_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage); i {
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
	file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TestMessage_Intchoice)(nil),
		(*TestMessage_Strchoice)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto = out.File
	file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_rawDesc = nil
	file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_goTypes = nil
	file_go_chromium_org_luci_common_proto_msgpackpb_msg_test_proto_depIdxs = nil
}