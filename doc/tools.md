# Tools

## git-chglog

```shell
# Install
go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
# Init
git-chglog --init
# Output 
git-chglog -o CHANGELOG.md
```

### gotests

```shell
# Install
go get -u github.com/cweill/gotests/...
# Output
gotests -w -all PATH
```

## go-callvis

```shell
# Install
go get -u github.com/ofabry/go-callvis
# Output
go-callvis ./cmd/app/main.go
```

## golangci-lint

```shell
# Install
go get github.com/golangci/golangci-lint/cmd/golangci-lint
# Run
golangci-lint run
```

## gomock

```shell
# Install
go install github.com/golang/mock/mockgen@v1.6.0
# Output
mockgen -source=./internal/app/todo/repository/todo.go -destination=./mock/todo_repository.go -package=mock

mockgen -source=./api/pb/todo.pb.go -destination=./mock/todo_client.go -package=mock
```
