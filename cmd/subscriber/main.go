package main

import (
	"log"

	"chi_nats/internal/messaging"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := broker.InitNATS(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subscriber := broker.NewSubscriber(nc)
	subscriber.SubscribeToSubject("example", func(msg *nats.Msg) {
		log.Printf("Received message: %s\n", string(msg.Data))
	})

	log.Println("Subscriber started. Waiting for messages...")

	select {}
}
