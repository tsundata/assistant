FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/storage_linux_amd64_v1/storage assistant-storage
CMD ["./assistant-storage"]
