// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pangea.proto

package pangea

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Profile struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	IdentityPubKey       []byte   `protobuf:"bytes,4,opt,name=identityPubKey,proto3" json:"identityPubKey,omitempty"`
	EthereumPubKey       []byte   `protobuf:"bytes,5,opt,name=ethereumPubKey,proto3" json:"ethereumPubKey,omitempty"`
	ChatIdentityPubKey   []byte   `protobuf:"bytes,6,opt,name=chatIdentityPubKey,proto3" json:"chatIdentityPubKey,omitempty"`
	EthereumAddress      []byte   `protobuf:"bytes,7,opt,name=ethereumAddress,proto3" json:"ethereumAddress,omitempty"`
	Timestamp            string   `protobuf:"bytes,8,opt,name=timestamp" json:"timestamp,omitempty"`
	Version              uint32   `protobuf:"varint,9,opt,name=version" json:"version,omitempty"`
	IdentityKeySignature []byte   `protobuf:"bytes,10,opt,name=identityKeySignature,proto3" json:"identityKeySignature,omitempty"`
	EthereumKeySignature []byte   `protobuf:"bytes,11,opt,name=ethereumKeySignature,proto3" json:"ethereumKeySignature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_pangea_0977aa02535408a9, []int{0}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (dst *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(dst, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Profile) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Profile) GetIdentityPubKey() []byte {
	if m != nil {
		return m.IdentityPubKey
	}
	return nil
}

func (m *Profile) GetEthereumPubKey() []byte {
	if m != nil {
		return m.EthereumPubKey
	}
	return nil
}

func (m *Profile) GetChatIdentityPubKey() []byte {
	if m != nil {
		return m.ChatIdentityPubKey
	}
	return nil
}

func (m *Profile) GetEthereumAddress() []byte {
	if m != nil {
		return m.EthereumAddress
	}
	return nil
}

func (m *Profile) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Profile) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Profile) GetIdentityKeySignature() []byte {
	if m != nil {
		return m.IdentityKeySignature
	}
	return nil
}

func (m *Profile) GetEthereumKeySignature() []byte {
	if m != nil {
		return m.EthereumKeySignature
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "pangea.Profile")
}

func init() { proto.RegisterFile("pangea.proto", fileDescriptor_pangea_0977aa02535408a9) }

var fileDescriptor_pangea_0977aa02535408a9 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x6d, 0x93, 0x66, 0xac, 0x0a, 0x43, 0x0f, 0x8b, 0x78, 0x08, 0x1e, 0x24, 0xa7,
	0x1e, 0xf4, 0x09, 0x3c, 0x4a, 0x2f, 0x25, 0x3e, 0xc1, 0xb6, 0x19, 0xd3, 0x85, 0xee, 0x6e, 0xd8,
	0x4c, 0x84, 0x3c, 0x8a, 0x6f, 0x2b, 0xd9, 0xb8, 0x96, 0x84, 0xdc, 0xf6, 0xff, 0xff, 0x6f, 0x66,
	0xf8, 0x59, 0xd8, 0xd4, 0xd2, 0x54, 0x24, 0x77, 0xb5, 0xb3, 0x6c, 0x31, 0x1e, 0xd4, 0xf3, 0xcf,
	0x02, 0x92, 0x83, 0xb3, 0x5f, 0xea, 0x42, 0x88, 0xb0, 0x34, 0x52, 0x93, 0x88, 0xb2, 0x28, 0x4f,
	0x0b, 0xff, 0xc6, 0x47, 0x58, 0x5f, 0xec, 0x49, 0xb2, 0xb2, 0x46, 0xdc, 0x78, 0xff, 0x5f, 0xe3,
	0x16, 0x56, 0x4a, 0xcb, 0x8a, 0xc4, 0xc2, 0x07, 0x83, 0xc0, 0x17, 0xb8, 0x57, 0x25, 0x19, 0x56,
	0xdc, 0x1d, 0xda, 0xe3, 0x9e, 0x3a, 0xb1, 0xcc, 0xa2, 0x7c, 0x53, 0x4c, 0xdc, 0x9e, 0x23, 0x3e,
	0x93, 0xa3, 0x56, 0xff, 0x71, 0xab, 0x81, 0x1b, 0xbb, 0xb8, 0x03, 0x3c, 0x9d, 0x25, 0x7f, 0x8c,
	0x77, 0xc6, 0x9e, 0x9d, 0x49, 0x30, 0x87, 0x87, 0xb0, 0xe1, 0xbd, 0x2c, 0x1d, 0x35, 0x8d, 0x48,
	0x3c, 0x3c, 0xb5, 0xf1, 0x09, 0x52, 0x56, 0x9a, 0x1a, 0x96, 0xba, 0x16, 0x6b, 0xdf, 0xe1, 0x6a,
	0xa0, 0x80, 0xe4, 0x9b, 0x5c, 0xd3, 0x17, 0x4f, 0xb3, 0x28, 0xbf, 0x2b, 0x82, 0xc4, 0x57, 0xd8,
	0x86, 0x2e, 0x7b, 0xea, 0x3e, 0x55, 0x65, 0x24, 0xb7, 0x8e, 0x04, 0xf8, 0x33, 0xb3, 0x59, 0x3f,
	0x13, 0xce, 0x8f, 0x66, 0x6e, 0x87, 0x99, 0xb9, 0xec, 0x18, 0xfb, 0xaf, 0x7a, 0xfb, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0xc5, 0xd9, 0x1f, 0xba, 0x01, 0x00, 0x00,
}
