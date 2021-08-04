FROM amd64/golang:1.15.3-alpine3.12 as build-env
RUN apk add --no-cache git make curl

WORKDIR /app
COPY app .

RUN GOFLAGS=-mod=vendor go build -o main

ENTRYPOINT [ "./main" ]