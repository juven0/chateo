package router

import (
	"chat/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func messageRoutes(r fiber.Router){
	r.Post("/message/create/",handler.CreateMessage)
}