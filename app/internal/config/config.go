package config

import (
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  `yaml:"listen"`
	MongoDB `yaml:"mongodb" env-required:"true"`
	Redis   `yaml:"redis" env-required:"true"`
	Minio   `yaml:"minio" env-required:"true"`
}

type Listen struct {
	Type   string `yaml:"type" env-default:"port"`
	BindIP string `yaml:"bind_ip" env-default:"localhost"`
	Port   string `yaml:"port" env-default:"8080"`
}

type MongoDB struct {
	Host       string `yaml:"host" env-required:"true"`
	Port       string `yaml:"port" env-required:"true"`
	Username   string `yaml:"username"`
	Password   string `yaml:"-" env:"MONGODB_PASSWORD"`
	AuthDB     string `yaml:"auth_db" env-required:"true"`
	Database   string `yaml:"database" env-required:"true"`
	Collection string `yaml:"collection" env-required:"true"`
}

type Redis struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Password string `yaml:"-" env:"REDIS_PASSWORD"`
}

type Minio struct {
	Host       string `yaml:"host" env-required:"true"`
	Port       string `yaml:"port" env-required:"true"`
	Username   string `yaml:"username" env-required:"true"`
	Password   string `yaml:"-" env:"MINIO_PASSWORD"`
	SSL        bool   `yaml:"ssl" env-default:"false"`
	BucketName string `yaml:"bucket_name" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")
		instance = &Config{}

		if err := godotenv.Load(".env"); err != nil {
			logger.Fatal("Error loading .env file")
		}

		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}

		instance.MongoDB.Password = os.Getenv("MONGODB_PASSWORD")
		instance.Redis.Password = os.Getenv("REDIS_PASSWORD")
		instance.Minio.Password = os.Getenv("MINIO_PASSWORD")
	})

	return instance
}
