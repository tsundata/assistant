FROM alpine:latest

RUN apk update && apk add bash

COPY cron app
CMD ["./app"]
