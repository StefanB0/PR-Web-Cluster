# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

ARG port
ARG main
ARG leader

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY client ./client
COPY databse ./databse
COPY web-server ./web-server

RUN go mod download

COPY *.go ./

RUN go build -o /docker-web-cluster

EXPOSE 8881

CMD [ "/docker-web-cluster", "-address=${port}", "-mainAddress=${main}", "-leader=${leader}" ]