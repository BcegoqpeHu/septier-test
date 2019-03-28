package handlers

import (
	"server/sockets"
)

type Handler struct {
	Hub *sockets.Hub
}

func NewHandler() *Handler {
	return &Handler{
		Hub: sockets.NewHub(),
	}
}
