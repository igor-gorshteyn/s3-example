FROM golang:1.22.6 as builder

WORKDIR /app

COPY . .

RUN go build -o /app/storageService ./cmd/storageService

FROM gcr.io/distroless/base

COPY --from=builder /app/storageService /storageService

CMD ["/storageService"]