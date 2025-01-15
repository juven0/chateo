package mongoservice

import (
	mongomodels "chat/internal/models/mongo"
	mongorepository "chat/internal/repository/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

func setupMessageRepository() {

	mongorepository.InitMessagRepository(mongoClient.Database(tableName))
}

func CreateMessage(messge *mongomodels.Message) error {
	setupMessageRepository()
	return mongorepository.CreateMessage(messge)
}

func GetAllMessageConversation(idConversation string)([]mongomodels.Message,error){
	setupMessageRepository()
	return mongorepository.GetAllMessageConversation(idConversation)
}

func GetOneMessage(messageId string)(mongomodels.Message, error){
	setupMessageRepository()
	return mongorepository.GetOneMessage(messageId)
}

func DeleteMessage(messageId string)(*mongo.DeleteResult, error){
	setupMessageRepository()
	return mongorepository.DeletMessage(messageId)
}