package main

import (
	"log"
	"context"

	"chi_nats/internal/config"
	"chi_nats/internal/messaging"
	"chi_nats/internal/httpserver"
)

func main() {
	cfg := config.MustLoad()

	nc, err := broker.InitNATS(cfg.NatsAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	publisher := broker.NewPublisher(nc)
	server := httpserver.New(cfg.PublisherAddr, *publisher)

	ctx := context.Background()
	if err := server.Listen(ctx); err != nil {
		log.Fatal(err)
	}
}
