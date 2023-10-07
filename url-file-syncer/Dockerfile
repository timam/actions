FROM golang:1.20 AS build
WORKDIR /app
COPY ./app /app
RUN go build -o /url-file-syncer
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/file-watcher"]


