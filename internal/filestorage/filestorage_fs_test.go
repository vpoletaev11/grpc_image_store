package filestorage_test

import (
	"bytes"
	"context"
	"grpc_file_storage/internal/filestorage"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	datasetFolder = "testdata"
)

var testData = []byte("test data")

func TestSave(t *testing.T) {
	ctx := context.Background()

	storage, err := filestorage.NewFileStorageFS(datasetFolder)
	require.NoError(t, err)

	err = storage.Save(ctx, "test_file_save", bytes.NewBuffer(testData))
	require.NoError(t, err)
	defer os.Remove(filepath.Join(datasetFolder, "test_file_save"))

	require.FileExists(t, filepath.Join(datasetFolder, "test_file_save"))

	f, err := os.Open(filepath.Join(datasetFolder, "test_file_save"))
	require.NoError(t, err)
	defer f.Close()

	fileContent := make([]byte, len(testData))
	_, err = f.Read(fileContent)
	require.NoError(t, err)

	assert.Equal(t, testData, fileContent)
}

func TestDownload(t *testing.T) {
	ctx := context.Background()

	storage, err := filestorage.NewFileStorageFS(datasetFolder)
	require.NoError(t, err)

	reader, err := storage.Download(ctx, "file_with_data")
	require.NoError(t, err)
	defer reader.Close()

	res, err := io.ReadAll(reader)
	require.NoError(t, err)
	assert.Equal(t, testData, res)
}

func TestDelete(t *testing.T) {
	ctx := context.Background()

	fileToDeletePath := filepath.Join(datasetFolder, "file_to_delete")

	storage, err := filestorage.NewFileStorageFS(datasetFolder)
	require.NoError(t, err)

	f, err := os.Create(fileToDeletePath)
	require.NoError(t, err)
	f.Close()
	defer os.Remove(fileToDeletePath)

	err = storage.Delete(ctx, "file_to_delete")
	assert.NoError(t, err)

	assert.NoFileExists(t, fileToDeletePath)
}

func TestRename(t *testing.T) {
	ctx := context.Background()

	fileToRenamePath := filepath.Join(datasetFolder, "file_to_rename")
	renamedFilePath := filepath.Join(datasetFolder, "renamed_file")

	storage, err := filestorage.NewFileStorageFS(datasetFolder)
	require.NoError(t, err)

	f, err := os.Create(fileToRenamePath)
	require.NoError(t, err)
	f.Close()
	defer os.Remove(fileToRenamePath)
	defer os.Remove(renamedFilePath)

	err = storage.Rename(ctx, "file_to_rename", "renamed_file")
	assert.NoError(t, err)

	assert.NoFileExists(t, fileToRenamePath)
	assert.FileExists(t, renamedFilePath)
}
