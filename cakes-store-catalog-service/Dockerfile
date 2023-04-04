FROM golang:alpine


WORKDIR /app

COPY ./ /app


RUN go mod download


ENTRYPOINT go run main.go server.go

EXPOSE 8080