package mongorepository

import (
	mongomodels "chat/internal/models/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection
var ctx = context.TODO()

func InitUserRepository(db *mongo.Database){
	userCollection = db.Collection("user")
}

func InsertUser(user *mongomodels.User)error{
	_, err := userCollection.InsertOne(ctx, user)
	return err
}

func GetAllUser (){
	filter := bson.D{}
	userCollection.Find(ctx, filter)
}



