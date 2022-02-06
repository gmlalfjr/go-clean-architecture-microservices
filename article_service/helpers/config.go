package helpers

import (
	"github.com/joho/godotenv"
	"os"
)
type Configuration interface {
	Get(key string) string
}

type ConfigurationImpl struct {
}

func (config *ConfigurationImpl) Get(key string) string {
	return os.Getenv(key)
}

func NewConfiguration() Configuration {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return &ConfigurationImpl{}
}
