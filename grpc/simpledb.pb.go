// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simpledb.proto

package simpledb

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type HostStatsReplyMsg struct {
	TotalMemory          string   `protobuf:"bytes,1,opt,name=totalMemory,proto3" json:"totalMemory,omitempty"`
	UsedMemory           string   `protobuf:"bytes,2,opt,name=usedMemory,proto3" json:"usedMemory,omitempty"`
	TotalSpace           string   `protobuf:"bytes,3,opt,name=totalSpace,proto3" json:"totalSpace,omitempty"`
	UsedSpace            string   `protobuf:"bytes,4,opt,name=usedSpace,proto3" json:"usedSpace,omitempty"`
	NumCores             string   `protobuf:"bytes,5,opt,name=numCores,proto3" json:"numCores,omitempty"`
	CpuPercent           string   `protobuf:"bytes,6,opt,name=cpuPercent,proto3" json:"cpuPercent,omitempty"`
	Os                   string   `protobuf:"bytes,7,opt,name=os,proto3" json:"os,omitempty"`
	Hostname             string   `protobuf:"bytes,8,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Uptime               string   `protobuf:"bytes,9,opt,name=uptime,proto3" json:"uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostStatsReplyMsg) Reset()         { *m = HostStatsReplyMsg{} }
func (m *HostStatsReplyMsg) String() string { return proto.CompactTextString(m) }
func (*HostStatsReplyMsg) ProtoMessage()    {}
func (*HostStatsReplyMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{1}
}

func (m *HostStatsReplyMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostStatsReplyMsg.Unmarshal(m, b)
}
func (m *HostStatsReplyMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostStatsReplyMsg.Marshal(b, m, deterministic)
}
func (m *HostStatsReplyMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatsReplyMsg.Merge(m, src)
}
func (m *HostStatsReplyMsg) XXX_Size() int {
	return xxx_messageInfo_HostStatsReplyMsg.Size(m)
}
func (m *HostStatsReplyMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatsReplyMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatsReplyMsg proto.InternalMessageInfo

func (m *HostStatsReplyMsg) GetTotalMemory() string {
	if m != nil {
		return m.TotalMemory
	}
	return ""
}

func (m *HostStatsReplyMsg) GetUsedMemory() string {
	if m != nil {
		return m.UsedMemory
	}
	return ""
}

func (m *HostStatsReplyMsg) GetTotalSpace() string {
	if m != nil {
		return m.TotalSpace
	}
	return ""
}

func (m *HostStatsReplyMsg) GetUsedSpace() string {
	if m != nil {
		return m.UsedSpace
	}
	return ""
}

func (m *HostStatsReplyMsg) GetNumCores() string {
	if m != nil {
		return m.NumCores
	}
	return ""
}

func (m *HostStatsReplyMsg) GetCpuPercent() string {
	if m != nil {
		return m.CpuPercent
	}
	return ""
}

func (m *HostStatsReplyMsg) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *HostStatsReplyMsg) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HostStatsReplyMsg) GetUptime() string {
	if m != nil {
		return m.Uptime
	}
	return ""
}

type RpcOkayMsg struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcOkayMsg) Reset()         { *m = RpcOkayMsg{} }
func (m *RpcOkayMsg) String() string { return proto.CompactTextString(m) }
func (*RpcOkayMsg) ProtoMessage()    {}
func (*RpcOkayMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{2}
}

func (m *RpcOkayMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcOkayMsg.Unmarshal(m, b)
}
func (m *RpcOkayMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcOkayMsg.Marshal(b, m, deterministic)
}
func (m *RpcOkayMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcOkayMsg.Merge(m, src)
}
func (m *RpcOkayMsg) XXX_Size() int {
	return xxx_messageInfo_RpcOkayMsg.Size(m)
}
func (m *RpcOkayMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcOkayMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RpcOkayMsg proto.InternalMessageInfo

func (m *RpcOkayMsg) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "simpledb.Empty")
	proto.RegisterType((*HostStatsReplyMsg)(nil), "simpledb.HostStatsReplyMsg")
	proto.RegisterType((*RpcOkayMsg)(nil), "simpledb.RpcOkayMsg")
}

func init() { proto.RegisterFile("simpledb.proto", fileDescriptor_748391160b9263c4) }

