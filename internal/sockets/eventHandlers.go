package sockets

import (
	mongomodels "chat/internal/models/mongo"
	mongoservice "chat/internal/service/mongo"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/contrib/socketio"
)

func (m *WebSocketManager) handleConnect(ep *socketio.EventPayload) error {
	userID := ep.Kws.GetStringAttribute("user_id")
	log.Printf("Connection event - User: %s", userID)
	return nil
}

func (m *WebSocketManager) handleMessage(ep *socketio.EventPayload) error {
	var message MessageObject
	if err := json.Unmarshal(ep.Data, &message); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	if message.Event != "" {
		ep.Kws.Fire(message.Event, []byte(message.Data))
	}
	return nil
}

func (m *WebSocketManager) handleSendMessageTo(ep *socketio.EventPayload) error {
	ctx := context.Background()

	var newMessage mongomodels.Message
	if err := json.Unmarshal([]byte(ep.Data), &newMessage); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	if err := mongoservice.CreateMessage(&newMessage); err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}

	recipientID := newMessage.To.Hex()
	val, err := m.redisClient.Get(ctx, recipientID)
	if err != nil {
		return fmt.Errorf("failed to get recipient connection: %w", err)
	}

	ep.Kws.EmitTo(val, []byte(ep.Data), socketio.TextMessage)
	return nil
}