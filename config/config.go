package config

import (
	"dbKursach/internal/pkg/utils"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Postgres Postgres
	Logger   Logger
}

type Logger struct {
	Level zerolog.Level `validate:"required"`
}

type Server struct {
	Host   string `validate:"required"`
	Port   string `validate:"required"`
	APIKey string `validate:"required"`
}

type Postgres struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	DBName   string `validate:"required"`
	SSLMode  string `validate:"required"`
	PGDriver string `validate:"required"`
	Settings struct {
		MaxOpenConns    int           `validate:"required,min=1"`
		ConnMaxLifetime time.Duration `validate:"required,min=1"`
		MaxIdleConns    int           `validate:"required,min=1"`
		ConnMaxIdleTime time.Duration `validate:"required,min=1"`
	}
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	path := os.Getenv("CONFIG")

	if len(path) != 0 {
		v.AddConfigPath(path)
	}

	v.AddConfigPath("config")

	v.SetConfigName(utils.ConfigFileName)
	v.SetConfigType(utils.ConfigExtension)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	err = validator.New().Struct(c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
