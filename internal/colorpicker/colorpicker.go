package colorpicker

import (
	"image/color"

	"github.com/dfirebaugh/spritely-go/internal/message"
	"github.com/dfirebaugh/spritely-go/internal/palette"
	"github.com/dfirebaugh/spritely-go/internal/topic"
	"github.com/dfirebaugh/spritely-go/pkg/broker"
	"github.com/dfirebaugh/spritely-go/pkg/geom"
	"github.com/dfirebaugh/spritely-go/pkg/widget"
)

type ColorPicker struct {
	broker  *broker.Broker
	palette palette.Palette
	Widget  *widget.Widget
}

func New(b *broker.Broker, offset geom.Offset, pixelSize int) *ColorPicker {
	cp := ColorPicker{
		broker:  b,
		palette: palette.DefaultColors,
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
	for y, v := range c.palette.To2D() {
		var pixelRow []color.Color
		for x := range v {
			pixelRow = append(pixelRow, c.palette.To2D()[x][y])
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
		Payload: c.palette.To2D()[local.X][local.Y],
	})
}

func (c *ColorPicker) colorToCoord(color color.Color) geom.Coordinate {
	var coord geom.Coordinate
	for x, row := range c.palette.To2D() {
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
