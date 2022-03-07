package sprite

import (
	"image/color"
	"spritely/internal/shared/request"
	"spritely/internal/shared/topic"
	"spritely/internal/widgetmediator"
	"spritely/pkg/actor"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	actorSystem   *actor.ActorSystem
	mediator      actor.Address
	offset        geom.Offset
	address       actor.Address
	widgetAddress actor.Address
}

// as *actor.ActorSystem, spriteSheetOffset geom.Offset, spriteSize widget.Size
func New(as *actor.ActorSystem, mediator actor.Address, offset geom.Offset, elementSize widget.Size) actor.Address {
	sprite := Sprite{
		actorSystem: as,
		mediator:    mediator,
	}

	sprite.widgetAddress = widgetmediator.NewWithColorElements(as, sprite.initWidget(8), elementSize)

	as.Lookup(sprite.widgetAddress).Message(actor.Message{
		Topic:   topic.SET_OFFSET,
		Payload: offset,
	})

	return as.Register(&sprite)
}

func (s *Sprite) initWidget(rowSize int) [][]color.Color {
	var elements [][]color.Color
	for i := 0; i < rowSize; i++ {
		var pixelRow []color.Color
		for j := 0; j < rowSize; j++ {
			// pixelRow = append(pixelRow, color.RGBA{0, uint8(j), uint8(i * j * 10), 255})
			pixelRow = append(pixelRow, color.Black)
		}
		elements = append(elements, pixelRow)
	}

	return elements
}

func (s *Sprite) update() {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return
	}
	x, y := ebiten.CursorPosition()
	s.actorSystem.Lookup(s.widgetAddress).Message(actor.Message{
		Topic:     topic.HANDLE_CLICK,
		Requestor: s.address,
		Payload: geom.Coordinate{
			X: x,
			Y: y,
		},
	})
}

func (s *Sprite) handleClick(coord geom.Coordinate) {
	s.actorSystem.Lookup(s.mediator).Message(actor.Message{
		Topic:     topic.SET_PIXEL,
		Requestor: s.widgetAddress,
		Payload: request.SetPixel{
			Coordinate: coord,
		},
	})
}
