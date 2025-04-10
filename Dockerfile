FROM golang:1.23-alpine

ARG CGO_ENABLED=0
ARG GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY src/ ./src/

RUN go build -o main ./src

EXPOSE 3000

CMD ["./main"]
