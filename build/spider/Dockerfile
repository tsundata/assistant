FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/spider-linux-amd64/spider app
CMD ["./app"]
