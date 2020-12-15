FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/middle-linux-amd64 app
ADD ./configs/middle.yml .
ADD ./scripts/wait-for .
