// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: internal/protoschema/server.proto

package protoschema

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

// ImageStorageClient is the client API for ImageStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageStorageClient interface {
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (ImageStorage_UploadImageClient, error)
	DownloadImage(ctx context.Context, in *DownloadImageRequest, opts ...grpc.CallOption) (ImageStorage_DownloadImageClient, error)
	ImageInfoList(ctx context.Context, in *ImageInfoListRequest, opts ...grpc.CallOption) (ImageStorage_ImageInfoListClient, error)
}

type imageStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewImageStorageClient(cc grpc.ClientConnInterface) ImageStorageClient {
	return &imageStorageClient{cc}
}

func (c *imageStorageClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (ImageStorage_UploadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageStorage_ServiceDesc.Streams[0], "/ImageStorage/UploadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageStorageUploadImageClient{stream}
	return x, nil
}

type ImageStorage_UploadImageClient interface {
	Send(*UploadImageRequest) error
	CloseAndRecv() (*UploadImageResponse, error)
	grpc.ClientStream
}

type imageStorageUploadImageClient struct {
	grpc.ClientStream
}

func (x *imageStorageUploadImageClient) Send(m *UploadImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageStorageUploadImageClient) CloseAndRecv() (*UploadImageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageStorageClient) DownloadImage(ctx context.Context, in *DownloadImageRequest, opts ...grpc.CallOption) (ImageStorage_DownloadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageStorage_ServiceDesc.Streams[1], "/ImageStorage/DownloadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageStorageDownloadImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImageStorage_DownloadImageClient interface {
	Recv() (*DownloadImageResponse, error)
	grpc.ClientStream
}

type imageStorageDownloadImageClient struct {
	grpc.ClientStream
}

func (x *imageStorageDownloadImageClient) Recv() (*DownloadImageResponse, error) {
	m := new(DownloadImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageStorageClient) ImageInfoList(ctx context.Context, in *ImageInfoListRequest, opts ...grpc.CallOption) (ImageStorage_ImageInfoListClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageStorage_ServiceDesc.Streams[2], "/ImageStorage/ImageInfoList", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageStorageImageInfoListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImageStorage_ImageInfoListClient interface {
	Recv() (*ImageInfoListResponse, error)
	grpc.ClientStream
}

type imageStorageImageInfoListClient struct {
	grpc.ClientStream
}

func (x *imageStorageImageInfoListClient) Recv() (*ImageInfoListResponse, error) {
	m := new(ImageInfoListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageStorageServer is the server API for ImageStorage service.
// All implementations should embed UnimplementedImageStorageServer
// for forward compatibility
type ImageStorageServer interface {
	UploadImage(ImageStorage_UploadImageServer) error
	DownloadImage(*DownloadImageRequest, ImageStorage_DownloadImageServer) error
	ImageInfoList(*ImageInfoListRequest, ImageStorage_ImageInfoListServer) error
}

// UnimplementedImageStorageServer should be embedded to have forward compatible implementations.
type UnimplementedImageStorageServer struct {
}

func (UnimplementedImageStorageServer) UploadImage(ImageStorage_UploadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedImageStorageServer) DownloadImage(*DownloadImageRequest, ImageStorage_DownloadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadImage not implemented")
}
func (UnimplementedImageStorageServer) ImageInfoList(*ImageInfoListRequest, ImageStorage_ImageInfoListServer) error {
	return status.Errorf(codes.Unimplemented, "method ImageInfoList not implemented")
}

// UnsafeImageStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageStorageServer will
// result in compilation errors.
type UnsafeImageStorageServer interface {
	mustEmbedUnimplementedImageStorageServer()
}

func RegisterImageStorageServer(s grpc.ServiceRegistrar, srv ImageStorageServer) {
	s.RegisterService(&ImageStorage_ServiceDesc, srv)
}

func _ImageStorage_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageStorageServer).UploadImage(&imageStorageUploadImageServer{stream})
}

type ImageStorage_UploadImageServer interface {
	SendAndClose(*UploadImageResponse) error
	Recv() (*UploadImageRequest, error)
	grpc.ServerStream
}

type imageStorageUploadImageServer struct {
	grpc.ServerStream
}

func (x *imageStorageUploadImageServer) SendAndClose(m *UploadImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageStorageUploadImageServer) Recv() (*UploadImageRequest, error) {
	m := new(UploadImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ImageStorage_DownloadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadImageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageStorageServer).DownloadImage(m, &imageStorageDownloadImageServer{stream})
}

type ImageStorage_DownloadImageServer interface {
	Send(*DownloadImageResponse) error
	grpc.ServerStream
}

type imageStorageDownloadImageServer struct {
	grpc.ServerStream
}

func (x *imageStorageDownloadImageServer) Send(m *DownloadImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ImageStorage_ImageInfoList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ImageInfoListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageStorageServer).ImageInfoList(m, &imageStorageImageInfoListServer{stream})
}

type ImageStorage_ImageInfoListServer interface {
	Send(*ImageInfoListResponse) error
	grpc.ServerStream
}

type imageStorageImageInfoListServer struct {
	grpc.ServerStream
}

func (x *imageStorageImageInfoListServer) Send(m *ImageInfoListResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ImageStorage_ServiceDesc is the grpc.ServiceDesc for ImageStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ImageStorage",
	HandlerType: (*ImageStorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _ImageStorage_UploadImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadImage",
			Handler:       _ImageStorage_DownloadImage_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ImageInfoList",
			Handler:       _ImageStorage_ImageInfoList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/protoschema/server.proto",
}