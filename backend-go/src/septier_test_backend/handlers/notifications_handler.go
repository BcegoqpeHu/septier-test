package handlers

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"log"
	"septier_test_backend/sockets"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func (h *Handler) NotificationsHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Printf("Websocket upgrading error: %v", err)
		return err
	}

	client := &sockets.WsClient{Hub: h.Hub, Conn: ws, Send: make(chan []byte, 256)}
	h.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

	return nil
}
