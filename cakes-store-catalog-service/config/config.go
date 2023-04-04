package config

import (
	"os"
)

func GetPortEnv() string {
	port := os.Getenv("PORT_CAKE_USER")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

type ConnectionMongo struct {
	MongoHost       string
	MongoPort       string
	MongoDBName     string
	MongoCollection string
}

func NewConnectionMongo() ConnectionMongo {
	return ConnectionMongo{
		MongoHost:       GetEnv("HOST_MONGO", "mongo"),
		MongoPort:       GetEnv("PORT_MONGO", "27017"),
		MongoDBName:     GetEnv("DBNAME_MONGO", "mongo"),
		MongoCollection: GetEnv("COLLECTION_MONGO", "catalog"),
	}
}
