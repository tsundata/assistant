FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/storage_linux_amd64/storage app
CMD ["./app"]
