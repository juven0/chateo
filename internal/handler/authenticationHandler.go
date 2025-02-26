package handler

import (
	mongomodels "chat/internal/models/mongo"
	"chat/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func Register(c *fiber.Ctx) error{
	user := &mongomodels.User{} 

	if err := c.BodyParser(user); err!= nil{
		return c.SendString("error parse user")
	}
	
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
		
	}
	user.Password = hashedPwd
	err = UserService.CreatUser(user)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
		
	}
	return c.Status(200).SendString("user created with succes")
}

func Login(c *fiber.Ctx)error{
	var userLogin any
	if err := c.BodyParser(userLogin);err != nil{
		return c.SendString("error parse user")
	}

	userFound ,err := UserService.GetUserByEmail(userLogin.email)
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
		
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Email), []byte(userLogin.password))
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"error": err})
		
	}
}