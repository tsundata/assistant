FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/message-linux-amd64 app
ADD ./configs/message.yml .
ADD ./scripts/wait-for .
