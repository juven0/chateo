package main

import (
	"chat/configs"
	"chat/internal/router"
	"chat/internal/sockets"
	"log"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


func main() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }
    mongoURI := os.Getenv("mongoUri")
    configs.MongoConnection(mongoURI)
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