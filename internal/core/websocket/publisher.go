package core_websocket

type EventPublisher interface {
	Publish(userID int64, event Event)
}

type WSPublisher struct {
	hub *Hub
}

func NewWSPublisher(hub *Hub) *WSPublisher {
	return &WSPublisher{hub: hub}
}

func (p *WSPublisher) Publish(userID int64, event Event) {
	p.hub.SendEventToUser(userID, event)
}
