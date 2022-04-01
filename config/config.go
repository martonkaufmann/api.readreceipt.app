package config

import (
	"os"
	"strconv"
)

func URL() string {
	return os.Getenv("URL")
}

func Secret() string {
	return os.Getenv("SECRET")
}

func Env() string {
	return os.Getenv("ENV")
}

func IsLocal() bool {
	return os.Getenv("ENV") == "local"
}

func Sentry() string {
	return os.Getenv("SENTRY_DSN")
}

func MongoURL() string {
	return os.Getenv("MONGO_URL")
}

func MongoDatabase() string {
	return os.Getenv("MONGO_DATABASE")
}

func RedisURL() string {
	return os.Getenv("REDIS_URL")
}

func RedisDatabase() int {
	db, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))

	if err != nil {
		panic(err)
	}

	return db
}
