package grpcserver

import (
	"grpc_file_storage/internal/protoschema"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GRPCServer) DownloadImage(req *protoschema.DownloadImageRequest, resp protoschema.ImageStorage_DownloadImageServer) error {
	g.downloadUploadParallelOPLimit <- struct{}{}
	defer func() {
		<-g.downloadUploadParallelOPLimit
	}()

	fileExists, err := g.fileInfoStorage.IsFileExists(g.ctx, req.Filename)
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot check file existing: %q", err.Error())
	}
	if !fileExists {
		return status.Errorf(codes.InvalidArgument, "file %q is not exists", req.Filename)
	}

	fileReader, err := g.fileStorage.Download(g.ctx, req.Filename)
	if err != nil {
		return status.Errorf(codes.Unknown, "read file from storage: %q", err.Error())
	}
	defer fileReader.Close()

	chunk := make([]byte, 1024)
	for {
		n, err := fileReader.Read(chunk)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return status.Errorf(codes.Unknown, "send download image response error: %q", err.Error())
		}
		if n < 1024 {
			chunk = chunk[:n]
		}
		err = resp.Send(&protoschema.DownloadImageResponse{
			ChunkData: chunk,
		})
		if err != nil {
			return status.Errorf(codes.Unknown, "send download image response error: %q", err.Error())
		}
	}
}
