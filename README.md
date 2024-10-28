# S3-подобный сервис хранения файлов

Этот проект реализует S3-подобный сервис для хранения и передачи файлов. Система состоит из сервиса хранения и сервиса передачи файлов, использует gRPC для коммуникации, PostgreSQL для метаданных и Redis для кэширования.

Основные функции: загрузка файлов с разделением на чанки, скачивание файлов, хранение метаданных в PostgreSQL, кэширование в Redis.

Требования: Go 1.22, PostgreSQL, Redis, Docker и Docker Compose.

Установка:
1. git clone [URL репозитория]
2. cd s3-task
3. docker-compose up -d
4. make migrate

API Endpoints:

1. Загрузка файла:
   POST /upload
   curl -X POST -F "file=@/path/to/your/file.txt" http://localhost:8080/upload

2. Скачивание файла:
   GET /download
   curl -O http://localhost:8080/download?filename=example.txt

3. Регистрация клиента:
   POST /register
   curl -X POST http://localhost:8080/register

4. Получение списка клиентов:
   GET /clients
   curl http://localhost:8080/clients

Разработка:
- Сборка: make build
- Тесты: make test
