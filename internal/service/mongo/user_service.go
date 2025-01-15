package mongoservice

import (
	mongomodels "chat/internal/models/mongo"
	mongorepository "chat/internal/repository/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)



func setupUserService() {
	mongorepository.InitUserRepository(mongoClient.Database(tableName))
}

func CreatUser(user *mongomodels.User) error{
	setupUserService()
	return mongorepository.InsertUser(user)
}

func GetUser(id string)(mongomodels.User, error){
	setupUserService()
	return mongorepository.GetUser(id)
}

func UpdateUser(id string, user *mongomodels.User)(*mongo.UpdateResult, error){
	setupUserService()
	return mongorepository.UpdateUser(id, user)
}

func DeleteUser(id string)(*mongo.DeleteResult, error){
	setupUserService()
	return mongorepository.DeleteUser(id)
}