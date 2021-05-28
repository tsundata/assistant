FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/worker-linux-amd64/worker app
CMD ["./app"]
