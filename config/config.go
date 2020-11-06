package config

import "funcext/application/service/storage"

type Config struct {
	Debug   string         `yaml:"debug"`
	Listen  string         `yaml:"listen"`
	Storage storage.Option `yaml:"storage"`
}
