FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/task_linux_amd64_v1/task assistant-task
CMD ["./assistant-task"]
