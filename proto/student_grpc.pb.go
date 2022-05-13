// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// StudentClient is the client API for Student service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentClient interface {
	GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error)
}

type studentClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentClient(cc grpc.ClientConnInterface) StudentClient {
	return &studentClient{cc}
}

func (c *studentClient) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*GetStudentResponse, error) {
	out := new(GetStudentResponse)
	err := c.cc.Invoke(ctx, "/proto.Student/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServer is the server API for Student service.
// All implementations must embed UnimplementedStudentServer
// for forward compatibility
type StudentServer interface {
	GetStudent(context.Context, *GetStudentRequest) (*GetStudentResponse, error)
	mustEmbedUnimplementedStudentServer()
}

// UnimplementedStudentServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServer struct {
}

func (UnimplementedStudentServer) GetStudent(context.Context, *GetStudentRequest) (*GetStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (UnimplementedStudentServer) mustEmbedUnimplementedStudentServer() {}

// UnsafeStudentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServer will
// result in compilation errors.
type UnsafeStudentServer interface {
	mustEmbedUnimplementedStudentServer()
}

func RegisterStudentServer(s grpc.ServiceRegistrar, srv StudentServer) {
	s.RegisterService(&Student_ServiceDesc, srv)
}

func _Student_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Student/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).GetStudent(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Student_ServiceDesc is the grpc.ServiceDesc for Student service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Student_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Student",
	HandlerType: (*StudentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudent",
			Handler:    _Student_GetStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/student.proto",
}
