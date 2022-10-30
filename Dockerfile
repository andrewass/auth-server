# syntax=docker/dockerfile:1

##Build
FROM golang:1.19.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /docker-auth-server

EXPOSE 8089

CMD ["/docker-auth-server"]
