package mongorepository

import (
	mongomodels "chat/internal/models/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var conversationCollection *mongo.Collection

const COLLECTION_NAME = "conversation"

func InitConversationRepository(db *mongo.Database){
	conversationCollection = db.Collection(COLLECTION_NAME)
}

func CreatConvesation(converastion *mongomodels.Conversation) error{
	_, err := conversationCollection.InsertOne(ctx, converastion)
	if err != nil{
		return err
	}
	return nil
}

func DeleteConversation(idConversation string) (*mongo.DeleteResult, error){
	id, _:= primitive.ObjectIDFromHex(idConversation)
	
	filter := bson.D{{Key: "_id", Value: id}}

	result, err :=  conversationCollection.DeleteOne(ctx, filter)
	if err != nil{
		return &mongo.DeleteResult{}, err
	}
	return result, nil
}