package mediator

import (
	"image/color"
	"spritely/internal/shared/request"
	"spritely/internal/shared/topic"
	"spritely/internal/tool"
	"spritely/pkg/actor"
	"spritely/pkg/geom"
)

var currentColor color.Color = color.Black
var currentTool = tool.Pencil
var currentSprite = geom.Coordinate{
	X: 0,
	Y: 0,
} // local coordinate on the spritesheet

func (m *Mediator) Message(msg actor.Message) {
	if msg.Hash() == m.lastMsg {
		return
	}
	m.lastMsg = msg.Hash()
	switch msg.Topic {
	case topic.SET_CURRENT_COLOR:
		currentColor = msg.Payload.(color.Color)
	case topic.SET_PIXEL:
		req := msg.Payload.(request.SetPixel)
		setPixel := request.SetPixel{
			Coordinate:    req.Coordinate,
			Color:         currentColor,
			CurrentSprite: currentSprite,
		}
		m.actorSystem.Lookup(msg.Requestor).Message(actor.Message{
			Topic:   msg.Topic,
			Payload: setPixel,
		})
		m.actorSystem.Lookup(m.spriteSheet).Message(actor.Message{
			Topic:   msg.Topic,
			Payload: setPixel,
		})
	case topic.SET_CURRENT_TOOL:
		currentTool = msg.Payload.(tool.Tool)
		m.actorSystem.Lookup(m.toolBar).Message(msg)
	case topic.SET_CURRENT_SPRITE:
		currentSprite = geom.Coordinate{
			X: msg.Payload.(geom.Coordinate).X,
			Y: msg.Payload.(geom.Coordinate).Y,
		}

		// request that the current sprite send it's pixels to the canvas
		m.actorSystem.Lookup(m.spriteSheet).Message(actor.Message{
			Topic:     topic.GET_PIXELS,
			Requestor: m.canvas,
		})
	case topic.GET_SELECTED_SPRITE:
		m.actorSystem.Lookup(m.spriteSheet).Message(actor.Message{
			Topic:     topic.GET_PIXELS,
			Requestor: msg.Requestor,
		})
	case topic.PASTE:
		m.actorSystem.Lookup(m.spriteSheet).Message(actor.Message{
			Topic:   topic.PUSH_PIXELS,
			Payload: msg.Payload,
		})
	}
}

func (m *Mediator) SetAddress(addr actor.Address) {
	m.address = addr
}
