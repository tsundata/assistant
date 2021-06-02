FROM alpine:latest

RUN apk update && apk add bash

COPY gateway app
CMD ["./app"]
