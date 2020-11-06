package types

type StorageOption struct {
	Drive  string                 `yaml:"drive"`
	Option map[string]interface{} `yaml:"option"`
}
