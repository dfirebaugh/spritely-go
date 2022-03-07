package spritesheet

import (
	"spritely/internal/shared/topic"
	"spritely/pkg/actor"
	"spritely/pkg/geom"

	"github.com/hajimehoshi/ebiten/v2"
)

func (ss *SpriteSheet) Message(msg actor.Message) {
	switch msg.Topic {
	case topic.SET_OFFSET:
		ss.actorSystem.Lookup(ss.widgetAddress).Message(msg)
		for _, s := range ss.spriteAddresses {
			ss.actorSystem.Lookup(s).Message(msg)
		}
	case topic.RENDER:
		ss.render(msg.Payload.(*ebiten.Image))
	case topic.UPDATE:
		ss.update()
	case topic.SAVE:
		ss.save()
	case topic.HANDLE_CLICK:
		ss.handleClick(msg.Payload.(geom.Coordinate))
	case topic.SET_PIXEL:
		ss.actorSystem.Lookup(ss.spriteAddresses[ss.coordToIndex(ss.selected)]).Message(msg)
	case topic.GET_PIXELS:
		// ask the selected sprite to send it's pixels to the requestor
		ss.actorSystem.Lookup(ss.spriteAddresses[ss.coordToIndex(ss.selected)]).Message(msg)
	}
}

func (s *SpriteSheet) SetAddress(address actor.Address) {
	s.address = address
}
