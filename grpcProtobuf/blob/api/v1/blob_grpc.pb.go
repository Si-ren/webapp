// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: blob.proto

package blobpb

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

// BlobServiceClient is the client API for BlobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlobServiceClient interface {
	CreateBolb(ctx context.Context, in *CreateBolbRequest, opts ...grpc.CallOption) (*CreateBolbResponse, error)
	GetBolb(ctx context.Context, in *GetBolbRequest, opts ...grpc.CallOption) (*GetBolbResponse, error)
	GetBolbUrl(ctx context.Context, in *GetBolbUrlRequest, opts ...grpc.CallOption) (*GetBolbUrlResponse, error)
}

type blobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlobServiceClient(cc grpc.ClientConnInterface) BlobServiceClient {
	return &blobServiceClient{cc}
}

func (c *blobServiceClient) CreateBolb(ctx context.Context, in *CreateBolbRequest, opts ...grpc.CallOption) (*CreateBolbResponse, error) {
	out := new(CreateBolbResponse)
	err := c.cc.Invoke(ctx, "/blob.v1.BlobService/CreateBolb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobServiceClient) GetBolb(ctx context.Context, in *GetBolbRequest, opts ...grpc.CallOption) (*GetBolbResponse, error) {
	out := new(GetBolbResponse)
	err := c.cc.Invoke(ctx, "/blob.v1.BlobService/GetBolb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobServiceClient) GetBolbUrl(ctx context.Context, in *GetBolbUrlRequest, opts ...grpc.CallOption) (*GetBolbUrlResponse, error) {
	out := new(GetBolbUrlResponse)
	err := c.cc.Invoke(ctx, "/blob.v1.BlobService/GetBolbUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlobServiceServer is the server API for BlobService service.
// All implementations should embed UnimplementedBlobServiceServer
// for forward compatibility
type BlobServiceServer interface {
	CreateBolb(context.Context, *CreateBolbRequest) (*CreateBolbResponse, error)
	GetBolb(context.Context, *GetBolbRequest) (*GetBolbResponse, error)
	GetBolbUrl(context.Context, *GetBolbUrlRequest) (*GetBolbUrlResponse, error)
}

// UnimplementedBlobServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBlobServiceServer struct {
}

func (UnimplementedBlobServiceServer) CreateBolb(context.Context, *CreateBolbRequest) (*CreateBolbResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBolb not implemented")
}
func (UnimplementedBlobServiceServer) GetBolb(context.Context, *GetBolbRequest) (*GetBolbResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBolb not implemented")
}
func (UnimplementedBlobServiceServer) GetBolbUrl(context.Context, *GetBolbUrlRequest) (*GetBolbUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBolbUrl not implemented")
}

// UnsafeBlobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlobServiceServer will
// result in compilation errors.
type UnsafeBlobServiceServer interface {
	mustEmbedUnimplementedBlobServiceServer()
}

func RegisterBlobServiceServer(s grpc.ServiceRegistrar, srv BlobServiceServer) {
	s.RegisterService(&BlobService_ServiceDesc, srv)
}

func _BlobService_CreateBolb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBolbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).CreateBolb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blob.v1.BlobService/CreateBolb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).CreateBolb(ctx, req.(*CreateBolbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobService_GetBolb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBolbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).GetBolb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blob.v1.BlobService/GetBolb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).GetBolb(ctx, req.(*GetBolbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobService_GetBolbUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBolbUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobServiceServer).GetBolbUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blob.v1.BlobService/GetBolbUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobServiceServer).GetBolbUrl(ctx, req.(*GetBolbUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlobService_ServiceDesc is the grpc.ServiceDesc for BlobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blob.v1.BlobService",
	HandlerType: (*BlobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBolb",
			Handler:    _BlobService_CreateBolb_Handler,
		},
		{
			MethodName: "GetBolb",
			Handler:    _BlobService_GetBolb_Handler,
		},
		{
			MethodName: "GetBolbUrl",
			Handler:    _BlobService_GetBolbUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blob.proto",
}
