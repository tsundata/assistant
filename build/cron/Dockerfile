FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/cron-linux-amd64/cron app
CMD ["./app"]
