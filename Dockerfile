FROM golang:1.20.3-alpine AS builder

COPY . github.com/Shemistan/grpc_user_api
WORKDIR github.com/Shemistan/grpc_user_api

RUN go mod download
RUN go build -o ./bin/user_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /go/github.com/Shemistan/grpc_user_api/bin/user_server .
COPY --from=builder /go/github.com/Shemistan/grpc_user_api/prod.env .

CMD ["./user_server"]
