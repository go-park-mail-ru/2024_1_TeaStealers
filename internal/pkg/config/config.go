package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload" // Load enviroment from .env
	"log"
	"os"
	"time"
)

type Config struct {
	ConfigPath string `env:"CONFIG_PATH" env-default:"config/config.yaml"`
	HTTPServer `yaml:"httpServer"`
	GRPC       GRPC
	Database   Database
}

type HTTPServer struct {
	Address           string        `yaml:"address" yaml-default:"localhost:8080"`
	Timeout           time.Duration `yaml:"timeout" yaml-default:"10s"`
	IdleTimeout       time.Duration `yaml:"idleTimeout" yaml-default:"60s"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout" yaml-default:"10s"`
}

type Database struct {
	DBName string `env:"POSTGRES_DB" env-required:"true"`
	DBPass string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DBHost string `env:"DB_HOST" env-default:"0.0.0.0"`
	DBPort int    `env:"DB_PORT" env-required:"true"`
	DBUser string `env:"POSTGRES_USER" env-required:"true"`
}

type GRPC struct {
	AuthPort   int `env:"GRPC_AUTH_PORT" env-default:"8081"`
	UserPort   int `env:"GRPC_USER_PORT" env-default:"8082"`
	AdvertPort int `env:"GRPC_ADVERT_PORT" env-default:"8083"`

	AuthContainerIP   string `env:"GRPC_AUTH_CONTAINER_IP" env-default:"localhost"`
	UsersContainerIP  string `env:"GRPC_USER_CONTAINER_IP" env-default:"localhost"`
	AdvertContainerIP string `env:"GRPC_ADVERT_CONTAINER_IP" env-default:"localhost"`
}

func MustLoad() *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Printf("cannot read .env file: %s\n (fix: you need to put .env file in main dir)", err)
		os.Exit(1)
	}

	if err := cleanenv.ReadConfig(cfg.ConfigPath, &cfg); err != nil {
		log.Printf("cannot read %s: %v", cfg.ConfigPath, err)
		os.Exit(1)
	}

	return &cfg
}
