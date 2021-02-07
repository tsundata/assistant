GitCommit=`git rev-parse --short HEAD`
BuildTime=`date +'%Y.%m.%d.%H%M%S'`
GoVersion=`go version`
LDFlags=" \
  -X 'github.com/tsundata/assistant/internal/pkg/version.Version=' \
  -X 'github.com/tsundata/assistant/internal/pkg/version.GitCommit=${GitCommit}' \
  -X 'github.com/tsundata/assistant/internal/pkg/version.BuildTime=${BuildTime}' \
  -X 'github.com/tsundata/assistant/internal/pkg/version.GoVersion=${GoVersion}' \
"

GOARCH="amd64" go build -v -ldflags "${LDFlags}" -o dist/gateway-linux-amd64 ./cmd/gateway/

for app in 'gateway' 'message' 'subscribe' 'web' 'middle' 'spider' 'cron' 'workflow' ;\
	do \
		GOOS=linux GOARCH="amd64" go build -v -ldflags "${LDFlags}" -o dist/${app}-linux-amd64 ./cmd/${app}/; \
	done

for agent in 'server' 'redis' ;\
	do \
		GOOS=linux GOARCH="amd64" go build -v -ldflags "${LDFlags}" -o dist/${agent}-linux-amd64 ./cmd/agent/${agent}/; \
	done