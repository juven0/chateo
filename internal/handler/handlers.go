package handler

import (
	mongomodels "chat/internal/models/mongo"
	mongoservice "chat/internal/service/mongo"

	"github.com/gofiber/fiber/v2"
)

func WellcomeHandler(c *fiber.Ctx)error{
	return c.SendString("firste go fiber apps")
}

func CreatUserHandler(c *fiber.Ctx)error{
	user := &mongomodels.User{} 

	if err := c.BodyParser(user); err!= nil{
		return c.SendString("error parse user")
	}

	err := mongoservice.CreatUser(user)
	if err != nil{
		return c.Status(500).SendString("error to create user")
		
	}
	return c.Status(200).SendString("user created with succes")
}