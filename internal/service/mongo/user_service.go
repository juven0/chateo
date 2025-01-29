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
		repository := mongorepository.InitUserRepository(mongoClient.Database(tableName))
		userServiceInstance = &UserSevice{
			repo: repository,
		}
	})
	return userServiceInstance
}

func setupUserService() {
	mongorepository.InitUserRepository(mongoClient.Database(tableName))
}

func (r *UserService)CreatUser(user *mongomodels.User) error{
	setupUserService()
	return r.repo.InsertUser(user)
}

func (r *UserService)GetUser(id string)(mongomodels.User, error){
	setupUserService()
	return r.repo.GetUser(id)
}

func (r *UserService)UpdateUser(id string, user *mongomodels.User)(*mongo.UpdateResult, error){
	setupUserService()
	return r.repo.UpdateUser(id, user)
}

func (r *UserService)DeleteUser(id string)(*mongo.DeleteResult, error){
	setupUserService()
	return r.repo.DeleteUser(id)
}