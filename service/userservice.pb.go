// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.2
// source: userservice.proto

package userService

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LoginResponse_LoginStatus int32

const (
	LoginResponse_SUCCESS             LoginResponse_LoginStatus = 0
	LoginResponse_INVALID_CREDENTIALS LoginResponse_LoginStatus = 1
	LoginResponse_NOT_ACTIVATED       LoginResponse_LoginStatus = 2
)

// Enum value maps for LoginResponse_LoginStatus.
var (
	LoginResponse_LoginStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "INVALID_CREDENTIALS",
		2: "NOT_ACTIVATED",
	}
	LoginResponse_LoginStatus_value = map[string]int32{
		"SUCCESS":             0,
		"INVALID_CREDENTIALS": 1,
		"NOT_ACTIVATED":       2,
	}
)

func (x LoginResponse_LoginStatus) Enum() *LoginResponse_LoginStatus {
	p := new(LoginResponse_LoginStatus)
	*p = x
	return p
}

func (x LoginResponse_LoginStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginResponse_LoginStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_userservice_proto_enumTypes[0].Descriptor()
}

func (LoginResponse_LoginStatus) Type() protoreflect.EnumType {
	return &file_userservice_proto_enumTypes[0]
}

func (x LoginResponse_LoginStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginResponse_LoginStatus.Descriptor instead.
func (LoginResponse_LoginStatus) EnumDescriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{3, 0}
}

type ActivateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ActivateUserRequest) Reset() {
	*x = ActivateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateUserRequest) ProtoMessage() {}

func (x *ActivateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateUserRequest.ProtoReflect.Descriptor instead.
func (*ActivateUserRequest) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{0}
}

func (x *ActivateUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ActivateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ActivateUserResponse) Reset() {
	*x = ActivateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateUserResponse) ProtoMessage() {}

func (x *ActivateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateUserResponse.ProtoReflect.Descriptor instead.
func (*ActivateUserResponse) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{1}
}

func (x *ActivateUserResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ActivateUserResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	// Types that are assignable to UserCredential:
	//	*LoginRequest_Username
	//	*LoginRequest_Email
	UserCredential isLoginRequest_UserCredential `protobuf_oneof:"userCredential"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (m *LoginRequest) GetUserCredential() isLoginRequest_UserCredential {
	if m != nil {
		return m.UserCredential
	}
	return nil
}

func (x *LoginRequest) GetUsername() string {
	if x, ok := x.GetUserCredential().(*LoginRequest_Username); ok {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetEmail() string {
	if x, ok := x.GetUserCredential().(*LoginRequest_Email); ok {
		return x.Email
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

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{3}
}

var File_userservice_proto protoreflect.FileDescriptor

var file_userservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x2d, 0x0a, 0x13, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x44, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x72, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1c, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x10, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x57, 0x0a, 0x0d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a, 0x0b, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x53, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x02, 0x32, 0xaa, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userservice_proto_rawDescOnce sync.Once
	file_userservice_proto_rawDescData = file_userservice_proto_rawDesc
)

func file_userservice_proto_rawDescGZIP() []byte {
	file_userservice_proto_rawDescOnce.Do(func() {
		file_userservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_userservice_proto_rawDescData)
	})
	return file_userservice_proto_rawDescData
}

var file_userservice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_userservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_userservice_proto_goTypes = []interface{}{
	(LoginResponse_LoginStatus)(0), // 0: userService.LoginResponse.LoginStatus
	(*ActivateUserRequest)(nil),    // 1: userService.ActivateUserRequest
	(*ActivateUserResponse)(nil),   // 2: userService.ActivateUserResponse
	(*LoginRequest)(nil),           // 3: userService.LoginRequest
	(*LoginResponse)(nil),          // 4: userService.LoginResponse
}
var file_userservice_proto_depIdxs = []int32{
	1, // 0: userService.UserService.ActivateUser:input_type -> userService.ActivateUserRequest
	3, // 1: userService.UserService.LoginUser:input_type -> userService.LoginRequest
	2, // 2: userService.UserService.ActivateUser:output_type -> userService.ActivateUserResponse
	4, // 3: userService.UserService.LoginUser:output_type -> userService.LoginResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userservice_proto_init() }
func file_userservice_proto_init() {
	if File_userservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_userservice_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*LoginRequest_Username)(nil),
		(*LoginRequest_Email)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_userservice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userservice_proto_goTypes,
		DependencyIndexes: file_userservice_proto_depIdxs,
		EnumInfos:         file_userservice_proto_enumTypes,
		MessageInfos:      file_userservice_proto_msgTypes,
	}.Build()
	File_userservice_proto = out.File
	file_userservice_proto_rawDesc = nil
	file_userservice_proto_goTypes = nil
	file_userservice_proto_depIdxs = nil
}
