// Copyright(c) 2020 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.6.1
// source: auth/auth.proto

package auth

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

//This is same as certHashAlgorithm in certs/certs.proto
//Keep these two definitions in sync
type HashAlgorithm int32

const (
	HashAlgorithm_HASH_NONE           HashAlgorithm = 0
	HashAlgorithm_HASH_SHA256_16bytes HashAlgorithm = 1 // hash with sha256, the 1st 16 bytes of result in 'senderCertHash'
	HashAlgorithm_HASH_SHA256_32bytes HashAlgorithm = 2 // hash with sha256, the whole 32 bytes of hash result
)

// Enum value maps for HashAlgorithm.
var (
	HashAlgorithm_name = map[int32]string{
		0: "HASH_NONE",
		1: "HASH_SHA256_16bytes",
		2: "HASH_SHA256_32bytes",
	}
	HashAlgorithm_value = map[string]int32{
		"HASH_NONE":           0,
		"HASH_SHA256_16bytes": 1,
		"HASH_SHA256_32bytes": 2,
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
	return file_auth_auth_proto_enumTypes[0].Descriptor()
}

func (HashAlgorithm) Type() protoreflect.EnumType {
	return &file_auth_auth_proto_enumTypes[0]
}

func (x HashAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashAlgorithm.Descriptor instead.
func (HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

type AuthBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *AuthBody) Reset() {
	*x = AuthBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthBody) ProtoMessage() {}

func (x *AuthBody) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthBody.ProtoReflect.Descriptor instead.
func (*AuthBody) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthBody) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type AuthContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthPayload *AuthBody `protobuf:"bytes,1,opt,name=authPayload,proto3" json:"authPayload,omitempty"` // envolope body, a marshalled protobuf data or it can be null
	// if the length of senderCertHash received is not N bytes, as described in hashAlgorithm, then the protobuf
	// message either is not AuthContainer type, or is corrupted. Otherwise, the
	// reciever may not have the sender's signing certificate
<<<<<<< HEAD
	Algo           HashAlgorithm `protobuf:"varint,2,opt,name=algo,proto3,enum=HashAlgorithm" json:"algo,omitempty"` // hash algorithm used by sender Cert
	SenderCertHash []byte        `protobuf:"bytes,3,opt,name=senderCertHash,proto3" json:"senderCertHash,omitempty"` // N bytes in length, 1st N bytes of sender siging cert sha256 hash
	SignatureHash  []byte        `protobuf:"bytes,4,opt,name=signatureHash,proto3" json:"signatureHash,omitempty"`   // signature of the sha256 hash of the payload
	SenderCert     []byte        `protobuf:"bytes,5,opt,name=senderCert,proto3" json:"senderCert,omitempty"`         // full senderCert needed for some payloads
=======
	Algo           common.HashAlgorithm `protobuf:"varint,2,opt,name=algo,proto3,enum=common.HashAlgorithm" json:"algo,omitempty"` // hash algorithm used by sender Cert
	SenderCertHash []byte               `protobuf:"bytes,3,opt,name=senderCertHash,proto3" json:"senderCertHash,omitempty"`        // N bytes in length, 1st N bytes of sender siging cert sha256 hash
	SignatureHash  []byte               `protobuf:"bytes,4,opt,name=signatureHash,proto3" json:"signatureHash,omitempty"`          // signature of the sha256 hash of the payload
	SenderCert     []byte               `protobuf:"bytes,5,opt,name=senderCert,proto3" json:"senderCert,omitempty"`                // full senderCert needed for some payloads
>>>>>>> 47bb3b73... fixup! Removed DupliCate enum definitions and moved common nes to proto/common/common.go
}

func (x *AuthContainer) Reset() {
	*x = AuthContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthContainer) ProtoMessage() {}

func (x *AuthContainer) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthContainer.ProtoReflect.Descriptor instead.
func (*AuthContainer) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthContainer) GetAuthPayload() *AuthBody {
	if x != nil {
		return x.AuthPayload
	}
	return nil
}

func (x *AuthContainer) GetAlgo() HashAlgorithm {
	if x != nil {
		return x.Algo
	}
	return HashAlgorithm_HASH_NONE
}

func (x *AuthContainer) GetSenderCertHash() []byte {
	if x != nil {
		return x.SenderCertHash
	}
	return nil
}

func (x *AuthContainer) GetSignatureHash() []byte {
	if x != nil {
		return x.SignatureHash
	}
	return nil
}

func (x *AuthContainer) GetSenderCert() []byte {
	if x != nil {
		return x.SenderCert
	}
	return nil
}

var File_auth_auth_proto protoreflect.FileDescriptor

var file_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
<<<<<<< HEAD
	0x6f, 0x22, 0x24, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xce, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0b, 0x61, 0x75, 0x74,
	0x68, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x2a, 0x50, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x68,
	0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x41, 0x53,
	0x48, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x41, 0x53, 0x48,
	0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x5f, 0x31, 0x36, 0x62, 0x79, 0x74, 0x65, 0x73, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36,
	0x5f, 0x33, 0x32, 0x62, 0x79, 0x74, 0x65, 0x73, 0x10, 0x02, 0x42, 0x45, 0x0a, 0x1f, 0x63, 0x6f,
	0x6d, 0x2e, 0x7a, 0x65, 0x64, 0x65, 0x64, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x22, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67,
	0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
=======
	0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xd5, 0x01, 0x0a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x2b,
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0b,
	0x61, 0x75, 0x74, 0x68, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x61,
	0x6c, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x24,
	0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x65,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x43, 0x65, 0x72, 0x74, 0x42, 0x45, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x65, 0x64, 0x65,
	0x64, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
>>>>>>> 47bb3b73... fixup! Removed DupliCate enum definitions and moved common nes to proto/common/common.go
}

var (
	file_auth_auth_proto_rawDescOnce sync.Once
	file_auth_auth_proto_rawDescData = file_auth_auth_proto_rawDesc
)

func file_auth_auth_proto_rawDescGZIP() []byte {
	file_auth_auth_proto_rawDescOnce.Do(func() {
		file_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_auth_proto_rawDescData)
	})
	return file_auth_auth_proto_rawDescData
}

var file_auth_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_auth_proto_goTypes = []interface{}{
<<<<<<< HEAD
	(HashAlgorithm)(0),    // 0: hashAlgorithm
	(*AuthBody)(nil),      // 1: AuthBody
	(*AuthContainer)(nil), // 2: AuthContainer
}
var file_auth_auth_proto_depIdxs = []int32{
	1, // 0: AuthContainer.authPayload:type_name -> AuthBody
	0, // 1: AuthContainer.algo:type_name -> hashAlgorithm
=======
	(*AuthBody)(nil),          // 0: AuthBody
	(*AuthContainer)(nil),     // 1: AuthContainer
	(common.HashAlgorithm)(0), // 2: common.HashAlgorithm
}
var file_auth_auth_proto_depIdxs = []int32{
	0, // 0: AuthContainer.authPayload:type_name -> AuthBody
	2, // 1: AuthContainer.algo:type_name -> common.HashAlgorithm
>>>>>>> 47bb3b73... fixup! Removed DupliCate enum definitions and moved common nes to proto/common/common.go
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthBody); i {
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
		file_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthContainer); i {
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
			RawDescriptor: file_auth_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
		EnumInfos:         file_auth_auth_proto_enumTypes,
		MessageInfos:      file_auth_auth_proto_msgTypes,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_rawDesc = nil
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}
