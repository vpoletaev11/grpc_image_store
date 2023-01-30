package filestorage

import (
	"context"
	"io"
)

type FileStorage interface {
	Save(ctx context.Context, filename string, data io.Reader) error
	Download(ctx context.Context, filename string) (io.ReadCloser, error)
	Delete(ctx context.Context, filename string) error
	Rename(ctx context.Context, oldName, newName string) error
}
