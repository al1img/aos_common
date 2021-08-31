// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package updatemanager

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

// UMControllerClient is the client API for UMController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UMControllerClient interface {
	RegisterUM(ctx context.Context, opts ...grpc.CallOption) (UMController_RegisterUMClient, error)
}

type uMControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewUMControllerClient(cc grpc.ClientConnInterface) UMControllerClient {
	return &uMControllerClient{cc}
}

func (c *uMControllerClient) RegisterUM(ctx context.Context, opts ...grpc.CallOption) (UMController_RegisterUMClient, error) {
	stream, err := c.cc.NewStream(ctx, &UMController_ServiceDesc.Streams[0], "/updatemanager.UMController/RegisterUM", opts...)
	if err != nil {
		return nil, err
	}
	x := &uMControllerRegisterUMClient{stream}
	return x, nil
}

type UMController_RegisterUMClient interface {
	Send(*UpdateStatus) error
	Recv() (*CMMessages, error)
	grpc.ClientStream
}

type uMControllerRegisterUMClient struct {
	grpc.ClientStream
}

func (x *uMControllerRegisterUMClient) Send(m *UpdateStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uMControllerRegisterUMClient) Recv() (*CMMessages, error) {
	m := new(CMMessages)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UMControllerServer is the server API for UMController service.
// All implementations must embed UnimplementedUMControllerServer
// for forward compatibility
type UMControllerServer interface {
	RegisterUM(UMController_RegisterUMServer) error
	mustEmbedUnimplementedUMControllerServer()
}

// UnimplementedUMControllerServer must be embedded to have forward compatible implementations.
type UnimplementedUMControllerServer struct {
}

func (UnimplementedUMControllerServer) RegisterUM(UMController_RegisterUMServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterUM not implemented")
}
func (UnimplementedUMControllerServer) mustEmbedUnimplementedUMControllerServer() {}

// UnsafeUMControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UMControllerServer will
// result in compilation errors.
type UnsafeUMControllerServer interface {
	mustEmbedUnimplementedUMControllerServer()
}

func RegisterUMControllerServer(s grpc.ServiceRegistrar, srv UMControllerServer) {
	s.RegisterService(&UMController_ServiceDesc, srv)
}

func _UMController_RegisterUM_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UMControllerServer).RegisterUM(&uMControllerRegisterUMServer{stream})
}

type UMController_RegisterUMServer interface {
	Send(*CMMessages) error
	Recv() (*UpdateStatus, error)
	grpc.ServerStream
}

type uMControllerRegisterUMServer struct {
	grpc.ServerStream
}

func (x *uMControllerRegisterUMServer) Send(m *CMMessages) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uMControllerRegisterUMServer) Recv() (*UpdateStatus, error) {
	m := new(UpdateStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UMController_ServiceDesc is the grpc.ServiceDesc for UMController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UMController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "updatemanager.UMController",
	HandlerType: (*UMControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterUM",
			Handler:       _UMController_RegisterUM_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "updatemanager/updatemanager.proto",
}
