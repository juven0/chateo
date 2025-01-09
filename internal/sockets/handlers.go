package sockets

import (
	"fmt"

	"github.com/gofiber/contrib/socketio"
)

func Setup(kws *socketio.Websocket, clients *map[string]string) {

	userId := kws.Params("id")

	// clients[userId] = kws.UUID

	kws.SetAttribute("user_id", userId)

	kws.Broadcast([]byte(fmt.Sprintf("New user connected:  and UUID: %s", kws.UUID)), true, socketio.TextMessage)

	kws.Emit([]byte(fmt.Sprintf("Hello user: with UUID: %s", kws.UUID)), socketio.TextMessage)
}