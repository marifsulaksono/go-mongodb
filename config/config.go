package config

import (
	"fmt"
	"os"
)

const (
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBName     = "DB_NAME"
	ServerPort = "SERVER_PORT"
)

type PaymentConfig struct {
	MidtranasServerKey   string
	MidtranasSandboxLink string
}

type DBConfig struct {
	ServerPort   string
	DatabaseURL  string
	DatabaseName string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		ServerPort:   os.Getenv(ServerPort),
		DatabaseURL:  fmt.Sprintf("%v:%v", os.Getenv(DBHost), os.Getenv(DBPort)),
		DatabaseName: os.Getenv(DBName),
	}
}
