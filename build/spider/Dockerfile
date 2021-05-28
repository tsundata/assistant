FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/spider_linux_amd64/spider app
CMD ["./app"]
