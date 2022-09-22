# 1. build executable binary
FROM golang:1.16-alpine AS builder
LABEL maintener="Reza Irfan Wijaya<rezairfanwijaya23@gmail.com>"
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o \myapp

EXPOSE 8080

CMD ["./myapp"]

# FROM alpine:3.16.0
# WORKDIR /app
# COPY --from=builder employee-service .
# COPY .env .
# EXPOSE 8080
# CMD ["./myapp", "-m=migrate"]