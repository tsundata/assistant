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
	MessageId int64 `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
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

func (m *ChatbotRequest) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type ChatbotReply struct {
	State bool `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
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

func (m *ChatbotReply) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

type GroupBotRequest struct {
	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Bot     *Bot  `protobuf:"bytes,2,opt,name=bot,proto3" json:"bot,omitempty"`
}

func (m *GroupBotRequest) Reset()         { *m = GroupBotRequest{} }
func (m *GroupBotRequest) String() string { return proto.CompactTextString(m) }
func (*GroupBotRequest) ProtoMessage()    {}
func (*GroupBotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{2}
}
func (m *GroupBotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupBotRequest.Unmarshal(m, b)
}
func (m *GroupBotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupBotRequest.Marshal(b, m, deterministic)
}
func (m *GroupBotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupBotRequest.Merge(m, src)
}
func (m *GroupBotRequest) XXX_Size() int {
	return xxx_messageInfo_GroupBotRequest.Size(m)
}
func (m *GroupBotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupBotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GroupBotRequest proto.InternalMessageInfo

func (m *GroupBotRequest) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupBotRequest) GetBot() *Bot {
	if m != nil {
		return m.Bot
	}
	return nil
}

type BotSettingRequest struct {
	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	BotId   int64 `protobuf:"varint,2,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	Kvs     []*KV `protobuf:"bytes,3,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (m *BotSettingRequest) Reset()         { *m = BotSettingRequest{} }
func (m *BotSettingRequest) String() string { return proto.CompactTextString(m) }
func (*BotSettingRequest) ProtoMessage()    {}
func (*BotSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{3}
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

func (m *BotSettingRequest) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *BotSettingRequest) GetBotId() int64 {
	if m != nil {
		return m.BotId
	}
	return 0
}

func (m *BotSettingRequest) GetKvs() []*KV {
	if m != nil {
		return m.Kvs
	}
	return nil
}

type BotSettingReply struct {
	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	BotId   int64 `protobuf:"varint,2,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	Kvs     []*KV `protobuf:"bytes,3,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (m *BotSettingReply) Reset()         { *m = BotSettingReply{} }
func (m *BotSettingReply) String() string { return proto.CompactTextString(m) }
func (*BotSettingReply) ProtoMessage()    {}
func (*BotSettingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{4}
}
func (m *BotSettingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotSettingReply.Unmarshal(m, b)
}
func (m *BotSettingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotSettingReply.Marshal(b, m, deterministic)
}
func (m *BotSettingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotSettingReply.Merge(m, src)
}
func (m *BotSettingReply) XXX_Size() int {
	return xxx_messageInfo_BotSettingReply.Size(m)
}
func (m *BotSettingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BotSettingReply.DiscardUnknown(m)
}

var xxx_messageInfo_BotSettingReply proto.InternalMessageInfo

func (m *BotSettingReply) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *BotSettingReply) GetBotId() int64 {
	if m != nil {
		return m.BotId
	}
	return 0
}

func (m *BotSettingReply) GetKvs() []*KV {
	if m != nil {
		return m.Kvs
	}
	return nil
}

type GroupSettingRequest struct {
	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Kvs     []*KV `protobuf:"bytes,3,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (m *GroupSettingRequest) Reset()         { *m = GroupSettingRequest{} }
func (m *GroupSettingRequest) String() string { return proto.CompactTextString(m) }
func (*GroupSettingRequest) ProtoMessage()    {}
func (*GroupSettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{5}
}
func (m *GroupSettingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupSettingRequest.Unmarshal(m, b)
}
func (m *GroupSettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupSettingRequest.Marshal(b, m, deterministic)
}
func (m *GroupSettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupSettingRequest.Merge(m, src)
}
func (m *GroupSettingRequest) XXX_Size() int {
	return xxx_messageInfo_GroupSettingRequest.Size(m)
}
func (m *GroupSettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupSettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GroupSettingRequest proto.InternalMessageInfo

func (m *GroupSettingRequest) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupSettingRequest) GetKvs() []*KV {
	if m != nil {
		return m.Kvs
	}
	return nil
}

type GroupSettingReply struct {
	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Kvs     []*KV `protobuf:"bytes,3,rep,name=kvs,proto3" json:"kvs,omitempty"`
}

func (m *GroupSettingReply) Reset()         { *m = GroupSettingReply{} }
func (m *GroupSettingReply) String() string { return proto.CompactTextString(m) }
func (*GroupSettingReply) ProtoMessage()    {}
func (*GroupSettingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{6}
}
func (m *GroupSettingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupSettingReply.Unmarshal(m, b)
}
func (m *GroupSettingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupSettingReply.Marshal(b, m, deterministic)
}
func (m *GroupSettingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupSettingReply.Merge(m, src)
}
func (m *GroupSettingReply) XXX_Size() int {
	return xxx_messageInfo_GroupSettingReply.Size(m)
}
func (m *GroupSettingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupSettingReply.DiscardUnknown(m)
}

var xxx_messageInfo_GroupSettingReply proto.InternalMessageInfo

func (m *GroupSettingReply) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupSettingReply) GetKvs() []*KV {
	if m != nil {
		return m.Kvs
	}
	return nil
}

type Group struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: db:"sequence"
	Sequence int64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// @inject_tag: db:"type"
	Type int32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	// @inject_tag: db:"uuid"
	Uuid string `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// @inject_tag: db:"user_id"
	UserId int64 `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// @inject_tag: db:"name"
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: db:"avatar"
	Avatar string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{7}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Group) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Group) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Group) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Group) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Group) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Group) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type GroupBot struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: db:"group_id"
	GroupId int64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// @inject_tag: db:"bot_id"
	BotId int64 `protobuf:"varint,3,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *GroupBot) Reset()         { *m = GroupBot{} }
func (m *GroupBot) String() string { return proto.CompactTextString(m) }
func (*GroupBot) ProtoMessage()    {}
func (*GroupBot) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{8}
}
func (m *GroupBot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupBot.Unmarshal(m, b)
}
func (m *GroupBot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupBot.Marshal(b, m, deterministic)
}
func (m *GroupBot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupBot.Merge(m, src)
}
func (m *GroupBot) XXX_Size() int {
	return xxx_messageInfo_GroupBot.Size(m)
}
func (m *GroupBot) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupBot.DiscardUnknown(m)
}

