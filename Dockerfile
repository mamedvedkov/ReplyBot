FROM golang:1.18-alpine AS builder
LABEL stage=builder

ENV CGO_ENABLED 0

ENV TZ=Europe/Moscow

RUN apk --no-cache add ca-certificates tzdata && \
    cp -r -f /usr/share/zoneinfo/$TZ /etc/localtime

WORKDIR /app

COPY ./cmd ./cmd
COPY ./vendor ./vendor
COPY ./go.mod ./
COPY ./go.sum ./
COPY ./internal ./internal

RUN GOOS=linux GOARCH=amd64 go build -mod=vendor -o /replybot ./cmd/replybot

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /replybot /replybot

ENTRYPOINT ["/replybot"]