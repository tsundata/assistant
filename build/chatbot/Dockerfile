FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/chatbot_linux_amd64/chatbot app
CMD ["./app"]
