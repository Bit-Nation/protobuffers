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
	Timestamp            string   `protobuf:"bytes,7,opt,name=timestamp" json:"timestamp,omitempty"`
	Version              uint32   `protobuf:"varint,8,opt,name=version" json:"version,omitempty"`
	IdentityKeySignature []byte   `protobuf:"bytes,9,opt,name=identityKeySignature,proto3" json:"identityKeySignature,omitempty"`
	EthereumKeySignature []byte   `protobuf:"bytes,10,opt,name=ethereumKeySignature,proto3" json:"ethereumKeySignature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_pangea_3a4517dd4a07306d, []int{0}
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

func init() { proto.RegisterFile("pangea.proto", fileDescriptor_pangea_3a4517dd4a07306d) }

var fileDescriptor_pangea_3a4517dd4a07306d = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x96, 0x4b, 0x9b, 0x34, 0xa7, 0xc2, 0x70, 0xea, 0x60, 0x21, 0x86, 0x88, 0x01, 0x65, 0xea,
	0x00, 0x4f, 0x81, 0xba, 0x54, 0xe1, 0x09, 0xae, 0xe5, 0x48, 0x2d, 0xd5, 0x76, 0xe4, 0x5e, 0x90,
	0xf2, 0xd4, 0xbc, 0x02, 0x8a, 0x4d, 0x40, 0x41, 0xd9, 0xfc, 0xfd, 0xf9, 0xbe, 0xd3, 0xc1, 0xa6,
	0x25, 0xd7, 0x30, 0xed, 0xda, 0xe0, 0xc5, 0x63, 0x96, 0xd0, 0xe3, 0xd7, 0x02, 0xf2, 0x43, 0xf0,
	0x1f, 0xe6, 0xc2, 0x88, 0xb0, 0x74, 0x64, 0x59, 0xab, 0x52, 0x55, 0x45, 0x1d, 0xdf, 0x78, 0x0f,
	0xeb, 0x8b, 0x3f, 0x91, 0x18, 0xef, 0xf4, 0x22, 0xf2, 0xbf, 0x18, 0xb7, 0xb0, 0x32, 0x96, 0x1a,
	0xd6, 0x37, 0x51, 0x48, 0x00, 0x9f, 0xe0, 0xce, 0xbc, 0xb3, 0x13, 0x23, 0xfd, 0xa1, 0x3b, 0xee,
	0xb9, 0xd7, 0xcb, 0x52, 0x55, 0x9b, 0xfa, 0x1f, 0x3b, 0xf8, 0x58, 0xce, 0x1c, 0xb8, 0xb3, 0x3f,
	0xbe, 0x55, 0xf2, 0x4d, 0x59, 0xdc, 0x01, 0x9e, 0xce, 0x24, 0xaf, 0xd3, 0x3f, 0xb3, 0xe8, 0x9d,
	0x51, 0xf0, 0x01, 0x0a, 0x31, 0x96, 0xaf, 0x42, 0xb6, 0xd5, 0x79, 0x6c, 0xf6, 0x47, 0xa0, 0x86,
	0xfc, 0x93, 0xc3, 0x75, 0x58, 0x67, 0x5d, 0xaa, 0xea, 0xb6, 0x1e, 0x21, 0x3e, 0xc3, 0x76, 0x6c,
	0xb8, 0xe7, 0xfe, 0xcd, 0x34, 0x8e, 0xa4, 0x0b, 0xac, 0x8b, 0x38, 0x69, 0x56, 0x1b, 0x32, 0x63,
	0xdb, 0x49, 0x06, 0x52, 0x66, 0x4e, 0x3b, 0x66, 0xf1, 0x00, 0x2f, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x65, 0xfd, 0x94, 0x1e, 0x90, 0x01, 0x00, 0x00,
}
