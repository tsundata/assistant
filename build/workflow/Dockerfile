FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/workflow-linux-amd64/workflow app
CMD ["./app"]
