-- +goose Up
-- +goose StatementBegin

-- Создаем таблицу files с колонкой id как первичным ключом
CREATE TABLE IF NOT EXISTS files (
                                     id SERIAL PRIMARY KEY,
                                     filename TEXT NOT NULL UNIQUE,
                                     total_chunks INTEGER NOT NULL,
                                     total_size BIGINT NOT NULL
);

-- Создаем таблицу chunks, связывая её с таблицей files через file_id
CREATE TABLE IF NOT EXISTS chunks (
                                      id SERIAL PRIMARY KEY,
                                      file_id INTEGER NOT NULL,
                                      chunk_number INTEGER NOT NULL,
                                      service_name TEXT NOT NULL,
                                      chunk_size BIGINT NOT NULL,
                                      chunk_hash TEXT NOT NULL,
                                      FOREIGN KEY (file_id) REFERENCES files (id) ON DELETE CASCADE
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Удаляем таблицы chunks и files
DROP TABLE IF EXISTS chunks;
DROP TABLE IF EXISTS files;

-- +goose StatementEnd
