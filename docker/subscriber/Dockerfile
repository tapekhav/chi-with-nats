FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY ./go.mod .
COPY ./go.sum .

RUN go clean -modcache
RUN go mod download

COPY . .

RUN go build -o subscriber cmd/subscriber/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /build/subscriber /app/subscriber

RUN adduser -D nonrootuser
USER nonrootuser

ENV CFG_PATH ""

CMD ["./subscriber"]