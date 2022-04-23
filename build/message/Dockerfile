FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/message_linux_amd64_v1/message assistant-message
CMD ["./assistant-message"]
