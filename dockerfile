# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

ARG id=0
ARG address="http:minion""
ARG port=":3000"
ARG leader="http:leader0:3000"
ARG isLeader=false

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY client ./client
COPY databse ./databse
COPY web-server ./web-server

RUN go mod download

COPY *.go ./

RUN go build -o /docker-web-cluster

EXPOSE 3000

CMD [ "/docker-web-cluster", "-id:${id}", "-address=${address}", "-port=${port}", "-leader=${leader}", "-isLeader=${isLeader}" ]