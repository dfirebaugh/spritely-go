package colorpicker

import (
	"spritely/internal/shared/topic"
	"spritely/pkg/actor"
	"spritely/pkg/geom"
)

func (c *ColorPicker) Message(msg actor.Message) {
	switch msg.Topic {
	case topic.SET_OFFSET:
		c.actor.Lookup(c.widgetAddress).Message(msg)
	case topic.RENDER:
		c.actor.Lookup(c.widgetAddress).Message(msg)
	case topic.UPDATE:
		c.update()
	case topic.SET_CURRENT_COLOR:
		c.pick(msg.Payload.(geom.Coordinate))
	case topic.HANDLE_CLICK:
		c.pick(msg.Payload.(geom.Coordinate))
	}
}

func (c *ColorPicker) SetAddress(address actor.Address) {
	c.address = address
}
