FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/web-linux-amd64/web app
CMD ["./app"]
