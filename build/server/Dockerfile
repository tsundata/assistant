FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/server_linux_amd64_v1/server assistant-server
CMD ["./assistant-server"]
