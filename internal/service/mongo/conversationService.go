package mongoservice

import (
	mongomodels "chat/internal/models/mongo"
	mongorepository "chat/internal/repository/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

func setupConversationRepository() {
	mongorepository.InitConversationRepository(mongoClient.Database(tableName))
}

func CreatConvesation(conversation *mongomodels.Conversation) error{
	setupConversationRepository()
	return mongorepository.CreatConvesation(conversation)
}

func DeleteConverstion(idConversation string)(*mongo.DeleteResult, error){
	setupConversationRepository()
	return mongorepository.DeleteConversation(idConversation)
}