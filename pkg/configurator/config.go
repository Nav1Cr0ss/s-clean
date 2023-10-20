package configurator

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

func NewConfig[T any]() *T {
	var config T
	err := loadDotEnv()
	if err != nil {
		panic(fmt.Errorf("error on parsing env file :%w", err))
	}

	err = envconfig.Process("", &config)
	if err != nil {
		panic(fmt.Errorf("error loading environment variables :%w", err))
	}
	return &config
}

func loadDotEnv() error {
	envPath := os.Getenv("ENV_FILE")

	var err error
	if envPath == "" {
		_ = godotenv.Load(".env") // ignore error by default
	} else {
		err = godotenv.Load(envPath) // if path to env file defined, check error
	}

	if err != nil {
		return err
	}

	return nil
}
