FROM golang:1.22.6 as builder

WORKDIR /app

COPY . .

RUN go build -o /app/transferService ./cmd/transferService

FROM gcr.io/distroless/base

COPY --from=builder /app/transferService /transferService
COPY --from=builder /app/migrations /migrations

CMD ["/transferService"]