// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simpledb.proto

package simpledb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Attribute_Type int32

const (
	Attribute_BOOL   Attribute_Type = 0
	Attribute_INT    Attribute_Type = 1
	Attribute_UINT   Attribute_Type = 2
	Attribute_FLOAT  Attribute_Type = 3
	Attribute_STRING Attribute_Type = 4
	Attribute_BYTES  Attribute_Type = 5
)

var Attribute_Type_name = map[int32]string{
	0: "BOOL",
	1: "INT",
	2: "UINT",
	3: "FLOAT",
	4: "STRING",
	5: "BYTES",
}

var Attribute_Type_value = map[string]int32{
	"BOOL":   0,
	"INT":    1,
	"UINT":   2,
	"FLOAT":  3,
	"STRING": 4,
	"BYTES":  5,
}

func (x Attribute_Type) String() string {
	return proto.EnumName(Attribute_Type_name, int32(x))
}

func (Attribute_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{3, 0}
}

type ReadMsg struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Attributes           []string `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMsg) Reset()         { *m = ReadMsg{} }
func (m *ReadMsg) String() string { return proto.CompactTextString(m) }
func (*ReadMsg) ProtoMessage()    {}
func (*ReadMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{0}
}

func (m *ReadMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMsg.Unmarshal(m, b)
}
func (m *ReadMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMsg.Marshal(b, m, deterministic)
}
func (m *ReadMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMsg.Merge(m, src)
}
func (m *ReadMsg) XXX_Size() int {
	return xxx_messageInfo_ReadMsg.Size(m)
}
func (m *ReadMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMsg proto.InternalMessageInfo

func (m *ReadMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReadMsg) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ScanMsg struct {
	StartKey             string   `protobuf:"bytes,1,opt,name=startKey,proto3" json:"startKey,omitempty"`
	EndKey               string   `protobuf:"bytes,2,opt,name=endKey,proto3" json:"endKey,omitempty"`
	Attributes           []string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanMsg) Reset()         { *m = ScanMsg{} }
func (m *ScanMsg) String() string { return proto.CompactTextString(m) }
func (*ScanMsg) ProtoMessage()    {}
func (*ScanMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{1}
}

func (m *ScanMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanMsg.Unmarshal(m, b)
}
func (m *ScanMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanMsg.Marshal(b, m, deterministic)
}
func (m *ScanMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanMsg.Merge(m, src)
}
func (m *ScanMsg) XXX_Size() int {
	return xxx_messageInfo_ScanMsg.Size(m)
}
func (m *ScanMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ScanMsg proto.InternalMessageInfo

func (m *ScanMsg) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *ScanMsg) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

func (m *ScanMsg) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type EntriesMsg struct {
	Entries              []*Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntriesMsg) Reset()         { *m = EntriesMsg{} }
func (m *EntriesMsg) String() string { return proto.CompactTextString(m) }
func (*EntriesMsg) ProtoMessage()    {}
func (*EntriesMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{2}
}

func (m *EntriesMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntriesMsg.Unmarshal(m, b)
}
func (m *EntriesMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntriesMsg.Marshal(b, m, deterministic)
}
func (m *EntriesMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntriesMsg.Merge(m, src)
}
func (m *EntriesMsg) XXX_Size() int {
	return xxx_messageInfo_EntriesMsg.Size(m)
}
func (m *EntriesMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EntriesMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EntriesMsg proto.InternalMessageInfo

func (m *EntriesMsg) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Attribute struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 Attribute_Type `protobuf:"varint,2,opt,name=type,proto3,enum=simpledb.Attribute_Type" json:"type,omitempty"`
	Value                []byte         `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Attribute) Reset()         { *m = Attribute{} }
