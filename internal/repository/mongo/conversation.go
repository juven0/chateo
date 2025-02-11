package mongorepository

import (
	mongomodels "chat/internal/models/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConvesationRepository struct{
	collection *mongo.Collection
}

var conversationCollection *mongo.Collection



const COLLECTION_NAME = "conversation"

func NewConversationRepository(db *mongo.Database) *MongoConvesationRepository{
	return &MongoConvesationRepository{
		collection: db.Collection(COLLECTION_NAME),
	}
}

func (r *MongoConvesationRepository)CreatConvesation(converastion *mongomodels.Conversation) error{
	_, err := conversationCollection.InsertOne(ctx, converastion)
	if err != nil{
		return err
	}
	return nil
}

func (r *MongoConvesationRepository)DeleteConversation(idConversation string) (*mongo.DeleteResult, error){
	id, _:= primitive.ObjectIDFromHex(idConversation)
	
	filter := bson.D{{Key: "_id", Value: id}}

	result, err :=  conversationCollection.DeleteOne(ctx, filter)
	if err != nil{
		return &mongo.DeleteResult{}, err
	}
	return result, nil
}