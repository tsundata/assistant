// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task.proto

package pb

import (
	fmt "fmt"
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

type JobRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args                 string   `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobRequest) Reset()         { *m = JobRequest{} }
func (m *JobRequest) String() string { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()    {}
func (*JobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}
func (m *JobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobRequest.Unmarshal(m, b)
}
func (m *JobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobRequest.Marshal(b, m, deterministic)
}
func (m *JobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobRequest.Merge(m, src)
}
func (m *JobRequest) XXX_Size() int {
	return xxx_messageInfo_JobRequest.Size(m)
}
func (m *JobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobRequest proto.InternalMessageInfo

func (m *JobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobRequest) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *JobRequest) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func init() {
	proto.RegisterType((*JobRequest)(nil), "pb.JobRequest")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0x85, 0xf0, 0x95, 0x3c, 0xb8, 0xb8, 0xbc, 0xf2, 0x93, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0,
	0x6c, 0x90, 0x58, 0x62, 0x51, 0x7a, 0xb1, 0x04, 0x13, 0x44, 0x0c, 0xc4, 0x06, 0x89, 0x95, 0x64,
	0xe6, 0xa6, 0x4a, 0x30, 0x43, 0xc4, 0x40, 0x6c, 0x23, 0x13, 0x2e, 0xf6, 0x90, 0xc4, 0xe2, 0xec,
	0xe0, 0xb2, 0x64, 0x21, 0x4d, 0x2e, 0x56, 0x97, 0xd4, 0x9c, 0xc4, 0x4a, 0x21, 0x3e, 0xbd, 0x82,
	0x24, 0x3d, 0x84, 0xf9, 0x52, 0x60, 0x7e, 0x70, 0x49, 0x62, 0x49, 0x6a, 0x50, 0x6a, 0x41, 0x4e,
	0xa5, 0x12, 0x83, 0x13, 0x47, 0x14, 0x5b, 0x62, 0x41, 0xa6, 0x7e, 0x41, 0x52, 0x12, 0x1b, 0xd8,
	0x41, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x56, 0x5d, 0xd6, 0xae, 0x00, 0x00, 0x00,
}
