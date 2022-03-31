# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Build stage
FROM golang:1.18.0-alpine as base
LABEL maintainer="FrasNym <frastyawan.nym@gmail.com>"
RUN apk update && apk upgrade && \
    apk add --no-cache g++
WORKDIR /app
COPY . .
RUN go build -o main main.go

# [Prod] Run stage
FROM alpine:3.15 as prod
WORKDIR /app
COPY --from=base /app/main .
COPY --from=base /app/config-local.yaml .
ENTRYPOINT /app/main