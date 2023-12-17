# side-hub-be
side hub backend

Generate [Swagger](https://github.com/swaggo/swag#declarative-comments-format) Docs

```
    swag init -g cmd/main.go
```

Test
```
    go test ./...
```

Build
```
    go build -o cmd/main.go
```

Docker
```
    docker build -t side-hub-be:dev .
```