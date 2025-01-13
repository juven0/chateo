package mongorepository

import (
	mongomodels "chat/internal/models/mongo"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetUser(id string)(mongomodels.User, error){
	filter := bson.D{{Key: "_id", Value: id}}
	var userFind mongomodels.User
	err := userCollection.FindOne(ctx, filter).Decode(&userFind)
	if err != nil {
		return mongomodels.User{}, err 
	}
	return userFind, nil
}

func UpdateUser(id string, user *mongomodels.User)(*mongo.UpdateResult, error){
	userId , _:= primitive.ObjectIDFromHex(id)

	filtre := bson.D{{Key: "_id", Value: userId}}

	update := bson.M{
		"$set": bson.M{
			"username":  user.UserName,
			"email":     user.Email,
			"password":  user.Password,
			"status":    user.Status,
			"lastseen":  user.LastSeen,
			"updateat":  time.Now().Format(time.RFC3339),
		},
	}

	result, err := userCollection.UpdateOne(ctx, filtre, update)
	if err != nil {
		return &mongo.UpdateResult{}, err 
	}

	return result, nil
}

func DeleteUser(){
	
}



