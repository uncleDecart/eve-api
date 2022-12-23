// Copyright(c) 2017-2020 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: evecommon/devmodelcommon.proto

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

type PhyIoType int32

const (
	PhyIoType_PhyIoNoop    PhyIoType = 0
	PhyIoType_PhyIoNetEth  PhyIoType = 1
	PhyIoType_PhyIoUSB     PhyIoType = 2
	PhyIoType_PhyIoCOM     PhyIoType = 3
	PhyIoType_PhyIoAudio   PhyIoType = 4
	PhyIoType_PhyIoNetWLAN PhyIoType = 5
	PhyIoType_PhyIoNetWWAN PhyIoType = 6
	PhyIoType_PhyIoHDMI    PhyIoType = 7
	// enum 8 is reserved for backward compatibility with controller API
	PhyIoType_PhyIoNVMEStorage PhyIoType = 9
	PhyIoType_PhyIoSATAStorage PhyIoType = 10
	PhyIoType_PhyIoNetEthPF    PhyIoType = 11
	PhyIoType_PhyIoNetEthVF    PhyIoType = 12
	PhyIoType_PhyIoOther       PhyIoType = 255
)

// Enum value maps for PhyIoType.
var (
	PhyIoType_name = map[int32]string{
		0:   "PhyIoNoop",
		1:   "PhyIoNetEth",
		2:   "PhyIoUSB",
		3:   "PhyIoCOM",
		4:   "PhyIoAudio",
		5:   "PhyIoNetWLAN",
		6:   "PhyIoNetWWAN",
		7:   "PhyIoHDMI",
		9:   "PhyIoNVMEStorage",
		10:  "PhyIoSATAStorage",
		11:  "PhyIoNetEthPF",
		12:  "PhyIoNetEthVF",
		255: "PhyIoOther",
	}
	PhyIoType_value = map[string]int32{
		"PhyIoNoop":        0,
		"PhyIoNetEth":      1,
		"PhyIoUSB":         2,
		"PhyIoCOM":         3,
		"PhyIoAudio":       4,
		"PhyIoNetWLAN":     5,
		"PhyIoNetWWAN":     6,
		"PhyIoHDMI":        7,
		"PhyIoNVMEStorage": 9,
		"PhyIoSATAStorage": 10,
		"PhyIoNetEthPF":    11,
		"PhyIoNetEthVF":    12,
		"PhyIoOther":       255,
	}
)

func (x PhyIoType) Enum() *PhyIoType {
	p := new(PhyIoType)
	*p = x
	return p
}

func (x PhyIoType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhyIoType) Descriptor() protoreflect.EnumDescriptor {
	return file_evecommon_devmodelcommon_proto_enumTypes[0].Descriptor()
}

func (PhyIoType) Type() protoreflect.EnumType {
	return &file_evecommon_devmodelcommon_proto_enumTypes[0]
}

