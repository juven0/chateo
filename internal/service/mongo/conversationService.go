package mongoservice

import (
	mongomodels "chat/internal/models/mongo"
	mongorepository "chat/internal/repository/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

type ConvesationService struct{
	repo mongorepository.MongoConvesationRepository
}

var(
	ConvesationServiceInstance *ConvesationService
)

func GetConversationServiceInstance()*ConvesationService{
	once.Do(func() {
		repository := mongorepository.NewConversationRepository(mongoClient.Database(tableName))
		ConvesationServiceInstance = &ConvesationService{
			repo: *repository,
		}
	})
	return ConvesationServiceInstance
}

func (s *ConvesationService)CreatConvesation(conversation *mongomodels.Conversation) error{
	return s.repo.CreatConvesation(conversation)
}

func (s *ConvesationService)DeleteConverstion(idConversation string)(*mongo.DeleteResult, error){

	return s.repo.DeleteConversation(idConversation)
}