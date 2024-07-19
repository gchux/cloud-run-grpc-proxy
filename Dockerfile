# syntax=docker/dockerfile:1.4

FROM golang:1.22.4-bookworm AS builder

WORKDIR /app

ADD cmd /app/cmd
ADD pkg /app/pkg
COPY ./go.mod /app
COPY ./go.sum /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -v -o /app/grpc_proxy /app/cmd/grpc_proxy.go

FROM scratch AS releaser
COPY --link --from=builder /app/grpc_proxy /

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/grpc_proxy /grpc_proxy
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER proxy
CMD ["/grpc_proxy"]
