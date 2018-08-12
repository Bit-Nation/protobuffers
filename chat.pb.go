// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

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

type MessageStateChange_NewState int32

const (
	MessageStateChange_SENT             MessageStateChange_NewState = 0
	MessageStateChange_DELIVERED        MessageStateChange_NewState = 1
	MessageStateChange_FAILED_TO_HANDLE MessageStateChange_NewState = 2
)

var MessageStateChange_NewState_name = map[int32]string{
	0: "SENT",
	1: "DELIVERED",
	2: "FAILED_TO_HANDLE",
}
var MessageStateChange_NewState_value = map[string]int32{
	"SENT":             0,
	"DELIVERED":        1,
	"FAILED_TO_HANDLE": 2,
}

func (x MessageStateChange_NewState) String() string {
	return proto.EnumName(MessageStateChange_NewState_name, int32(x))
}
func (MessageStateChange_NewState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{4, 0}
}

type DoubleRatchetMsg struct {
	DoubleRatchetPK      []byte   `protobuf:"bytes,1,opt,name=doubleRatchetPK,proto3" json:"doubleRatchetPK,omitempty"`
	N                    uint32   `protobuf:"varint,2,opt,name=n,proto3" json:"n,omitempty"`
	Pn                   uint32   `protobuf:"varint,3,opt,name=pn,proto3" json:"pn,omitempty"`
	CipherText           []byte   `protobuf:"bytes,4,opt,name=cipherText,proto3" json:"cipherText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRatchetMsg) Reset()         { *m = DoubleRatchetMsg{} }
func (m *DoubleRatchetMsg) String() string { return proto.CompactTextString(m) }
func (*DoubleRatchetMsg) ProtoMessage()    {}
func (*DoubleRatchetMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{0}
}
func (m *DoubleRatchetMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRatchetMsg.Unmarshal(m, b)
}
func (m *DoubleRatchetMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRatchetMsg.Marshal(b, m, deterministic)
}
func (dst *DoubleRatchetMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRatchetMsg.Merge(dst, src)
}
func (m *DoubleRatchetMsg) XXX_Size() int {
	return xxx_messageInfo_DoubleRatchetMsg.Size(m)
}
func (m *DoubleRatchetMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRatchetMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRatchetMsg proto.InternalMessageInfo

func (m *DoubleRatchetMsg) GetDoubleRatchetPK() []byte {
	if m != nil {
		return m.DoubleRatchetPK
	}
	return nil
}

func (m *DoubleRatchetMsg) GetN() uint32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *DoubleRatchetMsg) GetPn() uint32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *DoubleRatchetMsg) GetCipherText() []byte {
	if m != nil {
		return m.CipherText
	}
	return nil
}

type PlainChatMessage struct {
	DAppPublicKey            []byte   `protobuf:"bytes,1,opt,name=dAppPublicKey,proto3" json:"dAppPublicKey,omitempty"`
	Type                     string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Params                   []byte   `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	Message                  []byte   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt                int64    `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	MessageID                string   `protobuf:"bytes,6,opt,name=messageID,proto3" json:"messageID,omitempty"`
	SharedSecretBaseID       []byte   `protobuf:"bytes,7,opt,name=sharedSecretBaseID,proto3" json:"sharedSecretBaseID,omitempty"`
	SharedSecretCreationDate int64    `protobuf:"varint,8,opt,name=sharedSecretCreationDate,proto3" json:"sharedSecretCreationDate,omitempty"`
	Version                  uint32   `protobuf:"varint,9,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *PlainChatMessage) Reset()         { *m = PlainChatMessage{} }
func (m *PlainChatMessage) String() string { return proto.CompactTextString(m) }
func (*PlainChatMessage) ProtoMessage()    {}
func (*PlainChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{1}
}
func (m *PlainChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainChatMessage.Unmarshal(m, b)
}
func (m *PlainChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainChatMessage.Marshal(b, m, deterministic)
}
func (dst *PlainChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainChatMessage.Merge(dst, src)
}
func (m *PlainChatMessage) XXX_Size() int {
	return xxx_messageInfo_PlainChatMessage.Size(m)
}
func (m *PlainChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PlainChatMessage proto.InternalMessageInfo

func (m *PlainChatMessage) GetDAppPublicKey() []byte {
	if m != nil {
		return m.DAppPublicKey
	}
	return nil
}

func (m *PlainChatMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PlainChatMessage) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *PlainChatMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PlainChatMessage) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *PlainChatMessage) GetMessageID() string {
	if m != nil {
		return m.MessageID
	}
	return ""
}

