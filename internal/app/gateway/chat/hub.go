package chat

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type message struct {
	data   []byte
	room   string
	userId int64
}

type subscription struct {
	conn   *connection
	room   string
	userId int64
	h      *Hub
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
	messageSvc pb.MessageSvcClient
}

func NewHub(bus event.Bus, logger log.Logger, messageSvc pb.MessageSvcClient) *Hub {
	return &Hub{
		broadcast:  make(chan message, 1024),
		incoming:   make(chan message, 1024),
		register:   make(chan subscription),
		unregister: make(chan subscription),
		rooms:      make(map[string]map[*connection]bool),
		bus:        bus,
		logger:     logger,
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
			uuid := util.UUID()
			_, err := h.messageSvc.Create(md.BuildAuthContext(m.userId), &pb.MessageRequest{
				Message: &pb.Message{
					Uuid:         uuid,
					Text:         util.ByteToString(m.data),
					ReceiverType: m.room, // FIXME
				},
			})
			if err != nil {
				h.logger.Error(err)
				h.broadcast <- message{
					data: util.StringToByte(err.Error()),
					room: m.room,
				}
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
			room: m.ReceiverType, // FIXME
		}
	})
	if err != nil {
		h.logger.Error(err)
	}
}
