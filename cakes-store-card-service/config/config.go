package config

import (
	"fmt"
	"os"
)

func GetPortEnv() string {
	port := os.Getenv("PORT_CAKE_USER")
	if port == "" {
		port = "8002"
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

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBName     string
	DBSslmode  string
	DBPassword string
}

func NewDBConfig() DBConfig {
	return DBConfig{
		DBHost:     GetEnv("HOST_DB", "postgres"),
		DBPort:     GetEnv("PORT_DB", "5432"),
		DBUsername: GetEnv("USERNAME_DB", "postgres"),
		DBName:     GetEnv("DBNAME_DB", "postgres"),
		DBSslmode:  GetEnv("SSLMODE_DB", "disable"),
		DBPassword: GetEnv("PASSWORD_DB", "qwerty"),
	}
}
func (c *DBConfig) ConnectionDbData() string {
	return fmt.Sprintf("%s://%s:%s@postgres:%s/%s?sslmode=%s",
		c.DBHost, c.DBUsername, c.DBPassword, c.DBPort, c.DBName, c.DBSslmode)
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
