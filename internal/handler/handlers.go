package handler

import "github.com/gofiber/fiber/v2"

func WellcomeHandler(c *fiber.Ctx)error{
	return c.SendString("firste go fiber apps")
}