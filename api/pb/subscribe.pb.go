// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: subscribe.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SubscribeRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d2980c9543da44, []int{0}
}
func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type SubscribeReply struct {
	Text                 []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeReply) Reset()         { *m = SubscribeReply{} }
func (m *SubscribeReply) String() string { return proto.CompactTextString(m) }
func (*SubscribeReply) ProtoMessage()    {}
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_38d2980c9543da44, []int{1}
}
func (m *SubscribeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeReply.Unmarshal(m, b)
}
func (m *SubscribeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeReply.Marshal(b, m, deterministic)
}
func (m *SubscribeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeReply.Merge(m, src)
}
func (m *SubscribeReply) XXX_Size() int {
	return xxx_messageInfo_SubscribeReply.Size(m)
}
func (m *SubscribeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeReply proto.InternalMessageInfo

func (m *SubscribeReply) GetText() []string {
	if m != nil {
		return m.Text
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscribeRequest)(nil), "pb.SubscribeRequest")
	proto.RegisterType((*SubscribeReply)(nil), "pb.SubscribeReply")
}

func init() { proto.RegisterFile("subscribe.proto", fileDescriptor_38d2980c9543da44) }

var fileDescriptor_38d2980c9543da44 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2e, 0x4d, 0x2a,
	0x4e, 0x2e, 0xca, 0x4c, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92,
	0xe2, 0x4a, 0x4a, 0x2c, 0x86, 0xf2, 0x95, 0xd4, 0xb8, 0x04, 0x82, 0x61, 0x4a, 0x82, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x15, 0x2e, 0x3e, 0x24, 0x75, 0x05, 0x39, 0x95, 0x48, 0xaa,
	0x98, 0x61, 0xaa, 0x8c, 0xda, 0x98, 0xb8, 0x38, 0xe1, 0xca, 0x84, 0x8c, 0xb8, 0x58, 0x7c, 0x32,
	0x8b, 0x4b, 0x84, 0x44, 0xf4, 0x0a, 0x92, 0xf4, 0xd0, 0x6d, 0x91, 0x12, 0x42, 0x13, 0x2d, 0xc8,
	0xa9, 0x54, 0x62, 0x10, 0x32, 0xe2, 0xe2, 0x08, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0xc2,
	0xa1, 0x8f, 0x0f, 0x2c, 0x5a, 0x92, 0x58, 0x02, 0xd7, 0xa3, 0xc7, 0xc5, 0xe2, 0x5f, 0x90, 0x9a,
	0x47, 0xb4, 0x7a, 0x7d, 0x2e, 0x56, 0xe7, 0x9c, 0xfc, 0xe2, 0x54, 0xa2, 0x35, 0x18, 0x70, 0xb1,
	0x81, 0xf8, 0xa5, 0xc5, 0xc4, 0xea, 0x70, 0xe2, 0x88, 0x62, 0x4b, 0x2c, 0xc8, 0xd4, 0x2f, 0x48,
	0x4a, 0x62, 0x03, 0x87, 0xb3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x64, 0x93, 0xbc, 0x81, 0x8a,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SubscribeClient is the client API for Subscribe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscribeClient interface {
	List(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeReply, error)
	Register(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*StateReply, error)
	Open(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*StateReply, error)
	Close(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*StateReply, error)
	Status(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*StateReply, error)
}

type subscribeClient struct {
	cc *grpc.ClientConn
}

func NewSubscribeClient(cc *grpc.ClientConn) SubscribeClient {
	return &subscribeClient{cc}
}

func (c *subscribeClient) List(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeReply, error) {
	out := new(SubscribeReply)
	err := c.cc.Invoke(ctx, "/pb.Subscribe/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Register(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.Subscribe/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Open(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.Subscribe/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Close(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.Subscribe/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Status(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.Subscribe/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeServer is the server API for Subscribe service.
type SubscribeServer interface {
	List(context.Context, *SubscribeRequest) (*SubscribeReply, error)
	Register(context.Context, *SubscribeRequest) (*StateReply, error)
	Open(context.Context, *SubscribeRequest) (*StateReply, error)
	Close(context.Context, *SubscribeRequest) (*StateReply, error)
	Status(context.Context, *SubscribeRequest) (*StateReply, error)
}

// UnimplementedSubscribeServer can be embedded to have forward compatible implementations.
type UnimplementedSubscribeServer struct {
}

func (*UnimplementedSubscribeServer) List(ctx context.Context, req *SubscribeRequest) (*SubscribeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedSubscribeServer) Register(ctx context.Context, req *SubscribeRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedSubscribeServer) Open(ctx context.Context, req *SubscribeRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (*UnimplementedSubscribeServer) Close(ctx context.Context, req *SubscribeRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (*UnimplementedSubscribeServer) Status(ctx context.Context, req *SubscribeRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

func RegisterSubscribeServer(s *grpc.Server, srv SubscribeServer) {
	s.RegisterService(&_Subscribe_serviceDesc, srv)
}

func _Subscribe_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Subscribe/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).List(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Subscribe/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Register(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Subscribe/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Open(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Subscribe/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Close(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Subscribe/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Status(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Subscribe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Subscribe",
	HandlerType: (*SubscribeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Subscribe_List_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Subscribe_Register_Handler,
		},
		{
			MethodName: "Open",
			Handler:    _Subscribe_Open_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Subscribe_Close_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Subscribe_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscribe.proto",
}
