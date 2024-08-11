package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	MongoURI    string
	MongoDBName string
	ServerPort  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		MongoURI:    os.Getenv("MONGO_URI"),
		MongoDBName: os.Getenv("MONGO_DB_NAME"),
		ServerPort:  os.Getenv("SERVER_PORT"),
	}, nil
}