package filestorage

import (
	"bytes"
	"context"
)

type FileStorage interface {
	Save(ctx context.Context, filename string, data bytes.Buffer) error
	Download(ctx context.Context, filename string) (bytes.Buffer, error)
	Delete(ctx context.Context, filename string) error
	Rename(ctx context.Context, oldName, newName string) error
}
