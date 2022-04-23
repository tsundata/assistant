FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/user_linux_amd64_v1/user assistant-user
CMD ["./assistant-user"]
