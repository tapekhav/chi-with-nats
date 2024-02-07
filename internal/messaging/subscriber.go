package broker

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Subsriber struct {
	nc *nats.Conn
}

func (s *Publisher) SubsribeToSubject(
	subject string, 
	handleMsg func(msg *nats.Msg),
) {
	_, err := s.nc.Subscribe(subject, handleMsg)

	if err != nil {
		log.Printf("Error publishing message: %v", err)
	}
}

func NewSubsciber(nc *nats.Conn) *Subsriber {
	return &Subsriber{
		nc: nc,
	}
}
