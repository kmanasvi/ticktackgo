package index

import (
	"os"
)

type Environment struct {
	ROUTER_PORT     string
	ENV             string
	DB_ADDR         string
	DB_USER         string
	DB_PASSWORD     string
	DB_DATABASE     string
	REDIS_CLIENT    string
	RELEASE_VERSION string
}

var EnvironmentConfig = Environment{
	ROUTER_PORT:     os.Getenv("ROUTER_PORT"),
	ENV:             os.Getenv("ENV"),
	DB_ADDR:         os.Getenv("DB_ADDR"),
	DB_USER:         os.Getenv("DB_USER"),
	DB_PASSWORD:     os.Getenv("DB_PASSWORD"),
	DB_DATABASE:     os.Getenv("DB_DATABASE"),
	REDIS_CLIENT:    os.Getenv("REDIS_CLIENT"),
	RELEASE_VERSION: os.Getenv("RELEASE_VERSION"),
}
