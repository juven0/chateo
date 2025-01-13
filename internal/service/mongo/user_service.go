package mongoservice

import (
	"chat/configs"
	mongomodels "chat/internal/models/mongo"
	mongorepository "chat/internal/repository/mongo"
	"os"

	"github.com/joho/godotenv"
)

func CreatUser(user mongomodels.User){
	err := godotenv.Load()
    if err != nil {
        panic(err)
    }
    mongoURI := os.Getenv("mongoUri")
	mongoClient := configs.MongoConnection(mongoURI)
	mongorepository.InitUserRepository(mongoClient.Database("chat"))
}