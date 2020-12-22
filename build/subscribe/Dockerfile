FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/subscribe-linux-amd64 app
ADD ./configs/subscribe.yml .
