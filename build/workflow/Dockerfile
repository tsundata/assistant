FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/workflow-linux-amd64/workflow app
CMD ["./app"]
