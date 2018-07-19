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
	return fileDescriptor_chat_83c6f2403e0938a9, []int{5, 0}
}

type DoubleRatchedMsg struct {
	DoubleRatchetPK      []byte   `protobuf:"bytes,1,opt,name=doubleRatchetPK,proto3" json:"doubleRatchetPK,omitempty"`
	N                    uint32   `protobuf:"varint,2,opt,name=n" json:"n,omitempty"`
	Pn                   uint32   `protobuf:"varint,3,opt,name=pn" json:"pn,omitempty"`
	CipherText           []byte   `protobuf:"bytes,4,opt,name=cipherText,proto3" json:"cipherText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleRatchedMsg) Reset()         { *m = DoubleRatchedMsg{} }
func (m *DoubleRatchedMsg) String() string { return proto.CompactTextString(m) }
func (*DoubleRatchedMsg) ProtoMessage()    {}
func (*DoubleRatchedMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_83c6f2403e0938a9, []int{0}
}
func (m *DoubleRatchedMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRatchedMsg.Unmarshal(m, b)
}
func (m *DoubleRatchedMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRatchedMsg.Marshal(b, m, deterministic)
}
func (dst *DoubleRatchedMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRatchedMsg.Merge(dst, src)
}
func (m *DoubleRatchedMsg) XXX_Size() int {
	return xxx_messageInfo_DoubleRatchedMsg.Size(m)
}
func (m *DoubleRatchedMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRatchedMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRatchedMsg proto.InternalMessageInfo

func (m *DoubleRatchedMsg) GetDoubleRatchetPK() []byte {
	if m != nil {
		return m.DoubleRatchetPK
	}
	return nil
}

func (m *DoubleRatchedMsg) GetN() uint32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *DoubleRatchedMsg) GetPn() uint32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *DoubleRatchedMsg) GetCipherText() []byte {
	if m != nil {
		return m.CipherText
	}
	return nil
}

type PlainChatMessage struct {
	DAppPublicKey            []byte   `protobuf:"bytes,1,opt,name=dAppPublicKey,proto3" json:"dAppPublicKey,omitempty"`
	Type                     string   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Params                   []byte   `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	Message                  []byte   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt                int64    `protobuf:"varint,5,opt,name=createdAt" json:"createdAt,omitempty"`
	MessageID                string   `protobuf:"bytes,6,opt,name=messageID" json:"messageID,omitempty"`
	SharedSecretBaseID       []byte   `protobuf:"bytes,7,opt,name=sharedSecretBaseID,proto3" json:"sharedSecretBaseID,omitempty"`
	SharedSecretCreationDate int64    `protobuf:"varint,8,opt,name=sharedSecretCreationDate" json:"sharedSecretCreationDate,omitempty"`
	Version                  uint32   `protobuf:"varint,9,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *PlainChatMessage) Reset()         { *m = PlainChatMessage{} }
func (m *PlainChatMessage) String() string { return proto.CompactTextString(m) }
func (*PlainChatMessage) ProtoMessage()    {}
func (*PlainChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_83c6f2403e0938a9, []int{1}
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
	Message                  *DoubleRatchedMsg `protobuf:"bytes,8,opt,name=message" json:"message,omitempty"`
	Receiver                 []byte            `protobuf:"bytes,9,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender                   []byte            `protobuf:"bytes,10,opt,name=sender,proto3" json:"sender,omitempty"`
	MessageID                []byte            `protobuf:"bytes,11,opt,name=messageID,proto3" json:"messageID,omitempty"`
	UsedSharedSecret         []byte            `protobuf:"bytes,12,opt,name=usedSharedSecret,proto3" json:"usedSharedSecret,omitempty"`
	Version                  uint32            `protobuf:"varint,13,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_83c6f2403e0938a9, []int{2}
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

func (m *ChatMessage) GetMessage() *DoubleRatchedMsg {
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
	TimeStamp            int64    `protobuf:"varint,4,opt,name=timeStamp" json:"timeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreKey) Reset()         { *m = PreKey{} }
func (m *PreKey) String() string { return proto.CompactTextString(m) }
func (*PreKey) ProtoMessage()    {}
func (*PreKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_83c6f2403e0938a9, []int{3}
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

type PreKeyBundle struct {
	PreKey               *PreKey  `protobuf:"bytes,2,opt,name=preKey" json:"preKey,omitempty"`
	OneTimePreKey        *PreKey  `protobuf:"bytes,3,opt,name=oneTimePreKey" json:"oneTimePreKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreKeyBundle) Reset()         { *m = PreKeyBundle{} }
func (m *PreKeyBundle) String() string { return proto.CompactTextString(m) }
func (*PreKeyBundle) ProtoMessage()    {}
func (*PreKeyBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_83c6f2403e0938a9, []int{4}
}
func (m *PreKeyBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreKeyBundle.Unmarshal(m, b)
}
func (m *PreKeyBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreKeyBundle.Marshal(b, m, deterministic)
}
func (dst *PreKeyBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreKeyBundle.Merge(dst, src)
}
func (m *PreKeyBundle) XXX_Size() int {
	return xxx_messageInfo_PreKeyBundle.Size(m)
}
func (m *PreKeyBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_PreKeyBundle.DiscardUnknown(m)
}

var xxx_messageInfo_PreKeyBundle proto.InternalMessageInfo

func (m *PreKeyBundle) GetPreKey() *PreKey {
	if m != nil {
		return m.PreKey
	}
	return nil
}

func (m *PreKeyBundle) GetOneTimePreKey() *PreKey {
	if m != nil {
		return m.OneTimePreKey
	}
	return nil
}

type MessageStateChange struct {
	NewState             MessageStateChange_NewState `protobuf:"varint,1,opt,name=newState,enum=pangea.MessageStateChange_NewState" json:"newState,omitempty"`
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
	return fileDescriptor_chat_83c6f2403e0938a9, []int{5}
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
	Request              *BackendMessage_Request  `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Response             *BackendMessage_Response `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BackendMessage) Reset()         { *m = BackendMessage{} }
func (m *BackendMessage) String() string { return proto.CompactTextString(m) }
func (*BackendMessage) ProtoMessage()    {}
func (*BackendMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_83c6f2403e0938a9, []int{6}
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
	return fileDescriptor_chat_83c6f2403e0938a9, []int{6, 0}
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
	RequestID            []byte               `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Message              *ChatMessage         `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	MessageStateChange   *MessageStateChange  `protobuf:"bytes,3,opt,name=messageStateChange" json:"messageStateChange,omitempty"`
	PreKey               *PreKey              `protobuf:"bytes,4,opt,name=preKey" json:"preKey,omitempty"`
	NewOneTimePreKeys    uint32               `protobuf:"varint,5,opt,name=newOneTimePreKeys" json:"newOneTimePreKeys,omitempty"`
	PreKeyBundle         []byte               `protobuf:"bytes,6,opt,name=preKeyBundle,proto3" json:"preKeyBundle,omitempty"`
	Auth                 *BackendMessage_Auth `protobuf:"bytes,7,opt,name=auth" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BackendMessage_Request) Reset()         { *m = BackendMessage_Request{} }
func (m *BackendMessage_Request) String() string { return proto.CompactTextString(m) }
func (*BackendMessage_Request) ProtoMessage()    {}
func (*BackendMessage_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_83c6f2403e0938a9, []int{6, 1}
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

func (m *BackendMessage_Request) GetRequestID() []byte {
	if m != nil {
		return m.RequestID
	}
	return nil
}

func (m *BackendMessage_Request) GetMessage() *ChatMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *BackendMessage_Request) GetMessageStateChange() *MessageStateChange {
	if m != nil {
		return m.MessageStateChange
	}
	return nil
}

func (m *BackendMessage_Request) GetPreKey() *PreKey {
	if m != nil {
		return m.PreKey
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
	RequestID            []byte                       `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
	PreKeyBundle         *BackendMessage_PreKeyBundle `protobuf:"bytes,2,opt,name=preKeyBundle" json:"preKeyBundle,omitempty"`
	OneTimePrekeys       []*PreKey                    `protobuf:"bytes,3,rep,name=oneTimePrekeys" json:"oneTimePrekeys,omitempty"`
	Auth                 *BackendMessage_Auth         `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	Error                string                       `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BackendMessage_Response) Reset()         { *m = BackendMessage_Response{} }
func (m *BackendMessage_Response) String() string { return proto.CompactTextString(m) }
func (*BackendMessage_Response) ProtoMessage()    {}
func (*BackendMessage_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_83c6f2403e0938a9, []int{6, 2}
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

func (m *BackendMessage_Response) GetRequestID() []byte {
	if m != nil {
		return m.RequestID
	}
	return nil
}

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

func (m *BackendMessage_Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type BackendMessage_PreKeyBundle struct {
	Profile              *Profile `protobuf:"bytes,1,opt,name=profile" json:"profile,omitempty"`
	SignedPreKey         *PreKey  `protobuf:"bytes,2,opt,name=signedPreKey" json:"signedPreKey,omitempty"`
	OneTimePreKey        *PreKey  `protobuf:"bytes,3,opt,name=oneTimePreKey" json:"oneTimePreKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackendMessage_PreKeyBundle) Reset()         { *m = BackendMessage_PreKeyBundle{} }
func (m *BackendMessage_PreKeyBundle) String() string { return proto.CompactTextString(m) }
func (*BackendMessage_PreKeyBundle) ProtoMessage()    {}
func (*BackendMessage_PreKeyBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_83c6f2403e0938a9, []int{6, 3}
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
	proto.RegisterType((*DoubleRatchedMsg)(nil), "pangea.DoubleRatchedMsg")
	proto.RegisterType((*PlainChatMessage)(nil), "pangea.PlainChatMessage")
	proto.RegisterType((*ChatMessage)(nil), "pangea.ChatMessage")
	proto.RegisterType((*PreKey)(nil), "pangea.PreKey")
	proto.RegisterType((*PreKeyBundle)(nil), "pangea.PreKeyBundle")
	proto.RegisterType((*MessageStateChange)(nil), "pangea.MessageStateChange")
	proto.RegisterType((*BackendMessage)(nil), "pangea.BackendMessage")
	proto.RegisterType((*BackendMessage_Auth)(nil), "pangea.BackendMessage.Auth")
	proto.RegisterType((*BackendMessage_Request)(nil), "pangea.BackendMessage.Request")
	proto.RegisterType((*BackendMessage_Response)(nil), "pangea.BackendMessage.Response")
	proto.RegisterType((*BackendMessage_PreKeyBundle)(nil), "pangea.BackendMessage.PreKeyBundle")
	proto.RegisterEnum("pangea.MessageStateChange_NewState", MessageStateChange_NewState_name, MessageStateChange_NewState_value)
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_chat_83c6f2403e0938a9) }

var fileDescriptor_chat_83c6f2403e0938a9 = []byte{
	// 935 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x8e, 0xe3, 0x44,
	0x10, 0xc5, 0x49, 0x26, 0xc9, 0x54, 0x2e, 0x9b, 0x6d, 0x06, 0x64, 0x05, 0xb4, 0x44, 0x59, 0x84,
	0x06, 0xb4, 0x04, 0x29, 0xac, 0x10, 0x97, 0x07, 0x94, 0x19, 0x07, 0x08, 0x3b, 0x3b, 0x3b, 0xea,
	0x8c, 0x78, 0x5d, 0xf5, 0xc4, 0xb5, 0xb1, 0x99, 0xa4, 0x6d, 0xda, 0x9d, 0x5d, 0x86, 0x8f, 0xe0,
	0x27, 0xf8, 0x11, 0x7e, 0x80, 0x3f, 0xe0, 0x8d, 0x17, 0x1e, 0xf9, 0x04, 0xd4, 0x17, 0xc7, 0x97,
	0x24, 0x80, 0x78, 0x8a, 0xeb, 0xd4, 0x69, 0x57, 0xd7, 0xa9, 0xd3, 0xed, 0x00, 0x2c, 0x02, 0x26,
	0x47, 0xb1, 0x88, 0x64, 0x44, 0xea, 0x31, 0xe3, 0x4b, 0x64, 0xfd, 0xb6, 0xf9, 0x35, 0xe8, 0xf0,
	0x27, 0xe8, 0x79, 0xd1, 0xe6, 0x66, 0x85, 0x94, 0xc9, 0x45, 0x80, 0xfe, 0xd3, 0x64, 0x49, 0x4e,
	0xe1, 0x9e, 0x9f, 0xc3, 0xe4, 0xd5, 0x13, 0xd7, 0x19, 0x38, 0xa7, 0x6d, 0x5a, 0x86, 0x49, 0x1b,
	0x1c, 0xee, 0x56, 0x06, 0xce, 0x69, 0x87, 0x3a, 0x9c, 0x74, 0xa1, 0x12, 0x73, 0xb7, 0xaa, 0xc3,
	0x4a, 0xcc, 0xc9, 0x03, 0x80, 0x45, 0x18, 0x07, 0x28, 0xae, 0xf1, 0x47, 0xe9, 0xd6, 0xf4, 0x2b,
	0x72, 0xc8, 0xf0, 0xd7, 0x0a, 0xf4, 0xae, 0x56, 0x2c, 0xe4, 0xe7, 0x01, 0x93, 0x4f, 0x31, 0x49,
	0xd8, 0x12, 0xc9, 0xbb, 0xd0, 0xf1, 0x27, 0x71, 0x7c, 0xb5, 0xb9, 0x59, 0x85, 0x8b, 0x27, 0x78,
	0x67, 0x4b, 0x17, 0x41, 0x42, 0xa0, 0x26, 0xef, 0x62, 0xd4, 0xb5, 0x8f, 0xa9, 0x7e, 0x26, 0x6f,
	0x42, 0x3d, 0x66, 0x82, 0xad, 0x13, 0xbd, 0x85, 0x36, 0xb5, 0x11, 0x71, 0xa1, 0xb1, 0x36, 0x2f,
	0xb7, 0x7b, 0x48, 0x43, 0xf2, 0x36, 0x1c, 0x2f, 0x04, 0x32, 0x89, 0xfe, 0x44, 0xba, 0x47, 0x03,
	0xe7, 0xb4, 0x4a, 0x33, 0x40, 0x65, 0x2d, 0x71, 0xe6, 0xb9, 0x75, 0x5d, 0x28, 0x03, 0xc8, 0x08,
	0x48, 0x12, 0x30, 0x81, 0xfe, 0x1c, 0x17, 0x02, 0xe5, 0x19, 0x4b, 0x14, 0xad, 0xa1, 0x0b, 0xec,
	0xc9, 0x90, 0xcf, 0xc1, 0xcd, 0xa3, 0xe7, 0xaa, 0x4c, 0x18, 0x71, 0x8f, 0x49, 0x74, 0x9b, 0xba,
	0xf4, 0xc1, 0xbc, 0xea, 0xe0, 0x25, 0x8a, 0x24, 0x8c, 0xb8, 0x7b, 0xac, 0xd5, 0x4d, 0xc3, 0xe1,
	0x1f, 0x55, 0x68, 0x95, 0xd4, 0x8b, 0x38, 0x5e, 0x87, 0x6b, 0xbc, 0x12, 0x98, 0x53, 0xaf, 0x00,
	0x92, 0x21, 0xb4, 0x93, 0x70, 0xc9, 0xd1, 0xb7, 0xa4, 0x8a, 0x26, 0x15, 0x30, 0xc5, 0xc1, 0x38,
	0xc0, 0x35, 0x0a, 0xb6, 0x52, 0x1c, 0xa3, 0x69, 0x01, 0x23, 0x8f, 0xe1, 0x8d, 0x7c, 0x3c, 0x0f,
	0x97, 0x9c, 0xc9, 0x8d, 0x48, 0x75, 0xde, 0x9f, 0x54, 0xf6, 0x4a, 0x90, 0xfb, 0x28, 0xd4, 0xc6,
	0x67, 0x9e, 0x7a, 0xf9, 0x91, 0xb1, 0x57, 0x09, 0xd6, 0x9a, 0x15, 0xa1, 0xac, 0x44, 0x5d, 0x2f,
	0x39, 0x98, 0x27, 0xe3, 0x6c, 0xea, 0x4a, 0xde, 0xd6, 0xd8, 0x1d, 0x59, 0xe3, 0x97, 0xfd, 0x9e,
	0xf9, 0xa1, 0x0f, 0x4d, 0x81, 0x0b, 0x0c, 0x5f, 0xa2, 0xd0, 0x42, 0xb7, 0xe9, 0x36, 0x56, 0xee,
	0x32, 0xb5, 0x5c, 0x30, 0xee, 0x32, 0x51, 0xd1, 0x25, 0x2d, 0x9d, 0xca, 0xb9, 0xe4, 0x03, 0xe8,
	0x6d, 0x12, 0xf4, 0xe7, 0xb9, 0xc9, 0xba, 0x6d, 0x4d, 0xda, 0xc1, 0xf3, 0x53, 0xee, 0x14, 0xa7,
	0xfc, 0xb3, 0x03, 0x75, 0x3b, 0x96, 0x1e, 0x54, 0x6f, 0xb7, 0x63, 0x55, 0x8f, 0x64, 0x00, 0xad,
	0xd0, 0x47, 0x2e, 0x43, 0x79, 0x97, 0xcd, 0x32, 0x0f, 0x91, 0x31, 0x9c, 0xe4, 0xc2, 0x4c, 0x42,
	0x33, 0xd2, 0xbd, 0x39, 0xd5, 0x96, 0x0c, 0xd7, 0x38, 0x97, 0x6c, 0x1d, 0xeb, 0x71, 0x56, 0x69,
	0x06, 0x0c, 0x57, 0xd0, 0x36, 0xfb, 0x39, 0xdb, 0x70, 0x7f, 0x85, 0xe4, 0x3d, 0xa8, 0xc7, 0x99,
	0x95, 0x5a, 0xe3, 0x6e, 0xaa, 0xb5, 0x61, 0x51, 0x9b, 0x25, 0x8f, 0xcb, 0xf6, 0xac, 0xee, 0xa5,
	0x17, 0x49, 0xc3, 0x3f, 0x1d, 0x20, 0xd6, 0xe0, 0x73, 0xc9, 0x24, 0x9e, 0x07, 0x8a, 0x4d, 0xbe,
	0x84, 0x26, 0xc7, 0x57, 0x1a, 0xd1, 0x7a, 0x74, 0xc7, 0x0f, 0xd3, 0xf7, 0xec, 0xb2, 0x47, 0x97,
	0x96, 0x4a, 0xb7, 0x8b, 0x8a, 0xa3, 0xab, 0x94, 0x47, 0x47, 0xa0, 0xf6, 0x42, 0x44, 0x6b, 0xab,
	0x92, 0x7e, 0x56, 0xc7, 0x4b, 0xfd, 0x96, 0x8d, 0x5e, 0x04, 0x87, 0x9f, 0x41, 0x33, 0xad, 0x46,
	0x9a, 0x50, 0x9b, 0x4f, 0x2f, 0xaf, 0x7b, 0xaf, 0x91, 0x0e, 0x1c, 0x7b, 0xd3, 0x8b, 0xd9, 0x77,
	0x53, 0x3a, 0xf5, 0x7a, 0x0e, 0x39, 0x81, 0xde, 0x57, 0x93, 0xd9, 0xc5, 0xd4, 0x7b, 0x7e, 0xfd,
	0xec, 0xf9, 0x37, 0x93, 0x4b, 0xef, 0x62, 0xda, 0xab, 0x0c, 0x7f, 0x6f, 0x40, 0xf7, 0x8c, 0x2d,
	0x6e, 0x91, 0xfb, 0xe9, 0x91, 0xfe, 0x14, 0x1a, 0x02, 0x7f, 0xd8, 0x60, 0x22, 0x75, 0x97, 0xad,
	0xf1, 0x83, 0xb4, 0xcb, 0x22, 0x71, 0x44, 0x0d, 0x8b, 0xa6, 0x74, 0xf2, 0x85, 0xb2, 0x73, 0x12,
	0x47, 0x3c, 0x41, 0x3b, 0x97, 0x77, 0x0e, 0x2e, 0x35, 0x34, 0xba, 0x5d, 0xd0, 0xff, 0x1e, 0x6a,
	0x93, 0x8d, 0x0c, 0x94, 0xef, 0x65, 0xa4, 0x7a, 0xb3, 0x9e, 0xb3, 0x11, 0x79, 0x04, 0xf7, 0x53,
	0xe3, 0x64, 0x77, 0xb5, 0x11, 0x71, 0x37, 0xa1, 0xa4, 0x4e, 0x4a, 0xbe, 0xcb, 0x80, 0xfe, 0x6f,
	0x15, 0x68, 0xd8, 0xdd, 0x2b, 0xa6, 0xdd, 0xff, 0xcc, 0xb3, 0x25, 0x33, 0x80, 0x7c, 0x98, 0x9d,
	0x6a, 0xd3, 0xd1, 0xeb, 0x69, 0x47, 0xb9, 0x5b, 0x30, 0x3b, 0xd0, 0xdf, 0x02, 0x59, 0xef, 0x58,
	0xc1, 0x9a, 0xae, 0x7f, 0xd8, 0x2c, 0x74, 0xcf, 0xaa, 0x9c, 0xc7, 0x6b, 0xff, 0xe8, 0xf1, 0x47,
	0x70, 0x9f, 0xe3, 0xab, 0x67, 0x79, 0x07, 0x27, 0xfa, 0x82, 0xeb, 0xd0, 0xdd, 0x84, 0xba, 0x66,
	0xe3, 0xdc, 0x49, 0xb2, 0xd7, 0x5a, 0x01, 0x23, 0x1f, 0x41, 0x8d, 0x6d, 0x64, 0xa0, 0x3f, 0x2e,
	0xad, 0xf1, 0x5b, 0x07, 0x66, 0xa8, 0xa6, 0x45, 0x35, 0xb1, 0xff, 0x97, 0x03, 0xcd, 0x74, 0xa4,
	0xff, 0x22, 0xe8, 0xd7, 0xa5, 0xfa, 0x46, 0xd5, 0x87, 0x07, 0x6a, 0xe4, 0x0f, 0x7d, 0x69, 0x93,
	0x9f, 0x40, 0x37, 0x3b, 0xb5, 0xb7, 0xaa, 0xe7, 0xea, 0xa0, 0xba, 0x47, 0xa6, 0x12, 0x6b, 0xdb,
	0x5c, 0xed, 0x3f, 0x36, 0x47, 0x4e, 0xe0, 0x08, 0x85, 0x88, 0x84, 0xd6, 0xf4, 0x98, 0x9a, 0xa0,
	0xff, 0x8b, 0x53, 0xba, 0x92, 0xde, 0x87, 0x46, 0x2c, 0xa2, 0x17, 0xe1, 0x0a, 0xed, 0xb1, 0xb9,
	0x97, 0x6d, 0x44, 0xc3, 0x34, 0xcd, 0x93, 0xf1, 0x9e, 0xcf, 0xe1, 0xee, 0xc6, 0x8b, 0x9f, 0xc7,
	0xff, 0x75, 0x93, 0xdd, 0xd4, 0xf5, 0x9f, 0xae, 0x8f, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x09,
	0xdd, 0x28, 0x29, 0x98, 0x09, 0x00, 0x00,
}
