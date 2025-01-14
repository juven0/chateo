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

func GetOneMessage(idConversation string, idMessage string){

}

func DeletMessage(idConversation string, idMessage string){

}
