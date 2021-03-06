// Code generated by protoc-gen-go. DO NOT EDIT.
// source: struct.proto

package common

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_struct_14940efc5789c3fe, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// UserManager相关///////////////////////////////////////////////////////
type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Company              string   `protobuf:"bytes,2,opt,name=Company,proto3" json:"Company,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_struct_14940efc5789c3fe, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

type UserList struct {
	List                 []*User  `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_struct_14940efc5789c3fe, []int{2}
}
func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (dst *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(dst, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetList() []*User {
	if m != nil {
		return m.List
	}
	return nil
}

type DispatchRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Company              string   `protobuf:"bytes,2,opt,name=Company,proto3" json:"Company,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DispatchRequest) Reset()         { *m = DispatchRequest{} }
func (m *DispatchRequest) String() string { return proto.CompactTextString(m) }
func (*DispatchRequest) ProtoMessage()    {}
func (*DispatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_struct_14940efc5789c3fe, []int{3}
}
func (m *DispatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DispatchRequest.Unmarshal(m, b)
}
func (m *DispatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DispatchRequest.Marshal(b, m, deterministic)
}
func (dst *DispatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DispatchRequest.Merge(dst, src)
}
func (m *DispatchRequest) XXX_Size() int {
	return xxx_messageInfo_DispatchRequest.Size(m)
}
func (m *DispatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DispatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DispatchRequest proto.InternalMessageInfo

func (m *DispatchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DispatchRequest) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

// Department相关///////////////////////////////////////////////////////
type Department struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=Users,proto3" json:"Users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Department) Reset()         { *m = Department{} }
func (m *Department) String() string { return proto.CompactTextString(m) }
func (*Department) ProtoMessage()    {}
func (*Department) Descriptor() ([]byte, []int) {
	return fileDescriptor_struct_14940efc5789c3fe, []int{4}
}
func (m *Department) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Department.Unmarshal(m, b)
}
func (m *Department) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Department.Marshal(b, m, deterministic)
}
func (dst *Department) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Department.Merge(dst, src)
}
func (m *Department) XXX_Size() int {
	return xxx_messageInfo_Department.Size(m)
}
func (m *Department) XXX_DiscardUnknown() {
	xxx_messageInfo_Department.DiscardUnknown(m)
}

var xxx_messageInfo_Department proto.InternalMessageInfo

func (m *Department) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Department) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type DepartmentList struct {
	List                 []*Department `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DepartmentList) Reset()         { *m = DepartmentList{} }
func (m *DepartmentList) String() string { return proto.CompactTextString(m) }
func (*DepartmentList) ProtoMessage()    {}
func (*DepartmentList) Descriptor() ([]byte, []int) {
	return fileDescriptor_struct_14940efc5789c3fe, []int{5}
}
func (m *DepartmentList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepartmentList.Unmarshal(m, b)
}
func (m *DepartmentList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepartmentList.Marshal(b, m, deterministic)
}
func (dst *DepartmentList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepartmentList.Merge(dst, src)
}
func (m *DepartmentList) XXX_Size() int {
	return xxx_messageInfo_DepartmentList.Size(m)
}
func (m *DepartmentList) XXX_DiscardUnknown() {
	xxx_messageInfo_DepartmentList.DiscardUnknown(m)
}

var xxx_messageInfo_DepartmentList proto.InternalMessageInfo

func (m *DepartmentList) GetList() []*Department {
	if m != nil {
		return m.List
	}
	return nil
}

type PersonnelChangeRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Company              string   `protobuf:"bytes,2,opt,name=Company,proto3" json:"Company,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonnelChangeRequest) Reset()         { *m = PersonnelChangeRequest{} }
func (m *PersonnelChangeRequest) String() string { return proto.CompactTextString(m) }
func (*PersonnelChangeRequest) ProtoMessage()    {}
func (*PersonnelChangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_struct_14940efc5789c3fe, []int{6}
}
func (m *PersonnelChangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonnelChangeRequest.Unmarshal(m, b)
}
func (m *PersonnelChangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonnelChangeRequest.Marshal(b, m, deterministic)
}
func (dst *PersonnelChangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonnelChangeRequest.Merge(dst, src)
}
func (m *PersonnelChangeRequest) XXX_Size() int {
	return xxx_messageInfo_PersonnelChangeRequest.Size(m)
}
func (m *PersonnelChangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonnelChangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PersonnelChangeRequest proto.InternalMessageInfo

func (m *PersonnelChangeRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *PersonnelChangeRequest) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "common.Empty")
	proto.RegisterType((*User)(nil), "common.User")
	proto.RegisterType((*UserList)(nil), "common.UserList")
	proto.RegisterType((*DispatchRequest)(nil), "common.DispatchRequest")
	proto.RegisterType((*Department)(nil), "common.Department")
	proto.RegisterType((*DepartmentList)(nil), "common.DepartmentList")
	proto.RegisterType((*PersonnelChangeRequest)(nil), "common.PersonnelChangeRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserManagerServiceClient is the client API for UserManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserManagerServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserList, error)
	AddUser(ctx context.Context, in *UserList, opts ...grpc.CallOption) (*Empty, error)
	Dispatch(ctx context.Context, in *DispatchRequest, opts ...grpc.CallOption) (*Empty, error)
}

type userManagerServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserManagerServiceClient(cc *grpc.ClientConn) UserManagerServiceClient {
	return &userManagerServiceClient{cc}
}

func (c *userManagerServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/common.UserManagerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerServiceClient) AddUser(ctx context.Context, in *UserList, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/common.UserManagerService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerServiceClient) Dispatch(ctx context.Context, in *DispatchRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/common.UserManagerService/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagerServiceServer is the server API for UserManagerService service.
type UserManagerServiceServer interface {
	List(context.Context, *Empty) (*UserList, error)
	AddUser(context.Context, *UserList) (*Empty, error)
	Dispatch(context.Context, *DispatchRequest) (*Empty, error)
}

func RegisterUserManagerServiceServer(s *grpc.Server, srv UserManagerServiceServer) {
	s.RegisterService(&_UserManagerService_serviceDesc, srv)
}

func _UserManagerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.UserManagerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagerService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.UserManagerService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServiceServer).AddUser(ctx, req.(*UserList))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagerService_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServiceServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.UserManagerService/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServiceServer).Dispatch(ctx, req.(*DispatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.UserManagerService",
	HandlerType: (*UserManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserManagerService_List_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _UserManagerService_AddUser_Handler,
		},
		{
			MethodName: "Dispatch",
			Handler:    _UserManagerService_Dispatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "struct.proto",
}

// DepartmentManagerServiceClient is the client API for DepartmentManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DepartmentManagerServiceClient interface {
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DepartmentList, error)
	Create(ctx context.Context, in *Department, opts ...grpc.CallOption) (*Empty, error)
	PersonnelChange(ctx context.Context, in *PersonnelChangeRequest, opts ...grpc.CallOption) (*Empty, error)
}

type departmentManagerServiceClient struct {
	cc *grpc.ClientConn
}

func NewDepartmentManagerServiceClient(cc *grpc.ClientConn) DepartmentManagerServiceClient {
	return &departmentManagerServiceClient{cc}
}

func (c *departmentManagerServiceClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DepartmentList, error) {
	out := new(DepartmentList)
	err := c.cc.Invoke(ctx, "/common.DepartmentManagerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentManagerServiceClient) Create(ctx context.Context, in *Department, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/common.DepartmentManagerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentManagerServiceClient) PersonnelChange(ctx context.Context, in *PersonnelChangeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/common.DepartmentManagerService/PersonnelChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentManagerServiceServer is the server API for DepartmentManagerService service.
type DepartmentManagerServiceServer interface {
	List(context.Context, *Empty) (*DepartmentList, error)
	Create(context.Context, *Department) (*Empty, error)
	PersonnelChange(context.Context, *PersonnelChangeRequest) (*Empty, error)
}

func RegisterDepartmentManagerServiceServer(s *grpc.Server, srv DepartmentManagerServiceServer) {
	s.RegisterService(&_DepartmentManagerService_serviceDesc, srv)
}

func _DepartmentManagerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentManagerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.DepartmentManagerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentManagerServiceServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentManagerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Department)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentManagerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.DepartmentManagerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentManagerServiceServer).Create(ctx, req.(*Department))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentManagerService_PersonnelChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonnelChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentManagerServiceServer).PersonnelChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.DepartmentManagerService/PersonnelChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentManagerServiceServer).PersonnelChange(ctx, req.(*PersonnelChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DepartmentManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.DepartmentManagerService",
	HandlerType: (*DepartmentManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _DepartmentManagerService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DepartmentManagerService_Create_Handler,
		},
		{
			MethodName: "PersonnelChange",
			Handler:    _DepartmentManagerService_PersonnelChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "struct.proto",
}

func init() { proto.RegisterFile("struct.proto", fileDescriptor_struct_14940efc5789c3fe) }

var fileDescriptor_struct_14940efc5789c3fe = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x54, 0x4a, 0x5f, 0x2c, 0x85, 0xa2, 0x3d, 0x94, 0xa8, 0x07, 0x14, 0xf9, 0x00, 0x15, 0x8f,
	0x1e, 0x02, 0x07, 0x6e, 0x80, 0x5a, 0x6e, 0x80, 0x50, 0x81, 0x0f, 0x30, 0xe9, 0xaa, 0xad, 0x84,
	0xed, 0x60, 0xbb, 0x48, 0xfd, 0x18, 0xfe, 0x83, 0xcf, 0x43, 0x71, 0x30, 0xa5, 0x69, 0x38, 0x70,
	0x4a, 0xec, 0x9d, 0xd9, 0x9d, 0x99, 0x35, 0xb4, 0x8c, 0xd5, 0xf3, 0xc4, 0xf6, 0x53, 0xad, 0xac,
	0xc2, 0x7a, 0xa2, 0x84, 0x50, 0x92, 0x35, 0xa0, 0x76, 0x23, 0x52, 0xbb, 0x60, 0xe7, 0x50, 0x7d,
	0x36, 0xa4, 0x11, 0xa1, 0x7a, 0xcf, 0x05, 0x85, 0x41, 0x14, 0xf4, 0x36, 0x47, 0xee, 0x1f, 0x43,
	0x68, 0x0c, 0x94, 0x48, 0xb9, 0x5c, 0x84, 0x15, 0x77, 0xed, 0x8f, 0xec, 0x04, 0x9a, 0x19, 0xeb,
	0x76, 0x66, 0x2c, 0x46, 0x50, 0xcd, 0xbe, 0x61, 0x10, 0x6d, 0xf4, 0xb6, 0xe2, 0x56, 0x3f, 0x9f,
	0xd0, 0xcf, 0xea, 0x23, 0x57, 0x61, 0x97, 0xd0, 0x1e, 0xce, 0x4c, 0xca, 0x6d, 0x32, 0x1d, 0xd1,
	0xdb, 0x9c, 0x8c, 0xfd, 0xe7, 0xb8, 0x21, 0xc0, 0x90, 0x52, 0xae, 0xad, 0x20, 0x59, 0xce, 0x65,
	0x50, 0xcb, 0x06, 0x9a, 0xb0, 0x52, 0xa2, 0x22, 0x2f, 0xb1, 0x0b, 0xd8, 0x59, 0x76, 0x71, 0xd2,
	0x0f, 0x56, 0xa4, 0xa3, 0x27, 0x2d, 0x51, 0xdf, 0x06, 0x9e, 0xa0, 0xf3, 0x40, 0xda, 0x28, 0x29,
	0xe9, 0x75, 0x30, 0xe5, 0x72, 0x42, 0xde, 0x47, 0x94, 0xc7, 0xe7, 0xb4, 0xac, 0x99, 0x77, 0xc1,
	0xfe, 0xe9, 0x2a, 0xfe, 0x08, 0x00, 0x33, 0xc8, 0x1d, 0x97, 0x7c, 0x42, 0xfa, 0x91, 0xf4, 0xfb,
	0x2c, 0x21, 0x3c, 0xcc, 0x45, 0xe1, 0xb6, 0x6f, 0xe6, 0x16, 0xd5, 0xdd, 0xfd, 0xdd, 0xdb, 0x01,
	0x8e, 0xa0, 0x71, 0x3d, 0x1e, 0xbb, 0x21, 0x6b, 0xc5, 0xee, 0x2a, 0x1b, 0x63, 0x68, 0xfa, 0x15,
	0xe0, 0xde, 0x8f, 0xcf, 0xd5, 0xa5, 0x14, 0x38, 0xf1, 0x67, 0x00, 0xe1, 0x32, 0x8a, 0x82, 0xca,
	0xd3, 0x72, 0x95, 0x9d, 0xf5, 0x0c, 0x1d, 0xec, 0x18, 0xea, 0x03, 0x4d, 0xdc, 0x12, 0x96, 0xa4,
	0x5c, 0x14, 0x7b, 0x05, 0xed, 0x42, 0xdc, 0xb8, 0xef, 0x11, 0xe5, 0x7b, 0x28, 0x74, 0x78, 0xa9,
	0xbb, 0xd7, 0x7e, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x1c, 0x0c, 0xb7, 0xfd, 0x02, 0x00,
	0x00,
}
