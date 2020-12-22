FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/gateway-linux-amd64 app
ADD ./configs/gateway.yml .
