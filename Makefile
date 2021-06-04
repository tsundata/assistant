.PHONY: build
build:
	GOVERSION=`go version` goreleaser --snapshot --skip-publish --rm-dist

.PHONY: lint
lint:
	golangci-lint run ./... --timeout=5m

.PHONY: docker
docker-compose: build
	docker-compose -f deployments/docker-compose/docker-compose.yml up --build -d

.PHONY: proto
proto:
	rm api/pb/*.pb.go
	protoc -I api/pb ./api/pb/* --gogo_out=plugins=grpc:.
	protoc-go-inject-tag -input=./api/pb/middle.pb.go
	git add api/pb/*.pb.go
all: lint build

.PHONY: release
release:
	goreleaser release --snapshot --rm-dist

.PHONY: wire
wire:
	wire ./...
