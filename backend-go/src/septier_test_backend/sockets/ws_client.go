package sockets

import (
	"github.com/gorilla/websocket"
	"io"
	"log"
	"time"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512
)

type WsClient struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *WsClient) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		if err := c.Conn.Close(); err != nil {
			log.Printf("Websocket connection closing (read pump) error: %v", err)
		}
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Printf("Websocket read deadline setting error: %v", err)
	}

	c.Conn.SetPongHandler(func(string) error {
		if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			log.Printf("Websocket read deadline setting in PONG handler error: %v", err)
		}
		return nil
	})

	for {
		if _, _, err := c.Conn.NextReader(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Connection closed... Error: %v", err)
			}
			break
		}
	}
}

func (c *WsClient) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.Conn.Close(); err != nil {
			log.Printf("Websocket connection closing (write pump) error: %v", err)
		}
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("Websocket write deadline setting error: %v", err)
			}
			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Printf("Websocket write close message error: %v", err)
				}
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("Websocket next writer getting error: %v", err)
				return
			}

			writeMessage(w, message)
			n := len(c.Send)
			for i := 0; i < n; i++ {
				writeMessage(w, []byte{'\n'})
				writeMessage(w, <-c.Send)
			}

			if err := w.Close(); err != nil {
				log.Printf("Websocket writer closing error: %v", err)
				return
			}
		case <-ticker.C:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("Websocket write deadline setting (ticker) error: %v", err)
			}
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Websocket ping message writing error: %v", err)
				return
			}
		}
	}
}

func writeMessage(w io.WriteCloser, message []byte) {
	if _, err := w.Write(message); err != nil {
		log.Println(err)
	}
}
