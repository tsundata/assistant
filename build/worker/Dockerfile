FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/worker_linux_amd64/worker app
CMD ["./app"]
