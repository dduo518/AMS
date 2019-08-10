// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sms.proto

// 定义包名

package probuf

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

type SendCodeReq struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendCodeReq) Reset()         { *m = SendCodeReq{} }
func (m *SendCodeReq) String() string { return proto.CompactTextString(m) }
func (*SendCodeReq) ProtoMessage()    {}
func (*SendCodeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{0}
}

func (m *SendCodeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendCodeReq.Unmarshal(m, b)
}
func (m *SendCodeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendCodeReq.Marshal(b, m, deterministic)
}
func (m *SendCodeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCodeReq.Merge(m, src)
}
func (m *SendCodeReq) XXX_Size() int {
	return xxx_messageInfo_SendCodeReq.Size(m)
}
func (m *SendCodeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCodeReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendCodeReq proto.InternalMessageInfo

func (m *SendCodeReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type SendCodeRes struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendCodeRes) Reset()         { *m = SendCodeRes{} }
func (m *SendCodeRes) String() string { return proto.CompactTextString(m) }
func (*SendCodeRes) ProtoMessage()    {}
func (*SendCodeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{1}
}

func (m *SendCodeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendCodeRes.Unmarshal(m, b)
}
func (m *SendCodeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendCodeRes.Marshal(b, m, deterministic)
}
func (m *SendCodeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCodeRes.Merge(m, src)
}
func (m *SendCodeRes) XXX_Size() int {
	return xxx_messageInfo_SendCodeRes.Size(m)
}
func (m *SendCodeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCodeRes.DiscardUnknown(m)
}

var xxx_messageInfo_SendCodeRes proto.InternalMessageInfo

func (m *SendCodeRes) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type VerifyCodeReq struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyCodeReq) Reset()         { *m = VerifyCodeReq{} }
func (m *VerifyCodeReq) String() string { return proto.CompactTextString(m) }
func (*VerifyCodeReq) ProtoMessage()    {}
func (*VerifyCodeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{2}
}

func (m *VerifyCodeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyCodeReq.Unmarshal(m, b)
}
func (m *VerifyCodeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyCodeReq.Marshal(b, m, deterministic)
}
func (m *VerifyCodeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyCodeReq.Merge(m, src)
}
func (m *VerifyCodeReq) XXX_Size() int {
	return xxx_messageInfo_VerifyCodeReq.Size(m)
}
func (m *VerifyCodeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyCodeReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyCodeReq proto.InternalMessageInfo

func (m *VerifyCodeReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *VerifyCodeReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type VerifyCodeRes struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyCodeRes) Reset()         { *m = VerifyCodeRes{} }
func (m *VerifyCodeRes) String() string { return proto.CompactTextString(m) }
func (*VerifyCodeRes) ProtoMessage()    {}
func (*VerifyCodeRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{3}
}

func (m *VerifyCodeRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyCodeRes.Unmarshal(m, b)
}
func (m *VerifyCodeRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyCodeRes.Marshal(b, m, deterministic)
}
func (m *VerifyCodeRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyCodeRes.Merge(m, src)
}
func (m *VerifyCodeRes) XXX_Size() int {
	return xxx_messageInfo_VerifyCodeRes.Size(m)
}
func (m *VerifyCodeRes) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyCodeRes.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyCodeRes proto.InternalMessageInfo

func (m *VerifyCodeRes) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *VerifyCodeRes) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *VerifyCodeRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*SendCodeReq)(nil), "probuf.SendCodeReq")
	proto.RegisterType((*SendCodeRes)(nil), "probuf.SendCodeRes")
	proto.RegisterType((*VerifyCodeReq)(nil), "probuf.VerifyCodeReq")
	proto.RegisterType((*VerifyCodeRes)(nil), "probuf.VerifyCodeRes")
}

func init() { proto.RegisterFile("sms.proto", fileDescriptor_c8d8bdc537111860) }

