package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	flagConfig = "config"

	envConfigVar = "APP_CONFIG"
)

func MustLoad[T any]() *T {
	cfg, err := Load[T]()
	if err != nil {
		panic(err)
	}

	return cfg
}

func Load[T any]() (*T, error) {
	path := fetchConfigPath()
	if path == "" {
		return nil, ErrorConfigPathIsEmpty
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, ErrorConfigIsNotExist
	}

	var cfg T

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	return &cfg, nil
}

// fetchConfigPath fetches config path
// Priority: flag > env > default
// Default value is empty string
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, flagConfig, "", "config file path")

	if res == "" {
		res = os.Getenv(envConfigVar)
	}

	return res
}
