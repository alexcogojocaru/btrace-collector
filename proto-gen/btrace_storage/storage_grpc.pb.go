// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: v2/storage.proto

package btrace_storage

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

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	Store(ctx context.Context, in *StorageSpan, opts ...grpc.CallOption) (*StorageResponse, error)
	GetServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Storage_GetServicesClient, error)
	GetServiceData(ctx context.Context, in *ServiceName, opts ...grpc.CallOption) (Storage_GetServiceDataClient, error)
	GetMultipleServicesData(ctx context.Context, opts ...grpc.CallOption) (Storage_GetMultipleServicesDataClient, error)
	GetData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Storage_GetDataClient, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Store(ctx context.Context, in *StorageSpan, opts ...grpc.CallOption) (*StorageResponse, error) {
	out := new(StorageResponse)
	err := c.cc.Invoke(ctx, "/Storage/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Storage_GetServicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[0], "/Storage/GetServices", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetServicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_GetServicesClient interface {
	Recv() (*ServiceName, error)
	grpc.ClientStream
}

type storageGetServicesClient struct {
	grpc.ClientStream
}

func (x *storageGetServicesClient) Recv() (*ServiceName, error) {
	m := new(ServiceName)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) GetServiceData(ctx context.Context, in *ServiceName, opts ...grpc.CallOption) (Storage_GetServiceDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[1], "/Storage/GetServiceData", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetServiceDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_GetServiceDataClient interface {
	Recv() (*ServiceResponse, error)
	grpc.ClientStream
}

type storageGetServiceDataClient struct {
	grpc.ClientStream
}

func (x *storageGetServiceDataClient) Recv() (*ServiceResponse, error) {
	m := new(ServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) GetMultipleServicesData(ctx context.Context, opts ...grpc.CallOption) (Storage_GetMultipleServicesDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[2], "/Storage/GetMultipleServicesData", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetMultipleServicesDataClient{stream}
	return x, nil
}

type Storage_GetMultipleServicesDataClient interface {
	Send(*ServiceName) error
	CloseAndRecv() (*ServiceResponse, error)
	grpc.ClientStream
}

type storageGetMultipleServicesDataClient struct {
	grpc.ClientStream
}

func (x *storageGetMultipleServicesDataClient) Send(m *ServiceName) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageGetMultipleServicesDataClient) CloseAndRecv() (*ServiceResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) GetData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Storage_GetDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[3], "/Storage/GetData", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_GetDataClient interface {
	Recv() (*ServicePair, error)
	grpc.ClientStream
}

type storageGetDataClient struct {
	grpc.ClientStream
}

func (x *storageGetDataClient) Recv() (*ServicePair, error) {
	m := new(ServicePair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	Store(context.Context, *StorageSpan) (*StorageResponse, error)
	GetServices(*Empty, Storage_GetServicesServer) error
	GetServiceData(*ServiceName, Storage_GetServiceDataServer) error
	GetMultipleServicesData(Storage_GetMultipleServicesDataServer) error
	GetData(*Empty, Storage_GetDataServer) error
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) Store(context.Context, *StorageSpan) (*StorageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedStorageServer) GetServices(*Empty, Storage_GetServicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedStorageServer) GetServiceData(*ServiceName, Storage_GetServiceDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetServiceData not implemented")
}
func (UnimplementedStorageServer) GetMultipleServicesData(Storage_GetMultipleServicesDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMultipleServicesData not implemented")
}
func (UnimplementedStorageServer) GetData(*Empty, Storage_GetDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageSpan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Storage/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Store(ctx, req.(*StorageSpan))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetServices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).GetServices(m, &storageGetServicesServer{stream})
}

type Storage_GetServicesServer interface {
	Send(*ServiceName) error
	grpc.ServerStream
}

type storageGetServicesServer struct {
	grpc.ServerStream
}

func (x *storageGetServicesServer) Send(m *ServiceName) error {
	return x.ServerStream.SendMsg(m)
}

func _Storage_GetServiceData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServiceName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).GetServiceData(m, &storageGetServiceDataServer{stream})
}

type Storage_GetServiceDataServer interface {
	Send(*ServiceResponse) error
	grpc.ServerStream
}

type storageGetServiceDataServer struct {
	grpc.ServerStream
}

func (x *storageGetServiceDataServer) Send(m *ServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Storage_GetMultipleServicesData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).GetMultipleServicesData(&storageGetMultipleServicesDataServer{stream})
}

type Storage_GetMultipleServicesDataServer interface {
	SendAndClose(*ServiceResponse) error
	Recv() (*ServiceName, error)
	grpc.ServerStream
}

type storageGetMultipleServicesDataServer struct {
	grpc.ServerStream
}

func (x *storageGetMultipleServicesDataServer) SendAndClose(m *ServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageGetMultipleServicesDataServer) Recv() (*ServiceName, error) {
	m := new(ServiceName)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Storage_GetData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).GetData(m, &storageGetDataServer{stream})
}

type Storage_GetDataServer interface {
	Send(*ServicePair) error
	grpc.ServerStream
}

type storageGetDataServer struct {
	grpc.ServerStream
}

func (x *storageGetDataServer) Send(m *ServicePair) error {
	return x.ServerStream.SendMsg(m)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Storage_Store_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetServices",
			Handler:       _Storage_GetServices_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetServiceData",
			Handler:       _Storage_GetServiceData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMultipleServicesData",
			Handler:       _Storage_GetMultipleServicesData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetData",
			Handler:       _Storage_GetData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v2/storage.proto",
}
