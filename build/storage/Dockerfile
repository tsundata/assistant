FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/storage-linux-amd64/storage app
CMD ["./app"]
