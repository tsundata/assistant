FROM alpine:latest

RUN apk update && apk add bash

COPY ./dist/org_linux_amd64/org assistant-org
CMD ["./assistant-org"]
