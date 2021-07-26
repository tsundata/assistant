FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/middle_linux_amd64/middle assistant-middle
CMD ["./assistant-middle"]
