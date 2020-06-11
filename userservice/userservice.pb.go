// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userservice.proto

package userservice

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

type LoginResponse_LoginStatus int32

const (
	LoginResponse_SUCCESS             LoginResponse_LoginStatus = 0
	LoginResponse_INVALID_CREDENTIALS LoginResponse_LoginStatus = 1
	LoginResponse_NOT_ACTIVATED       LoginResponse_LoginStatus = 2
)

var LoginResponse_LoginStatus_name = map[int32]string{
	0: "SUCCESS",
	1: "INVALID_CREDENTIALS",
	2: "NOT_ACTIVATED",
}

var LoginResponse_LoginStatus_value = map[string]int32{
	"SUCCESS":             0,
	"INVALID_CREDENTIALS": 1,
	"NOT_ACTIVATED":       2,
}

func (x LoginResponse_LoginStatus) String() string {
	return proto.EnumName(LoginResponse_LoginStatus_name, int32(x))
}

func (LoginResponse_LoginStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{3, 0}
}

type CreateUserRequest_PrivacyLevel int32

const (
	CreateUserRequest_PUBLIC    CreateUserRequest_PrivacyLevel = 0
	CreateUserRequest_PROTECTED CreateUserRequest_PrivacyLevel = 1
	CreateUserRequest_PRIVATE   CreateUserRequest_PrivacyLevel = 2
)

var CreateUserRequest_PrivacyLevel_name = map[int32]string{
	0: "PUBLIC",
	1: "PROTECTED",
	2: "PRIVATE",
}

var CreateUserRequest_PrivacyLevel_value = map[string]int32{
	"PUBLIC":    0,
	"PROTECTED": 1,
	"PRIVATE":   2,
}

func (x CreateUserRequest_PrivacyLevel) String() string {
	return proto.EnumName(CreateUserRequest_PrivacyLevel_name, int32(x))
}

func (CreateUserRequest_PrivacyLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{4, 0}
}

type CreateUserResponse_CreationStatus int32

const (
	CreateUserResponse_SUCCESS         CreateUserResponse_CreationStatus = 0
	CreateUserResponse_EMAIL_EXISTS    CreateUserResponse_CreationStatus = 1
	CreateUserResponse_USERNAME_EXISTS CreateUserResponse_CreationStatus = 2
)

var CreateUserResponse_CreationStatus_name = map[int32]string{
	0: "SUCCESS",
	1: "EMAIL_EXISTS",
	2: "USERNAME_EXISTS",
}

var CreateUserResponse_CreationStatus_value = map[string]int32{
	"SUCCESS":         0,
	"EMAIL_EXISTS":    1,
	"USERNAME_EXISTS": 2,
}

func (x CreateUserResponse_CreationStatus) String() string {
	return proto.EnumName(CreateUserResponse_CreationStatus_name, int32(x))
}

func (CreateUserResponse_CreationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{5, 0}
}

type ActivateUserRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivateUserRequest) Reset()         { *m = ActivateUserRequest{} }
func (m *ActivateUserRequest) String() string { return proto.CompactTextString(m) }
func (*ActivateUserRequest) ProtoMessage()    {}
func (*ActivateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{0}
}

