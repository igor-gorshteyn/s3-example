# S3-подобный сервис хранения файлов

Этот проект реализует S3-подобный сервис для хранения и передачи файлов. Система состоит из сервиса хранения и сервиса передачи файлов, использует gRPC для коммуникации, PostgreSQL для метаданных и Redis для кэширования.

Основные функции: загрузка файлов с разделением на чанки, скачивание файлов, хранение метаданных в PostgreSQL, кэширование в Redis.

Требования: Go 1.22, PostgreSQL, Redis, Docker и Docker Compose.

Установка:
1. git clone [URL репозитория]
2. cd s3-example
3. docker-compose up -d
4. make migrate

API Endpoints:

1. Загрузка файла:
   POST /upload
   curl -X POST -F "file=@/path/to/your/file.txt" http://localhost:8080/upload

2. Скачивание файла:
   GET /download/{filename}
   curl -O http://localhost:8080/download/example.txt

3. Информация о файле:
   GET /info/{filename}
   curl http://localhost:8080/info/example.txt

4. Удаление файла:
   DELETE /delete/{filename}
   curl -X DELETE http://localhost:8080/delete/example.txt

Разработка:
- Сборка: make build
- Тесты: make test
