package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type dbConfig struct {
	// Port     string
	MongoURI string
	// Database string
}

func Loadconfig() *dbConfig {
	err := godotenv.Load()
    if err != nil {
        panic(err)
    }
	return &dbConfig{
		// Port: os.Getenv("mongoPort"),
		MongoURI: os.Getenv("mongoUri"),
		// Database: os.Getenv("mongoName"),
	}
}

func MongoConnection() *mongo.Client{
	mongoConfig := Loadconfig()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConfig.MongoURI))
	if err != nil{
		panic(err)
	}
	fmt.Println("connected to mongodb")
	return client
}