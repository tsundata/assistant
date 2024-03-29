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
const _ = grpc.SupportPackageIsVersion7

// OkrSvcClient is the client API for OkrSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OkrSvcClient interface {
	CreateObjective(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*StateReply, error)
	UpdateObjective(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*StateReply, error)
	GetObjective(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*ObjectiveReply, error)
	GetObjectives(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*ObjectivesReply, error)
	DeleteObjective(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*StateReply, error)
	CreateKeyResult(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*StateReply, error)
	UpdateKeyResult(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*StateReply, error)
	GetKeyResult(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*KeyResultReply, error)
	GetKeyResults(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*KeyResultsReply, error)
	DeleteKeyResult(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*StateReply, error)
	CreateKeyResultValue(ctx context.Context, in *KeyResultValueRequest, opts ...grpc.CallOption) (*StateReply, error)
	GetKeyResultsByTag(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*KeyResultsReply, error)
	GetKeyResultValues(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*KeyResultValuesReply, error)
}

type okrSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewOkrSvcClient(cc grpc.ClientConnInterface) OkrSvcClient {
	return &okrSvcClient{cc}
}

func (c *okrSvcClient) CreateObjective(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/CreateObjective", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) UpdateObjective(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/UpdateObjective", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) GetObjective(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*ObjectiveReply, error) {
	out := new(ObjectiveReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/GetObjective", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) GetObjectives(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*ObjectivesReply, error) {
	out := new(ObjectivesReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/GetObjectives", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) DeleteObjective(ctx context.Context, in *ObjectiveRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/DeleteObjective", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) CreateKeyResult(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/CreateKeyResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) UpdateKeyResult(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/UpdateKeyResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) GetKeyResult(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*KeyResultReply, error) {
	out := new(KeyResultReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/GetKeyResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) GetKeyResults(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*KeyResultsReply, error) {
	out := new(KeyResultsReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/GetKeyResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) DeleteKeyResult(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/DeleteKeyResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) CreateKeyResultValue(ctx context.Context, in *KeyResultValueRequest, opts ...grpc.CallOption) (*StateReply, error) {
	out := new(StateReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/CreateKeyResultValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) GetKeyResultsByTag(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*KeyResultsReply, error) {
	out := new(KeyResultsReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/GetKeyResultsByTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *okrSvcClient) GetKeyResultValues(ctx context.Context, in *KeyResultRequest, opts ...grpc.CallOption) (*KeyResultValuesReply, error) {
	out := new(KeyResultValuesReply)
	err := c.cc.Invoke(ctx, "/pb.OkrSvc/GetKeyResultValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OkrSvcServer is the server API for OkrSvc service.
// All implementations should embed UnimplementedOkrSvcServer
// for forward compatibility
type OkrSvcServer interface {
	CreateObjective(context.Context, *ObjectiveRequest) (*StateReply, error)
	UpdateObjective(context.Context, *ObjectiveRequest) (*StateReply, error)
	GetObjective(context.Context, *ObjectiveRequest) (*ObjectiveReply, error)
	GetObjectives(context.Context, *ObjectiveRequest) (*ObjectivesReply, error)
	DeleteObjective(context.Context, *ObjectiveRequest) (*StateReply, error)
	CreateKeyResult(context.Context, *KeyResultRequest) (*StateReply, error)
	UpdateKeyResult(context.Context, *KeyResultRequest) (*StateReply, error)
	GetKeyResult(context.Context, *KeyResultRequest) (*KeyResultReply, error)
	GetKeyResults(context.Context, *KeyResultRequest) (*KeyResultsReply, error)
	DeleteKeyResult(context.Context, *KeyResultRequest) (*StateReply, error)
	CreateKeyResultValue(context.Context, *KeyResultValueRequest) (*StateReply, error)
	GetKeyResultsByTag(context.Context, *KeyResultRequest) (*KeyResultsReply, error)
	GetKeyResultValues(context.Context, *KeyResultRequest) (*KeyResultValuesReply, error)
}

// UnimplementedOkrSvcServer should be embedded to have forward compatible implementations.
type UnimplementedOkrSvcServer struct {
}

func (UnimplementedOkrSvcServer) CreateObjective(context.Context, *ObjectiveRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjective not implemented")
}
func (UnimplementedOkrSvcServer) UpdateObjective(context.Context, *ObjectiveRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObjective not implemented")
}
func (UnimplementedOkrSvcServer) GetObjective(context.Context, *ObjectiveRequest) (*ObjectiveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjective not implemented")
}
func (UnimplementedOkrSvcServer) GetObjectives(context.Context, *ObjectiveRequest) (*ObjectivesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectives not implemented")
}
func (UnimplementedOkrSvcServer) DeleteObjective(context.Context, *ObjectiveRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjective not implemented")
}
func (UnimplementedOkrSvcServer) CreateKeyResult(context.Context, *KeyResultRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeyResult not implemented")
}
func (UnimplementedOkrSvcServer) UpdateKeyResult(context.Context, *KeyResultRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKeyResult not implemented")
}
func (UnimplementedOkrSvcServer) GetKeyResult(context.Context, *KeyResultRequest) (*KeyResultReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyResult not implemented")
}
func (UnimplementedOkrSvcServer) GetKeyResults(context.Context, *KeyResultRequest) (*KeyResultsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyResults not implemented")
}
func (UnimplementedOkrSvcServer) DeleteKeyResult(context.Context, *KeyResultRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKeyResult not implemented")
}
func (UnimplementedOkrSvcServer) CreateKeyResultValue(context.Context, *KeyResultValueRequest) (*StateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeyResultValue not implemented")
}
func (UnimplementedOkrSvcServer) GetKeyResultsByTag(context.Context, *KeyResultRequest) (*KeyResultsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyResultsByTag not implemented")
}
func (UnimplementedOkrSvcServer) GetKeyResultValues(context.Context, *KeyResultRequest) (*KeyResultValuesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeyResultValues not implemented")
}

// UnsafeOkrSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OkrSvcServer will
// result in compilation errors.
type UnsafeOkrSvcServer interface {
	mustEmbedUnimplementedOkrSvcServer()
}

func RegisterOkrSvcServer(s *grpc.Server, srv OkrSvcServer) {
	s.RegisterService(&_OkrSvc_serviceDesc, srv)
}

func _OkrSvc_CreateObjective_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).CreateObjective(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/CreateObjective",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).CreateObjective(ctx, req.(*ObjectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_UpdateObjective_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).UpdateObjective(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/UpdateObjective",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).UpdateObjective(ctx, req.(*ObjectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_GetObjective_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).GetObjective(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/GetObjective",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).GetObjective(ctx, req.(*ObjectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_GetObjectives_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).GetObjectives(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/GetObjectives",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).GetObjectives(ctx, req.(*ObjectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_DeleteObjective_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).DeleteObjective(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/DeleteObjective",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).DeleteObjective(ctx, req.(*ObjectiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_CreateKeyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).CreateKeyResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/CreateKeyResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).CreateKeyResult(ctx, req.(*KeyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_UpdateKeyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).UpdateKeyResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/UpdateKeyResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).UpdateKeyResult(ctx, req.(*KeyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_GetKeyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).GetKeyResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/GetKeyResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).GetKeyResult(ctx, req.(*KeyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_GetKeyResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).GetKeyResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/GetKeyResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).GetKeyResults(ctx, req.(*KeyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_DeleteKeyResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).DeleteKeyResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/DeleteKeyResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).DeleteKeyResult(ctx, req.(*KeyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_CreateKeyResultValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyResultValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).CreateKeyResultValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/CreateKeyResultValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).CreateKeyResultValue(ctx, req.(*KeyResultValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_GetKeyResultsByTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).GetKeyResultsByTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/GetKeyResultsByTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).GetKeyResultsByTag(ctx, req.(*KeyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OkrSvc_GetKeyResultValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OkrSvcServer).GetKeyResultValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OkrSvc/GetKeyResultValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OkrSvcServer).GetKeyResultValues(ctx, req.(*KeyResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OkrSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OkrSvc",
	HandlerType: (*OkrSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjective",
			Handler:    _OkrSvc_CreateObjective_Handler,
		},
		{
			MethodName: "UpdateObjective",
			Handler:    _OkrSvc_UpdateObjective_Handler,
		},
		{
			MethodName: "GetObjective",
			Handler:    _OkrSvc_GetObjective_Handler,
		},
		{
			MethodName: "GetObjectives",
			Handler:    _OkrSvc_GetObjectives_Handler,
		},
		{
			MethodName: "DeleteObjective",
			Handler:    _OkrSvc_DeleteObjective_Handler,
		},
		{
			MethodName: "CreateKeyResult",
			Handler:    _OkrSvc_CreateKeyResult_Handler,
		},
		{
			MethodName: "UpdateKeyResult",
			Handler:    _OkrSvc_UpdateKeyResult_Handler,
		},
		{
			MethodName: "GetKeyResult",
			Handler:    _OkrSvc_GetKeyResult_Handler,
		},
		{
			MethodName: "GetKeyResults",
			Handler:    _OkrSvc_GetKeyResults_Handler,
		},
		{
			MethodName: "DeleteKeyResult",
			Handler:    _OkrSvc_DeleteKeyResult_Handler,
		},
		{
			MethodName: "CreateKeyResultValue",
			Handler:    _OkrSvc_CreateKeyResultValue_Handler,
		},
		{
			MethodName: "GetKeyResultsByTag",
			Handler:    _OkrSvc_GetKeyResultsByTag_Handler,
		},
		{
			MethodName: "GetKeyResultValues",
			Handler:    _OkrSvc_GetKeyResultValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "okr.proto",
}
