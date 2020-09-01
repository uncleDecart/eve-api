// Copyright(c) 2020 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: eveuuid/eveuuid.proto

package uuid

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

type UuidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UuidRequest) Reset() {
	*x = UuidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eveuuid_eveuuid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UuidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UuidRequest) ProtoMessage() {}

func (x *UuidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eveuuid_eveuuid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UuidRequest.ProtoReflect.Descriptor instead.
func (*UuidRequest) Descriptor() ([]byte, []int) {
	return file_eveuuid_eveuuid_proto_rawDescGZIP(), []int{0}
}

// This is the response payload for POST /api/v2/edgeDevice/uuid
// The message is assumed to be protected by signing envelope
type UuidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`                                  //UUID of this edge device
	Manufacturer string `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`                  //Manufacturer, as per Controller
	ProductName  string `protobuf:"bytes,3,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"` //Product name, as per Controller
}

func (x *UuidResponse) Reset() {
	*x = UuidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eveuuid_eveuuid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UuidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UuidResponse) ProtoMessage() {}

func (x *UuidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eveuuid_eveuuid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UuidResponse.ProtoReflect.Descriptor instead.
func (*UuidResponse) Descriptor() ([]byte, []int) {
	return file_eveuuid_eveuuid_proto_rawDescGZIP(), []int{1}
}

func (x *UuidResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UuidResponse) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *UuidResponse) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

var File_eveuuid_eveuuid_proto protoreflect.FileDescriptor

var file_eveuuid_eveuuid_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x76, 0x65, 0x75, 0x75, 0x69, 0x64, 0x2f, 0x65, 0x76, 0x65, 0x75, 0x75, 0x69,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65,
	0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x22, 0x0d, 0x0a, 0x0b,
	0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x69, 0x0a, 0x0c, 0x55,
	0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x44, 0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x42, 0x07, 0x45,
	0x76, 0x65, 0x55, 0x75, 0x69, 0x64, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eveuuid_eveuuid_proto_rawDescOnce sync.Once
	file_eveuuid_eveuuid_proto_rawDescData = file_eveuuid_eveuuid_proto_rawDesc
)

func file_eveuuid_eveuuid_proto_rawDescGZIP() []byte {
	file_eveuuid_eveuuid_proto_rawDescOnce.Do(func() {
		file_eveuuid_eveuuid_proto_rawDescData = protoimpl.X.CompressGZIP(file_eveuuid_eveuuid_proto_rawDescData)
	})
	return file_eveuuid_eveuuid_proto_rawDescData
}

var file_eveuuid_eveuuid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eveuuid_eveuuid_proto_goTypes = []interface{}{
	(*UuidRequest)(nil),  // 0: org.lfedge.eve.uuid.UuidRequest
	(*UuidResponse)(nil), // 1: org.lfedge.eve.uuid.UuidResponse
}
var file_eveuuid_eveuuid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eveuuid_eveuuid_proto_init() }
func file_eveuuid_eveuuid_proto_init() {
	if File_eveuuid_eveuuid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eveuuid_eveuuid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UuidRequest); i {
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
		file_eveuuid_eveuuid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UuidResponse); i {
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
			RawDescriptor: file_eveuuid_eveuuid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eveuuid_eveuuid_proto_goTypes,
		DependencyIndexes: file_eveuuid_eveuuid_proto_depIdxs,
		MessageInfos:      file_eveuuid_eveuuid_proto_msgTypes,
	}.Build()
	File_eveuuid_eveuuid_proto = out.File
	file_eveuuid_eveuuid_proto_rawDesc = nil
	file_eveuuid_eveuuid_proto_goTypes = nil
	file_eveuuid_eveuuid_proto_depIdxs = nil
}
