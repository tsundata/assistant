FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/bot_linux_amd64/bot assistant-bot
CMD ["./assistant-bot"]
