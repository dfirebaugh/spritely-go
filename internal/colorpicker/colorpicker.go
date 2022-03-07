package colorpicker

import (
	"image/color"
	"spritely/internal/shared"
	"spritely/internal/shared/topic"
	"spritely/internal/widgetmediator"
	"spritely/pkg/actor"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type ColorPicker struct {
	actor            *actor.ActorSystem
	mediator         actor.Address
	address          actor.Address
	currentSelection geom.Coordinate
	palette          shared.Palette
	widgetAddress    actor.Address
}

func New(as *actor.ActorSystem, mediator actor.Address, offset geom.Offset, pixelSize int) actor.Address {
	cp := ColorPicker{
		actor:    as,
		mediator: mediator,
		palette:  shared.DefaultColors,
	}

	cp.widgetAddress = widgetmediator.NewSelectableColors(as, cp.initWidget(), widget.Size{
		Width:  pixelSize,
		Height: pixelSize,
	})

	as.Lookup(cp.widgetAddress).Message(actor.Message{
		Topic:   topic.SET_OFFSET,
		Payload: offset,
	})

	return as.Register(&cp)
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

func (c *ColorPicker) update() {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return
	}
	x, y := ebiten.CursorPosition()
	c.actor.Lookup(c.widgetAddress).Message(actor.Message{
		Topic:     topic.HANDLE_CLICK,
		Requestor: c.address,
		Payload: geom.Coordinate{
			X: x,
			Y: y,
		},
	})
}

func (c *ColorPicker) pick(coord geom.Coordinate) {
	c.actor.Lookup(c.mediator).Message(actor.Message{
		Topic:   topic.SET_CURRENT_COLOR,
		Payload: c.palette[coord.X][coord.Y],
	})
}
