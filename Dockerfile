# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY ./public ./public
RUN go mod download

COPY *.go ./
RUN go build -o /blog

EXPOSE 80

CMD [ "/blog" ]

