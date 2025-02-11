package handler

import (
	mongomodels "chat/internal/models/mongo"
	mongoservice "chat/internal/service/mongo"

	"github.com/gofiber/fiber/v2"
)

var convesationService = mongoservice.GetConversationServiceInstance()

func CreatConvesationHandler(c *fiber.Ctx)error{
	newConvesation := &mongomodels.Conversation{}

	if err := c.BodyParser(newConvesation); err != nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	err := convesationService.CreatConvesation(newConvesation)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).SendString("user created with succes")
}

func DeleteConversationHandler(c *fiber.Ctx)error{
	id := c.Params("id")

	result, err := convesationService.DeleteConverstion(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(fiber.Map{"result":result})
}