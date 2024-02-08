package broker

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	nc *nats.Conn
}

func (s *Subscriber) SubscribeToSubject(
	subject string,
	handleMsg func(msg *nats.Msg),
) {
	_, err := s.nc.Subscribe(subject, handleMsg)

	if err != nil {
		log.Printf("Error subscribing to subject: %v", err)
	}
}

func NewSubscriber(nc *nats.Conn) *Subscriber {
	return &Subscriber{
		nc: nc,
	}
}
