package router

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(a *fiber.App){
	route := a.Group("/api/v1")

	conversation := route.Group("/conversation")
	ConversationRoutes(conversation)

	user := route.Group("/user", )
	userRoutes(user)

	message := route.Group("/message")
	messageRoutes(message)
}