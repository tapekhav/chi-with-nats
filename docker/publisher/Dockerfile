FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go clean -modcache
RUN go mod download

COPY . .

RUN go build -o publisher cmd/publisher/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /build/publisher /app/publisher

RUN adduser -D nonrootuser
USER nonrootuser

ENV CFG_PATH ""

CMD ["./publisher"]
