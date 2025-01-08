package sockets

import (
	"fmt"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
)

func SoketsIO(app *fiber.App) {
	socketio.On(socketio.EventConnect, func (ep *socketio.EventPayload) {
		fmt.Println("connection event 1 -user: %s", ep.Kws.GetStringAttribute("user_id"))
	})

	socketio.On(socketio.EventDisconnect, func (ep *socketio.EventPayload){
		
		
	})
}
