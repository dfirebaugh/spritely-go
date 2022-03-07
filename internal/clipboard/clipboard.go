package clipboard

import (
	"spritely/internal/shared/topic"
	"spritely/pkg/actor"
	"spritely/pkg/widget"

	log "github.com/sirupsen/logrus"
)

type ClipBoardController struct {
	actorSystem *actor.ActorSystem
	address     actor.Address
	Pixels      [][]*widget.Element
}

func New(actorSystem *actor.ActorSystem) actor.Address {
	return actorSystem.Register(&ClipBoardController{
		actorSystem: actorSystem,
	})
}

func (c *ClipBoardController) Message(msg actor.Message) {
	switch msg.Topic {
	case topic.PUSH_PIXELS:
		// receive pixels into the clipboard
		if msg.Payload == nil {
			log.Errorf("clipboard message: expected pixels but none were received")
			return
		}

		c.Pixels = msg.Payload.([][]*widget.Element)
	case topic.GET_PIXELS:
		// send the clipboards pixels back to the requestor
		c.actorSystem.Lookup(msg.Requestor).Message(actor.Message{
			Topic:     topic.PUSH_PIXELS,
			Payload:   c.deepCopy(), // break the address linking
			Requestor: c.address,
		})
	case topic.PASTE:
		c.actorSystem.Lookup(msg.Requestor).Message(actor.Message{
			Topic:     topic.GET_SELECTED_SPRITE,
			Requestor: c.address,
		})
	}
}

func (c *ClipBoardController) deepCopy() [][]*widget.Element {
	var pixels [][]*widget.Element
	for _, row := range c.Pixels {
		var pixelRow []*widget.Element
		for _, p := range row {
			newPixel := &widget.Element{
				Graphic: p.Graphic,
			}
			pixelRow = append(pixelRow, newPixel)
		}
		pixels = append(pixels, pixelRow)
	}
	return pixels
}
func (c *ClipBoardController) SetAddress(address actor.Address) {
	c.address = address
}
