package sockets

type Hub struct {
	Clients    map[*WsClient]bool
	Broadcast  chan []byte
	Register   chan *WsClient
	Unregister chan *WsClient

	running bool
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*WsClient]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *WsClient),
		Unregister: make(chan *WsClient),
		running:    false,
	}
}

func (h *Hub) Run() {
	h.running = true
	for h.running {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

func (h *Hub) Stop() {
	h.running = false
}
