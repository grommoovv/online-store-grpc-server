package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Env  string     `yaml:"env" env-default:"local"`
		Grpc GrpcConfig `yaml:"grpc"`
		Psql PsqlConfig `yaml:"psql"`
	}

	GrpcConfig struct {
		Port    int           `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	}

	PsqlConfig struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	}
)

func MustLoadConfig() *Config {
	var configPath = "./config/local.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("can not load config file: " + err.Error())
	}

	return &cfg
}
