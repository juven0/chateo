package main

import (
	"chat/internal/router"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	clients := make(map[string]string)

	app.Use(func(c *fiber.Ctx) error{
		if websocket.IsWebSocketUpgrade(c){
			c.Locals("allowed", true)
			c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	router.Routes(app)

	if err:= app.Listen("localhost:1212"); err != nil{
		fmt.Println("error to start server : %v", err)
	}
}