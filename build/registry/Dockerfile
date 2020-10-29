FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/registry-linux-amd64 app
ADD ./configs/registry.yml .
ADD ./scripts/wait-for .
