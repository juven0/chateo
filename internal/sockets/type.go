package sockets

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/gofiber/contrib/socketio"
)

type MessageObject struct {
	Data  json.RawMessage `json:"data"`
	From  string          `json:"from"`
	Event string          `json:"event"`
	To    string          `json:"to"`
}

type WebSocketManager struct {
	clients     map[string]string
	eventHub    *EventHub
	redisClient RedisClient
}

type EventHub struct {
	handlers map[string]EventHandler
	mu       sync.RWMutex
}

type EventHandler func(*socketio.EventPayload) error

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
}

