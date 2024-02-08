package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	SubjectName   string `env:"SUBJECT_NAME" env-default:"example"`
	NatsAddr  	  string `env:"NATS_ADDR" env-default:":4222"`
	PublisherAddr string `env:"PUBLISHER_ADDR" env-default:":8080"`
}

func MustLoad() Config {
	var cfg Config

	err := cleanenv.ReadConfig("../../config.env", &cfg)

	if err != nil {
		panic(fmt.Sprintf("Cannot read config: %s", err))
	}

	return cfg
}