func (x PhyIoType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhyIoType.Descriptor instead.
func (PhyIoType) EnumDescriptor() ([]byte, []int) {
	return file_evecommon_devmodelcommon_proto_rawDescGZIP(), []int{0}
}

// PhyIoMemberUsage - Indicates how each adaptor must be used by Eve.
type PhyIoMemberUsage int32

const (
	PhyIoMemberUsage_PhyIoUsageNone PhyIoMemberUsage = 0
	// Used by both management and apps.
	PhyIoMemberUsage_PhyIoUsageMgmtAndApps PhyIoMemberUsage = 1
	// Shared by multiple apps
	PhyIoMemberUsage_PhyIoUsageShared PhyIoMemberUsage = 2
	// used by only one app
	PhyIoMemberUsage_PhyIoUsageDedicated PhyIoMemberUsage = 3
	// Adapter Blocked. Do not use the Adapter.
	PhyIoMemberUsage_PhyIoUsageDisabled PhyIoMemberUsage = 4
	// Used for Management traffic only. Cannot be used by Apps.
	PhyIoMemberUsage_PhyIoUsageMgmtOnly PhyIoMemberUsage = 5
)

// Enum value maps for PhyIoMemberUsage.
var (
	PhyIoMemberUsage_name = map[int32]string{
		0: "PhyIoUsageNone",
		1: "PhyIoUsageMgmtAndApps",
		2: "PhyIoUsageShared",
		3: "PhyIoUsageDedicated",
		4: "PhyIoUsageDisabled",
		5: "PhyIoUsageMgmtOnly",
	}
	PhyIoMemberUsage_value = map[string]int32{
		"PhyIoUsageNone":        0,
		"PhyIoUsageMgmtAndApps": 1,
		"PhyIoUsageShared":      2,
		"PhyIoUsageDedicated":   3,
		"PhyIoUsageDisabled":    4,
		"PhyIoUsageMgmtOnly":    5,
	}
)

func (x PhyIoMemberUsage) Enum() *PhyIoMemberUsage {
	p := new(PhyIoMemberUsage)
	*p = x
	return p
}

func (x PhyIoMemberUsage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhyIoMemberUsage) Descriptor() protoreflect.EnumDescriptor {
	return file_evecommon_devmodelcommon_proto_enumTypes[1].Descriptor()
}

func (PhyIoMemberUsage) Type() protoreflect.EnumType {
	return &file_evecommon_devmodelcommon_proto_enumTypes[1]
}

func (x PhyIoMemberUsage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhyIoMemberUsage.Descriptor instead.
func (PhyIoMemberUsage) EnumDescriptor() ([]byte, []int) {
	return file_evecommon_devmodelcommon_proto_rawDescGZIP(), []int{1}
}

var File_evecommon_devmodelcommon_proto protoreflect.FileDescriptor

var file_evecommon_devmodelcommon_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x76, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0xed, 0x01, 0x0a, 0x09, 0x50, 0x68, 0x79, 0x49,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x4e, 0x6f,
	0x6f, 0x70, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x4e, 0x65, 0x74,
	0x45, 0x74, 0x68, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x55, 0x53,
	0x42, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x43, 0x4f, 0x4d, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x10,
	0x04, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x4e, 0x65, 0x74, 0x57, 0x4c, 0x41,
	0x4e, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x4e, 0x65, 0x74, 0x57,
	0x57, 0x41, 0x4e, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x48, 0x44,
	0x4d, 0x49, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x4e, 0x56, 0x4d,
	0x45, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x68,
	0x79, 0x49, 0x6f, 0x53, 0x41, 0x54, 0x41, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x10, 0x0a,
	0x12, 0x11, 0x0a, 0x0d, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x4e, 0x65, 0x74, 0x45, 0x74, 0x68, 0x50,
	0x46, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x4e, 0x65, 0x74, 0x45,
	0x74, 0x68, 0x56, 0x46, 0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0a, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x10, 0xff, 0x01, 0x2a, 0xa0, 0x01, 0x0a, 0x10, 0x50, 0x68, 0x79, 0x49,
	0x6f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x0e,
	0x50, 0x68, 0x79, 0x49, 0x6f, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00,
	0x12, 0x19, 0x0a, 0x15, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x67,
	0x6d, 0x74, 0x41, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x73, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x50,
	0x68, 0x79, 0x49, 0x6f, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x10,
	0x02, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x55, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x68,
	0x79, 0x49, 0x6f, 0x55, 0x73, 0x61, 0x67, 0x65, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x68, 0x79, 0x49, 0x6f, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x4d, 0x67, 0x6d, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x10, 0x05, 0x42, 0x52, 0x0a, 0x15, 0x6f, 0x72,
	0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x42, 0x0e, 0x44, 0x65, 0x76, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evecommon_devmodelcommon_proto_rawDescOnce sync.Once
	file_evecommon_devmodelcommon_proto_rawDescData = file_evecommon_devmodelcommon_proto_rawDesc
)

func file_evecommon_devmodelcommon_proto_rawDescGZIP() []byte {
	file_evecommon_devmodelcommon_proto_rawDescOnce.Do(func() {
		file_evecommon_devmodelcommon_proto_rawDescData = protoimpl.X.CompressGZIP(file_evecommon_devmodelcommon_proto_rawDescData)
	})
	return file_evecommon_devmodelcommon_proto_rawDescData
}

var file_evecommon_devmodelcommon_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_evecommon_devmodelcommon_proto_goTypes = []interface{}{
	(PhyIoType)(0),        // 0: org.lfedge.eve.common.PhyIoType
	(PhyIoMemberUsage)(0), // 1: org.lfedge.eve.common.PhyIoMemberUsage
}
var file_evecommon_devmodelcommon_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_evecommon_devmodelcommon_proto_init() }
func file_evecommon_devmodelcommon_proto_init() {
	if File_evecommon_devmodelcommon_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_evecommon_devmodelcommon_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_evecommon_devmodelcommon_proto_goTypes,
		DependencyIndexes: file_evecommon_devmodelcommon_proto_depIdxs,
		EnumInfos:         file_evecommon_devmodelcommon_proto_enumTypes,
	}.Build()
	File_evecommon_devmodelcommon_proto = out.File
	file_evecommon_devmodelcommon_proto_rawDesc = nil
	file_evecommon_devmodelcommon_proto_goTypes = nil
	file_evecommon_devmodelcommon_proto_depIdxs = nil
}