func (m *Attribute) String() string { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()    {}
func (*Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{3}
}

func (m *Attribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attribute.Unmarshal(m, b)
}
func (m *Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attribute.Marshal(b, m, deterministic)
}
func (m *Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute.Merge(m, src)
}
func (m *Attribute) XXX_Size() int {
	return xxx_messageInfo_Attribute.Size(m)
}
func (m *Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute proto.InternalMessageInfo

func (m *Attribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Attribute) GetType() Attribute_Type {
	if m != nil {
		return m.Type
	}
	return Attribute_BOOL
}

func (m *Attribute) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Entry struct {
	Key                  string       `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Attributes           []*Attribute `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{4}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Entry) GetAttributes() []*Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type OkMsg struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OkMsg) Reset()         { *m = OkMsg{} }
func (m *OkMsg) String() string { return proto.CompactTextString(m) }
func (*OkMsg) ProtoMessage()    {}
func (*OkMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{5}
}

func (m *OkMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OkMsg.Unmarshal(m, b)
}
func (m *OkMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OkMsg.Marshal(b, m, deterministic)
}
func (m *OkMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OkMsg.Merge(m, src)
}
func (m *OkMsg) XXX_Size() int {
	return xxx_messageInfo_OkMsg.Size(m)
}
func (m *OkMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_OkMsg.DiscardUnknown(m)
}

var xxx_messageInfo_OkMsg proto.InternalMessageInfo

func (m *OkMsg) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type KeyMsg struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyMsg) Reset()         { *m = KeyMsg{} }
func (m *KeyMsg) String() string { return proto.CompactTextString(m) }
func (*KeyMsg) ProtoMessage()    {}
func (*KeyMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_748391160b9263c4, []int{6}
}

func (m *KeyMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyMsg.Unmarshal(m, b)
}
func (m *KeyMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyMsg.Marshal(b, m, deterministic)
}
func (m *KeyMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyMsg.Merge(m, src)
}
func (m *KeyMsg) XXX_Size() int {
	return xxx_messageInfo_KeyMsg.Size(m)
}
func (m *KeyMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyMsg.DiscardUnknown(m)
}

var xxx_messageInfo_KeyMsg proto.InternalMessageInfo

func (m *KeyMsg) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterEnum("simpledb.Attribute_Type", Attribute_Type_name, Attribute_Type_value)
	proto.RegisterType((*ReadMsg)(nil), "simpledb.ReadMsg")
	proto.RegisterType((*ScanMsg)(nil), "simpledb.ScanMsg")
	proto.RegisterType((*EntriesMsg)(nil), "simpledb.EntriesMsg")
	proto.RegisterType((*Attribute)(nil), "simpledb.Attribute")
	proto.RegisterType((*Entry)(nil), "simpledb.Entry")
	proto.RegisterType((*OkMsg)(nil), "simpledb.OkMsg")
	proto.RegisterType((*KeyMsg)(nil), "simpledb.KeyMsg")
}

func init() { proto.RegisterFile("simpledb.proto", fileDescriptor_748391160b9263c4) }

