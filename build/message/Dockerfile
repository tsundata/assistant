FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/message-linux-amd64/message app
CMD ["./app"]
