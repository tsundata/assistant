// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageSvcClient is the client API for MessageSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageSvcClient interface {
	List(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessagesReply, error)
	ListByGroup(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesReply, error)
	LastByGroup(ctx context.Context, in *LastByGroupRequest, opts ...grpc.CallOption) (*LastByGroupReply, error)
	GetByUuid(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*GetMessageReply, error)
	GetById(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*GetMessageReply, error)
	GetBySequence(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*GetMessageReply, error)
	Create(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
	Save(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
	Delete(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*TextReply, error)
	Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*StateReply, error)
	Run(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*TextReply, error)
	GetActionMessages(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ActionReply, error)
	CreateActionMessage(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*StateReply, error)
	DeleteWorkflowMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*StateReply, error)
	ListInbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxReply, error)
	LastInbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxReply, error)
	MarkSendInbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxReply, error)
	MarkReadInbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxReply, error)
}

type messageSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageSvcClient(cc grpc.ClientConnInterface) MessageSvcClient {
	return &messageSvcClient{cc}
}

func (c *messageSvcClient) List(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessagesReply, error) {
	out := new(MessagesReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) ListByGroup(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesReply, error) {
	out := new(GetMessagesReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/ListByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) LastByGroup(ctx context.Context, in *LastByGroupRequest, opts ...grpc.CallOption) (*LastByGroupReply, error) {
	out := new(LastByGroupReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/LastByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) GetByUuid(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*GetMessageReply, error) {
	out := new(GetMessageReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/GetByUuid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) GetById(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*GetMessageReply, error) {
	out := new(GetMessageReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) GetBySequence(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*GetMessageReply, error) {
	out := new(GetMessageReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/GetBySequence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) Create(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) Save(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) Delete(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*TextReply, error) {
	out := new(TextReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) Run(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*TextReply, error) {
	out := new(TextReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) GetActionMessages(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*ActionReply, error) {
	out := new(ActionReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/GetActionMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) CreateActionMessage(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/CreateActionMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) DeleteWorkflowMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/DeleteWorkflowMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) ListInbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxReply, error) {
	out := new(InboxReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/ListInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) LastInbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxReply, error) {
	out := new(InboxReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/LastInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) MarkSendInbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxReply, error) {
	out := new(InboxReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/MarkSendInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageSvcClient) MarkReadInbox(ctx context.Context, in *InboxRequest, opts ...grpc.CallOption) (*InboxReply, error) {
	out := new(InboxReply)
	err := c.cc.Invoke(ctx, "/pb.MessageSvc/MarkReadInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageSvcServer is the server API for MessageSvc service.
// All implementations should embed UnimplementedMessageSvcServer
// for forward compatibility
type MessageSvcServer interface {
	List(context.Context, *MessageRequest) (*MessagesReply, error)
	ListByGroup(context.Context, *GetMessagesRequest) (*GetMessagesReply, error)
	LastByGroup(context.Context, *LastByGroupRequest) (*LastByGroupReply, error)
	GetByUuid(context.Context, *MessageRequest) (*GetMessageReply, error)
	GetById(context.Context, *MessageRequest) (*GetMessageReply, error)
	GetBySequence(context.Context, *MessageRequest) (*GetMessageReply, error)
	Create(context.Context, *MessageRequest) (*MessageReply, error)
	Save(context.Context, *MessageRequest) (*MessageReply, error)
	Delete(context.Context, *MessageRequest) (*TextReply, error)
	Send(context.Context, *MessageRequest) (*StateReply, error)
	Run(context.Context, *MessageRequest) (*TextReply, error)
	GetActionMessages(context.Context, *TextRequest) (*ActionReply, error)
	CreateActionMessage(context.Context, *TextRequest) (*StateReply, error)
	DeleteWorkflowMessage(context.Context, *MessageRequest) (*StateReply, error)
	ListInbox(context.Context, *InboxRequest) (*InboxReply, error)
	LastInbox(context.Context, *InboxRequest) (*InboxReply, error)
	MarkSendInbox(context.Context, *InboxRequest) (*InboxReply, error)
	MarkReadInbox(context.Context, *InboxRequest) (*InboxReply, error)
}

// UnimplementedMessageSvcServer should be embedded to have forward compatible implementations.
type UnimplementedMessageSvcServer struct {
}

func (UnimplementedMessageSvcServer) List(context.Context, *MessageRequest) (*MessagesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMessageSvcServer) ListByGroup(context.Context, *GetMessagesRequest) (*GetMessagesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByGroup not implemented")
}
func (UnimplementedMessageSvcServer) LastByGroup(context.Context, *LastByGroupRequest) (*LastByGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastByGroup not implemented")
}
func (UnimplementedMessageSvcServer) GetByUuid(context.Context, *MessageRequest) (*GetMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUuid not implemented")
}
func (UnimplementedMessageSvcServer) GetById(context.Context, *MessageRequest) (*GetMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedMessageSvcServer) GetBySequence(context.Context, *MessageRequest) (*GetMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySequence not implemented")
}
func (UnimplementedMessageSvcServer) Create(context.Context, *MessageRequest) (*MessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMessageSvcServer) Save(context.Context, *MessageRequest) (*MessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedMessageSvcServer) Delete(context.Context, *MessageRequest) (*TextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMessageSvcServer) Send(context.Context, *MessageRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedMessageSvcServer) Run(context.Context, *MessageRequest) (*TextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedMessageSvcServer) GetActionMessages(context.Context, *TextRequest) (*ActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActionMessages not implemented")
}
func (UnimplementedMessageSvcServer) CreateActionMessage(context.Context, *TextRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActionMessage not implemented")
}
func (UnimplementedMessageSvcServer) DeleteWorkflowMessage(context.Context, *MessageRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflowMessage not implemented")
}
func (UnimplementedMessageSvcServer) ListInbox(context.Context, *InboxRequest) (*InboxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInbox not implemented")
}
func (UnimplementedMessageSvcServer) LastInbox(context.Context, *InboxRequest) (*InboxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastInbox not implemented")
}
func (UnimplementedMessageSvcServer) MarkSendInbox(context.Context, *InboxRequest) (*InboxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkSendInbox not implemented")
}
func (UnimplementedMessageSvcServer) MarkReadInbox(context.Context, *InboxRequest) (*InboxReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkReadInbox not implemented")
}

// UnsafeMessageSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageSvcServer will
// result in compilation errors.
type UnsafeMessageSvcServer interface {
	mustEmbedUnimplementedMessageSvcServer()
}

func RegisterMessageSvcServer(s grpc.ServiceRegistrar, srv MessageSvcServer) {
	s.RegisterService(&MessageSvc_ServiceDesc, srv)
}

func _MessageSvc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).List(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_ListByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).ListByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/ListByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).ListByGroup(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_LastByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastByGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).LastByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/LastByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).LastByGroup(ctx, req.(*LastByGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_GetByUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).GetByUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/GetByUuid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).GetByUuid(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).GetById(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_GetBySequence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).GetBySequence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/GetBySequence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).GetBySequence(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).Create(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).Save(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).Delete(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).Send(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).Run(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_GetActionMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).GetActionMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/GetActionMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).GetActionMessages(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_CreateActionMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).CreateActionMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/CreateActionMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).CreateActionMessage(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_DeleteWorkflowMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).DeleteWorkflowMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/DeleteWorkflowMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).DeleteWorkflowMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_ListInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).ListInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/ListInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).ListInbox(ctx, req.(*InboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_LastInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).LastInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/LastInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).LastInbox(ctx, req.(*InboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_MarkSendInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).MarkSendInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/MarkSendInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).MarkSendInbox(ctx, req.(*InboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageSvc_MarkReadInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InboxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSvcServer).MarkReadInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageSvc/MarkReadInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSvcServer).MarkReadInbox(ctx, req.(*InboxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageSvc_ServiceDesc is the grpc.ServiceDesc for MessageSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageSvc",
	HandlerType: (*MessageSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _MessageSvc_List_Handler,
		},
		{
			MethodName: "ListByGroup",
			Handler:    _MessageSvc_ListByGroup_Handler,
		},
		{
			MethodName: "LastByGroup",
			Handler:    _MessageSvc_LastByGroup_Handler,
		},
		{
			MethodName: "GetByUuid",
			Handler:    _MessageSvc_GetByUuid_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _MessageSvc_GetById_Handler,
		},
		{
			MethodName: "GetBySequence",
			Handler:    _MessageSvc_GetBySequence_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _MessageSvc_Create_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _MessageSvc_Save_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MessageSvc_Delete_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _MessageSvc_Send_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _MessageSvc_Run_Handler,
		},
		{
			MethodName: "GetActionMessages",
			Handler:    _MessageSvc_GetActionMessages_Handler,
		},
		{
			MethodName: "CreateActionMessage",
			Handler:    _MessageSvc_CreateActionMessage_Handler,
		},
		{
			MethodName: "DeleteWorkflowMessage",
			Handler:    _MessageSvc_DeleteWorkflowMessage_Handler,
		},
		{
			MethodName: "ListInbox",
			Handler:    _MessageSvc_ListInbox_Handler,
		},
		{
			MethodName: "LastInbox",
			Handler:    _MessageSvc_LastInbox_Handler,
		},
		{
			MethodName: "MarkSendInbox",
			Handler:    _MessageSvc_MarkSendInbox_Handler,
		},
		{
			MethodName: "MarkReadInbox",
			Handler:    _MessageSvc_MarkReadInbox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
