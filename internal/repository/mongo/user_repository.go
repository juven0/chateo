package mongo

import "go.mongodb.org/mongo-driver/mongo"

var userCollection *mongo.Collection

func InitUserRepository(db *mongo.Database){
	userCollection = db.Collection("user")
}

