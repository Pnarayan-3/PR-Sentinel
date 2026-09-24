FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build \
    -o /pr-sentinel \
    ./cmd/sentinel


FROM alpine:3.22

WORKDIR /workspace

COPY --from=builder /pr-sentinel /usr/local/bin/pr-sentinel

ENTRYPOINT ["pr-sentinel"]