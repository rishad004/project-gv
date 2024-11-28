FROM golang:1.23-alpine AS builder

WORKDIR /apiGateway

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /apiGateway/main .
COPY --from=builder /apiGateway/.env .env
COPY --from=builder /apiGateway/pkg/templates ./pkg/templates

CMD ["./main"]
