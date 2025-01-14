package mongoservice

import (
	"chat/configs"
	mongomodels "chat/internal/models/mongo"
	mongorepository "chat/internal/repository/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient = configs.MongoConnection()
const tableName = "chat"

func setup() {
	mongorepository.InitUserRepository(mongoClient.Database(tableName))
}


func CreatUser(user *mongomodels.User) error{
	setup()
	return mongorepository.InsertUser(user)
}

func GetUser(id string)(mongomodels.User, error){
	setup()
	return mongorepository.GetUser(id)
}

func UpdateUser(id string, user *mongomodels.User)(*mongo.UpdateResult, error){
	setup()
	return mongorepository.UpdateUser(id, user)
}

func DeleteUser(id string)(*mongo.DeleteResult, error){
	setup()
	return mongorepository.DeleteUser(id)
}