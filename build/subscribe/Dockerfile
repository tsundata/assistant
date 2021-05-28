FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/subscribe-linux-amd64/subscribe app
CMD ["./app"]
