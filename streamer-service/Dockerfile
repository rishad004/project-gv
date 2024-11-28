FROM golang:1.23-alpine AS builder

WORKDIR /streamer-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /streamer-service/main .
COPY --from=builder /streamer-service/.env .env

CMD ["./main"]
