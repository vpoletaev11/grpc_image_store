package filestorage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileStorageFS struct {
	filesPath string
}

func NewFileStorageFS(filesPath string) (*FileStorageFS, error) {
	fileInfo, err := os.Stat(filesPath)
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() {
		return nil, fmt.Errorf("path %q is not dir", filesPath)
	}

	return &FileStorageFS{filesPath: filesPath}, nil
}

func (f *FileStorageFS) Save(ctx context.Context, filename string, data io.Reader) error {
	file, err := os.Create(filepath.Join(f.filesPath, filename))
	if err != nil {
		return fmt.Errorf("create file error: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, data)
	if err != nil {
		return fmt.Errorf("copy data to file error: %w", err)
	}

	return nil
}

func (f *FileStorageFS) Download(ctx context.Context, filename string) (io.ReadCloser, error) {
	file, err := os.Open(filepath.Join(f.filesPath, filename))
	if err != nil {
		return nil, fmt.Errorf("open file error: %w", err)
	}

	return file, nil
}

func (f *FileStorageFS) Delete(ctx context.Context, filename string) error {
	return os.Remove(filepath.Join(f.filesPath, filename))
}

func (f *FileStorageFS) Rename(ctx context.Context, oldName, newName string) error {
	oldName = filepath.Join(f.filesPath, oldName)
	newName = filepath.Join(f.filesPath, newName)

	return os.Rename(oldName, newName)
}
