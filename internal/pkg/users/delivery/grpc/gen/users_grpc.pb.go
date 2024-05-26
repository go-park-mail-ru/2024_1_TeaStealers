// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: users.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Users_GetCurUser_FullMethodName     = "/users.Users/GetCurUser"
	Users_UpdateUserInfo_FullMethodName = "/users.Users/UpdateUserInfo"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	GetCurUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) GetCurUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, Users_GetCurUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, Users_UpdateUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	GetCurUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) GetCurUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurUser not implemented")
}
func (UnimplementedUsersServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_GetCurUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetCurUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_GetCurUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetCurUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurUser",
			Handler:    _Users_GetCurUser_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _Users_UpdateUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
