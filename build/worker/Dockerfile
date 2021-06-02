FROM alpine:latest

RUN apk update && apk add bash

COPY worker app
CMD ["./app"]
