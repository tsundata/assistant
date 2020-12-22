FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/cron-linux-amd64 app
ADD ./configs/cron.yml .
