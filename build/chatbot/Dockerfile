FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/chatbot_linux_amd64_v1/chatbot assistant-chatbot
CMD ["./assistant-chatbot"]