var fileDescriptor_748391160b9263c4 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x8f, 0x93, 0x50,
	0x10, 0x96, 0xdf, 0x30, 0x6b, 0x2a, 0x3e, 0x37, 0x4a, 0x38, 0x18, 0xc2, 0x09, 0x13, 0xcb, 0x81,
	0x3d, 0x78, 0xf0, 0xb4, 0xeb, 0x56, 0xd3, 0xec, 0x5a, 0xcc, 0x83, 0x3d, 0x78, 0xf0, 0x00, 0x32,
	0x31, 0x0d, 0x5d, 0x96, 0xc0, 0xab, 0x09, 0xff, 0x92, 0xff, 0xa2, 0x17, 0xf3, 0x1e, 0x94, 0xb6,
	0xb4, 0x07, 0x6f, 0xf3, 0xe3, 0x9b, 0x6f, 0x66, 0xbe, 0x37, 0x0f, 0x66, 0xed, 0xfa, 0xb1, 0xde,
	0x60, 0x91, 0x87, 0x75, 0xf3, 0xc4, 0x9e, 0x88, 0xb9, 0xf3, 0xfd, 0x8f, 0x60, 0x50, 0xcc, 0x8a,
	0xaf, 0xed, 0x2f, 0x62, 0x83, 0x52, 0x62, 0xe7, 0x48, 0x9e, 0x14, 0x58, 0x94, 0x9b, 0xe4, 0x2d,
	0x40, 0xc6, 0x58, 0xb3, 0xce, 0xb7, 0x0c, 0x5b, 0x47, 0xf6, 0x94, 0xc0, 0xa2, 0x07, 0x11, 0xff,
	0x07, 0x18, 0xc9, 0xcf, 0xac, 0xe2, 0xc5, 0x2e, 0x98, 0x2d, 0xcb, 0x1a, 0x76, 0x37, 0x32, 0x8c,
	0x3e, 0x79, 0x0d, 0x3a, 0x56, 0x05, 0xcf, 0xc8, 0x22, 0x33, 0x78, 0x13, 0x7a, 0xe5, 0x84, 0xfe,
	0x03, 0xc0, 0xa2, 0x62, 0xcd, 0x1a, 0x5b, 0xde, 0xe1, 0x1d, 0x18, 0xd8, 0x7b, 0x8e, 0xe4, 0x29,
	0xc1, 0x45, 0xf4, 0x22, 0x1c, 0xb7, 0xe2, 0xb0, 0x8e, 0xee, 0xf2, 0xfe, 0x1f, 0x09, 0xac, 0xeb,
	0x1d, 0x0f, 0x21, 0xa0, 0x56, 0xd9, 0x23, 0x0e, 0x63, 0x09, 0x9b, 0xbc, 0x07, 0x95, 0x75, 0x35,
	0x8a, 0x81, 0x66, 0x91, 0xb3, 0x67, 0x1a, 0xcb, 0xc2, 0xb4, 0xab, 0x91, 0x0a, 0x14, 0xb9, 0x04,
	0xed, 0x77, 0xb6, 0xd9, 0xa2, 0xa3, 0x78, 0x52, 0xf0, 0x9c, 0xf6, 0x8e, 0xbf, 0x00, 0x95, 0x63,
	0x88, 0x09, 0xea, 0x4d, 0x1c, 0xdf, 0xdb, 0xcf, 0x88, 0x01, 0xca, 0x72, 0x95, 0xda, 0x12, 0x0f,
	0x3d, 0x70, 0x4b, 0x26, 0x16, 0x68, 0x9f, 0xef, 0xe3, 0xeb, 0xd4, 0x56, 0x08, 0x80, 0x9e, 0xa4,
	0x74, 0xb9, 0xfa, 0x62, 0xab, 0x3c, 0x7c, 0xf3, 0x3d, 0x5d, 0x24, 0xb6, 0xe6, 0xaf, 0x40, 0x13,
	0xe3, 0x9f, 0xd1, 0xff, 0xea, 0x44, 0xff, 0x8b, 0xe8, 0xd5, 0x99, 0x59, 0x8f, 0x54, 0x7b, 0x03,
	0x5a, 0x5c, 0x72, 0xc1, 0x66, 0x20, 0xc7, 0xa5, 0xa0, 0x33, 0xa9, 0x1c, 0x97, 0xbe, 0x0b, 0xfa,
	0x1d, 0x76, 0x67, 0x5f, 0x3a, 0xfa, 0x2b, 0x81, 0x99, 0x08, 0xde, 0xdb, 0x9c, 0xcc, 0xfb, 0x9b,
	0xa0, 0xdf, 0x3e, 0x91, 0x97, 0xfb, 0x6e, 0xc3, 0x99, 0xb8, 0x53, 0xd9, 0x49, 0xd4, 0x5f, 0xc1,
	0x04, 0x3e, 0x1c, 0x86, 0x7b, 0x79, 0x0c, 0x1f, 0x1e, 0x73, 0x0e, 0xd6, 0x43, 0x5d, 0x64, 0x0c,
	0x79, 0xd5, 0x94, 0xf1, 0xb0, 0x45, 0xbf, 0xca, 0x1c, 0xac, 0x65, 0xd5, 0x62, 0xc3, 0xfe, 0x0f,
	0x1e, 0x82, 0x75, 0x8b, 0x1b, 0xec, 0xd9, 0xed, 0x7d, 0xb6, 0x5f, 0xff, 0x04, 0x9f, 0xeb, 0xe2,
	0x57, 0x5c, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x51, 0x1e, 0x74, 0x87, 0x27, 0x03, 0x00, 0x00,
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
	ReadRPC(ctx context.Context, in *ReadMsg, opts ...grpc.CallOption) (*Entry, error)
	ScanRPC(ctx context.Context, in *ScanMsg, opts ...grpc.CallOption) (*EntriesMsg, error)
	UpdateRPC(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*OkMsg, error)
	InsertRPC(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*OkMsg, error)
	DeleteRPC(ctx context.Context, in *KeyMsg, opts ...grpc.CallOption) (*OkMsg, error)
}

