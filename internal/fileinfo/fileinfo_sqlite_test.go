package fileinfo_test

import (
	"context"
	"grpc_file_storage/internal/fileinfo"
	"os"
	"testing"
	"time"

	sqlxmock "github.com/zhashkevych/go-sqlxmock"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testDBPath = "../../testdata/test.db"
)

func TestFileInfoStorageSqliteIntegratedSuccess(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Open("sqlite3", testDBPath)
	defer db.Close()
	require.NoError(t, err)

	fileInfoStorage, err := fileinfo.NewFileInfoStorageSqlite(ctx, db)
	require.NoError(t, err)
	defer os.Remove(testDBPath)

	err = fileInfoStorage.InsertFileInfo(ctx, "file1", time.Date(2001, time.January, 1, 1, 1, 1, 1, time.UTC))
	require.NoError(t, err)
	err = fileInfoStorage.InsertFileInfo(ctx, "file2", time.Date(2002, time.February, 2, 2, 2, 2, 2, time.UTC))
	require.NoError(t, err)
	err = fileInfoStorage.InsertFileInfo(ctx, "file3", time.Date(2003, time.March, 3, 3, 3, 3, 3, time.UTC))
	require.NoError(t, err)

	exists, err := fileInfoStorage.IsFileExists(ctx, "file1")
	assert.NoError(t, err)
	assert.True(t, exists)

	fileList1, err := fileInfoStorage.ListFileInfo(ctx)
	assert.NoError(t, err)
	assert.Equal(t, []fileinfo.FileInfo{
		{
			Name:       "file1",
			CreatedAt:  time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
			ModifiedAt: time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			Name:       "file2",
			CreatedAt:  time.Date(2002, time.February, 2, 2, 2, 2, 0, time.UTC),
			ModifiedAt: time.Date(2002, time.February, 2, 2, 2, 2, 0, time.UTC),
		},
		{
			Name:       "file3",
			CreatedAt:  time.Date(2003, time.March, 3, 3, 3, 3, 0, time.UTC),
			ModifiedAt: time.Date(2003, time.March, 3, 3, 3, 3, 0, time.UTC),
		},
	}, fileList1)

	err = fileInfoStorage.DeleteFileInfo(ctx, "file3")
	assert.NoError(t, err)

	exists, err = fileInfoStorage.IsFileExists(ctx, "file3")
	assert.NoError(t, err)
	assert.False(t, exists)

	err = fileInfoStorage.UpdateFileModifiedAt(ctx, "file2", time.Date(2020, time.December, 20, 20, 20, 20, 0, time.UTC))
	assert.NoError(t, err)

	fileList2, err := fileInfoStorage.ListFileInfo(ctx)
	assert.NoError(t, err)
	assert.Equal(t, []fileinfo.FileInfo{
		{
			Name:       "file1",
			CreatedAt:  time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
			ModifiedAt: time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			Name:       "file2",
			CreatedAt:  time.Date(2002, time.February, 2, 2, 2, 2, 0, time.UTC),
			ModifiedAt: time.Date(2020, time.December, 20, 20, 20, 20, 0, time.UTC),
		},
	}, fileList2)
}

func TestNewFileInfoStorageSqliteSuccess(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlxmock.Newx()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS file_info`).WillReturnResult(sqlxmock.NewResult(1, 1))

	_, err = fileinfo.NewFileInfoStorageSqlite(ctx, db)
	assert.NoError(t, err)
}

func TestFileInfoStorageSqliteInsert(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlxmock.Newx()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS file_info`).WillReturnResult(sqlxmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO file_info`).WithArgs("file1", 978310861).WillReturnResult(sqlxmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO file_info`).WithArgs("file2", 1012615322).WillReturnResult(sqlxmock.NewResult(1, 1))

	fileInfoStorage, err := fileinfo.NewFileInfoStorageSqlite(ctx, db)
	require.NoError(t, err)

	err = fileInfoStorage.InsertFileInfo(ctx, "file1", time.Date(2001, time.January, 1, 1, 1, 1, 1, time.UTC))
	assert.NoError(t, err)
	err = fileInfoStorage.InsertFileInfo(ctx, "file2", time.Date(2002, time.February, 2, 2, 2, 2, 2, time.UTC))
	assert.NoError(t, err)
}

