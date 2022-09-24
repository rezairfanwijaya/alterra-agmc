# 1. build executable binary
FROM golang:1.17-alpine AS builder
LABEL maintener="Reza Irfan Wijaya<rezairfanwijaya23@gmail.com>"
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build Day7/main.go
EXPOSE 8080
CMD ["main.exe"]