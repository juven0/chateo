package router

import (
	"chat/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func ConversationRoutes(r fiber.Router) {
	r.Post("/conversation/create/", handler.CreatConvesationHandler)
	r.Delete("/conversation/:id", handler.DeleteConversationHandler)
}