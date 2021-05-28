FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/middle-linux-amd64/middle app
CMD ["./app"]
