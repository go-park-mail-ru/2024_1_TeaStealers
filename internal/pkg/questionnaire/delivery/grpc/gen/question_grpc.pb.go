// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: question.proto

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
	Question_GetQuestionsByTheme_FullMethodName = "/questionnaire.Question/GetQuestionsByTheme"
)

// QuestionClient is the client API for Question service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionClient interface {
	GetQuestionsByTheme(ctx context.Context, in *GetQuestionsByThemeRequest, opts ...grpc.CallOption) (*GetQuestionsByThemeResponse, error)
}

type questionClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionClient(cc grpc.ClientConnInterface) QuestionClient {
	return &questionClient{cc}
}

func (c *questionClient) GetQuestionsByTheme(ctx context.Context, in *GetQuestionsByThemeRequest, opts ...grpc.CallOption) (*GetQuestionsByThemeResponse, error) {
	out := new(GetQuestionsByThemeResponse)
	err := c.cc.Invoke(ctx, Question_GetQuestionsByTheme_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServer is the server API for Question service.
// All implementations must embed UnimplementedQuestionServer
// for forward compatibility
type QuestionServer interface {
	GetQuestionsByTheme(context.Context, *GetQuestionsByThemeRequest) (*GetQuestionsByThemeResponse, error)
	mustEmbedUnimplementedQuestionServer()
}

// UnimplementedQuestionServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionServer struct {
}

func (UnimplementedQuestionServer) GetQuestionsByTheme(context.Context, *GetQuestionsByThemeRequest) (*GetQuestionsByThemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionsByTheme not implemented")
}
func (UnimplementedQuestionServer) mustEmbedUnimplementedQuestionServer() {}

// UnsafeQuestionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionServer will
// result in compilation errors.
type UnsafeQuestionServer interface {
	mustEmbedUnimplementedQuestionServer()
}

func RegisterQuestionServer(s grpc.ServiceRegistrar, srv QuestionServer) {
	s.RegisterService(&Question_ServiceDesc, srv)
}

func _Question_GetQuestionsByTheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionsByThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).GetQuestionsByTheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Question_GetQuestionsByTheme_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).GetQuestionsByTheme(ctx, req.(*GetQuestionsByThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Question_ServiceDesc is the grpc.ServiceDesc for Question service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Question_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "questionnaire.Question",
	HandlerType: (*QuestionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuestionsByTheme",
			Handler:    _Question_GetQuestionsByTheme_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "question.proto",
}