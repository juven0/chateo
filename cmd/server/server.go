package main

import (
	"chat/internal/router"
	"chat/internal/sockets"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// MessageObject Basic chat message object
type MessageObject struct {
    Data  string `json:"data"`
    From  string `json:"from"`
    Event string `json:"event"`
    To    string `json:"to"`
}

func main() {
    app := fiber.New()

    app.Use(func(c *fiber.Ctx) error {

        if websocket.IsWebSocketUpgrade(c) {
            c.Locals("allowed", true)
            return c.Next()
        }
        return fiber.ErrUpgradeRequired
    })

    sockets.SoketsIO(app)

    router.Routes(app)

    log.Fatal(app.Listen(":3000"))
}