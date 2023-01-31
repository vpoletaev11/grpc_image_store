package grpcserver

import (
	"context"
	"grpc_file_storage/internal/fileinfo"
	"grpc_file_storage/internal/filestorage"
)

type GRPCServer struct {
	ctx                           context.Context
	fileInfoStorage               fileinfo.FileInfoStorage
	fileStorage                   filestorage.FileStorage
	downloadUploadParallelOPLimit chan struct{}
	listFileInfoParallelOPLimit   chan struct{}
}

func NewGRPCServer(
	ctx context.Context,
	fileInfoStorage fileinfo.FileInfoStorage,
	fileStorage filestorage.FileStorage,
	downloadUploadParallelOPLimit int,
	listFileInfoParallelOPLimit int,
) *GRPCServer {
	return &GRPCServer{
		ctx:                           ctx,
		fileInfoStorage:               fileInfoStorage,
		fileStorage:                   fileStorage,
		downloadUploadParallelOPLimit: make(chan struct{}, downloadUploadParallelOPLimit),
		listFileInfoParallelOPLimit:   make(chan struct{}, listFileInfoParallelOPLimit),
	}
}
