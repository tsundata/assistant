FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/todo_linux_amd64/todo app
CMD ["./app"]
