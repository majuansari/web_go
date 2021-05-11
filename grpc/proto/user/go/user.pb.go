// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package userpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type UserIdRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdRequest) Reset()         { *m = UserIdRequest{} }
func (m *UserIdRequest) String() string { return proto.CompactTextString(m) }
func (*UserIdRequest) ProtoMessage()    {}
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdRequest.Unmarshal(m, b)
}
func (m *UserIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdRequest.Marshal(b, m, deterministic)
}
func (m *UserIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdRequest.Merge(m, src)
}
func (m *UserIdRequest) XXX_Size() int {
	return xxx_messageInfo_UserIdRequest.Size(m)
}
func (m *UserIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdRequest proto.InternalMessageInfo

func (m *UserIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserResponse) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func init() {
	proto.RegisterType((*UserIdRequest)(nil), "user.UserIdRequest")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xe4, 0xb9, 0x78, 0x43, 0x8b,
	0x53, 0x8b, 0x3c, 0x53, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32,
	0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x94, 0x7c, 0xb8, 0x78, 0x40,
	0x0a, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0xa1, 0x2a, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0x82, 0x8c, 0xfc, 0xbc, 0x54, 0x09, 0x26,
	0xb0, 0x20, 0x84, 0x03, 0x52, 0x99, 0x9c, 0x59, 0x52, 0x29, 0xc1, 0x0c, 0x51, 0x09, 0x62, 0x1b,
	0x85, 0x72, 0x09, 0x81, 0xac, 0x75, 0x49, 0x2d, 0x49, 0xcc, 0xcc, 0x29, 0x0e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x15, 0xb2, 0xc7, 0x2a, 0x2a, 0xac, 0x07, 0x76, 0x2d, 0x8a, 0xf3, 0xa4, 0x84,
	0x10, 0x82, 0x30, 0x27, 0x29, 0x31, 0x38, 0x71, 0x44, 0xb1, 0x81, 0x84, 0x0b, 0x92, 0x92, 0xd8,
	0xc0, 0x9e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x17, 0xfb, 0xc8, 0xea, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserDetailsServiceClient is the client API for UserDetailsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserDetailsServiceClient interface {
	UserDetailsService(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userDetailsServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserDetailsServiceClient(cc *grpc.ClientConn) UserDetailsServiceClient {
	return &userDetailsServiceClient{cc}
}

func (c *userDetailsServiceClient) UserDetailsService(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.userDetailsService/userDetailsService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDetailsServiceServer is the server API for UserDetailsService service.
type UserDetailsServiceServer interface {
	UserDetailsService(context.Context, *UserIdRequest) (*UserResponse, error)
}

func RegisterUserDetailsServiceServer(s *grpc.Server, srv UserDetailsServiceServer) {
	s.RegisterService(&_UserDetailsService_serviceDesc, srv)
}

func _UserDetailsService_UserDetailsService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDetailsServiceServer).UserDetailsService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userDetailsService/UserDetailsService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDetailsServiceServer).UserDetailsService(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserDetailsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.userDetailsService",
	HandlerType: (*UserDetailsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "userDetailsService",
			Handler:    _UserDetailsService_UserDetailsService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}