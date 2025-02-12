package main

import (
	"chat/internal/sockets"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber"
)

func SetupWebSocket(app *fiber.App, redisClient sockets.RedisClient) {
	wsManager := sockets.NewWebSocketManager(redisClient)
	wsManager.InitializeEvents()

	app.Get("/ws/:id",socketio.New(func(kws *socketio.Websocket) {
		wsManager.Setup(kws)
	}))
}