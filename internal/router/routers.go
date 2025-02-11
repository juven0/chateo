package router

import (
	"chat/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(a *fiber.App){
	route := a.Group("/api/v1")
	route.Group("/conversation",)
	route.Get("/",handler.WellcomeHandler)
	route.Post("/user/create/", handler.CreatUserHandler)
	route.Get("/user/:id", handler.GetOneUserHandler)
	route.Post("/message/create/",handler.CreateMessage)
	route.Post("/conversation/create/", handler.CreatConvesationHandler)
	route.Delete("/conversation/:id", handler.DeleteConversationHandler)
}