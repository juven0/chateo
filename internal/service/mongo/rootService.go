package mongoservice

import (
	"chat/configs"
	"sync"
)

var (
	mongoClient = configs.MongoConnection()
	once sync.Once
)

const tableName = "chat"