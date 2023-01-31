package grpcserver

import (
	"context"
	"grpc_file_storage/internal/fileinfo"
	"grpc_file_storage/internal/filestorage"
)

type GRPCServer struct {
	ctx             context.Context
	fileInfoStorage fileinfo.FileInfoStorage
	fileStorage     filestorage.FileStorage
}

func NewGRPCServer(
	ctx context.Context,
	fileInfoStorage fileinfo.FileInfoStorage,
	fileStorage filestorage.FileStorage,
) *GRPCServer {
	return &GRPCServer{
		ctx:             ctx,
		fileInfoStorage: fileInfoStorage,
		fileStorage:     fileStorage,
	}
}
