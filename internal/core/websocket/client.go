package core_websocket

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 4096
)

type Client struct {
	UserID int64
	conn   *websocket.Conn
	Send   chan []byte
}

func NewClient(userID int64, conn *websocket.Conn) *Client {
	return &Client{
		UserID: userID,
		conn:   conn,
		Send:   make(chan []byte, 256),
	}
}

func (c *Client) ReadPump(onDisconnect func()) {
	defer func() {
		onDisconnect()
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))

	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		if _, _, err := c.conn.ReadMessage(); err != nil {
			break
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
