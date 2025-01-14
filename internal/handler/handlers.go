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
		return c.Status(500).JSON(fiber.Map{"error": err})
		
	}
	return c.Status(200).SendString("user created with succes")
}

func GetOneUserHandler(c *fiber.Ctx)error{
	id:= c.Params("id")

	userFund, err:= mongoservice.GetUser(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return c.Status(200).JSON(fiber.Map{"user": userFund})
}

func UpdateUserHandler(c *fiber.Ctx)error{
	id:= c.Params("id")
	user := &mongomodels.User{}

	if err := c.BodyParser(user); err != nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	result, err := mongoservice.UpdateUser(id, user)
	if err!= nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	return c.Status(200).JSON(fiber.Map{"result":result})
}

func DeleteUserHanbler(c *fiber.Ctx)error{
	id := c.Params("id")

	result, err := mongoservice.DeleteUser(id)
	if err!= nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return c.Status(200).JSON(fiber.Map{"result":result})
}
