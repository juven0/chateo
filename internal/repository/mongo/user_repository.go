package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection
var ctx = context.TODO()

func InitUserRepository(db *mongo.Database){
	userCollection = db.Collection("user")
}

func GetAllUser (){
	filter := bson.D{}
	userCollection.Find(ctx, filter)
}

