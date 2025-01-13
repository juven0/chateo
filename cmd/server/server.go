package main

import (
	"chat/internal/router"
	"chat/internal/sockets"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)


func main() {
    app := fiber.New()

    app.Use(func(c *fiber.Ctx) error {

        if websocket.IsWebSocketUpgrade(c) {
            c.Locals("allowed", true)
            return c.Next()
        }
        return c.Next()
        // return fiber.ErrUpgradeRequired
    })

    sockets.SoketsIO(app)

    router.Routes(app)

    log.Fatal(app.Listen(":3000"))
}