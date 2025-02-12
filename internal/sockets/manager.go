package sockets

import (
	"chat/configs"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/contrib/socketio"
)

func NewWebSocketManager(redisClient RedisClient) *WebSocketManager {
	return &WebSocketManager{
		clients:     make(map[string]string),
		eventHub:    newEventHub(),
		redisClient: redisClient,
	}
}

func (m *WebSocketManager) InitializeEvents() {
	m.eventHub.RegisterHandler(socketio.EventConnect, m.handleConnect)
	m.eventHub.RegisterHandler(socketio.EventMessage, m.handleMessage)
	// m.eventHub.RegisterHandler(socketio.EventDisconnect, m.handleDisconnect)
	// m.eventHub.RegisterHandler(socketio.EventClose, m.handleClose)
	// m.eventHub.RegisterHandler(socketio.EventError, m.handleError)
	m.eventHub.RegisterHandler("SEND_MESSAGE_TO", m.handleSendMessageTo)
	// m.eventHub.RegisterHandler("CUSTOM_EVENT", m.handleCustomEvent)
}


func (m *WebSocketManager) Setup(ws *socketio.Websocket) {
    ctx := context.Background()
    userId := ws.Params("id")
    
    if err := m.registerUser(ctx, userId, ws.UUID); err != nil {
        log.Printf("Erreur lors de l'enregistrement de l'utilisateur: %v", err)
        return
    }

    ws.SetAttribute("user_id", userId)
    
    m.sendWelcomeMessages(ws, userId)
}

func (m *WebSocketManager) registerUser(ctx context.Context, userId, uuid string) error {
    // Définir un timeout pour l'opération Redis
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

	redisCilent := configs.RedisConnection()

	err := redisCilent.Set(ctx, userId, uuid, 0).Err()
	if err != nil {

	}

	val, err := redisCilent.Get(ctx, userId).Result()
	if err != nil {
		fmt.Println(err)
	}

    if val != uuid {
        return fmt.Errorf("incohérence UUID: attendu %s, reçu %s", uuid, val)
    }

    return nil
}


func (m *WebSocketManager) sendWelcomeMessages(ws *socketio.Websocket, userId string) {
    
    broadcastMsg := fmt.Sprintf("New user connected: %s with UUID: %s", userId, ws.UUID)
    ws.Broadcast([]byte(broadcastMsg), true, socketio.TextMessage)

    welcomeMsg := fmt.Sprintf("Hello user: %s with UUID: %s", userId, ws.UUID)
    ws.Emit([]byte(welcomeMsg), socketio.TextMessage)
}

