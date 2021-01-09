// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package userservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
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
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	ActivateUser(context.Context, *ActivateUserRequest) (*ActivateUserResponse, error)
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	RegisterUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) ActivateUser(context.Context, *ActivateUserRequest) (*ActivateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateUser not implemented")
}
func (UnimplementedUserServiceServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUserServiceServer) RegisterUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
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
