// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgs.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type RoundState int32

const (
	RoundState_NotStarted  RoundState = 0
	RoundState_PrePrepare  RoundState = 1
	RoundState_Prepare     RoundState = 2
	RoundState_Commit      RoundState = 3
	RoundState_ChangeRound RoundState = 4
	RoundState_Decided     RoundState = 5
	RoundState_Stopped     RoundState = 6
)

var RoundState_name = map[int32]string{
	0: "NotStarted",
	1: "PrePrepare",
	2: "Prepare",
	3: "Commit",
	4: "ChangeRound",
	5: "Decided",
	6: "Stopped",
}

var RoundState_value = map[string]int32{
	"NotStarted":  0,
	"PrePrepare":  1,
	"Prepare":     2,
	"Commit":      3,
	"ChangeRound": 4,
	"Decided":     5,
	"Stopped":     6,
}

func (x RoundState) String() string {
	return proto.EnumName(RoundState_name, int32(x))
}

func (RoundState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_952909143bb80d72, []int{0}
}

type Message struct {
	Type   RoundState `protobuf:"varint,1,opt,name=type,proto3,enum=proto.RoundState" json:"type,omitempty"`
	Round  uint64     `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Lambda []byte     `protobuf:"bytes,3,opt,name=lambda,proto3" json:"lambda,omitempty"`
	// sequence number is an incremental number for each instance, much like a block number would be in a blockchain
	SeqNumber            uint64   `protobuf:"varint,4,opt,name=seq_number,json=seqNumber,proto3" json:"seq_number,omitempty"`
	Value                []byte   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_952909143bb80d72, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() RoundState {
	if m != nil {
		return m.Type
	}
	return RoundState_NotStarted
}

func (m *Message) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *Message) GetLambda() []byte {
	if m != nil {
		return m.Lambda
	}
	return nil
}

func (m *Message) GetSeqNumber() uint64 {
	if m != nil {
		return m.SeqNumber
	}
	return 0
}

func (m *Message) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SignedMessage struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	SignerIds            []uint64 `protobuf:"varint,3,rep,packed,name=signer_ids,json=signerIds,proto3" json:"signer_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedMessage) Reset()         { *m = SignedMessage{} }
func (m *SignedMessage) String() string { return proto.CompactTextString(m) }
func (*SignedMessage) ProtoMessage()    {}
func (*SignedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_952909143bb80d72, []int{1}
}

func (m *SignedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedMessage.Unmarshal(m, b)
}
func (m *SignedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedMessage.Marshal(b, m, deterministic)
}
func (m *SignedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedMessage.Merge(m, src)
}
func (m *SignedMessage) XXX_Size() int {
	return xxx_messageInfo_SignedMessage.Size(m)
}
func (m *SignedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SignedMessage proto.InternalMessageInfo

func (m *SignedMessage) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SignedMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedMessage) GetSignerIds() []uint64 {
	if m != nil {
		return m.SignerIds
	}
	return nil
}

type ChangeRoundData struct {
	PreparedRound        uint64   `protobuf:"varint,1,opt,name=prepared_round,json=preparedRound,proto3" json:"prepared_round,omitempty"`
	PreparedValue        []byte   `protobuf:"bytes,2,opt,name=prepared_value,json=preparedValue,proto3" json:"prepared_value,omitempty"`
	JustificationMsg     *Message `protobuf:"bytes,3,opt,name=justification_msg,json=justificationMsg,proto3" json:"justification_msg,omitempty"`
	JustificationSig     []byte   `protobuf:"bytes,4,opt,name=justification_sig,json=justificationSig,proto3" json:"justification_sig,omitempty"`
	SignerIds            []uint64 `protobuf:"varint,5,rep,packed,name=signer_ids,json=signerIds,proto3" json:"signer_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeRoundData) Reset()         { *m = ChangeRoundData{} }
func (m *ChangeRoundData) String() string { return proto.CompactTextString(m) }
func (*ChangeRoundData) ProtoMessage()    {}
func (*ChangeRoundData) Descriptor() ([]byte, []int) {
	return fileDescriptor_952909143bb80d72, []int{2}
}

func (m *ChangeRoundData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeRoundData.Unmarshal(m, b)
}
func (m *ChangeRoundData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeRoundData.Marshal(b, m, deterministic)
}
func (m *ChangeRoundData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeRoundData.Merge(m, src)
}
func (m *ChangeRoundData) XXX_Size() int {
	return xxx_messageInfo_ChangeRoundData.Size(m)
}
func (m *ChangeRoundData) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeRoundData.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeRoundData proto.InternalMessageInfo

func (m *ChangeRoundData) GetPreparedRound() uint64 {
	if m != nil {
		return m.PreparedRound
	}
	return 0
}

func (m *ChangeRoundData) GetPreparedValue() []byte {
	if m != nil {
		return m.PreparedValue
	}
	return nil
}

func (m *ChangeRoundData) GetJustificationMsg() *Message {
	if m != nil {
		return m.JustificationMsg
	}
	return nil
}

func (m *ChangeRoundData) GetJustificationSig() []byte {
	if m != nil {
		return m.JustificationSig
	}
	return nil
}

func (m *ChangeRoundData) GetSignerIds() []uint64 {
	if m != nil {
		return m.SignerIds
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.RoundState", RoundState_name, RoundState_value)
	proto.RegisterType((*Message)(nil), "proto.Message")
	proto.RegisterType((*SignedMessage)(nil), "proto.SignedMessage")
	proto.RegisterType((*ChangeRoundData)(nil), "proto.ChangeRoundData")
}

func init() {
	proto.RegisterFile("msgs.proto", fileDescriptor_952909143bb80d72)
}

var fileDescriptor_952909143bb80d72 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x9d, 0xd7, 0xa4, 0x65, 0xb7, 0x5d, 0x97, 0x59, 0x08, 0x45, 0x48, 0x88, 0xaa, 0xd2, 0xa4,
	0x0a, 0x44, 0x26, 0xc6, 0x17, 0xb0, 0xed, 0x85, 0x87, 0x4d, 0x53, 0x22, 0xf1, 0xc0, 0x4b, 0xe5,
	0xd6, 0x77, 0x9e, 0xd1, 0x1c, 0x67, 0xb6, 0x83, 0xc4, 0x2b, 0xbf, 0xc0, 0x4f, 0xf1, 0x15, 0x7c,
	0x08, 0x4f, 0x28, 0xd7, 0x29, 0x2b, 0x15, 0x3c, 0xd9, 0xe7, 0xf8, 0x5c, 0xe7, 0x9c, 0x13, 0x03,
	0x18, 0xaf, 0x7c, 0xd1, 0x38, 0x1b, 0x2c, 0x4f, 0x69, 0x79, 0xfe, 0x46, 0xe9, 0x70, 0xd7, 0xae,
	0x8a, 0xb5, 0x35, 0xa7, 0xca, 0x2a, 0x7b, 0x4a, 0xf4, 0xaa, 0xbd, 0x25, 0x44, 0x80, 0x76, 0x71,
	0x6a, 0xfe, 0x9d, 0xc1, 0xe8, 0x0a, 0xbd, 0x17, 0x0a, 0xf9, 0x09, 0x24, 0xe1, 0x6b, 0x83, 0x39,
	0x9b, 0xb1, 0xc5, 0xf4, 0xec, 0x38, 0x2a, 0x8a, 0xd2, 0xb6, 0xb5, 0xac, 0x82, 0x08, 0x58, 0xd2,
	0x31, 0x7f, 0x0a, 0xa9, 0xeb, 0xb8, 0x7c, 0x7f, 0xc6, 0x16, 0x49, 0x19, 0x01, 0x7f, 0x06, 0xc3,
	0x7b, 0x61, 0x56, 0x52, 0xe4, 0x83, 0x19, 0x5b, 0x4c, 0xca, 0x1e, 0xf1, 0x17, 0x00, 0x1e, 0x1f,
	0x96, 0x75, 0x6b, 0x56, 0xe8, 0xf2, 0x84, 0x46, 0x0e, 0x3c, 0x3e, 0x5c, 0x13, 0xd1, 0x5d, 0xf6,
	0x45, 0xdc, 0xb7, 0x98, 0xa7, 0x34, 0x15, 0xc1, 0xfc, 0x1b, 0x83, 0xc3, 0x4a, 0xab, 0x1a, 0xe5,
	0xc6, 0x5b, 0x01, 0x23, 0x13, 0xb7, 0x64, 0x6f, 0x7c, 0x36, 0xed, 0xed, 0xf5, 0x82, 0xf3, 0xe4,
	0xc7, 0xcf, 0x97, 0x7b, 0xe5, 0x46, 0xc4, 0xe7, 0x70, 0xe0, 0xb5, 0xaa, 0x45, 0x68, 0x1d, 0x92,
	0xd1, 0x49, 0xaf, 0x78, 0xa4, 0xc9, 0x5a, 0xf7, 0x11, 0xb7, 0xd4, 0xd2, 0xe7, 0x83, 0xd9, 0x80,
	0xac, 0x11, 0xf3, 0x41, 0xfa, 0xf9, 0x2f, 0x06, 0x47, 0x17, 0x77, 0xa2, 0x56, 0x48, 0x15, 0x5c,
	0x8a, 0x20, 0xf8, 0x09, 0x4c, 0x1b, 0x87, 0x8d, 0x70, 0x28, 0x97, 0xb1, 0x04, 0x46, 0x89, 0x0e,
	0x37, 0x2c, 0x49, 0xf9, 0xeb, 0x2d, 0x59, 0x8c, 0xb7, 0x6d, 0xe1, 0x8f, 0xf8, 0x63, 0x77, 0xc4,
	0xdf, 0xc3, 0xf1, 0xe7, 0xd6, 0x07, 0x7d, 0xab, 0xd7, 0x22, 0x68, 0x5b, 0x2f, 0x8d, 0x57, 0x54,
	0xe2, 0xff, 0x42, 0x66, 0x7f, 0xc9, 0xaf, 0xbc, 0xe2, 0x6f, 0x77, 0xaf, 0xf0, 0x5a, 0x51, 0xd7,
	0x93, 0x7f, 0x8e, 0x54, 0x5a, 0xed, 0x84, 0x4f, 0x77, 0xc2, 0xbf, 0x6a, 0x00, 0x1e, 0x7f, 0x3c,
	0x9f, 0x02, 0x5c, 0xdb, 0x50, 0x05, 0xe1, 0x02, 0xca, 0x6c, 0xaf, 0xc3, 0x37, 0x0e, 0x6f, 0x62,
	0x8c, 0x8c, 0xf1, 0x31, 0x8c, 0x36, 0x60, 0x9f, 0x03, 0x0c, 0x2f, 0xac, 0x31, 0x3a, 0x64, 0x03,
	0x7e, 0x04, 0xe3, 0xad, 0x0a, 0xb3, 0xa4, 0x53, 0x5e, 0xe2, 0x5a, 0x4b, 0x94, 0x59, 0xda, 0x81,
	0x2a, 0xd8, 0xa6, 0x41, 0x99, 0x0d, 0xcf, 0xe1, 0xd3, 0x93, 0xa2, 0xe8, 0x5f, 0xeb, 0x90, 0x96,
	0x77, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x0e, 0xe1, 0x4e, 0xe0, 0x02, 0x00, 0x00,
}
