package grpcserver

import (
	"context"
	"fmt"
	"grpc_file_storage/internal/filestorage"
	"grpc_file_storage/internal/protoschema"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GRPCServer) UploadImage(stream protoschema.ImageStorage_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}

	filename := req.GetInfo().Filename
	if filename == "" {
		return status.Errorf(codes.InvalidArgument, "filename is empty")
	}

	fileExists, err := g.fileInfoStorage.IsFileExists(g.ctx, filename)
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot check file existing: %q", err.Error())
	}

	if !fileExists {
		err = saveFileToStorage(g.ctx, g.fileStorage, filename, stream)
		if err != nil {
			return status.Errorf(codes.Unknown, "save file to storage: %q", err.Error())
		}

		err = g.fileInfoStorage.InsertFileInfo(g.ctx, filename, time.Now())
		if err != nil {
			return status.Errorf(codes.Unknown, "save file info to storage: %q", err.Error())
		}

		return nil
	}

	uploadTmpFileName := filename + ".tmp"

	err = saveFileToStorage(g.ctx, g.fileStorage, uploadTmpFileName, stream)
	if err != nil {
		return status.Errorf(codes.Unknown, "save file to storage: %q", err.Error())
	}

	err = g.fileStorage.Delete(g.ctx, filename)
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot delete old file: %q", err.Error())
	}

	err = g.fileStorage.Rename(g.ctx, uploadTmpFileName, filename)
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot rename tmp file: %q", err.Error())
	}

	err = g.fileInfoStorage.UpdateFileModifiedAt(g.ctx, filename, time.Now())
	if err != nil {
		return status.Errorf(codes.Unknown, "update file modified at property: %q", err.Error())
	}

	return nil
}

func saveFileToStorage(
	ctx context.Context,
	fileStorage filestorage.FileStorage,
	fileName string,
	stream protoschema.ImageStorage_UploadImageServer,
) error {
	r, w := io.Pipe()
	defer w.Close()

	fileSaveErrCh := make(chan error)
	go func(errCh chan error) {
		fileSaveErrCh <- fileStorage.Save(ctx, fileName, r)
	}(fileSaveErrCh)

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			err = fileStorage.Delete(ctx, fileName)
			if err != nil {
				log.Printf("delete file %q error: %s", fileName, err.Error())
			}
			return fmt.Errorf("cannot receive chunk data: %w", err)
		}

		_, err = w.Write(chunk.GetChunkData())
		if err != nil {
			err = fileStorage.Delete(ctx, fileName)
			if err != nil {
				log.Printf("delete file %q error: %s", fileName, err.Error())
			}
			return fmt.Errorf("write data to pipe: %w", err)
		}
	}

	err := <-fileSaveErrCh
	if err != nil {
		return fmt.Errorf("storage error: %w", err)
	}

	return nil
}
