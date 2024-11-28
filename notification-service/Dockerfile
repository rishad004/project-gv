FROM golang:1.23-alpine AS builder

WORKDIR /notification-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /notification-service/main .
COPY --from=builder /notification-service/.env .env

CMD ["./main"]
