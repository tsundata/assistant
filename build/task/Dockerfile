FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/task-linux-amd64/task app
CMD ["./app"]
