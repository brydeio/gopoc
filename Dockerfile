# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app
RUN go mod init github.com/brydeio/gopoc
RUN go mod tidy
RUN ls -la
COPY *.go ./
RUN go build -o /hello
CMD ["/hello"]