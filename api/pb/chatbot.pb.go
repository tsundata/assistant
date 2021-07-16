// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chatbot.proto

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

type ChatbotRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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
	Text                 []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func init() {
	proto.RegisterType((*ChatbotRequest)(nil), "pb.ChatbotRequest")
	proto.RegisterType((*ChatbotReply)(nil), "pb.ChatbotReply")
}

func init() { proto.RegisterFile("chatbot.proto", fileDescriptor_acc44097314201ac) }

var fileDescriptor_acc44097314201ac = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x48, 0x2c,
	0x49, 0xca, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe1,
	0xe2, 0x73, 0x86, 0x08, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0x94,
	0xa4, 0x56, 0x94, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x4a, 0x5c, 0x3c,
	0x70, 0x55, 0x05, 0x39, 0x95, 0x48, 0x6a, 0x98, 0x61, 0x6a, 0x8c, 0xac, 0xb9, 0xd8, 0xa1, 0x6a,
	0x84, 0x0c, 0xb8, 0xd8, 0x3c, 0x12, 0xf3, 0x52, 0x72, 0x52, 0x85, 0x84, 0xf4, 0x0a, 0x92, 0xf4,
	0x50, 0x2d, 0x90, 0x12, 0x40, 0x11, 0x2b, 0xc8, 0xa9, 0x54, 0x62, 0x70, 0xe2, 0x88, 0x62, 0x4b,
	0x2c, 0xc8, 0xd4, 0x2f, 0x48, 0x4a, 0x62, 0x03, 0xbb, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x82, 0x90, 0xd1, 0x51, 0xac, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatbotClient is the client API for Chatbot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatbotClient interface {
	Handle(ctx context.Context, in *ChatbotRequest, opts ...grpc.CallOption) (*ChatbotReply, error)
}

type chatbotClient struct {
	cc *grpc.ClientConn
}

func NewChatbotClient(cc *grpc.ClientConn) ChatbotClient {
	return &chatbotClient{cc}
}

func (c *chatbotClient) Handle(ctx context.Context, in *ChatbotRequest, opts ...grpc.CallOption) (*ChatbotReply, error) {
	out := new(ChatbotReply)
	err := c.cc.Invoke(ctx, "/pb.Chatbot/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatbotServer is the server API for Chatbot service.
type ChatbotServer interface {
	Handle(context.Context, *ChatbotRequest) (*ChatbotReply, error)
}

// UnimplementedChatbotServer can be embedded to have forward compatible implementations.
type UnimplementedChatbotServer struct {
}

func (*UnimplementedChatbotServer) Handle(ctx context.Context, req *ChatbotRequest) (*ChatbotReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}

func RegisterChatbotServer(s *grpc.Server, srv ChatbotServer) {
	s.RegisterService(&_Chatbot_serviceDesc, srv)
}

func _Chatbot_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatbotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatbotServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Chatbot/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatbotServer).Handle(ctx, req.(*ChatbotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chatbot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Chatbot",
	HandlerType: (*ChatbotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Chatbot_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatbot.proto",
}
