// Copyright(c) 2017-2020 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: evecommon/evecommon.proto

package evecommon

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

type HashAlgorithm int32

const (
	HashAlgorithm_HASH_ALGORITHM_INVALID        HashAlgorithm = 0
	HashAlgorithm_HASH_ALGORITHM_SHA256_16BYTES HashAlgorithm = 1 // hash with sha256, the 1st 16 bytes of result
	HashAlgorithm_HASH_ALGORITHM_SHA256_32BYTES HashAlgorithm = 2 // hash with sha256, the whole 32 bytes of result
)

// Enum value maps for HashAlgorithm.
var (
	HashAlgorithm_name = map[int32]string{
		0: "HASH_ALGORITHM_INVALID",
		1: "HASH_ALGORITHM_SHA256_16BYTES",
		2: "HASH_ALGORITHM_SHA256_32BYTES",
	}
	HashAlgorithm_value = map[string]int32{
		"HASH_ALGORITHM_INVALID":        0,
		"HASH_ALGORITHM_SHA256_16BYTES": 1,
		"HASH_ALGORITHM_SHA256_32BYTES": 2,
	}
)

func (x HashAlgorithm) Enum() *HashAlgorithm {
	p := new(HashAlgorithm)
	*p = x
	return p
}

func (x HashAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_evecommon_evecommon_proto_enumTypes[0].Descriptor()
}

func (HashAlgorithm) Type() protoreflect.EnumType {
	return &file_evecommon_evecommon_proto_enumTypes[0]
}

func (x HashAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashAlgorithm.Descriptor instead.
func (HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_evecommon_evecommon_proto_rawDescGZIP(), []int{0}
}

type RadioAccessTechnology int32

const (
	RadioAccessTechnology_RADIO_ACCESS_TECHNOLOGY_UNSPECIFIED RadioAccessTechnology = 0
	// Global System for Mobile Communications (2G).
	RadioAccessTechnology_RADIO_ACCESS_TECHNOLOGY_GSM RadioAccessTechnology = 1
	// Universal Mobile Telecommunications System (3G).
	RadioAccessTechnology_RADIO_ACCESS_TECHNOLOGY_UMTS RadioAccessTechnology = 2
	// Long Term Evolution (4G).
	RadioAccessTechnology_RADIO_ACCESS_TECHNOLOGY_LTE RadioAccessTechnology = 3
	// 5G New Radio.
	RadioAccessTechnology_RADIO_ACCESS_TECHNOLOGY_5GNR RadioAccessTechnology = 4
)

// Enum value maps for RadioAccessTechnology.
var (
	RadioAccessTechnology_name = map[int32]string{
		0: "RADIO_ACCESS_TECHNOLOGY_UNSPECIFIED",
		1: "RADIO_ACCESS_TECHNOLOGY_GSM",
		2: "RADIO_ACCESS_TECHNOLOGY_UMTS",
		3: "RADIO_ACCESS_TECHNOLOGY_LTE",
		4: "RADIO_ACCESS_TECHNOLOGY_5GNR",
	}
	RadioAccessTechnology_value = map[string]int32{
		"RADIO_ACCESS_TECHNOLOGY_UNSPECIFIED": 0,
		"RADIO_ACCESS_TECHNOLOGY_GSM":         1,
		"RADIO_ACCESS_TECHNOLOGY_UMTS":        2,
		"RADIO_ACCESS_TECHNOLOGY_LTE":         3,
		"RADIO_ACCESS_TECHNOLOGY_5GNR":        4,
	}
)

func (x RadioAccessTechnology) Enum() *RadioAccessTechnology {
	p := new(RadioAccessTechnology)
	*p = x
	return p
}

func (x RadioAccessTechnology) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RadioAccessTechnology) Descriptor() protoreflect.EnumDescriptor {
	return file_evecommon_evecommon_proto_enumTypes[1].Descriptor()
}

func (RadioAccessTechnology) Type() protoreflect.EnumType {
	return &file_evecommon_evecommon_proto_enumTypes[1]
}

func (x RadioAccessTechnology) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RadioAccessTechnology.Descriptor instead.
func (RadioAccessTechnology) EnumDescriptor() ([]byte, []int) {
	return file_evecommon_evecommon_proto_rawDescGZIP(), []int{1}
}

