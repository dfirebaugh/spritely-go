package colorpicker

import (
	"image/color"
	"spritely/internal/topic"
	"spritely/pkg/geom"
)

func (c *ColorPicker) mailbox() {
	msg := c.broker.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.LEFT_CLICK.String():
			c.handleClick(m.GetPayload().(geom.Coordinate))
		case topic.SET_CURRENT_COLOR.String():
			c.selectCurrentColor(m.GetPayload().(color.Color))
		}
	}
}
