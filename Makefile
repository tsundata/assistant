apps = 'gateway' 'message' 'subscribe' 'web' 'middle' 'spider' 'cron'
agents = 'server'
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
	protoc -I api/pb ./api/pb/* --gogo_out=plugins=grpc:.
all: lint build
