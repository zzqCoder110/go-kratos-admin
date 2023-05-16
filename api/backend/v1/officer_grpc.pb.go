// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1--rc1
// source: backend/v1/officer.proto

package v1

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

// OfficerClient is the client API for Officer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OfficerClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRep, error)
}

type officerClient struct {
	cc grpc.ClientConnInterface
}

func NewOfficerClient(cc grpc.ClientConnInterface) OfficerClient {
	return &officerClient{cc}
}

func (c *officerClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRep, error) {
	out := new(CreateRep)
	err := c.cc.Invoke(ctx, "/backend.v1.Officer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OfficerServer is the server API for Officer service.
// All implementations must embed UnimplementedOfficerServer
// for forward compatibility
type OfficerServer interface {
	Create(context.Context, *CreateReq) (*CreateRep, error)
	mustEmbedUnimplementedOfficerServer()
}

// UnimplementedOfficerServer must be embedded to have forward compatible implementations.
type UnimplementedOfficerServer struct {
}

func (UnimplementedOfficerServer) Create(context.Context, *CreateReq) (*CreateRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOfficerServer) mustEmbedUnimplementedOfficerServer() {}

// UnsafeOfficerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OfficerServer will
// result in compilation errors.
type UnsafeOfficerServer interface {
	mustEmbedUnimplementedOfficerServer()
}

func RegisterOfficerServer(s grpc.ServiceRegistrar, srv OfficerServer) {
	s.RegisterService(&Officer_ServiceDesc, srv)
}

func _Officer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfficerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.v1.Officer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfficerServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Officer_ServiceDesc is the grpc.ServiceDesc for Officer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Officer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "backend.v1.Officer",
	HandlerType: (*OfficerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Officer_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend/v1/officer.proto",
}
