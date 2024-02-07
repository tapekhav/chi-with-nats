package broker

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	nc *nats.Conn
}

func (p *Publisher) PublishMsg(subject string, msg []byte) {
	err := p.nc.Publish(subject, msg)

	if err != nil {
		log.Printf("Error publishing message: %v", err)
	}
}

func NewPublisher(nc *nats.Conn) *Publisher {
	return &Publisher{
		nc: nc,
	}
}
