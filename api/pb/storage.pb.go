// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage.proto

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

type FileRequest struct {
	// Types that are valid to be assigned to Data:
	//	*FileRequest_Info
	//	*FileRequest_Chuck
	Data                 isFileRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}
func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

type isFileRequest_Data interface {
	isFileRequest_Data()
}

type FileRequest_Info struct {
	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof" json:"info,omitempty"`
}
type FileRequest_Chuck struct {
	Chuck []byte `protobuf:"bytes,2,opt,name=chuck,proto3,oneof" json:"chuck,omitempty"`
}

func (*FileRequest_Info) isFileRequest_Data()  {}
func (*FileRequest_Chuck) isFileRequest_Data() {}

func (m *FileRequest) GetData() isFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FileRequest) GetInfo() *FileInfo {
	if x, ok := m.GetData().(*FileRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *FileRequest) GetChuck() []byte {
	if x, ok := m.GetData().(*FileRequest_Chuck); ok {
		return x.Chuck
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileRequest_Info)(nil),
		(*FileRequest_Chuck)(nil),
	}
}

type FileInfo struct {
	FileType             string   `protobuf:"bytes,1,opt,name=fileType,proto3" json:"fileType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}
func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetFileType() string {
	if m != nil {
		return m.FileType
	}
	return ""
}

type FileReply struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileReply) Reset()         { *m = FileReply{} }
func (m *FileReply) String() string { return proto.CompactTextString(m) }
func (*FileReply) ProtoMessage()    {}
func (*FileReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}
func (m *FileReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileReply.Unmarshal(m, b)
}
func (m *FileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileReply.Marshal(b, m, deterministic)
}
func (m *FileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileReply.Merge(m, src)
}
func (m *FileReply) XXX_Size() int {
	return xxx_messageInfo_FileReply.Size(m)
}
func (m *FileReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FileReply.DiscardUnknown(m)
}

var xxx_messageInfo_FileReply proto.InternalMessageInfo

func (m *FileReply) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*FileRequest)(nil), "pb.FileRequest")
	proto.RegisterType((*FileInfo)(nil), "pb.FileInfo")
	proto.RegisterType((*FileReply)(nil), "pb.FileReply")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xe3, 0x2a, 0x84, 0xf4, 0xda, 0x0a, 0xe9, 0x06, 0x54, 0x65, 0xa1, 0xf2, 0x80, 0x32,
	0x05, 0x54, 0x46, 0xb6, 0x0e, 0xa8, 0x8c, 0x18, 0x58, 0xd8, 0xec, 0xd6, 0x21, 0x16, 0x56, 0x7c,
	0x24, 0xce, 0x90, 0x7f, 0x8f, 0xe2, 0x24, 0x74, 0xbb, 0x77, 0xef, 0xde, 0xdd, 0x77, 0xb0, 0x69,
	0xbd, 0x6b, 0xe4, 0xb7, 0x2e, 0xa8, 0x71, 0xde, 0xe1, 0x82, 0x54, 0x06, 0x4a, 0xb6, 0x93, 0xe6,
	0x6f, 0xb0, 0x7a, 0x31, 0x56, 0x0b, 0xfd, 0xdb, 0xe9, 0xd6, 0x23, 0x87, 0xd8, 0xd4, 0xa5, 0xdb,
	0xb2, 0x1d, 0xcb, 0x57, 0xfb, 0x75, 0x41, 0xaa, 0x18, 0xec, 0xd7, 0xba, 0x74, 0xc7, 0x48, 0x04,
	0x0f, 0x6f, 0xe1, 0xea, 0x54, 0x75, 0xa7, 0x9f, 0xed, 0x62, 0xc7, 0xf2, 0xf5, 0x31, 0x12, 0xa3,
	0x3c, 0x24, 0x10, 0x9f, 0xa5, 0x97, 0xfc, 0x1e, 0xd2, 0x39, 0x83, 0x19, 0xa4, 0xa5, 0xb1, 0xfa,
	0xa3, 0x27, 0x1d, 0x76, 0x2e, 0xc5, 0xbf, 0xe6, 0x77, 0xb0, 0x1c, 0x4f, 0x93, 0xed, 0x11, 0x21,
	0x26, 0xe9, 0xab, 0x69, 0x28, 0xd4, 0xfb, 0x67, 0xb8, 0x7e, 0x1f, 0xe1, 0xf1, 0x11, 0xe0, 0x93,
	0xac, 0x93, 0xe7, 0x21, 0x81, 0x37, 0x33, 0xd7, 0x84, 0x9d, 0x6d, 0x2e, 0x0d, 0xb2, 0x3d, 0x8f,
	0x72, 0x76, 0x48, 0xbf, 0x12, 0x49, 0xe6, 0x81, 0x94, 0x4a, 0xc2, 0xa7, 0x4f, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0xd7, 0x59, 0x9f, 0x0a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (Storage_UploadFileClient, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (Storage_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Storage_serviceDesc.Streams[0], "/pb.Storage/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageUploadFileClient{stream}
	return x, nil
}

type Storage_UploadFileClient interface {
	Send(*FileRequest) error
	CloseAndRecv() (*FileReply, error)
	grpc.ClientStream
}

type storageUploadFileClient struct {
	grpc.ClientStream
}

func (x *storageUploadFileClient) Send(m *FileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageUploadFileClient) CloseAndRecv() (*FileReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServer is the server API for Storage service.
type StorageServer interface {
	UploadFile(Storage_UploadFileServer) error
}

// UnimplementedStorageServer can be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (*UnimplementedStorageServer) UploadFile(srv Storage_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).UploadFile(&storageUploadFileServer{stream})
}

type Storage_UploadFileServer interface {
	SendAndClose(*FileReply) error
	Recv() (*FileRequest, error)
	grpc.ServerStream
}

type storageUploadFileServer struct {
	grpc.ServerStream
}

func (x *storageUploadFileServer) SendAndClose(m *FileReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageUploadFileServer) Recv() (*FileRequest, error) {
	m := new(FileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _Storage_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "storage.proto",
}