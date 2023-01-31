package grpcserver

import (
	"grpc_file_storage/internal/protoschema"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GRPCServer) ImageInfoList(req *protoschema.ImageInfoListRequest, resp protoschema.ImageStorage_ImageInfoListServer) error {
	fileInfoList, err := g.fileInfoStorage.ListFileInfo(g.ctx)
	if err != nil {
		return status.Errorf(codes.Unknown, "get file info list error: %q", err.Error())
	}

	for _, file := range fileInfoList {
		err = resp.Send(&protoschema.ImageInfoListResponse{
			Filename:   file.Name,
			CreatedAt:  file.CreatedAt.Unix(),
			ModifiedAt: file.ModifiedAt.Unix(),
		})
		if err != nil {
			return status.Errorf(codes.Unknown, "send image info resp: %q", err.Error())
		}
	}

	return nil
}
