// Code generated by protoc-gen-go.
// source: pkg/protobuf/build/build.proto
// DO NOT EDIT!

/*
Package build is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/build/build.proto

It has these top-level messages:
	BuildRequest
	BuildResponse
	ListRequest
	ListResponse
	RunRequest
	RunResponse
	DeleteRequest
	Empty
*/
package build

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BuildRequest struct {
	// Types that are valid to be assigned to Value:
	//	*BuildRequest_Info_
	//	*BuildRequest_File_
	Value isBuildRequest_Value `protobuf_oneof:"value"`
}

func (m *BuildRequest) Reset()                    { *m = BuildRequest{} }
func (m *BuildRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()               {}
func (*BuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isBuildRequest_Value interface{ isBuildRequest_Value() }

type BuildRequest_Info_ struct {
	Info *BuildRequest_Info `protobuf:"bytes,1,opt,name=info,oneof"`
}
type BuildRequest_File_ struct {
	File *BuildRequest_File `protobuf:"bytes,2,opt,name=file,oneof"`
}

func (*BuildRequest_Info_) isBuildRequest_Value() {}
func (*BuildRequest_File_) isBuildRequest_Value() {}

func (m *BuildRequest) GetValue() isBuildRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *BuildRequest) GetInfo() *BuildRequest_Info {
	if x, ok := m.GetValue().(*BuildRequest_Info_); ok {
		return x.Info
	}
	return nil
}

func (m *BuildRequest) GetFile() *BuildRequest_File {
	if x, ok := m.GetValue().(*BuildRequest_File_); ok {
		return x.File
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BuildRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BuildRequest_OneofMarshaler, _BuildRequest_OneofUnmarshaler, _BuildRequest_OneofSizer, []interface{}{
		(*BuildRequest_Info_)(nil),
		(*BuildRequest_File_)(nil),
	}
}

func _BuildRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BuildRequest)
	// value
	switch x := m.Value.(type) {
	case *BuildRequest_Info_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Info); err != nil {
			return err
		}
	case *BuildRequest_File_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.File); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BuildRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _BuildRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BuildRequest)
	switch tag {
	case 1: // value.info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildRequest_Info)
		err := b.DecodeMessage(msg)
		m.Value = &BuildRequest_Info_{msg}
		return true, err
	case 2: // value.file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuildRequest_File)
		err := b.DecodeMessage(msg)
		m.Value = &BuildRequest_File_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BuildRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BuildRequest)
	// value
	switch x := m.Value.(type) {
	case *BuildRequest_Info_:
		s := proto.Size(x.Info)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BuildRequest_File_:
		s := proto.Size(x.File)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BuildRequest_Info struct {
	App  string `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Run  bool   `protobuf:"varint,3,opt,name=run" json:"run,omitempty"`
}

func (m *BuildRequest_Info) Reset()                    { *m = BuildRequest_Info{} }
func (m *BuildRequest_Info) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest_Info) ProtoMessage()               {}
func (*BuildRequest_Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *BuildRequest_Info) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *BuildRequest_Info) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BuildRequest_Info) GetRun() bool {
	if m != nil {
		return m.Run
	}
	return false
}

type BuildRequest_File struct {
	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (m *BuildRequest_File) Reset()                    { *m = BuildRequest_File{} }
func (m *BuildRequest_File) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest_File) ProtoMessage()               {}
func (*BuildRequest_File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *BuildRequest_File) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type BuildResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *BuildResponse) Reset()                    { *m = BuildResponse{} }
func (m *BuildResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()               {}
func (*BuildResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BuildResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ListRequest struct {
	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName" json:"app_name,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

type ListResponse struct {
	Builds []*ListResponse_Build `protobuf:"bytes,1,rep,name=builds" json:"builds,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListResponse) GetBuilds() []*ListResponse_Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

type ListResponse_Build struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	LastModified string `protobuf:"bytes,2,opt,name=last_modified,json=lastModified" json:"last_modified,omitempty"`
}

func (m *ListResponse_Build) Reset()                    { *m = ListResponse_Build{} }
func (m *ListResponse_Build) String() string            { return proto.CompactTextString(m) }
func (*ListResponse_Build) ProtoMessage()               {}
func (*ListResponse_Build) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *ListResponse_Build) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListResponse_Build) GetLastModified() string {
	if m != nil {
		return m.LastModified
	}
	return ""
}

type RunRequest struct {
	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName" json:"app_name,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *RunRequest) Reset()                    { *m = RunRequest{} }
