FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/task_linux_amd64/task assistant-task
CMD ["./assistant-task"]
