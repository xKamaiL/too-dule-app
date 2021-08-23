package config

import (
	"github.com/acoshift/configfile"
)

type Config struct {
	// database
	DB struct {
		Host     string
		Port     int
		Username string
		Password string
		Name     string
	}
	// no password require
	RedisHost string
}

var config Config

func Init() {
	r := configfile.NewEnvReader()
	config.DB.Host = r.String("DB_HOST")
	config.DB.Port = r.Int("DB_PORT")
	config.DB.Username = r.String("DB_USERNAME")
	config.DB.Password = r.String("DB_PASSWORD")
	config.DB.Name = r.String("DB_NAME")

	config.RedisHost = r.String("REDIS_HOST")
}

func Load() *Config {
	return &config
}
