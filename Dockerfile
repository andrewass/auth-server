
FROM golang:1.19.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /auth-server-backend

CMD ["/auth-server-backend"]
