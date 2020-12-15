FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/web-linux-amd64 app
ADD ./configs/web.yml .
ADD ./scripts/wait-for .
