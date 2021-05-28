FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/middle-linux-amd64/middle app
CMD ["./app"]
