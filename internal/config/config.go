package config

import (
	"context"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/appengine/log"
)

type Config struct {
	NatsAddr  	  string `env:"NATS_ADDR" env-default:":4222"`
	SubjectName   string `env:"SUBJECT_NAME" env-default:"example"`
	PublisherAddr string `env:"PUBLISHER_ADDR" env-default:":8080"`
}

func MustLoad() Config {
	path := os.Getenv("CFG_PATH")

	var err error
	var cfg Config
	if path != "" {
		err = cleanenv.ReadConfig(path, &cfg)
	} else {
		err = cleanenv.ReadEnv(&cfg)
	}

	if err != nil {
		ctx := context.Background()
		log.Warningf(ctx, "Cannot read config: %s", err)
	}

	return cfg
}