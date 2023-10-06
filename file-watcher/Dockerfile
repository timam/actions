FROM golang:1.20 AS build
WORKDIR /app
COPY ./app /app
RUN go build -o /file-watcher
ENTRYPOINT ["/file-watcher"]


