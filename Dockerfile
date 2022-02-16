FROM golang:1.17.7-alpine3.15
RUN apk add build-base
LABEL maintainer="FrasNym <frastyawan.nym@gmail.com>"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build .
CMD [ "./go-boilerplate" ]