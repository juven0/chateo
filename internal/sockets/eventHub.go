package sockets

import (
	"fmt"

	"github.com/gofiber/contrib/socketio"
)



func newEventHub() *EventHub {
	return &EventHub{
		handlers: make(map[string]EventHandler),
	}
}

func (h *EventHub) RegisterHandler(event string, handler EventHandler) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.handlers[event] = handler
}

func (h *EventHub) HandleEvent(event string, ep *socketio.EventPayload) error {
	h.mu.RLock()
	handler, exists := h.handlers[event]
	h.mu.RUnlock()

	if !exists {
		return fmt.Errorf("no handler registered for event: %s", event)
	}
	return handler(ep)
}
