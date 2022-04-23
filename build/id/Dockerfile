FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/id_linux_amd64_v1/id assistant-id
CMD ["./assistant-id"]
