package config

import (
	"os"
	"strconv"
)

type (
	AppConfig struct {
		ContextTimeout int
		ServerPort     string
		DbHost         string
		DbPort         string
		DbUser         string
		DbPass         string
		DbName         string
	}
)

func InitConfig() AppConfig {
	timeOut, err := strconv.Atoi(getEnv("ContextTimeout", "10"))
	if err != nil {
		timeOut = 10
	}

	return AppConfig{
		ContextTimeout: timeOut,
		ServerPort:     getEnv("ServerAddress", "8080"),
		DbHost:         getEnv("DB_HOST", "mongodb://user:pass@mongo:27018/mezink"),
		DbPort:         getEnv("DB_PORT", "27018"),
		DbUser:         getEnv("DB_USER", "user"),
		DbPass:         getEnv("DB_PASS", "pass"),
		DbName:         getEnv("DB_NAME", "mezink"),
	}
}

func getEnv(env, def string) string {
	res, exists := os.LookupEnv(env)

	if exists {
		return res
	}
	return def
}
