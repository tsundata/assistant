package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type message struct {
	data []byte
	room string
}

type subscription struct {
	conn *connection
	room string
	h    *Hub
}

// Hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	rooms map[string]map[*connection]bool

	// Incoming message
	incoming chan message

	// Inbound messages from the connections.
	broadcast chan message

	// Register requests from the connections.
	register chan subscription

	// Unregister requests from connections.
	unregister chan subscription

	// system
	bus        event.Bus
	logger     log.Logger
	chatbotSvc pb.ChatbotSvcClient
	messageSvc pb.MessageSvcClient
}

func NewHub(bus event.Bus, logger log.Logger, chatbotSvc pb.ChatbotSvcClient, messageSvc pb.MessageSvcClient) *Hub {
	return &Hub{
		broadcast:  make(chan message, 1024),
		incoming:   make(chan message, 1024),
		register:   make(chan subscription),
		unregister: make(chan subscription),
		rooms:      make(map[string]map[*connection]bool),
		bus:        bus,
		logger:     logger,
		chatbotSvc: chatbotSvc,
		messageSvc: messageSvc,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true
		case s := <-h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		case m := <-h.incoming:
			// create message
			uuid:= util.UUID()
			_, err := h.messageSvc.Create(context.Background(), &pb.MessageRequest{
				Message: &pb.Message{
					Uuid: uuid,
					Text: util.ByteToString(m.data),
				},
			})
			if err != nil {
				h.logger.Error(err)
				continue
			}
		}
	}
}

func (h *Hub) EventHandle() {
	err := h.bus.Subscribe(context.Background(), event.MessageChannelSubject, func(msg *nats.Msg) {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			h.logger.Error(err)
			return
		}

		h.broadcast <- message{
			data: util.StringToByte(m.Text),
			room: fmt.Sprintf("group:%d", m.Receiver), // todo
		}
	})
	if err != nil {
		h.logger.Error(err)
	}
}