type simpleDbClient struct {
	cc *grpc.ClientConn
}

func NewSimpleDbClient(cc *grpc.ClientConn) SimpleDbClient {
	return &simpleDbClient{cc}
}

func (c *simpleDbClient) ReadRPC(ctx context.Context, in *ReadMsg, opts ...grpc.CallOption) (*Entry, error) {
	out := new(Entry)
	err := c.cc.Invoke(ctx, "/simpledb.SimpleDb/ReadRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleDbClient) ScanRPC(ctx context.Context, in *ScanMsg, opts ...grpc.CallOption) (*EntriesMsg, error) {
	out := new(EntriesMsg)
	err := c.cc.Invoke(ctx, "/simpledb.SimpleDb/ScanRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleDbClient) UpdateRPC(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*OkMsg, error) {
	out := new(OkMsg)
	err := c.cc.Invoke(ctx, "/simpledb.SimpleDb/UpdateRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleDbClient) InsertRPC(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*OkMsg, error) {
	out := new(OkMsg)
	err := c.cc.Invoke(ctx, "/simpledb.SimpleDb/InsertRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleDbClient) DeleteRPC(ctx context.Context, in *KeyMsg, opts ...grpc.CallOption) (*OkMsg, error) {
	out := new(OkMsg)
	err := c.cc.Invoke(ctx, "/simpledb.SimpleDb/DeleteRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleDbServer is the server API for SimpleDb service.
type SimpleDbServer interface {
	ReadRPC(context.Context, *ReadMsg) (*Entry, error)
	ScanRPC(context.Context, *ScanMsg) (*EntriesMsg, error)
	UpdateRPC(context.Context, *Entry) (*OkMsg, error)
	InsertRPC(context.Context, *Entry) (*OkMsg, error)
	DeleteRPC(context.Context, *KeyMsg) (*OkMsg, error)
}

// UnimplementedSimpleDbServer can be embedded to have forward compatible implementations.
type UnimplementedSimpleDbServer struct {
}

func (*UnimplementedSimpleDbServer) ReadRPC(ctx context.Context, req *ReadMsg) (*Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRPC not implemented")
}
func (*UnimplementedSimpleDbServer) ScanRPC(ctx context.Context, req *ScanMsg) (*EntriesMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanRPC not implemented")
}
func (*UnimplementedSimpleDbServer) UpdateRPC(ctx context.Context, req *Entry) (*OkMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRPC not implemented")
}
func (*UnimplementedSimpleDbServer) InsertRPC(ctx context.Context, req *Entry) (*OkMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertRPC not implemented")
}
func (*UnimplementedSimpleDbServer) DeleteRPC(ctx context.Context, req *KeyMsg) (*OkMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRPC not implemented")
}

func RegisterSimpleDbServer(s *grpc.Server, srv SimpleDbServer) {
	s.RegisterService(&_SimpleDb_serviceDesc, srv)
}

func _SimpleDb_ReadRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleDbServer).ReadRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpledb.SimpleDb/ReadRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleDbServer).ReadRPC(ctx, req.(*ReadMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleDb_ScanRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleDbServer).ScanRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpledb.SimpleDb/ScanRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleDbServer).ScanRPC(ctx, req.(*ScanMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleDb_UpdateRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleDbServer).UpdateRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpledb.SimpleDb/UpdateRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleDbServer).UpdateRPC(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleDb_InsertRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleDbServer).InsertRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpledb.SimpleDb/InsertRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleDbServer).InsertRPC(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleDb_DeleteRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleDbServer).DeleteRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpledb.SimpleDb/DeleteRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleDbServer).DeleteRPC(ctx, req.(*KeyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimpleDb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simpledb.SimpleDb",
	HandlerType: (*SimpleDbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadRPC",
			Handler:    _SimpleDb_ReadRPC_Handler,
		},
		{
			MethodName: "ScanRPC",
			Handler:    _SimpleDb_ScanRPC_Handler,
		},
		{
			MethodName: "UpdateRPC",
			Handler:    _SimpleDb_UpdateRPC_Handler,
		},
		{
			MethodName: "InsertRPC",
			Handler:    _SimpleDb_InsertRPC_Handler,
		},
		{
			MethodName: "DeleteRPC",
			Handler:    _SimpleDb_DeleteRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simpledb.proto",
}
