FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/task_linux_amd64/task app
CMD ["./app"]