// DiskDescription describes disk
// we can use different data to locate disk in the system
type DiskDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                  // bus-related name, for example: /dev/sdc
	LogicalName string `protobuf:"bytes,2,opt,name=logical_name,json=logicalName,proto3" json:"logical_name,omitempty"` // logical name, for example: disk3
	Serial      string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`                              // serial number of disk
}

func (x *DiskDescription) Reset() {
	*x = DiskDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evecommon_evecommon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskDescription) ProtoMessage() {}

func (x *DiskDescription) ProtoReflect() protoreflect.Message {
	mi := &file_evecommon_evecommon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskDescription.ProtoReflect.Descriptor instead.
func (*DiskDescription) Descriptor() ([]byte, []int) {
	return file_evecommon_evecommon_proto_rawDescGZIP(), []int{0}
}

func (x *DiskDescription) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiskDescription) GetLogicalName() string {
	if x != nil {
		return x.LogicalName
	}
	return ""
}

func (x *DiskDescription) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

var File_evecommon_evecommon_proto protoreflect.FileDescriptor

var file_evecommon_evecommon_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67,
	0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x6b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x2a, 0x71, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x41, 0x4c,
	0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x21, 0x0a, 0x1d, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49,
	0x54, 0x48, 0x4d, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x5f, 0x31, 0x36, 0x42, 0x59, 0x54,
	0x45, 0x53, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x41, 0x4c, 0x47,
	0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x5f, 0x33, 0x32,
	0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x02, 0x2a, 0xc6, 0x01, 0x0a, 0x15, 0x52, 0x61, 0x64, 0x69,
	0x6f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x12, 0x27, 0x0a, 0x23, 0x52, 0x41, 0x44, 0x49, 0x4f, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x54, 0x45, 0x43, 0x48, 0x4e, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x41,
	0x44, 0x49, 0x4f, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x45, 0x43, 0x48, 0x4e,
	0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x5f, 0x47, 0x53, 0x4d, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x52,
	0x41, 0x44, 0x49, 0x4f, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x45, 0x43, 0x48,
	0x4e, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x5f, 0x55, 0x4d, 0x54, 0x53, 0x10, 0x02, 0x12, 0x1f, 0x0a,
	0x1b, 0x52, 0x41, 0x44, 0x49, 0x4f, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x45,
	0x43, 0x48, 0x4e, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x5f, 0x4c, 0x54, 0x45, 0x10, 0x03, 0x12, 0x20,
	0x0a, 0x1c, 0x52, 0x41, 0x44, 0x49, 0x4f, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54,
	0x45, 0x43, 0x48, 0x4e, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x5f, 0x35, 0x47, 0x4e, 0x52, 0x10, 0x04,
	0x42, 0x4d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x09, 0x45, 0x76, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evecommon_evecommon_proto_rawDescOnce sync.Once
	file_evecommon_evecommon_proto_rawDescData = file_evecommon_evecommon_proto_rawDesc
)

func file_evecommon_evecommon_proto_rawDescGZIP() []byte {
	file_evecommon_evecommon_proto_rawDescOnce.Do(func() {
		file_evecommon_evecommon_proto_rawDescData = protoimpl.X.CompressGZIP(file_evecommon_evecommon_proto_rawDescData)
	})
	return file_evecommon_evecommon_proto_rawDescData
}

var file_evecommon_evecommon_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_evecommon_evecommon_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_evecommon_evecommon_proto_goTypes = []interface{}{
	(HashAlgorithm)(0),         // 0: org.lfedge.eve.common.HashAlgorithm
	(RadioAccessTechnology)(0), // 1: org.lfedge.eve.common.RadioAccessTechnology
	(*DiskDescription)(nil),    // 2: org.lfedge.eve.common.DiskDescription
}
var file_evecommon_evecommon_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_evecommon_evecommon_proto_init() }
func file_evecommon_evecommon_proto_init() {
	if File_evecommon_evecommon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_evecommon_evecommon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskDescription); i {
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
			RawDescriptor: file_evecommon_evecommon_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_evecommon_evecommon_proto_goTypes,
		DependencyIndexes: file_evecommon_evecommon_proto_depIdxs,
		EnumInfos:         file_evecommon_evecommon_proto_enumTypes,
		MessageInfos:      file_evecommon_evecommon_proto_msgTypes,
	}.Build()
	File_evecommon_evecommon_proto = out.File
	file_evecommon_evecommon_proto_rawDesc = nil
	file_evecommon_evecommon_proto_goTypes = nil
	file_evecommon_evecommon_proto_depIdxs = nil
}
