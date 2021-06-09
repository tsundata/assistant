FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/subscribe_linux_amd64/subscribe app
CMD ["./app"]
