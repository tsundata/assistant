// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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
	return m.Unmarshal(b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return m.Size()
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

type isFileRequest_Data interface {
	isFileRequest_Data()
	MarshalTo([]byte) (int, error)
	Size() int
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
	return m.Unmarshal(b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return m.Size()
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
	return m.Unmarshal(b)
}
func (m *FileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileReply.Merge(m, src)
}
func (m *FileReply) XXX_Size() int {
	return m.Size()
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
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x0a, 0xe4,
	0xe2, 0x76, 0xcb, 0xcc, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe2, 0x62,
	0xc9, 0xcc, 0x4b, 0xcb, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd1, 0x2b, 0x48, 0xd2,
	0x03, 0x49, 0x7b, 0xe6, 0xa5, 0xe5, 0x7b, 0x30, 0x04, 0x81, 0xe5, 0x84, 0xc4, 0xb8, 0x58, 0x93,
	0x33, 0x4a, 0x93, 0xb3, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x3c, 0x18, 0x82, 0x20, 0x5c, 0x27,
	0x36, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x25, 0x35, 0x2e, 0x0e, 0x98, 0x1e, 0x21, 0x29, 0x2e,
	0x8e, 0xb4, 0xcc, 0x9c, 0xd4, 0x90, 0xca, 0x82, 0x54, 0xb0, 0x99, 0x9c, 0x41, 0x70, 0xbe, 0x92,
	0x3c, 0x17, 0x27, 0xc4, 0xea, 0x82, 0x9c, 0x4a, 0x21, 0x21, 0x2e, 0x96, 0x82, 0xc4, 0x92, 0x0c,
	0xa8, 0x22, 0x30, 0xdb, 0xc8, 0x8e, 0x8b, 0x2b, 0x18, 0xe2, 0xe0, 0xe0, 0xb2, 0x64, 0x21, 0x03,
	0x2e, 0xae, 0xd0, 0x82, 0x9c, 0xfc, 0xc4, 0x14, 0x90, 0x26, 0x21, 0x7e, 0x98, 0xd3, 0xa0, 0x2e,
	0x97, 0xe2, 0x45, 0x08, 0x14, 0xe4, 0x54, 0x2a, 0x31, 0x68, 0x30, 0x3a, 0x49, 0x9c, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0xb1,
	0x25, 0x16, 0x64, 0xea, 0x17, 0x24, 0x25, 0xb1, 0x81, 0x03, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x15, 0x97, 0x63, 0x4d, 0x11, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageSvcClient is the client API for StorageSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageSvcClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (StorageSvc_UploadFileClient, error)
}

type storageSvcClient struct {
	cc *grpc.ClientConn
}

func NewStorageSvcClient(cc *grpc.ClientConn) StorageSvcClient {
	return &storageSvcClient{cc}
}

func (c *storageSvcClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (StorageSvc_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StorageSvc_serviceDesc.Streams[0], "/pb.StorageSvc/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageSvcUploadFileClient{stream}
	return x, nil
}

type StorageSvc_UploadFileClient interface {
	Send(*FileRequest) error
	CloseAndRecv() (*FileReply, error)
	grpc.ClientStream
}

type storageSvcUploadFileClient struct {
	grpc.ClientStream
}

func (x *storageSvcUploadFileClient) Send(m *FileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageSvcUploadFileClient) CloseAndRecv() (*FileReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageSvcServer is the server API for StorageSvc service.
type StorageSvcServer interface {
	UploadFile(StorageSvc_UploadFileServer) error
}

// UnimplementedStorageSvcServer can be embedded to have forward compatible implementations.
type UnimplementedStorageSvcServer struct {
}

func (*UnimplementedStorageSvcServer) UploadFile(srv StorageSvc_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}

func RegisterStorageSvcServer(s *grpc.Server, srv StorageSvcServer) {
	s.RegisterService(&_StorageSvc_serviceDesc, srv)
}

func _StorageSvc_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageSvcServer).UploadFile(&storageSvcUploadFileServer{stream})
}

type StorageSvc_UploadFileServer interface {
	SendAndClose(*FileReply) error
	Recv() (*FileRequest, error)
	grpc.ServerStream
}

type storageSvcUploadFileServer struct {
	grpc.ServerStream
}

func (x *storageSvcUploadFileServer) SendAndClose(m *FileReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageSvcUploadFileServer) Recv() (*FileRequest, error) {
	m := new(FileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StorageSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StorageSvc",
	HandlerType: (*StorageSvcServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _StorageSvc_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "storage.proto",
}

func (m *FileRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		{
			size := m.Data.Size()
			i -= size
			if _, err := m.Data.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *FileRequest_Info) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileRequest_Info) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *FileRequest_Chuck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileRequest_Chuck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Chuck != nil {
		i -= len(m.Chuck)
		copy(dAtA[i:], m.Chuck)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Chuck)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *FileInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.FileType) > 0 {
		i -= len(m.FileType)
		copy(dAtA[i:], m.FileType)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.FileType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FileReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintStorage(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FileRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		n += m.Data.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FileRequest_Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}
func (m *FileRequest_Chuck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chuck != nil {
		l = len(m.Chuck)
		n += 1 + l + sovStorage(uint64(l))
	}
	return n
}
func (m *FileInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FileType)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FileReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovStorage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStorage(x uint64) (n int) {
	return sovStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FileInfo{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Data = &FileRequest_Info{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chuck", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Data = &FileRequest_Chuck{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FileInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FileReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStorage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStorage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStorage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStorage = fmt.Errorf("proto: unexpected end of group")
)