func (m *RunRequest) String() string            { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()               {}
func (*RunRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RunRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *RunRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RunResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *RunResponse) Reset()                    { *m = RunResponse{} }
func (m *RunResponse) String() string            { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()               {}
func (*RunResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RunResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type DeleteRequest struct {
	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName" json:"app_name,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteRequest) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*BuildRequest)(nil), "build.BuildRequest")
	proto.RegisterType((*BuildRequest_Info)(nil), "build.BuildRequest.Info")
	proto.RegisterType((*BuildRequest_File)(nil), "build.BuildRequest.File")
	proto.RegisterType((*BuildResponse)(nil), "build.BuildResponse")
	proto.RegisterType((*ListRequest)(nil), "build.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "build.ListResponse")
	proto.RegisterType((*ListResponse_Build)(nil), "build.ListResponse.Build")
	proto.RegisterType((*RunRequest)(nil), "build.RunRequest")
	proto.RegisterType((*RunResponse)(nil), "build.RunResponse")
	proto.RegisterType((*DeleteRequest)(nil), "build.DeleteRequest")
	proto.RegisterType((*Empty)(nil), "build.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Build service

type BuildClient interface {
	Make(ctx context.Context, opts ...grpc.CallOption) (Build_MakeClient, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (Build_RunClient, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
}

type buildClient struct {
	cc *grpc.ClientConn
}

func NewBuildClient(cc *grpc.ClientConn) BuildClient {
	return &buildClient{cc}
}

func (c *buildClient) Make(ctx context.Context, opts ...grpc.CallOption) (Build_MakeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Build_serviceDesc.Streams[0], c.cc, "/build.Build/Make", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildMakeClient{stream}
	return x, nil
}

type Build_MakeClient interface {
	Send(*BuildRequest) error
	Recv() (*BuildResponse, error)
	grpc.ClientStream
}

type buildMakeClient struct {
	grpc.ClientStream
}

func (x *buildMakeClient) Send(m *BuildRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *buildMakeClient) Recv() (*BuildResponse, error) {
	m := new(BuildResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/build.Build/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (Build_RunClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Build_serviceDesc.Streams[1], c.cc, "/build.Build/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &buildRunClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Build_RunClient interface {
	Recv() (*RunResponse, error)
	grpc.ClientStream
}

type buildRunClient struct {
	grpc.ClientStream
}

func (x *buildRunClient) Recv() (*RunResponse, error) {
	m := new(RunResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *buildClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/build.Build/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Build service

type BuildServer interface {
	Make(Build_MakeServer) error
	List(context.Context, *ListRequest) (*ListResponse, error)
	Run(*RunRequest, Build_RunServer) error
	Delete(context.Context, *DeleteRequest) (*Empty, error)
}

func RegisterBuildServer(s *grpc.Server, srv BuildServer) {
	s.RegisterService(&_Build_serviceDesc, srv)
}

func _Build_Make_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BuildServer).Make(&buildMakeServer{stream})
}

type Build_MakeServer interface {
	Send(*BuildResponse) error
	Recv() (*BuildRequest, error)
	grpc.ServerStream
}

type buildMakeServer struct {
	grpc.ServerStream
}

func (x *buildMakeServer) Send(m *BuildResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *buildMakeServer) Recv() (*BuildRequest, error) {
	m := new(BuildRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Build_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build.Build/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Build_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuildServer).Run(m, &buildRunServer{stream})
}

type Build_RunServer interface {
	Send(*RunResponse) error
	grpc.ServerStream
}

type buildRunServer struct {
	grpc.ServerStream
}

func (x *buildRunServer) Send(m *RunResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Build_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build.Build/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Build_serviceDesc = grpc.ServiceDesc{
	ServiceName: "build.Build",
	HandlerType: (*BuildServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Build_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Build_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Make",
			Handler:       _Build_Make_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Run",
			Handler:       _Build_Run_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/protobuf/build/build.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/build/build.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x1d, 0x93, 0xf4, 0xcf, 0x49, 0x0a, 0x3a, 0xed, 0x22, 0x0d, 0x22, 0x35, 0xdd, 0x64,
	0x21, 0x69, 0xad, 0xb8, 0x12, 0x8a, 0x14, 0x15, 0x05, 0xeb, 0x62, 0x5e, 0xa0, 0x4c, 0xed, 0x44,
	0x43, 0x93, 0xc9, 0xd8, 0xcc, 0x88, 0xae, 0x7d, 0x3e, 0x97, 0xbe, 0x8f, 0xcc, 0x9f, 0x6a, 0x0a,
	0xf5, 0x5e, 0xb8, 0x9b, 0x70, 0xe6, 0xcb, 0xef, 0x9c, 0x39, 0xdf, 0x47, 0x02, 0x8f, 0xc5, 0xf1,
	0xf3, 0x42, 0x9c, 0x1a, 0xd9, 0xec, 0x55, 0xb1, 0xd8, 0xab, 0xb2, 0x3a, 0xd8, 0x67, 0x6e, 0x44,
	0x1c, 0x98, 0x43, 0xfa, 0x1b, 0x41, 0xb4, 0xd1, 0x15, 0x61, 0x5f, 0x15, 0x6b, 0x25, 0xce, 0xc1,
	0x2f, 0x79, 0xd1, 0xc4, 0x68, 0x86, 0xb2, 0x70, 0x15, 0xe7, 0xb6, 0xa7, 0x8b, 0xe4, 0xef, 0x79,
	0xd1, 0xbc, 0xbb, 0x47, 0x0c, 0xa7, 0xf9, 0xa2, 0xac, 0x58, 0x7c, 0xff, 0xff, 0xfc, 0xdb, 0xb2,
	0x62, 0x9a, 0xd7, 0x5c, 0xb2, 0x06, 0x5f, 0xf7, 0xe3, 0x07, 0xe0, 0x51, 0x21, 0xcc, 0x35, 0x43,
	0xa2, 0x4b, 0x8c, 0xc1, 0xe7, 0xb4, 0xb6, 0x93, 0x86, 0xc4, 0xd4, 0x9a, 0x3a, 0x29, 0x1e, 0x7b,
	0x33, 0x94, 0x0d, 0x88, 0x2e, 0x93, 0x47, 0xe0, 0xeb, 0x79, 0x78, 0x02, 0xc1, 0xa7, 0x2f, 0x8a,
	0x1f, 0xcd, 0x84, 0x88, 0xd8, 0xc3, 0xa6, 0x0f, 0xc1, 0x37, 0x5a, 0x29, 0x96, 0xce, 0x61, 0xe4,
	0x76, 0x68, 0x45, 0xc3, 0x5b, 0xa6, 0xa7, 0x4b, 0xf6, 0x5d, 0xba, 0x0b, 0x4d, 0x9d, 0x66, 0x10,
	0x7e, 0x28, 0x5b, 0x79, 0xb6, 0x3e, 0x85, 0x01, 0x15, 0x62, 0x67, 0x96, 0xb0, 0x58, 0x9f, 0x0a,
	0xf1, 0x91, 0xd6, 0x2c, 0xfd, 0x89, 0x20, 0xb2, 0xa8, 0x1b, 0xf7, 0x0c, 0x7a, 0xc6, 0x69, 0x1b,
	0xa3, 0x99, 0x97, 0x85, 0xab, 0xa9, 0x33, 0xde, 0x85, 0x5c, 0x0a, 0x0e, 0x4c, 0x5e, 0x41, 0x60,
	0x84, 0xbf, 0x46, 0x51, 0xc7, 0xe8, 0x1c, 0x46, 0x15, 0x6d, 0xe5, 0xae, 0x6e, 0x0e, 0x65, 0x51,
	0xb2, 0x83, 0x4b, 0x21, 0xd2, 0xe2, 0xd6, 0x69, 0xe9, 0x4b, 0x00, 0xa2, 0xf8, 0xed, 0xeb, 0x5e,
	0x8b, 0x32, 0x7d, 0x02, 0xa1, 0x69, 0xbe, 0x21, 0x8f, 0x35, 0x8c, 0x5e, 0xb3, 0x8a, 0x49, 0x76,
	0xc7, 0x2b, 0xfa, 0x10, 0xbc, 0xa9, 0x85, 0xfc, 0xb1, 0xfa, 0x85, 0xce, 0x5e, 0x5f, 0x80, 0xbf,
	0xa5, 0x47, 0x86, 0xc7, 0x57, 0x3e, 0x8c, 0x64, 0x72, 0x29, 0xda, 0xcd, 0x32, 0xb4, 0x44, 0x78,
	0x01, 0xbe, 0x4e, 0x12, 0xe3, 0x8b, 0x58, 0x6d, 0xd7, 0xf8, 0x4a, 0xd4, 0x38, 0x07, 0x8f, 0x28,
	0x8e, 0x1f, 0xba, 0x77, 0xff, 0x62, 0x4a, 0x70, 0x57, 0xb2, 0xf4, 0x12, 0xe1, 0xa7, 0xd0, 0xb3,
	0x56, 0xf1, 0x79, 0x89, 0x0b, 0xe7, 0x49, 0xe4, 0x54, 0xe3, 0x67, 0xdf, 0x33, 0xff, 0xcc, 0xf3,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x06, 0xb2, 0xaf, 0x55, 0x03, 0x00, 0x00,
}
