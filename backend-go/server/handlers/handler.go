package handlers

import "github.com/BcegoqpeHu/septier-test/backend-go/server/sockets"

type Handler struct {
	Hub *sockets.Hub
}

func NewHandler() *Handler {
	return &Handler{
		Hub: sockets.NewHub(),
	}
}
