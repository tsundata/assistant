FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/user_linux_amd64/user app
CMD ["./app"]
