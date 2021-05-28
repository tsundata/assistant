FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/cron-linux-amd64/cron app
CMD ["./app"]
