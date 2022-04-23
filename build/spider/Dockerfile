FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/spider_linux_amd64_v1/spider assistant-spider
CMD ["./assistant-spider"]