var fileDescriptor_c8d8bdc537111860 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xce, 0x2d, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x28, 0xca, 0x4f, 0x2a, 0x4d, 0x53, 0x52, 0xe6,
	0xe2, 0x0e, 0x4e, 0xcd, 0x4b, 0x71, 0xce, 0x4f, 0x49, 0x0d, 0x4a, 0x2d, 0x14, 0x12, 0xe1, 0x62,
	0x2d, 0xc8, 0xc8, 0xcf, 0x4b, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x54,
	0x91, 0x15, 0x15, 0x0b, 0x89, 0x71, 0xb1, 0x95, 0x16, 0xa7, 0x16, 0x79, 0xa6, 0x40, 0x55, 0x41,
	0x79, 0x4a, 0x96, 0x5c, 0xbc, 0x61, 0xa9, 0x45, 0x99, 0x69, 0x95, 0x78, 0x4d, 0x13, 0x12, 0xe2,
	0x62, 0x49, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x02, 0x0b, 0x82, 0xd9, 0x4a, 0xb1, 0xa8, 0x5a, 0x8b,
	0x85, 0x14, 0xb8, 0xb8, 0x13, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x43, 0xf2, 0xb3, 0x53, 0xf3, 0xa0,
	0x06, 0x20, 0x0b, 0x21, 0xb9, 0x82, 0x09, 0xd9, 0x15, 0x20, 0xe3, 0xf3, 0x12, 0x73, 0x53, 0x25,
	0x98, 0x21, 0xc6, 0x83, 0xd8, 0x46, 0xd5, 0x5c, 0xcc, 0xc1, 0xb9, 0xc5, 0x42, 0x66, 0x5c, 0x1c,
	0x30, 0x7f, 0x08, 0x09, 0xeb, 0x41, 0x42, 0x40, 0x0f, 0xc9, 0xfb, 0x52, 0x58, 0x04, 0x8b, 0x95,
	0x18, 0x84, 0x6c, 0xb8, 0xb8, 0x10, 0xae, 0x13, 0x12, 0x85, 0x29, 0x42, 0xf1, 0xac, 0x14, 0x56,
	0xe1, 0x62, 0x25, 0x86, 0x24, 0x36, 0x70, 0x88, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0,
	0xb8, 0x92, 0xf5, 0x7e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SmsClient is the client API for Sms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmsClient interface {
	SendCode(ctx context.Context, in *SendCodeReq, opts ...grpc.CallOption) (*SendCodeRes, error)
	VerifyCode(ctx context.Context, in *VerifyCodeReq, opts ...grpc.CallOption) (*VerifyCodeRes, error)
}

type smsClient struct {
	cc *grpc.ClientConn
}

func NewSmsClient(cc *grpc.ClientConn) SmsClient {
	return &smsClient{cc}
}

func (c *smsClient) SendCode(ctx context.Context, in *SendCodeReq, opts ...grpc.CallOption) (*SendCodeRes, error) {
	out := new(SendCodeRes)
	err := c.cc.Invoke(ctx, "/probuf.Sms/SendCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsClient) VerifyCode(ctx context.Context, in *VerifyCodeReq, opts ...grpc.CallOption) (*VerifyCodeRes, error) {
	out := new(VerifyCodeRes)
	err := c.cc.Invoke(ctx, "/probuf.Sms/VerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServer is the server API for Sms service.
type SmsServer interface {
	SendCode(context.Context, *SendCodeReq) (*SendCodeRes, error)
	VerifyCode(context.Context, *VerifyCodeReq) (*VerifyCodeRes, error)
}

// UnimplementedSmsServer can be embedded to have forward compatible implementations.
type UnimplementedSmsServer struct {
}

func (*UnimplementedSmsServer) SendCode(ctx context.Context, req *SendCodeReq) (*SendCodeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCode not implemented")
}
func (*UnimplementedSmsServer) VerifyCode(ctx context.Context, req *VerifyCodeReq) (*VerifyCodeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCode not implemented")
}

func RegisterSmsServer(s *grpc.Server, srv SmsServer) {
	s.RegisterService(&_Sms_serviceDesc, srv)
}

func _Sms_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/probuf.Sms/SendCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).SendCode(ctx, req.(*SendCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sms_VerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).VerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/probuf.Sms/VerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).VerifyCode(ctx, req.(*VerifyCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sms_serviceDesc = grpc.ServiceDesc{
	ServiceName: "probuf.Sms",
	HandlerType: (*SmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCode",
			Handler:    _Sms_SendCode_Handler,
		},
		{
			MethodName: "VerifyCode",
			Handler:    _Sms_VerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms.proto",
}