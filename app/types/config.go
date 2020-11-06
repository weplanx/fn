package types

type Config struct {
	Debug   string        `yaml:"debug"`
	Listen  string        `yaml:"listen"`
	Storage StorageOption `yaml:"storage"`
}
