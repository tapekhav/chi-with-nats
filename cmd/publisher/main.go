package main

import (
	"log"
	"context"

	"chi_nats/internal/httpserver"
	"chi_nats/internal/messaging"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := broker.InitNATS(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	publisher := broker.NewPublisher(nc)

	server := httpserver.New(":8080", *publisher)

	ctx := context.Background()

	if err := server.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}
