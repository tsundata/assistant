FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/worker-linux-amd64/worker app
CMD ["./app"]
