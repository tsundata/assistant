FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/workflow_linux_amd64/workflow assistant-workflow
CMD ["./assistant-workflow"]
