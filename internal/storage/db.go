package storage

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

type FileMetadata struct {
	ID          int64
	Filename    string
	TotalChunks int32
	TotalSize   int64
}

type ChunkMetadata struct {
	ID          int64
	FileID      int64
	ChunkNumber int32
	ServiceName string
	ChunkSize   int64
	ChunkHash   string
}

type Manager struct {
	DB *sql.DB
	mu sync.Mutex
}

func NewDBManager(pgHost, pgPort, pgUser, pgPassword, pgDBName string) (*Manager, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pgHost, pgPort, pgUser, pgPassword, pgDBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	manager := &Manager{
		DB: db,
	}

	return manager, nil
}

func (m *Manager) CreateFileMetadata(filename string, totalChunks int32, totalSize int64) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var fileID int64
	query := `INSERT INTO files (filename, total_chunks, total_size)
              VALUES ($1, $2, $3)
              RETURNING id;`
	err := m.DB.QueryRow(query, filename, totalChunks, totalSize).Scan(&fileID)
	if err != nil {
		return 0, err
	}
	return fileID, nil
}

func (m *Manager) UpdateFileMetadata(fileID int64, totalChunks int32, totalSize int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	query := `UPDATE files SET total_chunks = $1, total_size = $2 WHERE id = $3;`
	_, err := m.DB.Exec(query, totalChunks, totalSize, fileID)
	return err
}

func (m *Manager) GetFileID(filename string) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var fileID int64
	query := `SELECT id FROM files WHERE filename = $1;`
	err := m.DB.QueryRow(query, filename).Scan(&fileID)
	if err != nil {
		return 0, err
	}
	return fileID, nil
}

func (m *Manager) SaveChunkMetadata(metadata ChunkMetadata) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	query := `INSERT INTO chunks (file_id, chunk_number, service_name, chunk_size, chunk_hash)
              VALUES ($1, $2, $3, $4, $5);`
	_, err := m.DB.Exec(query, metadata.FileID, metadata.ChunkNumber, metadata.ServiceName, metadata.ChunkSize, metadata.ChunkHash)
	return err
}

func (m *Manager) GetChunkMetadata(fileID int64) ([]ChunkMetadata, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	query := `SELECT id, chunk_number, service_name, chunk_size, chunk_hash FROM chunks WHERE file_id = $1 ORDER BY chunk_number ASC;`
	rows, err := m.DB.Query(query, fileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metadataList []ChunkMetadata
	for rows.Next() {
		var metadata ChunkMetadata
		metadata.FileID = fileID
		err := rows.Scan(&metadata.ID, &metadata.ChunkNumber, &metadata.ServiceName, &metadata.ChunkSize, &metadata.ChunkHash)
		if err != nil {
			return nil, err
		}
		metadataList = append(metadataList, metadata)
	}

	return metadataList, nil
}

func (m *Manager) GetFileMetadata(filename string) (*FileMetadata, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	query := `SELECT id, total_chunks, total_size FROM files WHERE filename = $1;`
	row := m.DB.QueryRow(query, filename)

	var metadata FileMetadata
	metadata.Filename = filename
	err := row.Scan(&metadata.ID, &metadata.TotalChunks, &metadata.TotalSize)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

func (m *Manager) DeleteFileMetadata(fileID int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	queryChunks := `DELETE FROM chunks WHERE file_id = $1;`
	queryFiles := `DELETE FROM files WHERE id = $1;`

	_, err := m.DB.Exec(queryChunks, fileID)
	if err != nil {
		return err
	}
	_, err = m.DB.Exec(queryFiles, fileID)
	return err
}

func (m *Manager) Close() error {
	return m.DB.Close()
}
