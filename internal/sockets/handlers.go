package sockets

import (
	"chat/configs"
	"context"
	"fmt"

	"github.com/gofiber/contrib/socketio"
)



func Setup(kws *socketio.Websocket, clients *map[string]string) {

	userId := kws.Params("id")

	// clients[userId] = kws.UUID
	
	ctx := context.Background()
	redisCilent := configs.RedisConnection()

	err := redisCilent.Set(ctx, userId, kws.UUID, 0).Err()
	if err != nil {

	}

	val, err := redisCilent.Get(ctx, userId).Result()
	if err != nil {
		panic(err)
	}

	kws.SetAttribute("user_id", userId)

	kws.Broadcast([]byte(fmt.Sprintf("New user connected:  and UUID: %s", val)), true, socketio.TextMessage)

	kws.Emit([]byte(fmt.Sprintf("Hello user: with UUID: %s", val)), socketio.TextMessage)
}