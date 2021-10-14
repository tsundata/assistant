// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

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

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthRequest struct {
	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}
func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AuthRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthReply struct {
	State bool   `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	Id    int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *AuthReply) Reset()         { *m = AuthReply{} }
func (m *AuthReply) String() string { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()    {}
func (*AuthReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}
func (m *AuthReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReply.Unmarshal(m, b)
}
func (m *AuthReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReply.Marshal(b, m, deterministic)
}
func (m *AuthReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReply.Merge(m, src)
}
func (m *AuthReply) XXX_Size() int {
	return xxx_messageInfo_AuthReply.Size(m)
}
func (m *AuthReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReply proto.InternalMessageInfo

func (m *AuthReply) GetState() bool {
	if m != nil {
		return m.State
	}
	return false
}

func (m *AuthReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AuthReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RoleRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *RoleRequest) Reset()         { *m = RoleRequest{} }
func (m *RoleRequest) String() string { return proto.CompactTextString(m) }
func (*RoleRequest) ProtoMessage()    {}
func (*RoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}
func (m *RoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleRequest.Unmarshal(m, b)
}
func (m *RoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleRequest.Marshal(b, m, deterministic)
}
func (m *RoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleRequest.Merge(m, src)
}
func (m *RoleRequest) XXX_Size() int {
	return xxx_messageInfo_RoleRequest.Size(m)
}
func (m *RoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleRequest proto.InternalMessageInfo

func (m *RoleRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RoleReply struct {
	Role *Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *RoleReply) Reset()         { *m = RoleReply{} }
func (m *RoleReply) String() string { return proto.CompactTextString(m) }
func (*RoleReply) ProtoMessage()    {}
func (*RoleReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}
func (m *RoleReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleReply.Unmarshal(m, b)
}
func (m *RoleReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleReply.Marshal(b, m, deterministic)
}
func (m *RoleReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleReply.Merge(m, src)
}
func (m *RoleReply) XXX_Size() int {
	return xxx_messageInfo_RoleReply.Size(m)
}
func (m *RoleReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleReply.DiscardUnknown(m)
}

var xxx_messageInfo_RoleReply proto.InternalMessageInfo

func (m *RoleReply) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type RolesReply struct {
	Roles []*Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (m *RolesReply) Reset()         { *m = RolesReply{} }
func (m *RolesReply) String() string { return proto.CompactTextString(m) }
func (*RolesReply) ProtoMessage()    {}
func (*RolesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}
func (m *RolesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesReply.Unmarshal(m, b)
}
func (m *RolesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesReply.Marshal(b, m, deterministic)
}
func (m *RolesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesReply.Merge(m, src)
}
func (m *RolesReply) XXX_Size() int {
	return xxx_messageInfo_RolesReply.Size(m)
}
func (m *RolesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesReply.DiscardUnknown(m)
}

var xxx_messageInfo_RolesReply proto.InternalMessageInfo

func (m *RolesReply) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type UserRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}
func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UserReply struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *UserReply) Reset()         { *m = UserReply{} }
func (m *UserReply) String() string { return proto.CompactTextString(m) }
func (*UserReply) ProtoMessage()    {}
func (*UserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}
func (m *UserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReply.Unmarshal(m, b)
}
func (m *UserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReply.Marshal(b, m, deterministic)
}
func (m *UserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReply.Merge(m, src)
}
func (m *UserReply) XXX_Size() int {
	return xxx_messageInfo_UserReply.Size(m)
}
func (m *UserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReply.DiscardUnknown(m)
}

var xxx_messageInfo_UserReply proto.InternalMessageInfo

func (m *UserReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UsersReply struct {
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (m *UsersReply) Reset()         { *m = UsersReply{} }
func (m *UsersReply) String() string { return proto.CompactTextString(m) }
func (*UsersReply) ProtoMessage()    {}
func (*UsersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}
func (m *UsersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersReply.Unmarshal(m, b)
}
func (m *UsersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersReply.Marshal(b, m, deterministic)
}
func (m *UsersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersReply.Merge(m, src)
}
func (m *UsersReply) XXX_Size() int {
	return xxx_messageInfo_UsersReply.Size(m)
}
func (m *UsersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersReply.DiscardUnknown(m)
}

var xxx_messageInfo_UsersReply proto.InternalMessageInfo

func (m *UsersReply) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type BytesReply struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BytesReply) Reset()         { *m = BytesReply{} }
func (m *BytesReply) String() string { return proto.CompactTextString(m) }
func (*BytesReply) ProtoMessage()    {}
func (*BytesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}
func (m *BytesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesReply.Unmarshal(m, b)
}
func (m *BytesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesReply.Marshal(b, m, deterministic)
}
func (m *BytesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesReply.Merge(m, src)
}
func (m *BytesReply) XXX_Size() int {
	return xxx_messageInfo_BytesReply.Size(m)
}
func (m *BytesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesReply.DiscardUnknown(m)
}

var xxx_messageInfo_BytesReply proto.InternalMessageInfo

func (m *BytesReply) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Device struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" gorm:"primaryKey"`
	// @inject_tag: db:"user_id"
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" db:"user_id"`
	// @inject_tag: db:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" db:"name"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Device) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Device) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type User struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" gorm:"primaryKey"`
	// @inject_tag: db:"username"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" db:"username"`
	// @inject_tag: db:"password"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" db:"password"`
	// @inject_tag: db:"name"
	Nickname string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty" db:"name"`
	// @inject_tag: db:"mobile"
	Mobile string `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty" db:"mobile"`
	// @inject_tag: db:"remark"
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty" db:"remark"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at"`
	// @inject_tag: db:"role"
	Role *Role `protobuf:"bytes,9,opt,name=role,proto3" json:"role,omitempty" db:"role"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *User) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *User) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *User) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type Role struct {
	// @inject_tag: db:"id" gorm:"primaryKey"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" gorm:"primaryKey"`
	// @inject_tag: db:"user_id"
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" db:"user_id"`
	// @inject_tag: db:"profession"
	Profession string `protobuf:"bytes,3,opt,name=profession,proto3" json:"profession,omitempty" db:"profession"`
	// @inject_tag: db:"exp"
	Exp int64 `protobuf:"varint,4,opt,name=exp,proto3" json:"exp,omitempty" db:"exp"`
	// @inject_tag: db:"level"
	Level int64 `protobuf:"varint,5,opt,name=level,proto3" json:"level,omitempty" db:"level"`
	// @inject_tag: db:"strength"
	Strength int64 `protobuf:"varint,6,opt,name=strength,proto3" json:"strength,omitempty" db:"strength"`
	// @inject_tag: db:"culture"
	Culture int64 `protobuf:"varint,7,opt,name=culture,proto3" json:"culture,omitempty" db:"culture"`
	// @inject_tag: db:"environment"
	Environment int64 `protobuf:"varint,8,opt,name=environment,proto3" json:"environment,omitempty" db:"environment"`
	// @inject_tag: db:"charisma"
	Charisma int64 `protobuf:"varint,9,opt,name=charisma,proto3" json:"charisma,omitempty" db:"charisma"`
	// @inject_tag: db:"talent"
	Talent int64 `protobuf:"varint,10,opt,name=talent,proto3" json:"talent,omitempty" db:"talent"`
	// @inject_tag: db:"intellect"
	Intellect int64 `protobuf:"varint,11,opt,name=intellect,proto3" json:"intellect,omitempty" db:"intellect"`
	// @inject_tag: db:"created_at"
	CreatedAt int64 `protobuf:"varint,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at"`
	// @inject_tag: db:"updated_at"
	UpdatedAt int64 `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{12}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Role) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Role) GetProfession() string {
	if m != nil {
		return m.Profession
	}
	return ""
}

func (m *Role) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *Role) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Role) GetStrength() int64 {
	if m != nil {
		return m.Strength
	}
	return 0
}

func (m *Role) GetCulture() int64 {
	if m != nil {
		return m.Culture
	}
	return 0
}

func (m *Role) GetEnvironment() int64 {
	if m != nil {
		return m.Environment
	}
	return 0
}

func (m *Role) GetCharisma() int64 {
	if m != nil {
		return m.Charisma
	}
	return 0
}

func (m *Role) GetTalent() int64 {
	if m != nil {
		return m.Talent
	}
	return 0
}

func (m *Role) GetIntellect() int64 {
	if m != nil {
		return m.Intellect
	}
	return 0
}

func (m *Role) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Role) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type Equipment struct {
	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quality   string `protobuf:"bytes,3,opt,name=quality,proto3" json:"quality,omitempty"`
	Level     int64  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	Category  string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *Equipment) Reset()         { *m = Equipment{} }
func (m *Equipment) String() string { return proto.CompactTextString(m) }
func (*Equipment) ProtoMessage()    {}
func (*Equipment) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{13}
}
func (m *Equipment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Equipment.Unmarshal(m, b)
}
func (m *Equipment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Equipment.Marshal(b, m, deterministic)
}
func (m *Equipment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equipment.Merge(m, src)
}
func (m *Equipment) XXX_Size() int {
	return xxx_messageInfo_Equipment.Size(m)
}
func (m *Equipment) XXX_DiscardUnknown() {
	xxx_messageInfo_Equipment.DiscardUnknown(m)
}

var xxx_messageInfo_Equipment proto.InternalMessageInfo

func (m *Equipment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Equipment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Equipment) GetQuality() string {
	if m != nil {
		return m.Quality
	}
	return ""
}

func (m *Equipment) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Equipment) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Equipment) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type Quest struct {
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Exp           int64  `protobuf:"varint,3,opt,name=exp,proto3" json:"exp,omitempty"`
	AttrPoints    string `protobuf:"bytes,4,opt,name=attr_points,json=attrPoints,proto3" json:"attr_points,omitempty"`
	Preconditions string `protobuf:"bytes,5,opt,name=preconditions,proto3" json:"preconditions,omitempty"`
	CreatedAt     string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *Quest) Reset()         { *m = Quest{} }
func (m *Quest) String() string { return proto.CompactTextString(m) }
func (*Quest) ProtoMessage()    {}
func (*Quest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{14}
}
func (m *Quest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quest.Unmarshal(m, b)
}
func (m *Quest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quest.Marshal(b, m, deterministic)
}
func (m *Quest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quest.Merge(m, src)
}
func (m *Quest) XXX_Size() int {
	return xxx_messageInfo_Quest.Size(m)
}
func (m *Quest) XXX_DiscardUnknown() {
	xxx_messageInfo_Quest.DiscardUnknown(m)
}

var xxx_messageInfo_Quest proto.InternalMessageInfo

func (m *Quest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Quest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Quest) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *Quest) GetAttrPoints() string {
	if m != nil {
		return m.AttrPoints
	}
	return ""
}

func (m *Quest) GetPreconditions() string {
	if m != nil {
		return m.Preconditions
	}
	return ""
}

func (m *Quest) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type AttrChange struct {
	UserId  int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *AttrChange) Reset()         { *m = AttrChange{} }
func (m *AttrChange) String() string { return proto.CompactTextString(m) }
func (*AttrChange) ProtoMessage()    {}
func (*AttrChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{15}
}
func (m *AttrChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttrChange.Unmarshal(m, b)
}
func (m *AttrChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttrChange.Marshal(b, m, deterministic)
}
func (m *AttrChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttrChange.Merge(m, src)
}
func (m *AttrChange) XXX_Size() int {
	return xxx_messageInfo_AttrChange.Size(m)
}
func (m *AttrChange) XXX_DiscardUnknown() {
	xxx_messageInfo_AttrChange.DiscardUnknown(m)
}

var xxx_messageInfo_AttrChange proto.InternalMessageInfo

func (m *AttrChange) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AttrChange) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type DeviceRequest struct {
	Device *Device `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (m *DeviceRequest) Reset()         { *m = DeviceRequest{} }
func (m *DeviceRequest) String() string { return proto.CompactTextString(m) }
func (*DeviceRequest) ProtoMessage()    {}
func (*DeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{16}
}
func (m *DeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceRequest.Unmarshal(m, b)
}
func (m *DeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceRequest.Marshal(b, m, deterministic)
}
func (m *DeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceRequest.Merge(m, src)
}
func (m *DeviceRequest) XXX_Size() int {
	return xxx_messageInfo_DeviceRequest.Size(m)
}
func (m *DeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceRequest proto.InternalMessageInfo

func (m *DeviceRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*AuthRequest)(nil), "pb.AuthRequest")
	proto.RegisterType((*AuthReply)(nil), "pb.AuthReply")
	proto.RegisterType((*RoleRequest)(nil), "pb.RoleRequest")
	proto.RegisterType((*RoleReply)(nil), "pb.RoleReply")
	proto.RegisterType((*RolesReply)(nil), "pb.RolesReply")
	proto.RegisterType((*UserRequest)(nil), "pb.UserRequest")
	proto.RegisterType((*UserReply)(nil), "pb.UserReply")
	proto.RegisterType((*UsersReply)(nil), "pb.UsersReply")
	proto.RegisterType((*BytesReply)(nil), "pb.BytesReply")
	proto.RegisterType((*Device)(nil), "pb.Device")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*Role)(nil), "pb.Role")
	proto.RegisterType((*Equipment)(nil), "pb.Equipment")
	proto.RegisterType((*Quest)(nil), "pb.Quest")
	proto.RegisterType((*AttrChange)(nil), "pb.AttrChange")
	proto.RegisterType((*DeviceRequest)(nil), "pb.DeviceRequest")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 902 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcb, 0x6e, 0x24, 0x35,
	0x14, 0xed, 0x7a, 0xf4, 0xa3, 0x6e, 0x77, 0x27, 0x83, 0x35, 0x82, 0x56, 0x94, 0x69, 0xa2, 0x12,
	0x0b, 0x60, 0x86, 0x84, 0x99, 0x7c, 0x00, 0x4a, 0x06, 0x88, 0x46, 0x42, 0x08, 0x6a, 0x98, 0x0d,
	0x9b, 0xc8, 0x5d, 0x65, 0x3a, 0x56, 0xaa, 0xcb, 0x15, 0xdb, 0x15, 0x68, 0xb6, 0xfc, 0x00, 0x6b,
	0x24, 0x96, 0xfc, 0x0b, 0xcb, 0x59, 0xb2, 0x43, 0xa3, 0xfc, 0x03, 0x6b, 0x74, 0x6d, 0x57, 0x75,
	0x75, 0xe7, 0xb9, 0xab, 0x73, 0xee, 0xc3, 0xc7, 0xd7, 0xd7, 0xd7, 0x05, 0x50, 0x29, 0x26, 0xf7,
	0x4b, 0x29, 0xb4, 0x20, 0x7e, 0x39, 0xdb, 0x81, 0x19, 0x55, 0xcc, 0xe2, 0x9d, 0xc7, 0x73, 0x31,
	0x17, 0xe6, 0xf3, 0x00, 0xbf, 0x2c, 0x1b, 0x7f, 0x0d, 0xa3, 0x6f, 0xc4, 0x9c, 0x17, 0x09, 0xbb,
	0xa8, 0x98, 0xd2, 0x64, 0x07, 0x06, 0x98, 0xa3, 0xa0, 0x0b, 0x36, 0xf1, 0xf6, 0xbc, 0x8f, 0xa3,
	0xa4, 0xc1, 0x68, 0x2b, 0xa9, 0x52, 0x3f, 0x0b, 0x99, 0x4d, 0x7c, 0x6b, 0xab, 0x71, 0x7c, 0x08,
	0xc3, 0xa3, 0x4a, 0x9f, 0xd5, 0x69, 0xb6, 0xc0, 0xe7, 0x99, 0x49, 0x10, 0x24, 0x3e, 0xcf, 0xc8,
	0x63, 0xe8, 0x6a, 0x71, 0xce, 0x0a, 0x17, 0x67, 0x41, 0x7c, 0x02, 0x91, 0x0d, 0x2a, 0xf3, 0x25,
	0xba, 0x28, 0x4d, 0xb5, 0x5d, 0x76, 0x90, 0x58, 0xe0, 0x12, 0xf9, 0xd7, 0x13, 0x05, 0xed, 0x44,
	0x4f, 0x60, 0x98, 0x88, 0x9c, 0xdd, 0xb2, 0x7a, 0xfc, 0x09, 0x44, 0xd6, 0x8c, 0xeb, 0xec, 0x42,
	0x28, 0x45, 0x6e, 0x97, 0x19, 0xbe, 0x18, 0xec, 0x97, 0xb3, 0x7d, 0x63, 0x34, 0x6c, 0xfc, 0x0c,
	0x00, 0x91, 0xb2, 0xbe, 0x53, 0xe8, 0x22, 0xab, 0x26, 0xde, 0x5e, 0xb0, 0xe6, 0x6c, 0xe9, 0xf8,
	0x29, 0x0c, 0xdf, 0x28, 0x26, 0xeb, 0x75, 0x77, 0x21, 0xc4, 0x62, 0xb5, 0x53, 0x1b, 0xb3, 0x61,
	0x51, 0x85, 0x75, 0x76, 0x2a, 0xee, 0x70, 0x7d, 0x06, 0x80, 0x68, 0xa5, 0x02, 0xd9, 0x35, 0x15,
	0xc6, 0xd9, 0xd2, 0xf1, 0x1e, 0xc0, 0xf1, 0x52, 0xd7, 0x9a, 0x09, 0x84, 0x19, 0xd5, 0xd4, 0x64,
	0x1e, 0x25, 0xe6, 0x3b, 0xfe, 0xcd, 0x83, 0xde, 0x97, 0xec, 0x92, 0xa7, 0xec, 0xda, 0xc9, 0x7c,
	0x00, 0x7d, 0xcc, 0x72, 0xda, 0x54, 0xb9, 0x87, 0xf0, 0x55, 0x86, 0x79, 0x4c, 0x17, 0xd8, 0x42,
	0x9b, 0x6f, 0xf2, 0x04, 0x20, 0x95, 0x8c, 0x6a, 0x96, 0x9d, 0x52, 0x3d, 0xe9, 0x1a, 0xff, 0xc8,
	0x31, 0x47, 0x1a, 0xcd, 0x55, 0x99, 0xd5, 0xe6, 0x9e, 0x35, 0x3b, 0xe6, 0x48, 0xc7, 0xff, 0x79,
	0x10, 0xa2, 0xee, 0x6b, 0x1a, 0xda, 0x4d, 0xe7, 0xdf, 0xd1, 0x74, 0xc1, 0x7a, 0xd3, 0xa1, 0xad,
	0xe0, 0xe9, 0xb9, 0x89, 0x0b, 0xad, 0xad, 0xc6, 0xe4, 0x7d, 0xe8, 0x2d, 0xc4, 0x8c, 0xe7, 0xcc,
	0xc8, 0x8c, 0x12, 0x87, 0x90, 0x97, 0x6c, 0x41, 0xe5, 0xb9, 0xd1, 0x17, 0x25, 0x0e, 0x6d, 0x6c,
	0xad, 0x7f, 0xf7, 0xd6, 0x06, 0x1b, 0x5b, 0x6b, 0x9a, 0x2a, 0xba, 0xb1, 0xa9, 0xfe, 0xf5, 0x21,
	0x44, 0xf8, 0xf0, 0xe2, 0x4f, 0x01, 0x4a, 0x29, 0x7e, 0x62, 0x4a, 0x71, 0x51, 0xf7, 0x7a, 0x8b,
	0x21, 0x8f, 0x20, 0x60, 0xbf, 0x94, 0x66, 0xd3, 0x41, 0x82, 0x9f, 0x78, 0x31, 0x72, 0x76, 0xc9,
	0x72, 0x77, 0x2a, 0x16, 0x60, 0x85, 0x94, 0x96, 0xac, 0x98, 0xeb, 0x33, 0x77, 0x1e, 0x0d, 0x26,
	0x13, 0xe8, 0xa7, 0x55, 0xae, 0x2b, 0xc9, 0xdc, 0x76, 0x6b, 0x48, 0xf6, 0x60, 0xc8, 0x8a, 0x4b,
	0x2e, 0x45, 0xb1, 0x60, 0x45, 0xbd, 0xdb, 0x36, 0x85, 0x79, 0xd3, 0x33, 0x2a, 0xb9, 0x5a, 0x50,
	0xb3, 0xe7, 0x20, 0x69, 0x30, 0x56, 0x58, 0xd3, 0x1c, 0x03, 0xc1, 0xee, 0xc9, 0x22, 0xb2, 0x0b,
	0x11, 0x2f, 0x34, 0xcb, 0x73, 0x96, 0xea, 0xc9, 0xd0, 0x56, 0xb0, 0x21, 0x36, 0xea, 0x3f, 0xba,
	0xbb, 0xfe, 0xe3, 0xcd, 0xd6, 0xfa, 0xc3, 0x83, 0xe8, 0xab, 0x8b, 0x8a, 0x97, 0x46, 0xdd, 0x66,
	0x99, 0xeb, 0x56, 0xf6, 0x5b, 0xad, 0x3c, 0x81, 0xfe, 0x45, 0x45, 0x73, 0xae, 0x97, 0xae, 0xbc,
	0x35, 0x5c, 0x55, 0x32, 0xdc, 0xa8, 0x64, 0x4a, 0x35, 0x9b, 0x0b, 0xb9, 0x74, 0x1d, 0xd5, 0xe0,
	0x0d, 0xed, 0xb6, 0xaf, 0x56, 0xda, 0xe3, 0xbf, 0x3c, 0xe8, 0x7e, 0x7f, 0xeb, 0x58, 0xe4, 0x3a,
	0x67, 0xcd, 0x58, 0x44, 0x50, 0x1f, 0x6e, 0xb0, 0x3a, 0xdc, 0x0f, 0x61, 0x48, 0xb5, 0x96, 0xa7,
	0xa5, 0xe0, 0x85, 0x56, 0xae, 0xd7, 0x01, 0xa9, 0xef, 0x0c, 0x43, 0x3e, 0x82, 0x71, 0x29, 0x59,
	0x2a, 0x8a, 0x8c, 0x6b, 0x2e, 0x0a, 0xe5, 0x24, 0xae, 0x93, 0xf7, 0xe9, 0xfc, 0x02, 0xe0, 0x48,
	0x6b, 0xf9, 0xf2, 0x8c, 0x16, 0x73, 0xd6, 0xee, 0x4d, 0x6f, 0xad, 0x37, 0xb1, 0x6f, 0x44, 0xa1,
	0xf1, 0x80, 0xad, 0xec, 0x1a, 0xc6, 0x87, 0x30, 0xb6, 0x53, 0xa6, 0x1e, 0x88, 0x31, 0xf4, 0x32,
	0x43, 0xb8, 0x39, 0x07, 0x78, 0x31, 0x9c, 0x8b, 0xb3, 0xbc, 0xf8, 0x33, 0x84, 0x3e, 0x4e, 0x85,
	0xd7, 0x97, 0x29, 0xf9, 0x14, 0xba, 0xe6, 0x35, 0x22, 0x8f, 0xd0, 0xb1, 0xfd, 0x30, 0xed, 0x8c,
	0x91, 0x69, 0x5e, 0x8b, 0xb8, 0x43, 0x3e, 0x87, 0xd1, 0x09, 0xd3, 0xc8, 0xfc, 0x80, 0x6f, 0x00,
	0xd9, 0x5e, 0x39, 0xdc, 0x12, 0xf1, 0x1c, 0xc6, 0x08, 0x85, 0xe4, 0xbf, 0x52, 0x2c, 0xc8, 0x03,
	0x42, 0x9e, 0x42, 0xff, 0x84, 0x69, 0x73, 0x77, 0xb7, 0x9b, 0x4b, 0xdd, 0x76, 0x6e, 0xde, 0x15,
	0x93, 0x7f, 0xe4, 0x9c, 0x5f, 0x2d, 0xe8, 0xfc, 0x86, 0x88, 0x2d, 0x24, 0x56, 0xa3, 0x3a, 0xee,
	0x90, 0x7d, 0x80, 0x97, 0xa6, 0xfe, 0x66, 0x2e, 0x6e, 0x37, 0x93, 0xbd, 0xbd, 0x44, 0xf3, 0x68,
	0x34, 0x7a, 0x1e, 0xe8, 0xfc, 0x1c, 0xc6, 0xce, 0xf9, 0x78, 0xf9, 0x2d, 0xf6, 0xfc, 0xfd, 0x21,
	0x9f, 0xc1, 0xc0, 0x85, 0xa8, 0xeb, 0xde, 0x5b, 0x35, 0xd1, 0xc8, 0x3f, 0x00, 0x78, 0x63, 0xee,
	0xe0, 0xcd, 0x8a, 0x4c, 0xc0, 0x6b, 0x7c, 0xc7, 0x57, 0x92, 0xe0, 0x98, 0x17, 0x99, 0x7b, 0x8b,
	0xde, 0x6b, 0xb5, 0xc3, 0x6d, 0x21, 0xc7, 0xbb, 0x7f, 0xbf, 0x9b, 0x7a, 0x6f, 0xdf, 0x4d, 0x3b,
	0xbf, 0x5f, 0x4d, 0x3b, 0x6f, 0xaf, 0xa6, 0x9d, 0x7f, 0xae, 0xa6, 0x9d, 0x1f, 0x7b, 0xb4, 0xe4,
	0x07, 0xe5, 0x6c, 0xd6, 0x33, 0xbf, 0x31, 0x87, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xd7,
	0x58, 0x9d, 0xfa, 0x08, 0x00, 0x00,
}
