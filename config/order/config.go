package cfgorder

import (
	"flag"
	"log"
	"net"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

var (
	flagConfigPath string
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	LogLevel   string `yaml:"log_level" env-default:"info"`
	DBURI      string
	HTTPserver `yaml:"http_server"`
	Kafka      `yaml:"kafka"`
	GRPCServer `yaml:"grpc_server"`
}

type HTTPserver struct {
	Address    string        `yaml:"address" env-default:"localhost:8080"`
	Timeout    time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type GRPCServer struct {
	Port int    `yaml:"port" env-default:"50051"`
	Ip   net.IP `yaml:"ip" env-default:"127.0.0.1"`
}

type Kafka struct {
	Brokers []string `env-required:"true" yaml:"brokers" env:"KAFKA_BROKERS"`   // ["broker:9092"]
	Topic   string   `env-requeired:"true" yaml:"topic" env:"KAFKA_TOPIC"`      // "my_topic"
	GroupID string   `env-required:"true" yaml:"group_id" env:"KAFKA_GROUP_ID"` // "my_group"
}

func MustLoad() *Config {
	flag.StringVar(&flagConfigPath, "c", "config/order/local.yaml", "config path")
	flag.Parse()

	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file:", err)
	}

	cfg := Config{}
	if envDBURI := os.Getenv("DB_URI"); envDBURI != "" {
		cfg.DBURI = envDBURI
	}

	// Read the YAML configuration file
	if err := cleanenv.ReadConfig(flagConfigPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	// Read the environment variables from the loaded .env file
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
	}

	log.Print(cfg)

	return &cfg
}
