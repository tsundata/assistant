// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chatbot.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ChatbotRequest struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *ChatbotRequest) Reset()         { *m = ChatbotRequest{} }
func (m *ChatbotRequest) String() string { return proto.CompactTextString(m) }
func (*ChatbotRequest) ProtoMessage()    {}
func (*ChatbotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{0}
}
func (m *ChatbotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatbotRequest.Unmarshal(m, b)
}
func (m *ChatbotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatbotRequest.Marshal(b, m, deterministic)
}
func (m *ChatbotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatbotRequest.Merge(m, src)
}
func (m *ChatbotRequest) XXX_Size() int {
	return xxx_messageInfo_ChatbotRequest.Size(m)
}
func (m *ChatbotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatbotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChatbotRequest proto.InternalMessageInfo

func (m *ChatbotRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ChatbotReply struct {
	Text []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
}

func (m *ChatbotReply) Reset()         { *m = ChatbotReply{} }
func (m *ChatbotReply) String() string { return proto.CompactTextString(m) }
func (*ChatbotReply) ProtoMessage()    {}
func (*ChatbotReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{1}
}
func (m *ChatbotReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatbotReply.Unmarshal(m, b)
}
func (m *ChatbotReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatbotReply.Marshal(b, m, deterministic)
}
func (m *ChatbotReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatbotReply.Merge(m, src)
}
func (m *ChatbotReply) XXX_Size() int {
	return xxx_messageInfo_ChatbotReply.Size(m)
}
func (m *ChatbotReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatbotReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChatbotReply proto.InternalMessageInfo

func (m *ChatbotReply) GetText() []string {
	if m != nil {
		return m.Text
	}
	return nil
}

type BotSettingRequest struct {
	Kvs []*KV `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (m *BotSettingRequest) Reset()         { *m = BotSettingRequest{} }
func (m *BotSettingRequest) String() string { return proto.CompactTextString(m) }
func (*BotSettingRequest) ProtoMessage()    {}
func (*BotSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{2}
}
func (m *BotSettingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotSettingRequest.Unmarshal(m, b)
}
func (m *BotSettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotSettingRequest.Marshal(b, m, deterministic)
}
func (m *BotSettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotSettingRequest.Merge(m, src)
}
func (m *BotSettingRequest) XXX_Size() int {
	return xxx_messageInfo_BotSettingRequest.Size(m)
}
func (m *BotSettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BotSettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BotSettingRequest proto.InternalMessageInfo

func (m *BotSettingRequest) GetKvs() []*KV {
	if m != nil {
		return m.Kvs
	}
	return nil
}

type Bot struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: db:"uuid"
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// @inject_tag: db:"name"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: db:"avatar"
	Avatar string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *Bot) Reset()         { *m = Bot{} }
func (m *Bot) String() string { return proto.CompactTextString(m) }
func (*Bot) ProtoMessage()    {}
func (*Bot) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{3}
}
func (m *Bot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bot.Unmarshal(m, b)
}
func (m *Bot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bot.Marshal(b, m, deterministic)
}
func (m *Bot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bot.Merge(m, src)
}
func (m *Bot) XXX_Size() int {
	return xxx_messageInfo_Bot.Size(m)
}
func (m *Bot) XXX_DiscardUnknown() {
	xxx_messageInfo_Bot.DiscardUnknown(m)
}

var xxx_messageInfo_Bot proto.InternalMessageInfo

func (m *Bot) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Bot) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Bot) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bot) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Bot) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Bot) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type BotRequest struct {
	Bot *Bot `protobuf:"bytes,1,opt,name=bot,proto3" json:"bot,omitempty"`
}

func (m *BotRequest) Reset()         { *m = BotRequest{} }
func (m *BotRequest) String() string { return proto.CompactTextString(m) }
func (*BotRequest) ProtoMessage()    {}
func (*BotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{4}
}
func (m *BotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotRequest.Unmarshal(m, b)
}
func (m *BotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotRequest.Marshal(b, m, deterministic)
}
func (m *BotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotRequest.Merge(m, src)
}
func (m *BotRequest) XXX_Size() int {
	return xxx_messageInfo_BotRequest.Size(m)
}
func (m *BotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BotRequest proto.InternalMessageInfo

func (m *BotRequest) GetBot() *Bot {
	if m != nil {
		return m.Bot
	}
	return nil
}

type BotReply struct {
	Bot *Bot `protobuf:"bytes,1,opt,name=bot,proto3" json:"bot,omitempty"`
}

func (m *BotReply) Reset()         { *m = BotReply{} }
func (m *BotReply) String() string { return proto.CompactTextString(m) }
func (*BotReply) ProtoMessage()    {}
func (*BotReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{5}
}
func (m *BotReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotReply.Unmarshal(m, b)
}
func (m *BotReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotReply.Marshal(b, m, deterministic)
}
func (m *BotReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotReply.Merge(m, src)
}
func (m *BotReply) XXX_Size() int {
	return xxx_messageInfo_BotReply.Size(m)
}
func (m *BotReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BotReply.DiscardUnknown(m)
}

var xxx_messageInfo_BotReply proto.InternalMessageInfo

func (m *BotReply) GetBot() *Bot {
	if m != nil {
		return m.Bot
	}
	return nil
}

type BotsReply struct {
	Bots []*Bot `protobuf:"bytes,1,rep,name=bots,proto3" json:"bots,omitempty"`
}

func (m *BotsReply) Reset()         { *m = BotsReply{} }
func (m *BotsReply) String() string { return proto.CompactTextString(m) }
func (*BotsReply) ProtoMessage()    {}
func (*BotsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{6}
}
func (m *BotsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotsReply.Unmarshal(m, b)
}
func (m *BotsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotsReply.Marshal(b, m, deterministic)
}
func (m *BotsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotsReply.Merge(m, src)
}
func (m *BotsReply) XXX_Size() int {
	return xxx_messageInfo_BotsReply.Size(m)
}
func (m *BotsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BotsReply.DiscardUnknown(m)
}

var xxx_messageInfo_BotsReply proto.InternalMessageInfo

func (m *BotsReply) GetBots() []*Bot {
	if m != nil {
		return m.Bots
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatbotRequest)(nil), "pb.ChatbotRequest")
	proto.RegisterType((*ChatbotReply)(nil), "pb.ChatbotReply")
	proto.RegisterType((*BotSettingRequest)(nil), "pb.BotSettingRequest")
	proto.RegisterType((*Bot)(nil), "pb.Bot")
	proto.RegisterType((*BotRequest)(nil), "pb.BotRequest")
	proto.RegisterType((*BotReply)(nil), "pb.BotReply")
	proto.RegisterType((*BotsReply)(nil), "pb.BotsReply")
}

func init() { proto.RegisterFile("chatbot.proto", fileDescriptor_acc44097314201ac) }

var fileDescriptor_acc44097314201ac = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0xf3, 0x51, 0x53, 0x7b, 0xfa, 0x41, 0x1d, 0x54, 0x62, 0xd4, 0x20, 0x41, 0xb1, 0x08,
	0xb6, 0x52, 0x97, 0xae, 0x1a, 0x17, 0x0a, 0xee, 0x52, 0x74, 0xe1, 0x46, 0x66, 0x9a, 0xa1, 0x06,
	0x6b, 0x67, 0x6c, 0x4e, 0x8a, 0x7d, 0x0b, 0xc1, 0x97, 0xba, 0xcb, 0xde, 0xdd, 0xdd, 0x96, 0xbe,
	0xc8, 0x65, 0x3e, 0x92, 0xde, 0x72, 0xe1, 0xee, 0xce, 0xf9, 0x9d, 0xdf, 0x9c, 0x09, 0xff, 0x09,
	0xf4, 0x17, 0x3f, 0x29, 0x32, 0x81, 0x63, 0xb9, 0x11, 0x28, 0x88, 0x27, 0x59, 0x04, 0x8c, 0x96,
	0xdc, 0xf4, 0xd1, 0xc3, 0xa5, 0x58, 0x0a, 0x5d, 0x4e, 0x54, 0x65, 0x68, 0xf2, 0x12, 0x06, 0x1f,
	0xcd, 0xb1, 0x8c, 0xff, 0xa9, 0x78, 0x89, 0x84, 0x40, 0x0b, 0xf9, 0x5f, 0x0c, 0xdd, 0x17, 0xee,
	0xa8, 0x93, 0xe9, 0x3a, 0x49, 0xa0, 0xd7, 0x58, 0x72, 0xb5, 0xbb, 0xe1, 0xf8, 0x8d, 0xf3, 0x16,
	0x1e, 0xa4, 0x02, 0xe7, 0x1c, 0xb1, 0x58, 0x2f, 0xeb, 0x65, 0x21, 0xf8, 0xbf, 0xb6, 0xa5, 0xf6,
	0xba, 0xd3, 0x60, 0x2c, 0xd9, 0xf8, 0xcb, 0xb7, 0x4c, 0xa1, 0xe4, 0xbf, 0x0b, 0x7e, 0x2a, 0x90,
	0x0c, 0xc0, 0x2b, 0x72, 0x7d, 0x99, 0x9f, 0x79, 0x45, 0xae, 0x56, 0x57, 0x55, 0x91, 0x87, 0x9e,
	0xb9, 0x5e, 0xd5, 0x8a, 0xad, 0xe9, 0x6f, 0x1e, 0xb6, 0x0c, 0x53, 0x35, 0x79, 0x0c, 0x01, 0xdd,
	0x52, 0xa4, 0x9b, 0xf0, 0x9e, 0xa6, 0xb6, 0x23, 0xcf, 0x01, 0x16, 0x1b, 0x4e, 0x91, 0xe7, 0x3f,
	0x28, 0x86, 0x81, 0xde, 0xdb, 0xb1, 0x64, 0x86, 0x6a, 0x5c, 0xc9, 0xbc, 0x1e, 0xb7, 0xcd, 0xd8,
	0x92, 0x19, 0x26, 0xaf, 0x01, 0xd2, 0x53, 0x14, 0x4f, 0xc0, 0x67, 0xc2, 0x24, 0xd1, 0x9d, 0xb6,
	0xd5, 0xd7, 0xab, 0xa1, 0x62, 0xc9, 0x2b, 0xb8, 0x9f, 0xd6, 0x69, 0xdc, 0xa1, 0x8d, 0xa0, 0x93,
	0x0a, 0x2c, 0x8d, 0xf7, 0x14, 0x5a, 0x4c, 0x60, 0x9d, 0x46, 0x23, 0x6a, 0x38, 0xbd, 0x74, 0x01,
	0x6c, 0xc6, 0xf3, 0xed, 0x82, 0xbc, 0x83, 0xe0, 0x33, 0x5d, 0xe7, 0x2b, 0x4e, 0x88, 0xf2, 0xce,
	0xdf, 0x28, 0x1a, 0x9e, 0x31, 0xb9, 0xda, 0x25, 0x0e, 0x19, 0x41, 0xf0, 0x89, 0xa3, 0x8e, 0xb4,
	0xde, 0x6c, 0xed, 0x5e, 0xd3, 0x1b, 0xf3, 0x0d, 0xb4, 0x8d, 0x59, 0xde, 0x52, 0xfb, 0xb6, 0x2f,
	0x6b, 0xf7, 0x03, 0x0c, 0xbf, 0xea, 0x74, 0x4e, 0x6f, 0x4b, 0x1e, 0x59, 0xe9, 0xfc, 0xad, 0x23,
	0xbd, 0x6b, 0x8e, 0x14, 0xb9, 0x3d, 0x9c, 0x3e, 0xbb, 0x38, 0xc4, 0xee, 0xfe, 0x10, 0x3b, 0xff,
	0x8e, 0xb1, 0xb3, 0x3f, 0xc6, 0xce, 0xd5, 0x31, 0x76, 0xbe, 0x07, 0x54, 0x16, 0x13, 0xc9, 0x58,
	0xa0, 0xff, 0xc0, 0xf7, 0xd7, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0x5a, 0xd3, 0xcb, 0xb8, 0x02,
	0x00, 0x00,
}
