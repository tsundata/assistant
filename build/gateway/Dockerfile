FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/gateway_linux_amd64/gateway app
CMD ["./app"]
