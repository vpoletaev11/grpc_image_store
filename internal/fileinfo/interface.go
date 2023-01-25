package fileinfo

import (
	"context"
	"time"
)

type FileInfoStorage interface {
	IsFileExists(ctx context.Context, fileName string) bool
	InsertFileInfo(ctx context.Context, fileName string, currentTime time.Time) error
	UpdateFileInfo(ctx context.Context, fileName string, currentTime time.Time) error
	// TODO: add pagination (if count of files would be appropriate for this solution)
	ListFileInfo(ctx context.Context) ([]FileInfo, error)
}

type FileInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	FileName  string
}
