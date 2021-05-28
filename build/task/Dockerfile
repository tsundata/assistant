FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/task-linux-amd64/task app
CMD ["./app"]
