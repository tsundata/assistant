FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/web_linux_amd64/web assistant-web
CMD ["./assistant-web"]
