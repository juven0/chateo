package configs

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type dbConfig struct {
	Port     string
	MongoURI string
	Database string
}

func Loadconfig() *dbConfig {
	return &dbConfig{
		Port: os.Getenv("mongoPort"),
		MongoURI: os.Getenv("mongoUri"),
		Database: os.Getenv("mongoName"),
	}
}

func MongoConnection(uri string) *mongo.Client{
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil{
		panic(err)
	}
	return client
}