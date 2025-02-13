package handler

import (
	mongomodels "chat/internal/models/mongo"

	"github.com/gofiber/fiber/v2"
)


func Register(c *fiber.Ctx) error{
	user := &mongomodels.User{} 

	if err := c.BodyParser(user); err!= nil{
		return c.SendString("error parse user")
	}

	err := UserService.CreatUser(user)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
		
	}
	return c.Status(200).SendString("user created with succes")
}

func Login(c *fiber.Ctx)error{
	
}