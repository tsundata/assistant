apps = 'gateway' 'message' 'subscribe' 'web' 'middle' 'spider' 'cron' 'workflow'
agents = 'server' 'redis'

.PHONY: build
build:
	for app in $(apps) ;\
	do \
		GOOS=linux GOARCH="amd64" go build -o dist/$$app-linux-amd64 ./cmd/$$app/; \
	done
	for agent in $(agents) ;\
	do \
	    GOOS=linux GOARCH="amd64" go build -o dist/$$agent-agent-linux-amd64 ./cmd/agent/$$agent/; \
    done

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: docker
docker-compose: build
	docker-compose -f deployments/docker-compose.yml up --build -d

.PHONY: proto
proto:
	rm api/pb/*.pb.go
	protoc -I api/pb ./api/pb/* --gogo_out=plugins=grpc:.
	git add api/pb/*.pb.go
all: lint build

.PHONY: release
release:
	goreleaser release --snapshot --rm-dist
