package main

import (
	"context"
	"fmt"
	"grpc_file_storage/internal/fileinfo"
	"grpc_file_storage/internal/filestorage"
	"grpc_file_storage/internal/grpcserver"
	"grpc_file_storage/internal/protoschema"
	"net"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

// TODO: add configuration consumption from ENV.
const (
	dbPath                        = "test.db"
	fileStoragePath               = "testdata/images"
	grpcPort                      = 8080
	downloadUploadParallelOPLimit = 10
	listFileInfoParallelOPLimit   = 100
)

func main() {
	// TODO: add logging.
	// TODO: add graceful shutdown (and correct operations cancellations).
	// TODO: prevent more than one upload operation with the same file name at the same time.
	ctx := context.Background()

	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		panic(fmt.Errorf("open database error: %w", err))
	}
	defer db.Close()

	fileInfoStorage, err := fileinfo.NewFileInfoStorageSqlite(ctx, db)
	if err != nil {
		panic(fmt.Errorf("create file info storage error: %w", err))
	}
	fileStorage, err := filestorage.NewFileStorageFS(fileStoragePath)
	if err != nil {
		panic(fmt.Errorf("create file storage error: %w", err))
	}

	server := grpcserver.NewGRPCServer(
		ctx,
		fileInfoStorage,
		fileStorage,
		downloadUploadParallelOPLimit,
		listFileInfoParallelOPLimit,
	)

	fmt.Printf("Server is started on %d port\n", grpcPort)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protoschema.RegisterImageStorageServer(s, server)
	if err := s.Serve(listener); err != nil {
		panic(fmt.Errorf("failed to serve: %w", err))
	}
}