func (m *PlainChatMessage) GetSharedSecretBaseID() []byte {
	if m != nil {
		return m.SharedSecretBaseID
	}
	return nil
}

func (m *PlainChatMessage) GetSharedSecretCreationDate() int64 {
	if m != nil {
		return m.SharedSecretCreationDate
	}
	return 0
}

func (m *PlainChatMessage) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type ChatMessage struct {
	OneTimePreKey            []byte            `protobuf:"bytes,1,opt,name=oneTimePreKey,proto3" json:"oneTimePreKey,omitempty"`
	SignedPreKey             []byte            `protobuf:"bytes,2,opt,name=signedPreKey,proto3" json:"signedPreKey,omitempty"`
	EphemeralKey             []byte            `protobuf:"bytes,3,opt,name=ephemeralKey,proto3" json:"ephemeralKey,omitempty"`
	EphemeralKeySignature    []byte            `protobuf:"bytes,4,opt,name=ephemeralKeySignature,proto3" json:"ephemeralKeySignature,omitempty"`
	SenderChatIDKey          []byte            `protobuf:"bytes,5,opt,name=senderChatIDKey,proto3" json:"senderChatIDKey,omitempty"`
	SenderChatIDKeySignature []byte            `protobuf:"bytes,6,opt,name=senderChatIDKeySignature,proto3" json:"senderChatIDKeySignature,omitempty"`
	Message                  *DoubleRatchetMsg `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
	Receiver                 []byte            `protobuf:"bytes,8,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender                   []byte            `protobuf:"bytes,9,opt,name=sender,proto3" json:"sender,omitempty"`
	MessageID                []byte            `protobuf:"bytes,10,opt,name=messageID,proto3" json:"messageID,omitempty"`
	UsedSharedSecret         []byte            `protobuf:"bytes,11,opt,name=usedSharedSecret,proto3" json:"usedSharedSecret,omitempty"`
	Version                  uint32            `protobuf:"varint,12,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{2}
}
func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (dst *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(dst, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetOneTimePreKey() []byte {
	if m != nil {
		return m.OneTimePreKey
	}
	return nil
}

func (m *ChatMessage) GetSignedPreKey() []byte {
	if m != nil {
		return m.SignedPreKey
	}
	return nil
}

func (m *ChatMessage) GetEphemeralKey() []byte {
	if m != nil {
		return m.EphemeralKey
	}
	return nil
}

func (m *ChatMessage) GetEphemeralKeySignature() []byte {
	if m != nil {
		return m.EphemeralKeySignature
	}
	return nil
}

func (m *ChatMessage) GetSenderChatIDKey() []byte {
	if m != nil {
		return m.SenderChatIDKey
	}
	return nil
}

func (m *ChatMessage) GetSenderChatIDKeySignature() []byte {
	if m != nil {
		return m.SenderChatIDKeySignature
	}
	return nil
}

func (m *ChatMessage) GetMessage() *DoubleRatchetMsg {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ChatMessage) GetReceiver() []byte {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *ChatMessage) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *ChatMessage) GetMessageID() []byte {
	if m != nil {
		return m.MessageID
	}
	return nil
}

func (m *ChatMessage) GetUsedSharedSecret() []byte {
	if m != nil {
		return m.UsedSharedSecret
	}
	return nil
}

func (m *ChatMessage) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type PreKey struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	IdentityKey          []byte   `protobuf:"bytes,2,opt,name=identityKey,proto3" json:"identityKey,omitempty"`
	IdentityKeySignature []byte   `protobuf:"bytes,3,opt,name=identityKeySignature,proto3" json:"identityKeySignature,omitempty"`
	TimeStamp            int64    `protobuf:"varint,4,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreKey) Reset()         { *m = PreKey{} }
func (m *PreKey) String() string { return proto.CompactTextString(m) }
func (*PreKey) ProtoMessage()    {}
func (*PreKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{3}
}
func (m *PreKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreKey.Unmarshal(m, b)
}
func (m *PreKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreKey.Marshal(b, m, deterministic)
}
func (dst *PreKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreKey.Merge(dst, src)
}
func (m *PreKey) XXX_Size() int {
	return xxx_messageInfo_PreKey.Size(m)
}
func (m *PreKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PreKey.DiscardUnknown(m)
}

var xxx_messageInfo_PreKey proto.InternalMessageInfo

func (m *PreKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PreKey) GetIdentityKey() []byte {
	if m != nil {
		return m.IdentityKey
	}
	return nil
}

func (m *PreKey) GetIdentityKeySignature() []byte {
	if m != nil {
		return m.IdentityKeySignature
	}
	return nil
}

func (m *PreKey) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

type MessageStateChange struct {
	NewState             MessageStateChange_NewState `protobuf:"varint,1,opt,name=newState,proto3,enum=pangea.MessageStateChange_NewState" json:"newState,omitempty"`
	MessageID            []byte                      `protobuf:"bytes,2,opt,name=messageID,proto3" json:"messageID,omitempty"`
	From                 []byte                      `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	FromSignature        []byte                      `protobuf:"bytes,4,opt,name=fromSignature,proto3" json:"fromSignature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MessageStateChange) Reset()         { *m = MessageStateChange{} }
func (m *MessageStateChange) String() string { return proto.CompactTextString(m) }
func (*MessageStateChange) ProtoMessage()    {}
func (*MessageStateChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{4}
}
func (m *MessageStateChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageStateChange.Unmarshal(m, b)
}
func (m *MessageStateChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageStateChange.Marshal(b, m, deterministic)
}
func (dst *MessageStateChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageStateChange.Merge(dst, src)
}
func (m *MessageStateChange) XXX_Size() int {
	return xxx_messageInfo_MessageStateChange.Size(m)
}
func (m *MessageStateChange) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageStateChange.DiscardUnknown(m)
}

var xxx_messageInfo_MessageStateChange proto.InternalMessageInfo

func (m *MessageStateChange) GetNewState() MessageStateChange_NewState {
	if m != nil {
		return m.NewState
	}
	return MessageStateChange_SENT
}

func (m *MessageStateChange) GetMessageID() []byte {
	if m != nil {
		return m.MessageID
	}
	return nil
}

func (m *MessageStateChange) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *MessageStateChange) GetFromSignature() []byte {
	if m != nil {
		return m.FromSignature
	}
	return nil
}

