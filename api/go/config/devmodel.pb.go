// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devmodel.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// XXX duplicate of definition in zmet.proto with uniq names (ZCio vs Zio)
// Types of I/O adapters that can be assigned to applications
type ZCioType int32

const (
	ZCioType_ZCioNop   ZCioType = 0
	ZCioType_ZCioEth   ZCioType = 1
	ZCioType_ZCioUSB   ZCioType = 2
	ZCioType_ZCioCOM   ZCioType = 3
	ZCioType_ZCioHDMI  ZCioType = 4
	ZCioType_ZCioOther ZCioType = 255
)

var ZCioType_name = map[int32]string{
	0:   "ZCioNop",
	1:   "ZCioEth",
	2:   "ZCioUSB",
	3:   "ZCioCOM",
	4:   "ZCioHDMI",
	255: "ZCioOther",
}

var ZCioType_value = map[string]int32{
	"ZCioNop":   0,
	"ZCioEth":   1,
	"ZCioUSB":   2,
	"ZCioCOM":   3,
	"ZCioHDMI":  4,
	"ZCioOther": 255,
}

func (x ZCioType) String() string {
	return proto.EnumName(ZCioType_name, int32(x))
}

func (ZCioType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{0}
}

type SWAdapterType int32

const (
	SWAdapterType_IGNORE SWAdapterType = 0
	SWAdapterType_VLAN   SWAdapterType = 1
	SWAdapterType_BOND   SWAdapterType = 2
)

var SWAdapterType_name = map[int32]string{
	0: "IGNORE",
	1: "VLAN",
	2: "BOND",
}

var SWAdapterType_value = map[string]int32{
	"IGNORE": 0,
	"VLAN":   1,
	"BOND":   2,
}

func (x SWAdapterType) String() string {
	return proto.EnumName(SWAdapterType_name, int32(x))
}

func (SWAdapterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{1}
}

