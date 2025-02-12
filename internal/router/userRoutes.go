package router

import (
	"chat/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func userRoutes(r fiber.Router) {
	r.Get("/",handler.WellcomeHandler)
	r.Post("/user/create/", handler.CreatUserHandler)
	r.Get("/user/:id", handler.GetOneUserHandler)
}