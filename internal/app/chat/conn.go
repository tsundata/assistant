package chat

import (
	"encoding/json"
	"github.com/gofiber/websocket/v2"
	"github.com/tsundata/assistant/api/pb"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		s.h.unregister <- s
		_ = c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	_ = c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { _ = c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		var data wsData
		err = json.Unmarshal(msg, &data)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}
		m := message{pb.Message{Text: data.Data, Type: data.Type}, s.roomId, s.userId}
		s.h.incoming <- m
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[gateway] ws conn write", mt, payload, err)
		}
	}()
	_ = c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		// _ = c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				_ = c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func ServeWs(h *Hub, conn *websocket.Conn, roomId, userId int64) {
	c := &connection{send: make(chan []byte, 256), ws: conn}
	s := subscription{c, roomId, userId, h}
	h.register <- s
	go s.writePump()
	s.readPump()
}

type wsData struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
