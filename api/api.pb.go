// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Text struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Status struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Pattern struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pattern) Reset()         { *m = Pattern{} }
func (m *Pattern) String() string { return proto.CompactTextString(m) }
func (*Pattern) ProtoMessage()    {}
func (*Pattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Pattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pattern.Unmarshal(m, b)
}
func (m *Pattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pattern.Marshal(b, m, deterministic)
}
func (m *Pattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pattern.Merge(m, src)
}
func (m *Pattern) XXX_Size() int {
	return xxx_messageInfo_Pattern.Size(m)
}
func (m *Pattern) XXX_DiscardUnknown() {
	xxx_messageInfo_Pattern.DiscardUnknown(m)
}

var xxx_messageInfo_Pattern proto.InternalMessageInfo

func (m *Pattern) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Offsets struct {
	Value                []int32  `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Offsets) Reset()         { *m = Offsets{} }
func (m *Offsets) String() string { return proto.CompactTextString(m) }
func (*Offsets) ProtoMessage()    {}
func (*Offsets) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Offsets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Offsets.Unmarshal(m, b)
}
func (m *Offsets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Offsets.Marshal(b, m, deterministic)
}
func (m *Offsets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Offsets.Merge(m, src)
}
func (m *Offsets) XXX_Size() int {
	return xxx_messageInfo_Offsets.Size(m)
}
func (m *Offsets) XXX_DiscardUnknown() {
	xxx_messageInfo_Offsets.DiscardUnknown(m)
}

var xxx_messageInfo_Offsets proto.InternalMessageInfo

func (m *Offsets) GetValue() []int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Text)(nil), "Text")
	proto.RegisterType((*Status)(nil), "Status")
	proto.RegisterType((*Pattern)(nil), "Pattern")
	proto.RegisterType((*Offsets)(nil), "Offsets")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe1, 0x62, 0x09, 0x49, 0xad, 0x28, 0x11, 0x12, 0xe1,
	0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94,
	0xe4, 0xb8, 0xd8, 0x82, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x51, 0xe5, 0x59, 0x61, 0xf2, 0xf2, 0x5c,
	0xec, 0x01, 0x89, 0x25, 0x25, 0xa9, 0x45, 0x79, 0x38, 0x0c, 0x90, 0xe7, 0x62, 0xf7, 0x4f, 0x4b,
	0x2b, 0x4e, 0x2d, 0x41, 0x31, 0x81, 0x19, 0x6e, 0x82, 0x91, 0x07, 0x17, 0x9b, 0x5b, 0x66, 0x5e,
	0x4a, 0x6a, 0x91, 0x90, 0x1c, 0x17, 0x97, 0x53, 0x69, 0x66, 0x4e, 0x8a, 0x67, 0x5e, 0x4a, 0x6a,
	0x85, 0x10, 0xab, 0x1e, 0xc8, 0x59, 0x52, 0xec, 0x7a, 0x10, 0xfb, 0x95, 0x18, 0x84, 0xe4, 0xb8,
	0xd8, 0x7c, 0xf2, 0xf3, 0xb3, 0x4b, 0x0b, 0x84, 0x38, 0xf4, 0xa0, 0x96, 0x4a, 0x71, 0xe8, 0x41,
	0x4d, 0x57, 0x62, 0x48, 0x62, 0x03, 0x7b, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xfe,
	0xd1, 0x1f, 0xdd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FinderClient is the client API for Finder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FinderClient interface {
	BuildIndex(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Status, error)
	Lookup(ctx context.Context, in *Pattern, opts ...grpc.CallOption) (*Offsets, error)
}

type finderClient struct {
	cc *grpc.ClientConn
}

func NewFinderClient(cc *grpc.ClientConn) FinderClient {
	return &finderClient{cc}
}

func (c *finderClient) BuildIndex(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/Finder/BuildIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *finderClient) Lookup(ctx context.Context, in *Pattern, opts ...grpc.CallOption) (*Offsets, error) {
	out := new(Offsets)
	err := c.cc.Invoke(ctx, "/Finder/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinderServer is the server API for Finder service.
type FinderServer interface {
	BuildIndex(context.Context, *Text) (*Status, error)
	Lookup(context.Context, *Pattern) (*Offsets, error)
}

// UnimplementedFinderServer can be embedded to have forward compatible implementations.
type UnimplementedFinderServer struct {
}

func (*UnimplementedFinderServer) BuildIndex(ctx context.Context, req *Text) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildIndex not implemented")
}
func (*UnimplementedFinderServer) Lookup(ctx context.Context, req *Pattern) (*Offsets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}

func RegisterFinderServer(s *grpc.Server, srv FinderServer) {
	s.RegisterService(&_Finder_serviceDesc, srv)
}

func _Finder_BuildIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinderServer).BuildIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Finder/BuildIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinderServer).BuildIndex(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

func _Finder_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pattern)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinderServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Finder/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinderServer).Lookup(ctx, req.(*Pattern))
	}
	return interceptor(ctx, in, info, handler)
}

var _Finder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Finder",
	HandlerType: (*FinderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildIndex",
			Handler:    _Finder_BuildIndex_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _Finder_Lookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
