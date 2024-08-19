# syntax=docker/dockerfile:1.4
FROM --platform=linux/amd64 golang:1.22.4-bookworm

ARG BIN_NAME

WORKDIR /app

COPY _bin/${BIN_NAME} /app/grpc_proxy
COPY _bin/*.so /app/

CMD ["/app/grpc_proxy"]
