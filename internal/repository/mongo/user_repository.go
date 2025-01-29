package mongorepository

import (
	mongomodels "chat/internal/models/mongo"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	InsertUser(user *mongomodels.User) error
	GetUser(id string) (mongomodels.User, error)
	UpdateUser(id string, user *mongomodels.User) (*mongo.UpdateResult, error)
	DeleteUser(id string) (*mongo.DeleteResult, error)
}

type MongoUserRepository struct{
	collection *mongo.Collection
}

var ctx = context.TODO()

func InitUserRepository(db *mongo.Database) MongoUserRepository {
	return MongoUserRepository{
		collection: db.Collection("user"),
	}
}

func (r *MongoUserRepository)InsertUser(user *mongomodels.User)error{
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

// func  (r *MongoUserRepository)GetAllUser (){
// 	filter := bson.D{}
// 	r.collection.Find(ctx, filter)
// }

func  (r *MongoUserRepository)GetUser(id string)(mongomodels.User, error){
	filter := bson.D{{Key: "_id", Value: id}}
	var userFind mongomodels.User
	err := r.collection.FindOne(ctx, filter).Decode(&userFind)
	if err != nil {
		return mongomodels.User{}, err 
	}
	return userFind, nil
}

func  (r *MongoUserRepository)UpdateUser(id string, user *mongomodels.User)(*mongo.UpdateResult, error){
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

	result, err := r.collection.UpdateOne(ctx, filtre, update)
	if err != nil {
		return &mongo.UpdateResult{}, err 
	}

	return result, nil
}

func  (r *MongoUserRepository)DeleteUser(id string)(*mongo.DeleteResult, error){
	userId,_ :=primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: userId}}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return &mongo.DeleteResult{}, err
	}
	return result, nil
}



