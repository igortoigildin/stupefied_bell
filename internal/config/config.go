package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	flagConfigPath string
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	LogLevel   string `yaml:"log_level" env-default:"info"`
	DBURI      string
	HTTPserver `yaml:"http_server"`
}

type HTTPserver struct {
	Address    string        `yaml:"address" env-default:"localhost:8080"`
	Timeout    time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	flag.StringVar(&flagConfigPath, "c", "config/local.yaml", "config path")
	flag.Parse()

	if envConfigPath := os.Getenv("CONFIG_PATH"); envConfigPath != "" {
		flagConfigPath = envConfigPath
	}
	if _, err := os.Stat(flagConfigPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", flagConfigPath)
	}

	var cfg Config
	if envDBURI := os.Getenv("DB_URI"); envDBURI != "" {
		cfg.DBURI = envDBURI
	}

	if err := cleanenv.ReadConfig(flagConfigPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
