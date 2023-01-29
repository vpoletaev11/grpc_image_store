package fileinfo

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	createFileInfoTable = `CREATE TABLE IF NOT EXISTS file_info(
	name VARCHAR(150) PRIMARY KEY,
	created_at INTEGER,
	modified_at INTEGER
);`
	getFileCount = `SELECT COUNT(*) FROM file_info
	WHERE name = $1;`
	insertFileInfo = `INSERT INTO file_info (name, created_at, modified_at)
	VALUES ($1, $2, $2);`
	updateFileModifiedAt = `UPDATE file_info SET modified_at = $1
	WHERE name = $2;`
	deleteFileInfo  = `DELETE FROM file_info WHERE name = $1;`
	getFileInfoList = `SELECT name, created_at, modified_at FROM file_info
	ORDER BY name;`
)

type FileInfoStorageSqlite struct {
	db *sqlx.DB
}

func NewFileInfoStorageSqlite(ctx context.Context, db *sqlx.DB) (*FileInfoStorageSqlite, error) {
	_, err := db.ExecContext(ctx, createFileInfoTable)
	if err != nil {
		return nil, fmt.Errorf("create file_info table error: %w", err)
	}

	return &FileInfoStorageSqlite{db: db}, nil
}

func (f *FileInfoStorageSqlite) IsFileExists(ctx context.Context, fileName string) (bool, error) {
	cnt := 0
	err := f.db.GetContext(ctx, &cnt, getFileCount, fileName)
	if err != nil {
		return false, fmt.Errorf("get file count error: %w", err)
	}
	if cnt < 0 || cnt > 1 {
		return false, fmt.Errorf("get file count error: 'incorrect count of files [%d], expected 0 or 1'", cnt)
	}

	if cnt == 1 {
		return true, nil
	}
	return false, nil
}

func (f *FileInfoStorageSqlite) InsertFileInfo(ctx context.Context, fileName string, currentTime time.Time) error {
	_, err := f.db.ExecContext(ctx, insertFileInfo, fileName, currentTime.Unix())
	if err != nil {
		return fmt.Errorf("insert file info error: %w", err)
	}

	return nil
}

func (f *FileInfoStorageSqlite) UpdateFileModifiedAt(ctx context.Context, fileName string, currentTime time.Time) error {
	_, err := f.db.ExecContext(ctx, updateFileModifiedAt, currentTime.Unix(), fileName)
	if err != nil {
		return fmt.Errorf("update file info modified_at column error: %w", err)
	}

	return nil
}

func (f *FileInfoStorageSqlite) DeleteFileInfo(ctx context.Context, fileName string) error {
	_, err := f.db.ExecContext(ctx, deleteFileInfo, fileName)
	if err != nil {
		return fmt.Errorf("update file info modified_at column error: %w", err)
	}

	return nil
}

func (f *FileInfoStorageSqlite) ListFileInfo(ctx context.Context) ([]FileInfo, error) {
	fileInfoList := []FileInfo{}

	rows, err := f.db.QueryContext(ctx, getFileInfoList)
	if err != nil {
		return nil, fmt.Errorf("select file info list: %w", err)
	}

	fileInfo := FileInfo{}
	var createdAt, modifiedAt int64
	for rows.Next() {
		err = rows.Scan(&fileInfo.Name, &createdAt, &modifiedAt)
		if err != nil {
			return nil, fmt.Errorf("scan row error: %w", err)
		}
		fileInfo.CreatedAt = time.Unix(createdAt, 0).UTC()
		fileInfo.ModifiedAt = time.Unix(modifiedAt, 0).UTC()

		fileInfoList = append(fileInfoList, fileInfo)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return fileInfoList, nil
}
