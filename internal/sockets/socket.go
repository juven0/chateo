package sockets

import (
	"fmt"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
)

func SoketsIO(app *fiber.App) {

	clients := make(map[string]string)

	socketio.On(socketio.EventConnect, func (ep *socketio.EventPayload) {
		fmt.Printf("connection event 1 -user: %s", ep.Kws.GetStringAttribute("user_id"))
	})

	socketio.On(socketio.EventDisconnect, func (ep *socketio.EventPayload){
		delete(clients, ep.Kws.GetStringAttribute(("user_id")))

		fmt.Printf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id"))
	})

	app.Get("/ws/", socketio.New(func(kws *socketio.Websocket) {
		
	}))
}
