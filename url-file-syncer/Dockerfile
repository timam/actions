FROM golang:1.20 AS build
WORKDIR /app
COPY ./app /app
RUN go build -o /url-file-syncer
ENTRYPOINT ["/url-file-syncer"]

FROM alpine:edge
COPY --from=build /url-file-syncer /url-file-syncer
ENTRYPOINT ["/url-file-syncer"]

