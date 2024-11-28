FROM golang:1.23-alpine AS builder

WORKDIR /stream-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /stream-service/main .
COPY --from=builder /stream-service/.env .env

CMD ["./main"]
