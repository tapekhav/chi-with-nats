package broker

import (
	"log"

	"github.com/nats-io/nats.go"
)

func InitNATS(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return nc, nil
}
