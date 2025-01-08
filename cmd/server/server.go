package main

import (
	"chat/internal/router"
	"chat/internal/sockets"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use("/ws",func(c *fiber.Ctx) error{
		if websocket.IsWebSocketUpgrade(c){
			c.Locals("allowed", true)
			c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	sockets.SoketsIO(app)

	router.Routes(app)

	if err:= app.Listen("localhost:1212"); err != nil{
		fmt.Printf("error to start server : %v", err)
	}
}