var fileDescriptor_748391160b9263c4 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xe9, 0xfc, 0x7f, 0xdb, 0xf4, 0x0a, 0x91, 0xce, 0x42, 0x86, 0x5a, 0xa4, 0x64, 0xe5,
	0xaa, 0x0b, 0x7d, 0x02, 0xa9, 0xa2, 0x2e, 0x42, 0x24, 0x79, 0x82, 0x49, 0x7a, 0xd5, 0x92, 0x4c,
	0x66, 0xc8, 0xdc, 0x2c, 0xf2, 0x1c, 0xbe, 0xb0, 0xe4, 0xa6, 0x6d, 0x0a, 0x2e, 0xef, 0x77, 0x0e,
	0x67, 0xf1, 0x5d, 0x08, 0xfd, 0xc1, 0xb8, 0x0a, 0xf7, 0xf9, 0xd6, 0x35, 0x96, 0xac, 0x0c, 0x4e,
	0x77, 0x34, 0x87, 0xe9, 0x8b, 0x71, 0xd4, 0x45, 0x3f, 0x02, 0x96, 0x6f, 0xd6, 0x53, 0x46, 0x9a,
	0x7c, 0x8a, 0xae, 0xea, 0x62, 0xff, 0x25, 0x37, 0x70, 0x45, 0x96, 0x74, 0x15, 0xa3, 0xb1, 0x4d,
	0xa7, 0x26, 0x9b, 0xc9, 0xfd, 0x22, 0xbd, 0x44, 0xf2, 0x0e, 0xa0, 0xf5, 0xb8, 0x3f, 0x16, 0x04,
	0x17, 0x2e, 0x48, 0x9f, 0x73, 0x3d, 0x73, 0xba, 0x40, 0xf5, 0x6f, 0xc8, 0x47, 0x22, 0xd7, 0xb0,
	0xe8, 0xdb, 0x43, 0xfc, 0x9f, 0xe3, 0x11, 0xc8, 0x15, 0x04, 0x75, 0x6b, 0x76, 0xb6, 0x41, 0xaf,
	0xa6, 0x1c, 0x9e, 0xef, 0x7e, 0xb9, 0x70, 0xed, 0x07, 0x36, 0x05, 0xd6, 0xa4, 0x66, 0xc3, 0xf2,
	0x48, 0x64, 0x08, 0xc2, 0x7a, 0x35, 0x67, 0x2e, 0xac, 0xef, 0xb7, 0xbe, 0xad, 0xa7, 0x5a, 0x1b,
	0x54, 0xc1, 0xb0, 0x75, 0xba, 0xe5, 0x0d, 0xcc, 0x5a, 0x47, 0x07, 0x83, 0x6a, 0xc1, 0xc9, 0xf1,
	0x8a, 0xd6, 0x00, 0xa9, 0x2b, 0x92, 0x52, 0xb3, 0x8d, 0x10, 0x44, 0x52, 0xb2, 0x84, 0x20, 0x15,
	0x49, 0xf9, 0x10, 0x43, 0x90, 0xb1, 0xc8, 0xe7, 0x5c, 0x3e, 0xc1, 0xf2, 0x15, 0xa9, 0x37, 0xf8,
	0x5e, 0x7f, 0xda, 0x9d, 0xae, 0x2a, 0x6c, 0xe4, 0xf5, 0xf6, 0x2c, 0x9e, 0x2d, 0xaf, 0x6e, 0x47,
	0xf0, 0x47, 0x76, 0x3e, 0xe3, 0xe7, 0x3c, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x33, 0x6d, 0x8a,
	0xc2, 0xae, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimpleDbClient is the client API for SimpleDb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleDbClient interface {
	GetHostInfoCaller(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HostStatsReplyMsg, error)
}

type simpleDbClient struct {
	cc *grpc.ClientConn
}

func NewSimpleDbClient(cc *grpc.ClientConn) SimpleDbClient {
	return &simpleDbClient{cc}
}

func (c *simpleDbClient) GetHostInfoCaller(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HostStatsReplyMsg, error) {
	out := new(HostStatsReplyMsg)
	err := c.cc.Invoke(ctx, "/simpledb.SimpleDb/GetHostInfoCaller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleDbServer is the server API for SimpleDb service.
type SimpleDbServer interface {
	GetHostInfoCaller(context.Context, *Empty) (*HostStatsReplyMsg, error)
}

// UnimplementedSimpleDbServer can be embedded to have forward compatible implementations.
type UnimplementedSimpleDbServer struct {
}

func (*UnimplementedSimpleDbServer) GetHostInfoCaller(ctx context.Context, req *Empty) (*HostStatsReplyMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostInfoCaller not implemented")
}

func RegisterSimpleDbServer(s *grpc.Server, srv SimpleDbServer) {
	s.RegisterService(&_SimpleDb_serviceDesc, srv)
}

func _SimpleDb_GetHostInfoCaller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleDbServer).GetHostInfoCaller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpledb.SimpleDb/GetHostInfoCaller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleDbServer).GetHostInfoCaller(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimpleDb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simpledb.SimpleDb",
	HandlerType: (*SimpleDbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHostInfoCaller",
			Handler:    _SimpleDb_GetHostInfoCaller_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simpledb.proto",
}