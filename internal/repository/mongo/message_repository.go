package mongorepository

import (
	mongomodels "chat/internal/models/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var messageCollection *mongo.Collection

func InitMessagRepository(db *mongo.Database){
	messageCollection = db.Collection("message")
}

func CreateMessage(message *mongomodels.Message)error{
	_, err := messageCollection.InsertOne(ctx, message)
	if err!= nil{
		return err
	}
	return nil
}

func GetAllMessageConversation(idConversation string)([]mongomodels.Message, error){
	id , _ := primitive.ObjectIDFromHex(idConversation)

	filter := bson.D{{Key: "_id", Value: id}}
	cursor , err := messageCollection.Find(ctx, filter)
	if err!= nil{
		return []mongomodels.Message{},err
	}
	 var messages []mongomodels.Message
	if err =cursor.All(ctx, &messages); err!= nil{
		return  []mongomodels.Message{}, err
	}

	return messages, err
}

func GetOneMessage(idMessage string)(mongomodels.Message, error){
	idMess, _ := primitive.ObjectIDFromHex(idMessage)

	filter := bson.D{{Key: "_id", Value: idMess}}
	
	var message mongomodels.Message
	err := messageCollection.FindOne(ctx, filter).Decode(&message)
	if err != nil{
		return mongomodels.Message{}, err
	}
	return message , nil
}

func DeletMessage(idMessage string)(*mongo.DeleteResult, error){
	id , _:= primitive.ObjectIDFromHex(idMessage)

	filtre := bson.D{{Key: "_id", Value: id}}

	result, err :=messageCollection.DeleteOne(ctx, filtre)
	if err!= nil {
		return &mongo.DeleteResult{}, err
	}
	return result, nil
}
