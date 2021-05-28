FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/spider-linux-amd64/spider app
CMD ["./app"]
