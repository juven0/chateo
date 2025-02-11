package handler

import (
	mongomodels "chat/internal/models/mongo"
	mongoservice "chat/internal/service/mongo"

	"github.com/gofiber/fiber/v2"
)

func CreateMessage(c *fiber.Ctx)error{
	message := &mongomodels.Message{}

	if err:= c.BodyParser(message); err != nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	err := mongoservice.CreateMessage(message)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return c.Status(200).SendString("user created with succes")
}
