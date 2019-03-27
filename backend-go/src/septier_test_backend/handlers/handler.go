package handlers

import "septier_test_backend/sockets"

type Handler struct {
	Hub *sockets.Hub
}

func NewHandler() *Handler {
	return &Handler{
		Hub: sockets.NewHub(),
	}
}
