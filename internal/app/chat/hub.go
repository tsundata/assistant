package chat

import (
	"context"
	"encoding/json"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/service"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"go.uber.org/zap"
)

type message struct {
	data   pb.Message
	roomId int64
	userId int64
}

type subscription struct {
	conn   *connection
	roomId int64
	userId int64
	h      *Hub
}

// Hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	rooms map[int64]map[*connection]bool

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
	messageSvc service.MessageSvcClient
}

func NewHub(bus event.Bus, logger log.Logger, messageSvc service.MessageSvcClient) *Hub {
	return &Hub{
		broadcast:  make(chan message, 1024),
		incoming:   make(chan message, 1024),
		register:   make(chan subscription),
		unregister: make(chan subscription),
		rooms:      make(map[int64]map[*connection]bool),
		bus:        bus,
		logger:     logger,
		messageSvc: messageSvc,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.roomId]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.roomId] = connections
			}
			h.rooms[s.roomId][s.conn] = true
			h.logger.Info("[chat] hub register", zap.Any("room", s.roomId))
		case s := <-h.unregister:
			connections := h.rooms[s.roomId]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.roomId)
					}
				}
			}
			h.logger.Info("[chat] hub unregister", zap.Any("room", s.roomId))
		case m := <-h.broadcast:
			connections := h.rooms[m.roomId]
			data, err := json.Marshal(m.data)
			if err != nil {
				h.logger.Error(err)
				continue
			}
			for c := range connections {
				select {
				case c.send <- data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.roomId)
					}
				}
			}
			h.logger.Info("[chat] hub broadcast", zap.Any("room", m.roomId), zap.Any("data", m.data))
		case m := <-h.incoming:
			// create message
			_, err := h.messageSvc.Create(md.BuildAuthContext(m.userId), &pb.MessageRequest{
				Message: &pb.Message{
					UserId:  m.userId,
					Text:    m.data.Text,
					Type:    m.data.Type,
					GroupId: m.roomId,
				},
			})
			if err != nil {
				h.logger.Error(err)
				h.broadcast <- message{
					data:   pb.Message{Text: err.Error(), Direction: enum.MessageIncomingDirection},
					roomId: m.roomId,
				}
				continue
			}
			h.logger.Info("[chat] hub incoming", zap.Any("room", m.roomId), zap.Any("data", m.data))
		}
	}
}

func (h *Hub) EventHandle() {
	err := h.bus.Subscribe(context.Background(), enum.Message, event.MessageChannelSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		h.broadcast <- message{
			data:   m,
			roomId: m.GroupId,
		}
		return nil
	})
	if err != nil {
		h.logger.Error(err)
	}
}
