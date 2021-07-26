FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/finance_linux_amd64/finance assistant-finance
CMD ["./assistant-finance"]
