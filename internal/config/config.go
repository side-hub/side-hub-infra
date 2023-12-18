package config

import "github.com/side-hub-be/side-hub-be/internal/db"

type Config struct {
	DB db.Config `yaml:"db"`
}