func (m *ActivateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateUserRequest.Unmarshal(m, b)
}
func (m *ActivateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateUserRequest.Marshal(b, m, deterministic)
}
func (m *ActivateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateUserRequest.Merge(m, src)
}
func (m *ActivateUserRequest) XXX_Size() int {
	return xxx_messageInfo_ActivateUserRequest.Size(m)
}
func (m *ActivateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateUserRequest proto.InternalMessageInfo

func (m *ActivateUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ActivateUserResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivateUserResponse) Reset()         { *m = ActivateUserResponse{} }
func (m *ActivateUserResponse) String() string { return proto.CompactTextString(m) }
func (*ActivateUserResponse) ProtoMessage()    {}
func (*ActivateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{1}
}

func (m *ActivateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateUserResponse.Unmarshal(m, b)
}
func (m *ActivateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateUserResponse.Marshal(b, m, deterministic)
}
func (m *ActivateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateUserResponse.Merge(m, src)
}
func (m *ActivateUserResponse) XXX_Size() int {
	return xxx_messageInfo_ActivateUserResponse.Size(m)
}
func (m *ActivateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateUserResponse proto.InternalMessageInfo

func (m *ActivateUserResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ActivateUserResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LoginRequest struct {
	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	// Types that are valid to be assigned to UserCredential:
	//	*LoginRequest_Username
	//	*LoginRequest_Email
	UserCredential       isLoginRequest_UserCredential `protobuf_oneof:"userCredential"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type isLoginRequest_UserCredential interface {
	isLoginRequest_UserCredential()
}

type LoginRequest_Username struct {
	Username string `protobuf:"bytes,2,opt,name=username,proto3,oneof"`
}

type LoginRequest_Email struct {
	Email string `protobuf:"bytes,3,opt,name=email,proto3,oneof"`
}

func (*LoginRequest_Username) isLoginRequest_UserCredential() {}

func (*LoginRequest_Email) isLoginRequest_UserCredential() {}

func (m *LoginRequest) GetUserCredential() isLoginRequest_UserCredential {
	if m != nil {
		return m.UserCredential
	}
	return nil
}

func (m *LoginRequest) GetUsername() string {
	if x, ok := m.GetUserCredential().(*LoginRequest_Username); ok {
		return x.Username
	}
	return ""
}

func (m *LoginRequest) GetEmail() string {
	if x, ok := m.GetUserCredential().(*LoginRequest_Email); ok {
		return x.Email
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LoginRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LoginRequest_Username)(nil),
		(*LoginRequest_Email)(nil),
	}
}

type LoginResponse struct {
	Status               LoginResponse_LoginStatus `protobuf:"varint,1,opt,name=status,proto3,enum=userservice.LoginResponse_LoginStatus" json:"status,omitempty"`
	UserId               string                    `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Role                 string                    `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetStatus() LoginResponse_LoginStatus {
	if m != nil {
		return m.Status
	}
	return LoginResponse_SUCCESS
}

func (m *LoginResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LoginResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type CreateUserRequest struct {
	Username             string                         `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string                         `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Firstname            string                         `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string                         `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Level                CreateUserRequest_PrivacyLevel `protobuf:"varint,5,opt,name=level,proto3,enum=userservice.CreateUserRequest_PrivacyLevel" json:"level,omitempty"`
	Password             string                         `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{4}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *CreateUserRequest) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *CreateUserRequest) GetLevel() CreateUserRequest_PrivacyLevel {
	if m != nil {
		return m.Level
	}
	return CreateUserRequest_PUBLIC
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	Status               CreateUserResponse_CreationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=userservice.CreateUserResponse_CreationStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{5}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetStatus() CreateUserResponse_CreationStatus {
	if m != nil {
		return m.Status
	}
	return CreateUserResponse_SUCCESS
}

func init() {
	proto.RegisterEnum("userservice.LoginResponse_LoginStatus", LoginResponse_LoginStatus_name, LoginResponse_LoginStatus_value)
	proto.RegisterEnum("userservice.CreateUserRequest_PrivacyLevel", CreateUserRequest_PrivacyLevel_name, CreateUserRequest_PrivacyLevel_value)
	proto.RegisterEnum("userservice.CreateUserResponse_CreationStatus", CreateUserResponse_CreationStatus_name, CreateUserResponse_CreationStatus_value)
	proto.RegisterType((*ActivateUserRequest)(nil), "userservice.ActivateUserRequest")
	proto.RegisterType((*ActivateUserResponse)(nil), "userservice.ActivateUserResponse")
	proto.RegisterType((*LoginRequest)(nil), "userservice.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "userservice.LoginResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "userservice.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "userservice.CreateUserResponse")
}

func init() {
	proto.RegisterFile("userservice.proto", fileDescriptor_68a7ca558839fd2b)
}

var fileDescriptor_68a7ca558839fd2b = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x4d, 0xc2, 0x5a, 0xda, 0xdb, 0x0f, 0x52, 0xb7, 0x1a, 0xa5, 0x9a, 0x60, 0xf8, 0x01, 0x21,
	0x21, 0xfa, 0x30, 0x24, 0x1e, 0x91, 0xb2, 0x24, 0x13, 0x91, 0xb2, 0xae, 0x38, 0xc9, 0xc4, 0x5b,
	0x15, 0x3a, 0x33, 0x45, 0xca, 0x9a, 0x62, 0xa7, 0x45, 0xfc, 0x11, 0x7e, 0x00, 0x7f, 0x85, 0xdf,
	0x85, 0x84, 0x6c, 0x27, 0x5d, 0x52, 0xb6, 0xf1, 0xd6, 0x73, 0x7d, 0xee, 0xb9, 0x37, 0xe7, 0xd8,
	0x85, 0xc1, 0x86, 0x53, 0xc6, 0x29, 0xdb, 0x26, 0x4b, 0x3a, 0x5d, 0xb3, 0x2c, 0xcf, 0x50, 0xa7,
	0x52, 0xc2, 0x6f, 0x61, 0x68, 0x2d, 0xf3, 0x64, 0x1b, 0xe7, 0x34, 0xe2, 0x94, 0x11, 0xfa, 0x6d,
	0x43, 0x79, 0x8e, 0x0e, 0xa1, 0x29, 0x58, 0xde, 0xd5, 0x58, 0x3f, 0xd6, 0x5f, 0xb7, 0x49, 0x81,
	0xb0, 0x03, 0xa3, 0x3a, 0x9d, 0xaf, 0xb3, 0x15, 0xa7, 0x82, 0xcf, 0xf3, 0x38, 0xdf, 0x70, 0xc9,
	0x6f, 0x91, 0x02, 0xa1, 0x11, 0x34, 0x28, 0x63, 0x19, 0x1b, 0x1b, 0x52, 0x46, 0x01, 0xcc, 0xa0,
	0xeb, 0x67, 0xd7, 0xc9, 0xaa, 0x9c, 0x36, 0x81, 0xd6, 0x3a, 0xe6, 0xfc, 0x7b, 0xc6, 0xca, 0x79,
	0x3b, 0x8c, 0x8e, 0xa0, 0x25, 0x66, 0xaf, 0xe2, 0x1b, 0xaa, 0x44, 0x3e, 0x6a, 0x64, 0x57, 0x41,
	0x87, 0xd0, 0xa0, 0x37, 0x71, 0x92, 0x8e, 0x1f, 0x15, 0x47, 0x0a, 0x9e, 0x9a, 0xd0, 0x17, 0x1c,
	0x9b, 0xd1, 0x2b, 0xba, 0xca, 0x93, 0x38, 0xc5, 0xbf, 0x75, 0xe8, 0x15, 0x43, 0x8b, 0x9d, 0x3f,
	0xd4, 0x76, 0xee, 0x9f, 0xbc, 0x9a, 0x56, 0xbd, 0xaa, 0x71, 0x15, 0x0a, 0x24, 0x7b, 0xf7, 0x6d,
	0xb7, 0x1e, 0x19, 0x55, 0x8f, 0x10, 0x82, 0x03, 0x96, 0xa5, 0x54, 0xad, 0x44, 0xe4, 0x6f, 0x7c,
	0x06, 0x9d, 0x8a, 0x04, 0xea, 0xc0, 0xe3, 0x20, 0xb2, 0x6d, 0x37, 0x08, 0x4c, 0x0d, 0x3d, 0x85,
	0xa1, 0x37, 0xbb, 0xb4, 0x7c, 0xcf, 0x59, 0xd8, 0xc4, 0x75, 0xdc, 0x59, 0xe8, 0x59, 0x7e, 0x60,
	0xea, 0x68, 0x00, 0xbd, 0xd9, 0x45, 0xb8, 0xb0, 0xec, 0xd0, 0xbb, 0xb4, 0x42, 0xd7, 0x31, 0x0d,
	0xfc, 0xd3, 0x80, 0x81, 0xcd, 0xe8, 0x5e, 0x5a, 0x93, 0x8a, 0x47, 0x85, 0x7f, 0x3b, 0x87, 0x46,
	0xa5, 0x43, 0x65, 0x02, 0x02, 0xa0, 0x23, 0x68, 0x7f, 0x4d, 0x18, 0xcf, 0x65, 0x8b, 0x5a, 0xf4,
	0xb6, 0x20, 0xf4, 0xd2, 0xb8, 0x38, 0x3c, 0x50, 0x7a, 0x25, 0x46, 0x16, 0x34, 0x52, 0xba, 0xa5,
	0xe9, 0xb8, 0x21, 0x4d, 0x7b, 0x53, 0x33, 0xed, 0x9f, 0xd5, 0xa6, 0x73, 0x96, 0x6c, 0xe3, 0xe5,
	0x0f, 0x5f, 0xb4, 0x10, 0xd5, 0x59, 0x8b, 0xbb, 0x59, 0x8f, 0x1b, 0xbf, 0x87, 0x6e, 0xb5, 0x05,
	0x01, 0x34, 0xe7, 0xd1, 0xa9, 0xef, 0xd9, 0xa6, 0x86, 0x7a, 0xd0, 0x9e, 0x93, 0x8b, 0xd0, 0xb5,
	0x85, 0x17, 0xba, 0x30, 0x71, 0x4e, 0xa4, 0x33, 0xa6, 0x81, 0x7f, 0xe9, 0x80, 0xaa, 0xd3, 0x8b,
	0x8c, 0xcf, 0xf6, 0x32, 0x9e, 0xde, 0xbb, 0x6e, 0x11, 0xb4, 0x2c, 0x25, 0xd9, 0x5e, 0xd6, 0xd8,
	0x81, 0x7e, 0xfd, 0xa4, 0x1e, 0xa1, 0x09, 0x5d, 0xf7, 0xdc, 0xf2, 0xfc, 0x85, 0xfb, 0xd9, 0x0b,
	0x42, 0x91, 0xdd, 0x10, 0x9e, 0x44, 0x81, 0x4b, 0x66, 0xd6, 0xb9, 0x5b, 0x16, 0x8d, 0x93, 0x3f,
	0x3a, 0x74, 0xc4, 0xb4, 0x40, 0xcd, 0x47, 0x11, 0x74, 0xab, 0xaf, 0x09, 0x1d, 0xd7, 0xb6, 0xbb,
	0xe3, 0x5d, 0x4e, 0x5e, 0x3e, 0xc0, 0x50, 0x5f, 0x80, 0x35, 0xe4, 0x40, 0x5b, 0x5e, 0x36, 0xa9,
	0xf9, 0xec, 0xae, 0x5b, 0xad, 0xc4, 0x26, 0xf7, 0x5f, 0x78, 0xac, 0xa1, 0x4f, 0xd0, 0x25, 0xf4,
	0x3a, 0xe1, 0x39, 0x65, 0x52, 0xe8, 0xf9, 0xc3, 0x49, 0x4f, 0x5e, 0xfc, 0xc7, 0x5a, 0xac, 0x7d,
	0x69, 0xca, 0x3f, 0xa0, 0x77, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xa8, 0xc7, 0xbc, 0x95,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	ActivateUser(ctx context.Context, in *ActivateUserRequest, opts ...grpc.CallOption) (*ActivateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	RegisterUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) ActivateUser(ctx context.Context, in *ActivateUserRequest, opts ...grpc.CallOption) (*ActivateUserResponse, error) {
	out := new(ActivateUserResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/ActivateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	ActivateUser(context.Context, *ActivateUserRequest) (*ActivateUserResponse, error)
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	RegisterUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) ActivateUser(ctx context.Context, req *ActivateUserRequest) (*ActivateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateUser not implemented")
}
func (*UnimplementedUserServiceServer) LoginUser(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (*UnimplementedUserServiceServer) RegisterUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_ActivateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ActivateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/ActivateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ActivateUser(ctx, req.(*ActivateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userservice.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActivateUser",
			Handler:    _UserService_ActivateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserService_LoginUser_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userservice.proto",
}
