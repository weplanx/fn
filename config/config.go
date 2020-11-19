package config

import (
	"func-api/application/service/storage"
	"func-api/config/options"
)

type Config struct {
	Debug    string                 `yaml:"debug"`
	Listen   string                 `yaml:"listen"`
	Logger   bool                   `yaml:"logger"`
	Database options.DatabaseOption `yaml:"database"`
	Storage  storage.Option         `yaml:"storage"`
	Fonts    map[string]string      `yaml:"fonts"`
}
