package fileinfo

import (
	"context"
	"time"
)

type FileInfoStorage interface {
	IsFileExists(ctx context.Context, fileName string) (bool, error)
	InsertFileInfo(ctx context.Context, fileName string, currentTime time.Time) error
	UpdateFileModifiedAt(ctx context.Context, fileName string, currentTime time.Time) error
	// TODO: add pagination (if count of files would be appropriate for this solution)
	ListFileInfo(ctx context.Context) ([]FileInfo, error)
}

type FileInfo struct {
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
