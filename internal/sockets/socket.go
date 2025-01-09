package sockets

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
)

type MessageObject struct {
    Data  string `json:"data"`
    From  string `json:"from"`
    Event string `json:"event"`
    To    string `json:"to"`
}

func SoketsIO(app *fiber.App) {

	clients := make(map[string]string)

	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
        fmt.Printf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("user_id"))
    })

    // Custom event handling supported
    socketio.On("CUSTOM_EVENT", func(ep *socketio.EventPayload) {
        fmt.Printf("Custom event - User: %s", ep.Kws.GetStringAttribute("user_id"))
       
    })

    // On message event
    socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {

        fmt.Printf("Message event - User: %s - Message: %s", ep.Kws.GetStringAttribute("user_id"), string(ep.Data))

        message := MessageObject{}

        // Unmarshal the json message
        // {
        //  "from": "<user-id>",
        //  "to": "<recipient-user-id>",
        //  "event": "CUSTOM_EVENT",
        //  "data": "hello"
        //}
        err := json.Unmarshal(ep.Data, &message)
        if err != nil {
            fmt.Println(err)
            return
        }

        if message.Event != "" {
            ep.Kws.Fire(message.Event, []byte(message.Data))
        }

        err = ep.Kws.EmitTo(clients[message.To], ep.Data, socketio.TextMessage)
        if err != nil {
            fmt.Println(err)
        }
    })

    
    socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
       
        delete(clients, ep.Kws.GetStringAttribute("user_id"))
        fmt.Printf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id"))
    })

    
    socketio.On(socketio.EventClose, func(ep *socketio.EventPayload) {

        delete(clients, ep.Kws.GetStringAttribute("user_id"))
        fmt.Printf("Close event - User: %s", ep.Kws.GetStringAttribute("user_id"))
    })

    socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
        fmt.Printf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id"))
    })

    app.Get("/ws/:id", socketio.New(func(kws *socketio.Websocket){Setup(kws, &clients )}))

}
