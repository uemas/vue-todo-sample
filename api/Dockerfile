FROM golang:1.18

WORKDIR /usr/src/api

COPY . .
RUN go mod download && go mod verify
RUN go build ./cmd/go-todo-api/
