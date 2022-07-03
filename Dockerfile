FROM golang:1.18.2 as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/booklog-prometheus-exporter \
    -ldflags '-s -w'

FROM alpine:latest as runner
RUN apk add curl
COPY --from=builder /go/bin/booklog-prometheus-exporter /app/booklog-prometheus-exporter

# don't create homeDir.
RUN adduser -D -S -H exporter

USER exporter
ENTRYPOINT [ "/app/booklog-prometheus-exporter" ]
