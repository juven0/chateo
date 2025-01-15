package mongoservice

import "chat/configs"

var mongoClient = configs.MongoConnection()

const tableName = "chat"