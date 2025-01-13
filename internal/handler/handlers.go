package handler

import (
	mongomodels "chat/internal/models/mongo"
	"log"

	"github.com/gofiber/fiber/v2"
)

func WellcomeHandler(c *fiber.Ctx)error{
	return c.SendString("firste go fiber apps")
}

func CreatUserHandler(c *fiber.Ctx)error{
	user := new(mongomodels.User) 

	if err := c.BodyParser(user); err!= nil{
		return c.SendString("error parse user")
	}
	log.Println(user.UserName)
	return c.SendString("firste go fiber apps")
}