var xxx_messageInfo_GroupBot proto.InternalMessageInfo

func (m *GroupBot) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GroupBot) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupBot) GetBotId() int64 {
	if m != nil {
		return m.BotId
	}
	return 0
}

func (m *GroupBot) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *GroupBot) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type GroupBotSetting struct {
	// @inject_tag: db:"group_id"
	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// @inject_tag: db:"bot_id"
	BotId int64 `protobuf:"varint,2,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	// @inject_tag: db:"key"
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// @inject_tag: db:"value"
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *GroupBotSetting) Reset()         { *m = GroupBotSetting{} }
func (m *GroupBotSetting) String() string { return proto.CompactTextString(m) }
func (*GroupBotSetting) ProtoMessage()    {}
func (*GroupBotSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{9}
}
func (m *GroupBotSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupBotSetting.Unmarshal(m, b)
}
func (m *GroupBotSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupBotSetting.Marshal(b, m, deterministic)
}
func (m *GroupBotSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupBotSetting.Merge(m, src)
}
func (m *GroupBotSetting) XXX_Size() int {
	return xxx_messageInfo_GroupBotSetting.Size(m)
}
func (m *GroupBotSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupBotSetting.DiscardUnknown(m)
}

var xxx_messageInfo_GroupBotSetting proto.InternalMessageInfo

func (m *GroupBotSetting) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupBotSetting) GetBotId() int64 {
	if m != nil {
		return m.BotId
	}
	return 0
}

func (m *GroupBotSetting) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GroupBotSetting) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *GroupBotSetting) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *GroupBotSetting) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type GroupSetting struct {
	// @inject_tag: db:"group_id"
	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// @inject_tag: db:"key"
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// @inject_tag: db:"value"
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *GroupSetting) Reset()         { *m = GroupSetting{} }
func (m *GroupSetting) String() string { return proto.CompactTextString(m) }
func (*GroupSetting) ProtoMessage()    {}
func (*GroupSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{10}
}
func (m *GroupSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupSetting.Unmarshal(m, b)
}
func (m *GroupSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupSetting.Marshal(b, m, deterministic)
}
func (m *GroupSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupSetting.Merge(m, src)
}
func (m *GroupSetting) XXX_Size() int {
	return xxx_messageInfo_GroupSetting.Size(m)
}
func (m *GroupSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupSetting.DiscardUnknown(m)
}

var xxx_messageInfo_GroupSetting proto.InternalMessageInfo

func (m *GroupSetting) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupSetting) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GroupSetting) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *GroupSetting) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *GroupSetting) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type GroupTag struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: db:"group_id"
	GroupId int64 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// @inject_tag: db:"tag"
	Tag string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *GroupTag) Reset()         { *m = GroupTag{} }
func (m *GroupTag) String() string { return proto.CompactTextString(m) }
func (*GroupTag) ProtoMessage()    {}
func (*GroupTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{11}
}
func (m *GroupTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupTag.Unmarshal(m, b)
}
func (m *GroupTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupTag.Marshal(b, m, deterministic)
}
func (m *GroupTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupTag.Merge(m, src)
}
func (m *GroupTag) XXX_Size() int {
	return xxx_messageInfo_GroupTag.Size(m)
}
func (m *GroupTag) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupTag.DiscardUnknown(m)
}

var xxx_messageInfo_GroupTag proto.InternalMessageInfo

func (m *GroupTag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GroupTag) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *GroupTag) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *GroupTag) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *GroupTag) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type Bot struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: db:"uuid"
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// @inject_tag: db:"name"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: db:"identifier"
	Identifier string `protobuf:"bytes,5,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// @inject_tag: db:"detail"
	Detail string `protobuf:"bytes,6,opt,name=detail,proto3" json:"detail,omitempty"`
	// @inject_tag: db:"avatar"
	Avatar string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
	// @inject_tag: db:"extend"
	Extend string `protobuf:"bytes,8,opt,name=extend,proto3" json:"extend,omitempty"`
	// @inject_tag: db:"status"
	Status int32 `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *Bot) Reset()         { *m = Bot{} }
func (m *Bot) String() string { return proto.CompactTextString(m) }
func (*Bot) ProtoMessage()    {}
func (*Bot) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{12}
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

func (m *Bot) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Bot) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Bot) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Bot) GetExtend() string {
	if m != nil {
		return m.Extend
	}
	return ""
}

func (m *Bot) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
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
	return fileDescriptor_acc44097314201ac, []int{13}
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
	return fileDescriptor_acc44097314201ac, []int{14}
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

type BotsRequest struct {
	GroupId int64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (m *BotsRequest) Reset()         { *m = BotsRequest{} }
func (m *BotsRequest) String() string { return proto.CompactTextString(m) }
func (*BotsRequest) ProtoMessage()    {}
func (*BotsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{15}
}
func (m *BotsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BotsRequest.Unmarshal(m, b)
}
func (m *BotsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BotsRequest.Marshal(b, m, deterministic)
}
func (m *BotsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BotsRequest.Merge(m, src)
}
func (m *BotsRequest) XXX_Size() int {
	return xxx_messageInfo_BotsRequest.Size(m)
}
func (m *BotsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BotsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BotsRequest proto.InternalMessageInfo

func (m *BotsRequest) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

type BotsReply struct {
	Bots []*Bot `protobuf:"bytes,1,rep,name=bots,proto3" json:"bots,omitempty"`
}

func (m *BotsReply) Reset()         { *m = BotsReply{} }
func (m *BotsReply) String() string { return proto.CompactTextString(m) }
func (*BotsReply) ProtoMessage()    {}
func (*BotsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{16}
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

type GroupRequest struct {
	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (m *GroupRequest) Reset()         { *m = GroupRequest{} }
func (m *GroupRequest) String() string { return proto.CompactTextString(m) }
func (*GroupRequest) ProtoMessage()    {}
func (*GroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{17}
}
func (m *GroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupRequest.Unmarshal(m, b)
}
func (m *GroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupRequest.Marshal(b, m, deterministic)
}
func (m *GroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupRequest.Merge(m, src)
}
func (m *GroupRequest) XXX_Size() int {
	return xxx_messageInfo_GroupRequest.Size(m)
}
func (m *GroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GroupRequest proto.InternalMessageInfo

func (m *GroupRequest) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type GroupReply struct {
	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (m *GroupReply) Reset()         { *m = GroupReply{} }
func (m *GroupReply) String() string { return proto.CompactTextString(m) }
func (*GroupReply) ProtoMessage()    {}
func (*GroupReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{18}
}
func (m *GroupReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupReply.Unmarshal(m, b)
}
func (m *GroupReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupReply.Marshal(b, m, deterministic)
}
func (m *GroupReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupReply.Merge(m, src)
}
func (m *GroupReply) XXX_Size() int {
	return xxx_messageInfo_GroupReply.Size(m)
}
func (m *GroupReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupReply.DiscardUnknown(m)
}

var xxx_messageInfo_GroupReply proto.InternalMessageInfo

func (m *GroupReply) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type GroupsReply struct {
	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (m *GroupsReply) Reset()         { *m = GroupsReply{} }
func (m *GroupsReply) String() string { return proto.CompactTextString(m) }
func (*GroupsReply) ProtoMessage()    {}
func (*GroupsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_acc44097314201ac, []int{19}
}
func (m *GroupsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupsReply.Unmarshal(m, b)
}
func (m *GroupsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupsReply.Marshal(b, m, deterministic)
}
func (m *GroupsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupsReply.Merge(m, src)
}
func (m *GroupsReply) XXX_Size() int {
	return xxx_messageInfo_GroupsReply.Size(m)
}
func (m *GroupsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GroupsReply proto.InternalMessageInfo

func (m *GroupsReply) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatbotRequest)(nil), "pb.ChatbotRequest")
	proto.RegisterType((*ChatbotReply)(nil), "pb.ChatbotReply")
	proto.RegisterType((*GroupBotRequest)(nil), "pb.GroupBotRequest")
	proto.RegisterType((*BotSettingRequest)(nil), "pb.BotSettingRequest")
	proto.RegisterType((*BotSettingReply)(nil), "pb.BotSettingReply")
	proto.RegisterType((*GroupSettingRequest)(nil), "pb.GroupSettingRequest")
	proto.RegisterType((*GroupSettingReply)(nil), "pb.GroupSettingReply")
	proto.RegisterType((*Group)(nil), "pb.Group")
	proto.RegisterType((*GroupBot)(nil), "pb.GroupBot")
	proto.RegisterType((*GroupBotSetting)(nil), "pb.GroupBotSetting")
	proto.RegisterType((*GroupSetting)(nil), "pb.GroupSetting")
	proto.RegisterType((*GroupTag)(nil), "pb.GroupTag")
	proto.RegisterType((*Bot)(nil), "pb.Bot")
	proto.RegisterType((*BotRequest)(nil), "pb.BotRequest")
	proto.RegisterType((*BotReply)(nil), "pb.BotReply")
	proto.RegisterType((*BotsRequest)(nil), "pb.BotsRequest")
	proto.RegisterType((*BotsReply)(nil), "pb.BotsReply")
	proto.RegisterType((*GroupRequest)(nil), "pb.GroupRequest")
	proto.RegisterType((*GroupReply)(nil), "pb.GroupReply")
	proto.RegisterType((*GroupsReply)(nil), "pb.GroupsReply")
}

func init() { proto.RegisterFile("chatbot.proto", fileDescriptor_acc44097314201ac) }

var fileDescriptor_acc44097314201ac = []byte{
	// 898 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0x8e, 0xed, 0xd8, 0xb1, 0x2b, 0xb3, 0xc9, 0x6c, 0xcf, 0xce, 0xae, 0x37, 0x40, 0x08, 0x16,
	0x08, 0x4b, 0xc0, 0xcc, 0xec, 0x70, 0xd8, 0x23, 0x6c, 0x06, 0x69, 0x36, 0x70, 0xf3, 0x00, 0x07,
	0x24, 0x34, 0x6a, 0xc7, 0x8d, 0xb1, 0x36, 0x1b, 0x9b, 0xb8, 0x1d, 0x91, 0x07, 0xe0, 0xc2, 0x05,
	0x5e, 0x83, 0x77, 0xe0, 0x01, 0x38, 0xee, 0x91, 0x13, 0xd2, 0x6a, 0x1e, 0x83, 0x0b, 0xea, 0x3f,
	0xc7, 0xf1, 0xe4, 0x6f, 0x80, 0x5b, 0xd7, 0xd7, 0x55, 0x5f, 0x7f, 0xe5, 0xaa, 0xea, 0x36, 0xdc,
	0x1b, 0x7f, 0x8f, 0x69, 0x98, 0xd2, 0x93, 0x6c, 0x96, 0xd2, 0x14, 0xe9, 0x59, 0xd8, 0x83, 0x10,
	0xe7, 0x44, 0xd8, 0xbd, 0x07, 0x71, 0x1a, 0xa7, 0x7c, 0x79, 0xca, 0x56, 0x02, 0xf5, 0x4e, 0xa1,
	0x73, 0x21, 0xc2, 0x02, 0xf2, 0x43, 0x41, 0x72, 0x8a, 0xde, 0x02, 0x78, 0x49, 0xf2, 0x1c, 0xc7,
	0xe4, 0x3a, 0x89, 0x5c, 0x6d, 0xa0, 0xf9, 0x46, 0xe0, 0x48, 0x64, 0x14, 0x79, 0xef, 0xc2, 0x41,
	0x19, 0x90, 0x4d, 0x16, 0xe8, 0x01, 0x98, 0x39, 0xc5, 0x94, 0x70, 0x4f, 0x3b, 0x10, 0x86, 0x77,
	0x09, 0xdd, 0xcb, 0x59, 0x5a, 0x64, 0xc3, 0x25, 0xef, 0x63, 0xb0, 0x63, 0x06, 0x2d, 0x59, 0x5b,
	0xdc, 0x1e, 0x45, 0xe8, 0x31, 0x18, 0x61, 0x4a, 0x5d, 0x7d, 0xa0, 0xf9, 0xed, 0xf3, 0xd6, 0x49,
	0x16, 0x9e, 0xb0, 0x38, 0x86, 0x79, 0xd7, 0x70, 0x7f, 0x98, 0xd2, 0x2b, 0x42, 0x69, 0x32, 0x8d,
	0xf7, 0xa0, 0x3a, 0x06, 0x2b, 0x4c, 0x29, 0xdb, 0xd0, 0xf9, 0x86, 0x19, 0xa6, 0x74, 0x14, 0x21,
	0x17, 0x8c, 0x17, 0xf3, 0xdc, 0x35, 0x06, 0x86, 0xdf, 0x3e, 0xb7, 0xd8, 0x09, 0x5f, 0x7c, 0x1d,
	0x30, 0xc8, 0xfb, 0x16, 0xba, 0xd5, 0x03, 0x58, 0x4a, 0xff, 0x27, 0xfd, 0xe7, 0x70, 0xc4, 0x3f,
	0xc4, 0xfe, 0x19, 0x6c, 0xe6, 0x7a, 0x0e, 0xf7, 0x57, 0xb9, 0x76, 0x88, 0xdd, 0xcc, 0xf4, 0x97,
	0x06, 0x26, 0xa7, 0x42, 0x1d, 0xd0, 0xcb, 0x40, 0x3d, 0x89, 0x50, 0x0f, 0xec, 0x9c, 0x69, 0x9c,
	0x8e, 0x89, 0x4c, 0xb1, 0xb4, 0x11, 0x82, 0x26, 0x5d, 0x64, 0xc4, 0x35, 0x06, 0x9a, 0x6f, 0x06,
	0x7c, 0xcd, 0xb0, 0xa2, 0x48, 0x22, 0xb7, 0x39, 0xd0, 0x7c, 0x27, 0xe0, 0x6b, 0xf4, 0x08, 0x5a,
	0x45, 0x4e, 0x66, 0x4c, 0x91, 0xc9, 0x29, 0x2c, 0x66, 0x8e, 0x22, 0xe6, 0x3c, 0xc5, 0x2f, 0x89,
	0x6b, 0x09, 0x67, 0xb6, 0x46, 0x0f, 0xc1, 0xc2, 0x73, 0x4c, 0xf1, 0xcc, 0x6d, 0x71, 0x54, 0x5a,
	0xac, 0x0d, 0xc7, 0x33, 0x82, 0x29, 0x89, 0xae, 0x31, 0x75, 0x6d, 0xd1, 0x86, 0x12, 0x79, 0xc6,
	0xbb, 0xb4, 0xc8, 0x22, 0xb5, 0xed, 0x88, 0x6d, 0x89, 0x3c, 0xa3, 0xde, 0xcf, 0x1a, 0xd8, 0xaa,
	0x01, 0x6f, 0xe5, 0x58, 0xfd, 0x64, 0xfa, 0xa6, 0xfa, 0x1a, 0xd5, 0xfa, 0xae, 0x8a, 0x69, 0x6e,
	0x17, 0x63, 0xd6, 0xc5, 0xfc, 0xa6, 0x2d, 0xa7, 0x41, 0xd6, 0xee, 0x5f, 0xf4, 0xd8, 0x21, 0x18,
	0x2f, 0xc8, 0x82, 0xeb, 0x72, 0x02, 0xb6, 0x64, 0xa3, 0x37, 0xc7, 0x93, 0x82, 0xc8, 0x8f, 0x2f,
	0x8c, 0x9a, 0x56, 0x73, 0xbb, 0x56, 0xab, 0xae, 0xf5, 0x17, 0x0d, 0x0e, 0xaa, 0x4d, 0xb6, 0x4d,
	0xa8, 0x54, 0xa4, 0xaf, 0x51, 0x64, 0x6c, 0x56, 0x74, 0xd7, 0xaf, 0xf7, 0x93, 0x2a, 0xe5, 0x97,
	0x38, 0xbe, 0x4b, 0x29, 0x0f, 0xc1, 0xa0, 0x38, 0x56, 0xdf, 0x8b, 0xe2, 0xf8, 0x3f, 0xea, 0xf8,
	0x5b, 0x03, 0x63, 0x5d, 0x37, 0xa9, 0x09, 0xd0, 0x2b, 0x13, 0xa0, 0x1a, 0xbd, 0x59, 0x69, 0xf4,
	0x3e, 0x40, 0x12, 0x91, 0x29, 0x4d, 0xbe, 0x4b, 0xc8, 0x8c, 0xd3, 0x3b, 0x41, 0x05, 0x61, 0x83,
	0x10, 0x11, 0x8a, 0x93, 0x89, 0x1c, 0x0f, 0x69, 0x6d, 0x1c, 0x90, 0x87, 0x60, 0x91, 0x1f, 0x29,
	0x99, 0x46, 0x7c, 0x38, 0x9c, 0x40, 0x5a, 0x0c, 0x67, 0x77, 0x70, 0x91, 0xf3, 0xa9, 0x30, 0x03,
	0x69, 0xd5, 0xb2, 0x87, 0xed, 0xd9, 0xb7, 0xeb, 0xd9, 0xbf, 0x0f, 0xb0, 0x72, 0x97, 0xf3, 0x0b,
	0x5b, 0x5b, 0x73, 0x61, 0xbf, 0x07, 0xf6, 0x50, 0xbd, 0x0d, 0x5b, 0xdc, 0x7c, 0x68, 0x0f, 0x53,
	0x9a, 0xef, 0xbe, 0x0f, 0x3d, 0x1f, 0x1c, 0xe1, 0xc9, 0x18, 0xdf, 0x80, 0x66, 0x98, 0xd2, 0xdc,
	0xd5, 0xf8, 0x9d, 0x56, 0x52, 0x72, 0xd0, 0x3b, 0x95, 0xad, 0xab, 0x48, 0xdf, 0x06, 0x93, 0x93,
	0x48, 0x01, 0x0e, 0xf3, 0x16, 0x0e, 0x02, 0xf7, 0x3e, 0x02, 0x90, 0x01, 0x8c, 0x7b, 0xa7, 0xfb,
	0x19, 0xb4, 0xb9, 0x2d, 0xb5, 0xbc, 0x03, 0x16, 0xc7, 0x95, 0x9a, 0x4a, 0x80, 0xdc, 0x38, 0xff,
	0xdd, 0x02, 0x90, 0xaf, 0xe5, 0xd5, 0x7c, 0x8c, 0xce, 0xc0, 0x7a, 0x8e, 0xa7, 0xd1, 0x84, 0x20,
	0xc4, 0x7c, 0x57, 0x1f, 0xde, 0xde, 0xe1, 0x0a, 0x96, 0x4d, 0x16, 0x5e, 0x03, 0x7d, 0x08, 0x76,
	0x40, 0xe2, 0x24, 0xa7, 0x64, 0x86, 0x3a, 0x2a, 0x5b, 0xe9, 0xcf, 0xed, 0x2b, 0xf6, 0xdc, 0x2a,
	0x6f, 0x1f, 0xac, 0x4b, 0x42, 0x79, 0x93, 0xd6, 0x7c, 0x0f, 0x4a, 0x5b, 0x78, 0x7e, 0x00, 0x2d,
	0xe1, 0x99, 0xa3, 0xae, 0xdc, 0x52, 0xb5, 0xe8, 0xdd, 0x5b, 0x02, 0xc2, 0xf9, 0x29, 0x74, 0x2e,
	0x78, 0x9f, 0x94, 0x37, 0xea, 0x51, 0x99, 0xea, 0x56, 0x3d, 0x4f, 0xa1, 0xf3, 0x19, 0x99, 0x90,
	0xbb, 0x07, 0x7e, 0x0a, 0xc7, 0x5f, 0xf1, 0xd6, 0xab, 0x5f, 0x9b, 0xc7, 0x52, 0xdb, 0xea, 0x73,
	0xba, 0x86, 0xe1, 0x13, 0x40, 0x15, 0x06, 0x15, 0xfe, 0xa8, 0x3c, 0x7e, 0x27, 0xc1, 0x10, 0xd0,
	0x25, 0xa1, 0x7b, 0x9e, 0x7f, 0x54, 0x87, 0x05, 0xc7, 0x05, 0x74, 0x15, 0xc7, 0x4e, 0x05, 0xc7,
	0xb7, 0x37, 0x04, 0xc9, 0x19, 0x38, 0x8a, 0x24, 0x47, 0x87, 0xcb, 0x1e, 0x93, 0x71, 0xdd, 0x12,
	0x29, 0xeb, 0xf5, 0x04, 0xda, 0x95, 0x7a, 0xad, 0x89, 0xb9, 0x9d, 0xed, 0x09, 0xd8, 0xea, 0x90,
	0x4d, 0xfe, 0xcb, 0x49, 0x11, 0x47, 0x54, 0x2a, 0xbb, 0xd7, 0x11, 0x4f, 0xa0, 0x5d, 0xa9, 0xc8,
	0x3e, 0x21, 0xc3, 0x37, 0xff, 0x78, 0xdd, 0xd7, 0x5e, 0xbd, 0xee, 0x37, 0x7e, 0xbd, 0xe9, 0x37,
	0x5e, 0xdd, 0xf4, 0x1b, 0x7f, 0xde, 0xf4, 0x1b, 0xdf, 0x58, 0x38, 0x4b, 0x4e, 0xb3, 0x30, 0xb4,
	0xf8, 0x1f, 0xec, 0xc7, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x68, 0x2d, 0xa7, 0xf8, 0x0a,
	0x00, 0x00,
}
