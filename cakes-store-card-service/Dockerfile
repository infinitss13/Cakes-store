FROM golang:alpine

RUN apk update
RUN apk add postgresql-client


WORKDIR /app

COPY ./ /app


RUN go mod download


ENTRYPOINT go run main.go server.go

EXPOSE 8001