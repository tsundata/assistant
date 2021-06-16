FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/nlp_linux_amd64/nlp app
CMD ["./app"]
