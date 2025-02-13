package sockets

import (
	"chat/configs"
	mongomodels "chat/internal/models/mongo"
	mongoservice "chat/internal/service/mongo"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
)



func SoketsIO(app *fiber.App) {

	clients := make(map[string]string)

	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
        fmt.Printf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("user_id"))
    })

    socketio.On("CUSTOM_EVENT", func(ep *socketio.EventPayload) {
        fmt.Printf("Custom event - User: %s", ep.Kws.GetStringAttribute("user_id"))
    })

    socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {
        
        message := MessageObject{}

        err := json.Unmarshal(ep.Data, &message)
        if err != nil {
            fmt.Println(err)
            return
        }

        if message.Event != "" {
            ep.Kws.Fire(message.Event, []byte(message.Data))
        }

    })

    
    socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
       
        delete(clients, ep.Kws.GetStringAttribute("user_id"))
        fmt.Printf("Disconnection event - User: `%s", ep.Kws.GetStringAttribute("user_id"))
    })

    
    socketio.On(socketio.EventClose, func(ep *socketio.EventPayload) {

        delete(clients, ep.Kws.GetStringAttribute("user_id"))
        fmt.Printf("Close event - User: %s", ep.Kws.GetStringAttribute("user_id"))
    })

    socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
        fmt.Printf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id"))
    })

    socketio.On("SEND_MESSAGE_TO", func(ep *socketio.EventPayload){

        ctx := context.Background()
	    redisCilent := configs.RedisConnection()

        newMessage := mongomodels.Message{}
        err := json.Unmarshal([]byte(ep.Data), &newMessage)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("Unmarsahal JSON: %s\n", string(ep.Data))
        err = mongoservice.CreateMessage(&newMessage)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(newMessage.To.Hex())
        val, err := redisCilent.Get(ctx, newMessage.To.Hex()).Result()
        if err != nil {
            fmt.Println(err)
            return
        }

        ep.Kws.EmitTo(val, []byte(ep.Data),socketio.TextMessage )
    })

    app.Get("/ws/:id", socketio.New(func(kws *socketio.Websocket){Setup(kws, &clients )}))

}
