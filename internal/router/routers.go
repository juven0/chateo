package router

import (
	"chat/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(a *fiber.App){
	route := a.Group("/api/v1")

	route.Get("/",handler.WellcomeHandler)
	route.Post("/user/create/", handler.CreatUserHandler)
}