type BackendMessage struct {
	Request              *BackendMessage_Request  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response             *BackendMessage_Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	Error                string                   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	RequestID            string                   `protobuf:"bytes,4,opt,name=requestID,proto3" json:"requestID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BackendMessage) Reset()         { *m = BackendMessage{} }
func (m *BackendMessage) String() string { return proto.CompactTextString(m) }
func (*BackendMessage) ProtoMessage()    {}
func (*BackendMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{5}
}
func (m *BackendMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendMessage.Unmarshal(m, b)
}
func (m *BackendMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendMessage.Marshal(b, m, deterministic)
}
func (dst *BackendMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendMessage.Merge(dst, src)
}
func (m *BackendMessage) XXX_Size() int {
	return xxx_messageInfo_BackendMessage.Size(m)
}
func (m *BackendMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BackendMessage proto.InternalMessageInfo

func (m *BackendMessage) GetRequest() *BackendMessage_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *BackendMessage) GetResponse() *BackendMessage_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *BackendMessage) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *BackendMessage) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

type BackendMessage_Auth struct {
	ToSign               []byte   `protobuf:"bytes,1,opt,name=toSign,proto3" json:"toSign,omitempty"`
	IdentityPublicKey    []byte   `protobuf:"bytes,2,opt,name=identityPublicKey,proto3" json:"identityPublicKey,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackendMessage_Auth) Reset()         { *m = BackendMessage_Auth{} }
func (m *BackendMessage_Auth) String() string { return proto.CompactTextString(m) }
func (*BackendMessage_Auth) ProtoMessage()    {}
func (*BackendMessage_Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{5, 0}
}
func (m *BackendMessage_Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendMessage_Auth.Unmarshal(m, b)
}
func (m *BackendMessage_Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendMessage_Auth.Marshal(b, m, deterministic)
}
func (dst *BackendMessage_Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendMessage_Auth.Merge(dst, src)
}
func (m *BackendMessage_Auth) XXX_Size() int {
	return xxx_messageInfo_BackendMessage_Auth.Size(m)
}
func (m *BackendMessage_Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendMessage_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_BackendMessage_Auth proto.InternalMessageInfo

func (m *BackendMessage_Auth) GetToSign() []byte {
	if m != nil {
		return m.ToSign
	}
	return nil
}

func (m *BackendMessage_Auth) GetIdentityPublicKey() []byte {
	if m != nil {
		return m.IdentityPublicKey
	}
	return nil
}

func (m *BackendMessage_Auth) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type BackendMessage_Request struct {
	Messages             []*ChatMessage       `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	MessageStateChange   *MessageStateChange  `protobuf:"bytes,2,opt,name=messageStateChange,proto3" json:"messageStateChange,omitempty"`
	SignedPreKey         []byte               `protobuf:"bytes,3,opt,name=signedPreKey,proto3" json:"signedPreKey,omitempty"`
	NewOneTimePreKeys    uint32               `protobuf:"varint,4,opt,name=newOneTimePreKeys,proto3" json:"newOneTimePreKeys,omitempty"`
	PreKeyBundle         []byte               `protobuf:"bytes,5,opt,name=preKeyBundle,proto3" json:"preKeyBundle,omitempty"`
	Auth                 *BackendMessage_Auth `protobuf:"bytes,6,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BackendMessage_Request) Reset()         { *m = BackendMessage_Request{} }
func (m *BackendMessage_Request) String() string { return proto.CompactTextString(m) }
func (*BackendMessage_Request) ProtoMessage()    {}
func (*BackendMessage_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{5, 1}
}
func (m *BackendMessage_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendMessage_Request.Unmarshal(m, b)
}
func (m *BackendMessage_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendMessage_Request.Marshal(b, m, deterministic)
}
func (dst *BackendMessage_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendMessage_Request.Merge(dst, src)
}
func (m *BackendMessage_Request) XXX_Size() int {
	return xxx_messageInfo_BackendMessage_Request.Size(m)
}
func (m *BackendMessage_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendMessage_Request.DiscardUnknown(m)
}

var xxx_messageInfo_BackendMessage_Request proto.InternalMessageInfo

func (m *BackendMessage_Request) GetMessages() []*ChatMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *BackendMessage_Request) GetMessageStateChange() *MessageStateChange {
	if m != nil {
		return m.MessageStateChange
	}
	return nil
}

func (m *BackendMessage_Request) GetSignedPreKey() []byte {
	if m != nil {
		return m.SignedPreKey
	}
	return nil
}

func (m *BackendMessage_Request) GetNewOneTimePreKeys() uint32 {
	if m != nil {
		return m.NewOneTimePreKeys
	}
	return 0
}

func (m *BackendMessage_Request) GetPreKeyBundle() []byte {
	if m != nil {
		return m.PreKeyBundle
	}
	return nil
}

func (m *BackendMessage_Request) GetAuth() *BackendMessage_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

type BackendMessage_Response struct {
	PreKeyBundle         *BackendMessage_PreKeyBundle `protobuf:"bytes,1,opt,name=preKeyBundle,proto3" json:"preKeyBundle,omitempty"`
	OneTimePrekeys       []*PreKey                    `protobuf:"bytes,2,rep,name=oneTimePrekeys,proto3" json:"oneTimePrekeys,omitempty"`
	Auth                 *BackendMessage_Auth         `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	SignedPreKey         *PreKey                      `protobuf:"bytes,4,opt,name=signedPreKey,proto3" json:"signedPreKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BackendMessage_Response) Reset()         { *m = BackendMessage_Response{} }
func (m *BackendMessage_Response) String() string { return proto.CompactTextString(m) }
func (*BackendMessage_Response) ProtoMessage()    {}
func (*BackendMessage_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{5, 2}
}
func (m *BackendMessage_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendMessage_Response.Unmarshal(m, b)
}
func (m *BackendMessage_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendMessage_Response.Marshal(b, m, deterministic)
}
func (dst *BackendMessage_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendMessage_Response.Merge(dst, src)
}
func (m *BackendMessage_Response) XXX_Size() int {
	return xxx_messageInfo_BackendMessage_Response.Size(m)
}
func (m *BackendMessage_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendMessage_Response.DiscardUnknown(m)
}

var xxx_messageInfo_BackendMessage_Response proto.InternalMessageInfo

func (m *BackendMessage_Response) GetPreKeyBundle() *BackendMessage_PreKeyBundle {
	if m != nil {
		return m.PreKeyBundle
	}
	return nil
}

func (m *BackendMessage_Response) GetOneTimePrekeys() []*PreKey {
	if m != nil {
		return m.OneTimePrekeys
	}
	return nil
}

func (m *BackendMessage_Response) GetAuth() *BackendMessage_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *BackendMessage_Response) GetSignedPreKey() *PreKey {
	if m != nil {
		return m.SignedPreKey
	}
	return nil
}

type BackendMessage_PreKeyBundle struct {
	Profile              *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	SignedPreKey         *PreKey  `protobuf:"bytes,2,opt,name=signedPreKey,proto3" json:"signedPreKey,omitempty"`
	OneTimePreKey        *PreKey  `protobuf:"bytes,3,opt,name=oneTimePreKey,proto3" json:"oneTimePreKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackendMessage_PreKeyBundle) Reset()         { *m = BackendMessage_PreKeyBundle{} }
func (m *BackendMessage_PreKeyBundle) String() string { return proto.CompactTextString(m) }
func (*BackendMessage_PreKeyBundle) ProtoMessage()    {}
func (*BackendMessage_PreKeyBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_b5b0755f2470b837, []int{5, 3}
}
func (m *BackendMessage_PreKeyBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendMessage_PreKeyBundle.Unmarshal(m, b)
}
func (m *BackendMessage_PreKeyBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendMessage_PreKeyBundle.Marshal(b, m, deterministic)
}
func (dst *BackendMessage_PreKeyBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendMessage_PreKeyBundle.Merge(dst, src)
}
func (m *BackendMessage_PreKeyBundle) XXX_Size() int {
	return xxx_messageInfo_BackendMessage_PreKeyBundle.Size(m)
}
func (m *BackendMessage_PreKeyBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendMessage_PreKeyBundle.DiscardUnknown(m)
}

var xxx_messageInfo_BackendMessage_PreKeyBundle proto.InternalMessageInfo

func (m *BackendMessage_PreKeyBundle) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *BackendMessage_PreKeyBundle) GetSignedPreKey() *PreKey {
	if m != nil {
		return m.SignedPreKey
	}
	return nil
}

func (m *BackendMessage_PreKeyBundle) GetOneTimePreKey() *PreKey {
	if m != nil {
		return m.OneTimePreKey
	}
	return nil
}

func init() {
	proto.RegisterType((*DoubleRatchetMsg)(nil), "pangea.DoubleRatchetMsg")
	proto.RegisterType((*PlainChatMessage)(nil), "pangea.PlainChatMessage")
	proto.RegisterType((*ChatMessage)(nil), "pangea.ChatMessage")
	proto.RegisterType((*PreKey)(nil), "pangea.PreKey")
	proto.RegisterType((*MessageStateChange)(nil), "pangea.MessageStateChange")
	proto.RegisterType((*BackendMessage)(nil), "pangea.BackendMessage")
	proto.RegisterType((*BackendMessage_Auth)(nil), "pangea.BackendMessage.Auth")
	proto.RegisterType((*BackendMessage_Request)(nil), "pangea.BackendMessage.Request")
	proto.RegisterType((*BackendMessage_Response)(nil), "pangea.BackendMessage.Response")
	proto.RegisterType((*BackendMessage_PreKeyBundle)(nil), "pangea.BackendMessage.PreKeyBundle")
	proto.RegisterEnum("pangea.MessageStateChange_NewState", MessageStateChange_NewState_name, MessageStateChange_NewState_value)
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_chat_b5b0755f2470b837) }

var fileDescriptor_chat_b5b0755f2470b837 = []byte{
	// 925 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0x66, 0xc6, 0x5e, 0xff, 0x94, 0xbd, 0x8e, 0x53, 0x2c, 0x68, 0x64, 0x50, 0xb0, 0x1c, 0x0e,
	0x06, 0x45, 0x8e, 0x64, 0x22, 0xc4, 0xcf, 0x01, 0x79, 0xd7, 0x06, 0x4c, 0x36, 0x1b, 0xab, 0xbd,
	0xe2, 0x1a, 0xf5, 0x7a, 0x2a, 0xf6, 0xb0, 0xf6, 0xcc, 0xd0, 0xd3, 0x4e, 0x58, 0x1e, 0x82, 0x97,
	0xe0, 0xcc, 0x3b, 0xf0, 0x2e, 0x5c, 0x38, 0xee, 0x23, 0xa0, 0xee, 0xe9, 0xf1, 0xfc, 0xd8, 0x16,
	0x39, 0x79, 0xea, 0xab, 0xaf, 0xbb, 0xba, 0xbe, 0xfa, 0x52, 0x59, 0x80, 0xc5, 0x8a, 0xcb, 0x41,
	0x28, 0x02, 0x19, 0x60, 0x25, 0xe4, 0xfe, 0x92, 0x78, 0xa7, 0x19, 0xff, 0xc6, 0x68, 0xef, 0x77,
	0x68, 0x8f, 0x83, 0xed, 0xcd, 0x9a, 0x18, 0x97, 0x8b, 0x15, 0xc9, 0x17, 0xd1, 0x12, 0xfb, 0xf0,
	0xc0, 0xcd, 0x62, 0xb3, 0xe7, 0x8e, 0xd5, 0xb5, 0xfa, 0x4d, 0x56, 0x84, 0xb1, 0x09, 0x96, 0xef,
	0xd8, 0x5d, 0xab, 0x7f, 0xca, 0x2c, 0x1f, 0x5b, 0x60, 0x87, 0xbe, 0x53, 0xd2, 0xa1, 0x1d, 0xfa,
	0xf8, 0x08, 0x60, 0xe1, 0x85, 0x2b, 0x12, 0xd7, 0xf4, 0x9b, 0x74, 0xca, 0xfa, 0x8a, 0x0c, 0xd2,
	0xfb, 0xdb, 0x86, 0xf6, 0x6c, 0xcd, 0x3d, 0xff, 0x62, 0xc5, 0xe5, 0x0b, 0x8a, 0x22, 0xbe, 0x24,
	0xfc, 0x14, 0x4e, 0xdd, 0x51, 0x18, 0xce, 0xb6, 0x37, 0x6b, 0x6f, 0xf1, 0x9c, 0xee, 0x4c, 0xe9,
	0x3c, 0x88, 0x08, 0x65, 0x79, 0x17, 0x92, 0xae, 0x5d, 0x67, 0xfa, 0x1b, 0x3f, 0x84, 0x4a, 0xc8,
	0x05, 0xdf, 0x44, 0xfa, 0x09, 0x4d, 0x66, 0x22, 0x74, 0xa0, 0xba, 0x89, 0x2f, 0x37, 0x6f, 0x48,
	0x42, 0xfc, 0x18, 0xea, 0x0b, 0x41, 0x5c, 0x92, 0x3b, 0x92, 0xce, 0x49, 0xd7, 0xea, 0x97, 0x58,
	0x0a, 0xa8, 0xac, 0x21, 0x4e, 0xc7, 0x4e, 0x45, 0x17, 0x4a, 0x01, 0x1c, 0x00, 0x46, 0x2b, 0x2e,
	0xc8, 0x9d, 0xd3, 0x42, 0x90, 0x3c, 0xe7, 0x91, 0xa2, 0x55, 0x75, 0x81, 0x03, 0x19, 0xfc, 0x06,
	0x9c, 0x2c, 0x7a, 0xa1, 0xca, 0x78, 0x81, 0x3f, 0xe6, 0x92, 0x9c, 0x9a, 0x2e, 0x7d, 0x34, 0xaf,
	0x3a, 0x78, 0x43, 0x22, 0xf2, 0x02, 0xdf, 0xa9, 0x6b, 0x75, 0x93, 0xb0, 0xf7, 0x4f, 0x09, 0x1a,
	0x05, 0xf5, 0x02, 0x9f, 0xae, 0xbd, 0x0d, 0xcd, 0x04, 0x65, 0xd4, 0xcb, 0x81, 0xd8, 0x83, 0x66,
	0xe4, 0x2d, 0x7d, 0x72, 0x0d, 0xc9, 0xd6, 0xa4, 0x1c, 0xa6, 0x38, 0x14, 0xae, 0x68, 0x43, 0x82,
	0xaf, 0x15, 0x27, 0xd6, 0x34, 0x87, 0xe1, 0x33, 0xf8, 0x20, 0x1b, 0xcf, 0xbd, 0xa5, 0xcf, 0xe5,
	0x56, 0x24, 0x3a, 0x1f, 0x4e, 0x2a, 0x7b, 0x45, 0xe4, 0xbb, 0x24, 0xd4, 0xc3, 0xa7, 0x63, 0x75,
	0xf9, 0x49, 0x6c, 0xaf, 0x02, 0xac, 0x35, 0xcb, 0x43, 0x69, 0x89, 0x8a, 0x3e, 0x72, 0x34, 0x8f,
	0xc3, 0x74, 0xea, 0x6a, 0x28, 0x8d, 0xa1, 0x33, 0x30, 0xc6, 0x2f, 0xfa, 0x3d, 0xf5, 0x43, 0x07,
	0x6a, 0x82, 0x16, 0xe4, 0xbd, 0x21, 0xa1, 0x67, 0xd2, 0x64, 0xbb, 0x58, 0xb9, 0x2b, 0xae, 0xa5,
	0x47, 0xd0, 0x64, 0x26, 0xca, 0xbb, 0x04, 0x74, 0x2a, 0xe3, 0x92, 0xcf, 0xa1, 0xbd, 0x8d, 0xc8,
	0x9d, 0x67, 0x26, 0xeb, 0x34, 0x34, 0x69, 0x0f, 0xcf, 0x4e, 0xb9, 0x99, 0x9f, 0xf2, 0x1f, 0x16,
	0x54, 0xcc, 0x58, 0xda, 0x50, 0xba, 0xdd, 0x8d, 0x55, 0x7d, 0x62, 0x17, 0x1a, 0x9e, 0x4b, 0xbe,
	0xf4, 0xe4, 0x5d, 0x3a, 0xcb, 0x2c, 0x84, 0x43, 0x38, 0xcb, 0x84, 0xa9, 0x84, 0xf1, 0x48, 0x0f,
	0xe6, 0x54, 0x5b, 0xd2, 0xdb, 0xd0, 0x5c, 0xf2, 0x4d, 0xa8, 0xc7, 0x59, 0x62, 0x29, 0xd0, 0xfb,
	0xd7, 0x02, 0x34, 0x96, 0x9b, 0x4b, 0x2e, 0xe9, 0x62, 0xa5, 0xa4, 0xc5, 0xef, 0xa0, 0xe6, 0xd3,
	0x5b, 0x8d, 0xe8, 0x17, 0xb6, 0x86, 0x8f, 0x13, 0xd1, 0xf7, 0xd9, 0x83, 0x2b, 0x43, 0x65, 0xbb,
	0x43, 0x79, 0x31, 0xed, 0xa2, 0x98, 0x08, 0xe5, 0xd7, 0x22, 0xd8, 0x98, 0x77, 0xeb, 0x6f, 0x65,
	0x78, 0xf5, 0x5b, 0xb4, 0x5e, 0x1e, 0xec, 0x7d, 0x0d, 0xb5, 0xa4, 0x1a, 0xd6, 0xa0, 0x3c, 0x9f,
	0x5c, 0x5d, 0xb7, 0xdf, 0xc3, 0x53, 0xa8, 0x8f, 0x27, 0x97, 0xd3, 0x9f, 0x27, 0x6c, 0x32, 0x6e,
	0x5b, 0x78, 0x06, 0xed, 0xef, 0x47, 0xd3, 0xcb, 0xc9, 0xf8, 0xd5, 0xf5, 0xcb, 0x57, 0x3f, 0x8e,
	0xae, 0xc6, 0x97, 0x93, 0xb6, 0xdd, 0xbb, 0xaf, 0x42, 0xeb, 0x9c, 0x2f, 0x6e, 0xc9, 0x77, 0x93,
	0x7f, 0x64, 0x5f, 0x41, 0x55, 0xd0, 0xaf, 0x5b, 0x8a, 0xa4, 0xee, 0xb2, 0x31, 0x7c, 0x94, 0x74,
	0x99, 0x27, 0x0e, 0x58, 0xcc, 0x62, 0x09, 0x1d, 0xbf, 0x55, 0x06, 0x8b, 0xc2, 0xc0, 0x8f, 0xe2,
	0xd5, 0xd5, 0x18, 0x7e, 0x72, 0xf4, 0x68, 0x4c, 0x63, 0xbb, 0x03, 0x78, 0x06, 0x27, 0x24, 0x44,
	0x20, 0x74, 0xff, 0x75, 0x16, 0x07, 0x4a, 0x32, 0x73, 0xfb, 0x74, 0xac, 0x9b, 0xaf, 0xb3, 0x14,
	0xe8, 0xfc, 0x02, 0xe5, 0xd1, 0x56, 0xae, 0x94, 0x7b, 0x65, 0xa0, 0xf4, 0x30, 0xce, 0x31, 0x11,
	0x3e, 0x81, 0x87, 0xc9, 0xf8, 0xd3, 0x8d, 0x1b, 0x0b, 0xbf, 0x9f, 0x50, 0xb5, 0xa2, 0x82, 0x7b,
	0x52, 0xa0, 0xf3, 0x97, 0x0d, 0x55, 0xd3, 0x31, 0x3e, 0x85, 0x9a, 0x99, 0x5b, 0xe4, 0x58, 0xdd,
	0x52, 0xbf, 0x31, 0x7c, 0x3f, 0x69, 0x34, 0xb3, 0xae, 0xd8, 0x8e, 0x84, 0x3f, 0x01, 0x6e, 0xf6,
	0x2c, 0x62, 0x34, 0xea, 0x1c, 0x37, 0x11, 0x3b, 0x70, 0x6a, 0x6f, 0xbd, 0x95, 0x0e, 0xac, 0xb7,
	0x27, 0xf0, 0xd0, 0xa7, 0xb7, 0x2f, 0xb3, 0x6b, 0x31, 0xd2, 0xf2, 0x9d, 0xb2, 0xfd, 0x84, 0xba,
	0x31, 0xd4, 0x9f, 0xe7, 0x5b, 0xdf, 0x5d, 0x93, 0xd9, 0x57, 0x39, 0x0c, 0x9f, 0x42, 0x99, 0x6f,
	0xe5, 0x4a, 0x2f, 0xa6, 0xc6, 0xf0, 0xa3, 0x23, 0x73, 0x55, 0xd3, 0x60, 0x9a, 0xd8, 0xb9, 0xb7,
	0xa0, 0x96, 0x8c, 0x19, 0x7f, 0x28, 0x54, 0x88, 0x8d, 0xf5, 0xf8, 0xc8, 0x2d, 0xb3, 0x0c, 0xb5,
	0xf0, 0x8c, 0x2f, 0xa1, 0x95, 0x2e, 0xfb, 0x5b, 0xd5, 0x95, 0xad, 0xf5, 0x6f, 0x25, 0x57, 0xc5,
	0x67, 0x59, 0x81, 0xb5, 0x7b, 0x7e, 0xe9, 0x1d, 0x9f, 0x8f, 0xc3, 0x82, 0xca, 0x65, 0x7d, 0xb0,
	0x58, 0x26, 0xc7, 0xe9, 0xfc, 0x69, 0x41, 0x33, 0xfb, 0x76, 0xfc, 0x0c, 0xaa, 0xa1, 0x08, 0x5e,
	0x7b, 0xbb, 0x8e, 0x1f, 0xa4, 0xe7, 0x35, 0xcc, 0x92, 0xfc, 0x5e, 0x3d, 0xfb, 0xff, 0xeb, 0xe1,
	0xb3, 0xe2, 0x7f, 0x87, 0xa5, 0x83, 0x87, 0xf2, 0xa4, 0x9b, 0x8a, 0xfe, 0xd3, 0xe8, 0x8b, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x15, 0x2e, 0x15, 0xa5, 0x3e, 0x09, 0x00, 0x00,
}
