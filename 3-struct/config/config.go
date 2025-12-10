package config

import (
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("There are no KEY variable set!")
	}

	return &Config{
		Key: key,
	}
}
