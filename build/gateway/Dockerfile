FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/gateway_linux_amd64_v1/gateway assistant-gateway
CMD ["./assistant-gateway"]
