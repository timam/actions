FROM golang:1.20 AS build
WORKDIR /app
COPY ./app /app
RUN go build -o /bin/sync-guardian
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]


