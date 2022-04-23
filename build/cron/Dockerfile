FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/cron_linux_amd64_v1/cron assistant-cron
CMD ["./assistant-cron"]