func TestFileInfoStorageSqliteUpdate(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlxmock.Newx()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS file_info`).WillReturnResult(sqlxmock.NewResult(1, 1))
	mock.ExpectExec(`UPDATE file_info SET modified_at`).WithArgs(978310861, "file1").WillReturnResult(sqlxmock.NewResult(1, 1))
	mock.ExpectExec(`UPDATE file_info SET modified_at`).WithArgs(1012615322, "file2").WillReturnResult(sqlxmock.NewResult(1, 1))

	fileInfoStorage, err := fileinfo.NewFileInfoStorageSqlite(ctx, db)
	require.NoError(t, err)

	err = fileInfoStorage.UpdateFileModifiedAt(ctx, "file1", time.Date(2001, time.January, 1, 1, 1, 1, 1, time.UTC))
	assert.NoError(t, err)
	err = fileInfoStorage.UpdateFileModifiedAt(ctx, "file2", time.Date(2002, time.February, 2, 2, 2, 2, 2, time.UTC))
	assert.NoError(t, err)
}

func TestFileInfoStorageSqliteIsExists(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlxmock.Newx()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS file_info`).WillReturnResult(sqlxmock.NewResult(1, 1))
	mock.ExpectQuery(`SELECT COUNT`).WithArgs("file1").WillReturnRows(sqlxmock.NewRows([]string{"count"}).AddRow(1))
	mock.ExpectQuery(`SELECT COUNT`).WithArgs("file2").WillReturnRows(sqlxmock.NewRows([]string{"count"}).AddRow(0))

	fileInfoStorage, err := fileinfo.NewFileInfoStorageSqlite(ctx, db)
	require.NoError(t, err)

	f1Exists, err := fileInfoStorage.IsFileExists(ctx, "file1")
	assert.NoError(t, err)
	assert.True(t, f1Exists)

	f1Exists, err = fileInfoStorage.IsFileExists(ctx, "file2")
	assert.NoError(t, err)
	assert.False(t, f1Exists)
}

func TestFileInfoStorageSqliteDelete(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlxmock.Newx()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS file_info`).WillReturnResult(sqlxmock.NewResult(1, 1))
	mock.ExpectExec(`DELETE FROM file_info`).WithArgs("file1").WillReturnResult(sqlxmock.NewResult(1, 1))

	fileInfoStorage, err := fileinfo.NewFileInfoStorageSqlite(ctx, db)
	require.NoError(t, err)

	err = fileInfoStorage.DeleteFileInfo(ctx, "file1")
	assert.NoError(t, err)
}

func TestFileInfoStorageSqliteInfoList(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlxmock.Newx()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS file_info`).WillReturnResult(sqlxmock.NewResult(1, 1))
	mock.ExpectQuery(`SELECT name, created_at, modified_at FROM file_info`).WillReturnRows(
		sqlxmock.NewRows([]string{"name", "created_at", "modified_at"}).
			AddRow("file1", 978310861, 978310861).
			AddRow("file2", 1012615322, 1012615322),
	)

	fileInfoStorage, err := fileinfo.NewFileInfoStorageSqlite(ctx, db)
	require.NoError(t, err)

	fileList, err := fileInfoStorage.ListFileInfo(ctx)
	assert.NoError(t, err)
	assert.Equal(t, []fileinfo.FileInfo{
		{
			Name:       "file1",
			CreatedAt:  time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
			ModifiedAt: time.Date(2001, time.January, 1, 1, 1, 1, 0, time.UTC),
		},
		{
			Name:       "file2",
			CreatedAt:  time.Date(2002, time.February, 2, 2, 2, 2, 0, time.UTC),
			ModifiedAt: time.Date(2002, time.February, 2, 2, 2, 2, 0, time.UTC),
		},
	}, fileList)
}
