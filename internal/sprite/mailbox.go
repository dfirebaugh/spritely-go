package sprite

import (
	"spritely/internal/shared/topic"
	"spritely/pkg/actor"
	"spritely/pkg/geom"

	log "github.com/sirupsen/logrus"
)

func (s *Sprite) Message(msg actor.Message) {
	switch msg.Topic {
	case topic.RENDER:
		s.actorSystem.Lookup(s.widgetAddress).Message(msg)
	case topic.UPDATE:
		s.actorSystem.Lookup(s.widgetAddress).Message(msg)
		s.update()
	case topic.SET_OFFSET:
		s.actorSystem.Lookup(s.widgetAddress).Message(actor.Message{
			Topic: msg.Topic,
			Payload: geom.Offset{
				X: msg.Payload.(geom.Offset).X + s.offset.X,
				Y: msg.Payload.(geom.Offset).Y + s.offset.Y,
			},
		})
	case topic.PUSH_PIXELS:
		if msg.Payload == nil {
			log.Errorf("sprite message: expected pixels but none were received")
			return
		}
		// forward push pixels msg to the widget
		s.actorSystem.Lookup(s.widgetAddress).Message(msg)
	case topic.SET_PIXEL:
		s.actorSystem.Lookup(s.widgetAddress).Message(msg)
	case topic.GET_PIXELS:
		s.actorSystem.Lookup(s.widgetAddress).Message(msg)
	case topic.HANDLE_CLICK:
		s.handleClick(msg.Payload.(geom.Coordinate))
	}
}

func (s *Sprite) SetAddress(address actor.Address) {
	s.address = address
}
