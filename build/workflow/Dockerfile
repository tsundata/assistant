FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/workflow_linux_amd64/workflow app
CMD ["./app"]
