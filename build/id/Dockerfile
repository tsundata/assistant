FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/id_linux_amd64/id assistant-id
CMD ["./assistant-id"]
