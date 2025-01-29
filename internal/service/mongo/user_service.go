package mongoservice

import (
	mongomodels "chat/internal/models/mongo"
	mongorepository "chat/internal/repository/mongo"

	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	repo mongorepository.UserRepository
}

var(
	userServiceInstance *UserService
	once sync.Once
)

func GetUserServiceInstace()*UserService{
	once.Do(func(){
		repository := mongorepository.NewUserRepository(mongoClient.Database(tableName))
		userServiceInstance = &UserService{
			repo: repository,
		}
	})
	return userServiceInstance
}


func (r *UserService)CreatUser(user *mongomodels.User) error{

	return r.repo.InsertUser(user)
}

func (r *UserService)GetUser(id string)(mongomodels.User, error){
	return r.repo.GetUser(id)
}

func (r *UserService)UpdateUser(id string, user *mongomodels.User)(*mongo.UpdateResult, error){
	return r.repo.UpdateUser(id, user)
}

func (r *UserService)DeleteUser(id string)(*mongo.DeleteResult, error){
	return r.repo.DeleteUser(id)
}