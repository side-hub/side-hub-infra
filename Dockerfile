# syntax=docker/dockerfile:1

## Build
FROM golang:1.20 AS build

WORKDIR /app

# Dependency download caching
COPY go.mod go.sum ./
RUN go mod download

# Build 
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /app/cmd/main /app/cmd/main.go

## Deploy
FROM alpine:3.19

# Deploy Application
WORKDIR /app

RUN mkdir /app/log

COPY --from=build /app/cmd/main /app/main

ENTRYPOINT ["./main"]