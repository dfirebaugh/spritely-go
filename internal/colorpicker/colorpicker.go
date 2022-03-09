package colorpicker

import (
	"image/color"
	"spritely/internal/shared"
	"spritely/internal/shared/message"
	"spritely/internal/shared/topic"
	"spritely/pkg/broker"
	"spritely/pkg/geom"
	"spritely/pkg/widget"
)

type ColorPicker struct {
	broker  *broker.Broker
	palette shared.Palette
	Widget  *widget.Widget
}

func New(b *broker.Broker, offset geom.Offset, pixelSize int) *ColorPicker {
	cp := ColorPicker{
		broker:  b,
		palette: shared.DefaultColors,
	}

	cp.Widget = widget.NewSelectableColors(cp.initWidget(), geom.Size{
		Width:  pixelSize,
		Height: pixelSize,
	},
		offset,
	)

	go cp.mailbox()

	return &cp
}

func (c *ColorPicker) initWidget() [][]color.Color {
	var elements [][]color.Color
	for y, v := range c.palette {
		var pixelRow []color.Color
		for x := range v {
			pixelRow = append(pixelRow, c.palette[x][y])
		}
		elements = append(elements, pixelRow)
	}

	return elements
}

func (c *ColorPicker) handleClick(coord geom.Coordinate) {
	if !c.Widget.IsWithinBounds(coord) {
		return
	}
	local := c.Widget.ToLocalCoordinate(coord)
	c.broker.Publish(message.Message{
		Topic:   topic.SET_CURRENT_COLOR,
		Payload: c.palette[local.X][local.Y],
	})
}

func (c *ColorPicker) colorToCoord(color color.Color) geom.Coordinate {
	var coord geom.Coordinate
	for x, row := range c.palette {
		for y, clr := range row {
			if clr != color {
				continue
			}
			coord = geom.Coordinate{
				X: x,
				Y: y,
			}
			break
		}
	}
	return coord
}

func (c *ColorPicker) selectCurrentColor(clr color.Color) {
	local := c.colorToCoord(clr)
	c.Widget.SelectElement(local)
}
