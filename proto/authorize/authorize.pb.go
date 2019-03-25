// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authorize.proto

package authorize

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

type AuthorizeRequest struct {
	// request context
	Route string `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	// user context
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Groups               []string `protobuf:"bytes,4,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeRequest) Reset()         { *m = AuthorizeRequest{} }
func (m *AuthorizeRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizeRequest) ProtoMessage()    {}
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_authorize_dad4e29706fc340b, []int{0}
}
func (m *AuthorizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeRequest.Unmarshal(m, b)
}
func (m *AuthorizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeRequest.Marshal(b, m, deterministic)
}
func (dst *AuthorizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeRequest.Merge(dst, src)
}
func (m *AuthorizeRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizeRequest.Size(m)
}
func (m *AuthorizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeRequest proto.InternalMessageInfo

func (m *AuthorizeRequest) GetRoute() string {
	if m != nil {
		return m.Route
	}
	return ""
}

func (m *AuthorizeRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *AuthorizeRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthorizeRequest) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

type AuthorizeReply struct {
	IsValid              bool     `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeReply) Reset()         { *m = AuthorizeReply{} }
func (m *AuthorizeReply) String() string { return proto.CompactTextString(m) }
func (*AuthorizeReply) ProtoMessage()    {}
func (*AuthorizeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_authorize_dad4e29706fc340b, []int{1}
}
func (m *AuthorizeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeReply.Unmarshal(m, b)
}
func (m *AuthorizeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeReply.Marshal(b, m, deterministic)
}
func (dst *AuthorizeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeReply.Merge(dst, src)
}
func (m *AuthorizeReply) XXX_Size() int {
	return xxx_messageInfo_AuthorizeReply.Size(m)
}
func (m *AuthorizeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeReply proto.InternalMessageInfo

func (m *AuthorizeReply) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func init() {
	proto.RegisterType((*AuthorizeRequest)(nil), "authorize.AuthorizeRequest")
	proto.RegisterType((*AuthorizeReply)(nil), "authorize.AuthorizeReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizerClient is the client API for Authorizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizerClient interface {
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeReply, error)
}

type authorizerClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizerClient(cc *grpc.ClientConn) AuthorizerClient {
	return &authorizerClient{cc}
}

func (c *authorizerClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeReply, error) {
	out := new(AuthorizeReply)
	err := c.cc.Invoke(ctx, "/authorize.Authorizer/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizerServer is the server API for Authorizer service.
type AuthorizerServer interface {
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeReply, error)
}

func RegisterAuthorizerServer(s *grpc.Server, srv AuthorizerServer) {
	s.RegisterService(&_Authorizer_serviceDesc, srv)
}

func _Authorizer_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizerServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorize.Authorizer/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizerServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authorize.Authorizer",
	HandlerType: (*AuthorizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _Authorizer_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorize.proto",
}

func init() { proto.RegisterFile("authorize.proto", fileDescriptor_authorize_dad4e29706fc340b) }

var fileDescriptor_authorize_dad4e29706fc340b = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0x2d, 0xc9,
	0xc8, 0x2f, 0xca, 0xac, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x65, 0x71, 0x09, 0x38, 0xc2, 0x38, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c,
	0xac, 0x45, 0xf9, 0xa5, 0x25, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90,
	0x10, 0x17, 0x4b, 0x69, 0x71, 0x6a, 0x91, 0x04, 0x13, 0x58, 0x10, 0xcc, 0x06, 0xa9, 0x4c, 0xcd,
	0x4d, 0xcc, 0xcc, 0x91, 0x60, 0x86, 0xa8, 0x04, 0x73, 0x84, 0xc4, 0xb8, 0xd8, 0xd2, 0x8b, 0xf2,
	0x4b, 0x0b, 0x8a, 0x25, 0x58, 0x14, 0x98, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x25, 0x6d, 0x2e, 0x3e,
	0x24, 0xbb, 0x0a, 0x72, 0x2a, 0x85, 0x24, 0xb9, 0x38, 0x32, 0x8b, 0xe3, 0xcb, 0x12, 0x73, 0x32,
	0x53, 0xc0, 0x96, 0x71, 0x04, 0xb1, 0x67, 0x16, 0x87, 0x81, 0xb8, 0x46, 0xc1, 0x5c, 0x5c, 0x70,
	0xc5, 0x45, 0x42, 0xae, 0x5c, 0x9c, 0x70, 0x9e, 0x90, 0xb4, 0x1e, 0xc2, 0x43, 0xe8, 0x8e, 0x97,
	0x92, 0xc4, 0x2e, 0x59, 0x90, 0x53, 0xa9, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0xbf, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x28, 0xac, 0x76, 0x2d, 0x12, 0x01, 0x00, 0x00,
}