type SWAdapterParams struct {
	AType SWAdapterType `protobuf:"varint,1,opt,name=aType,proto3,enum=SWAdapterType" json:"aType,omitempty"`
	// vlan
	UnderlayInterface string `protobuf:"bytes,8,opt,name=underlayInterface,proto3" json:"underlayInterface,omitempty"`
	VlanId            uint32 `protobuf:"varint,9,opt,name=vlanId,proto3" json:"vlanId,omitempty"`
	// OR : repeated physical interfaces for bond0
	Bondgroup            []string `protobuf:"bytes,10,rep,name=bondgroup,proto3" json:"bondgroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SWAdapterParams) Reset()         { *m = SWAdapterParams{} }
func (m *SWAdapterParams) String() string { return proto.CompactTextString(m) }
func (*SWAdapterParams) ProtoMessage()    {}
func (*SWAdapterParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{0}
}

func (m *SWAdapterParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SWAdapterParams.Unmarshal(m, b)
}
func (m *SWAdapterParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SWAdapterParams.Marshal(b, m, deterministic)
}
func (m *SWAdapterParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SWAdapterParams.Merge(m, src)
}
func (m *SWAdapterParams) XXX_Size() int {
	return xxx_messageInfo_SWAdapterParams.Size(m)
}
func (m *SWAdapterParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SWAdapterParams.DiscardUnknown(m)
}

var xxx_messageInfo_SWAdapterParams proto.InternalMessageInfo

func (m *SWAdapterParams) GetAType() SWAdapterType {
	if m != nil {
		return m.AType
	}
	return SWAdapterType_IGNORE
}

func (m *SWAdapterParams) GetUnderlayInterface() string {
	if m != nil {
		return m.UnderlayInterface
	}
	return ""
}

func (m *SWAdapterParams) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

func (m *SWAdapterParams) GetBondgroup() []string {
	if m != nil {
		return m.Bondgroup
	}
	return nil
}

type SystemAdapter struct {
	// name of the adapter; hardware-specific e.g., eth0
	Name         string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AllocDetails *SWAdapterParams `protobuf:"bytes,20,opt,name=allocDetails,proto3" json:"allocDetails,omitempty"`
	// this is part of the freelink group
	FreeUplink bool `protobuf:"varint,2,opt,name=freeUplink,proto3" json:"freeUplink,omitempty"`
	// this is part of the uplink group
	Uplink bool `protobuf:"varint,3,opt,name=uplink,proto3" json:"uplink,omitempty"`
	// attach this network config for this adapter
	NetworkUUID string `protobuf:"bytes,4,opt,name=networkUUID,proto3" json:"networkUUID,omitempty"`
	// if its static network we need ip address
	Addr string `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	// alias/logical name which will be reported to zedcloud
	// and used for app instances
	LogicalName          string   `protobuf:"bytes,6,opt,name=logicalName,proto3" json:"logicalName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemAdapter) Reset()         { *m = SystemAdapter{} }
func (m *SystemAdapter) String() string { return proto.CompactTextString(m) }
func (*SystemAdapter) ProtoMessage()    {}
func (*SystemAdapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fb58492383773ea, []int{1}
}

func (m *SystemAdapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemAdapter.Unmarshal(m, b)
}
func (m *SystemAdapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemAdapter.Marshal(b, m, deterministic)
}
func (m *SystemAdapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemAdapter.Merge(m, src)
}
func (m *SystemAdapter) XXX_Size() int {
	return xxx_messageInfo_SystemAdapter.Size(m)
}
func (m *SystemAdapter) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemAdapter.DiscardUnknown(m)
}

var xxx_messageInfo_SystemAdapter proto.InternalMessageInfo

func (m *SystemAdapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SystemAdapter) GetAllocDetails() *SWAdapterParams {
	if m != nil {
		return m.AllocDetails
	}
	return nil
}

func (m *SystemAdapter) GetFreeUplink() bool {
	if m != nil {
		return m.FreeUplink
	}
	return false
}

func (m *SystemAdapter) GetUplink() bool {
	if m != nil {
		return m.Uplink
	}
	return false
}

func (m *SystemAdapter) GetNetworkUUID() string {
	if m != nil {
		return m.NetworkUUID
	}
	return ""
}

func (m *SystemAdapter) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SystemAdapter) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

func init() {
	proto.RegisterEnum("ZCioType", ZCioType_name, ZCioType_value)
	proto.RegisterEnum("SWAdapterType", SWAdapterType_name, SWAdapterType_value)
	proto.RegisterType((*SWAdapterParams)(nil), "sWAdapterParams")
	proto.RegisterType((*SystemAdapter)(nil), "SystemAdapter")
}

func init() { proto.RegisterFile("devmodel.proto", fileDescriptor_9fb58492383773ea) }

var fileDescriptor_9fb58492383773ea = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0xe7, 0xb6, 0x2b, 0xcd, 0xe9, 0x5a, 0x8c, 0x85, 0x50, 0x2e, 0x10, 0x44, 0xd3, 0x2e,
	0xa2, 0x09, 0x12, 0x69, 0xf0, 0x02, 0xeb, 0x3a, 0x8d, 0x48, 0x2c, 0x45, 0x19, 0x05, 0x69, 0xe2,
	0xc6, 0x8d, 0x4f, 0xd3, 0x68, 0x4e, 0x1c, 0x39, 0x49, 0x51, 0x79, 0x15, 0x9e, 0x90, 0xa7, 0x00,
	0x39, 0x69, 0xa1, 0x65, 0x77, 0xff, 0x8f, 0x73, 0xa4, 0x9f, 0x3f, 0x60, 0x2c, 0x70, 0x9d, 0x29,
	0x81, 0xd2, 0x2b, 0xb4, 0xaa, 0xd4, 0xe9, 0x4f, 0x02, 0x4f, 0xcb, 0xaf, 0x97, 0x82, 0x17, 0x15,
	0xea, 0x4f, 0x5c, 0xf3, 0xac, 0x64, 0x67, 0x70, 0xcc, 0x3f, 0x6f, 0x0a, 0xb4, 0x89, 0x43, 0xdc,
	0xf1, 0xc5, 0xd8, 0xfb, 0x3b, 0x60, 0xd2, 0xa8, 0x2d, 0xd9, 0x1b, 0x78, 0x56, 0xe7, 0x02, 0xb5,
	0xe4, 0x9b, 0x20, 0xaf, 0x50, 0x2f, 0x79, 0x8c, 0xf6, 0xc0, 0x21, 0xae, 0x15, 0x3d, 0x2e, 0xd8,
	0x0b, 0xe8, 0xaf, 0x25, 0xcf, 0x03, 0x61, 0x5b, 0x0e, 0x71, 0x47, 0xd1, 0xd6, 0xb1, 0x97, 0x60,
	0x2d, 0x54, 0x2e, 0x12, 0xad, 0xea, 0xc2, 0x06, 0xa7, 0xeb, 0x5a, 0xd1, 0xbf, 0xe0, 0xf4, 0x17,
	0x81, 0xd1, 0xdd, 0xa6, 0xac, 0x30, 0xdb, 0x02, 0x30, 0x06, 0xbd, 0x9c, 0x67, 0x2d, 0x9a, 0x15,
	0x35, 0x9a, 0xbd, 0x87, 0x13, 0x2e, 0xa5, 0x8a, 0xa7, 0x58, 0xf1, 0x54, 0x96, 0xf6, 0x73, 0x87,
	0xb8, 0xc3, 0x0b, 0xea, 0xfd, 0x77, 0xae, 0xe8, 0x60, 0x8a, 0xbd, 0x02, 0x58, 0x6a, 0xc4, 0x79,
	0x21, 0xd3, 0xfc, 0xc1, 0xee, 0x38, 0xc4, 0x1d, 0x44, 0x7b, 0x89, 0x21, 0xae, 0xdb, 0xae, 0xdb,
	0x74, 0x5b, 0xc7, 0x1c, 0x18, 0xe6, 0x58, 0x7d, 0x57, 0xfa, 0x61, 0x3e, 0x0f, 0xa6, 0x76, 0xaf,
	0x01, 0xd9, 0x8f, 0x0c, 0x23, 0x17, 0x42, 0xdb, 0xc7, 0x2d, 0xa3, 0xd1, 0x66, 0x4b, 0xaa, 0x24,
	0x8d, 0xb9, 0x0c, 0x0d, 0x7e, 0xbf, 0xdd, 0xda, 0x8b, 0xce, 0xbf, 0xc1, 0xe0, 0xfe, 0x2a, 0x55,
	0xcd, 0xdd, 0x0e, 0xe1, 0x89, 0xd1, 0xa1, 0x2a, 0xe8, 0xd1, 0xce, 0x5c, 0x57, 0x2b, 0x4a, 0x76,
	0x66, 0x7e, 0x37, 0xa1, 0x9d, 0x9d, 0xb9, 0x9a, 0xdd, 0xd2, 0x2e, 0x3b, 0x69, 0xf7, 0x3f, 0x4c,
	0x6f, 0x03, 0xda, 0x63, 0x63, 0xb0, 0x8c, 0x9b, 0x55, 0x2b, 0xd4, 0xf4, 0x37, 0x39, 0xf7, 0x61,
	0x74, 0xf0, 0x8a, 0x0c, 0xa0, 0x1f, 0xdc, 0x84, 0xb3, 0xe8, 0x9a, 0x1e, 0xb1, 0x01, 0xf4, 0xbe,
	0x7c, 0xbc, 0x0c, 0x29, 0x31, 0x6a, 0x32, 0x0b, 0xa7, 0xb4, 0x33, 0xb9, 0x81, 0xd7, 0xb1, 0xca,
	0xbc, 0x1f, 0x28, 0x50, 0x70, 0x2f, 0x96, 0xaa, 0x16, 0x5e, 0x5d, 0xa2, 0x5e, 0xa7, 0x31, 0xb6,
	0x7f, 0xe7, 0xfe, 0x2c, 0x49, 0xab, 0x55, 0xbd, 0xf0, 0x62, 0x95, 0xf9, 0x72, 0xf9, 0x16, 0x45,
	0x82, 0x3e, 0xae, 0xd1, 0xe7, 0x45, 0xea, 0x27, 0xca, 0x8f, 0x55, 0xbe, 0x4c, 0x93, 0x45, 0xbf,
	0x19, 0x7e, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x2d, 0x82, 0x50, 0x7a, 0x02, 0x00, 0x00,
}
