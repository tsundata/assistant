FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/message-linux-amd64/message app
CMD ["./app"]
