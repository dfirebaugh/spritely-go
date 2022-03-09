package sprite

import (
	"image/color"
	"spritely/internal/shared/message"
	"spritely/internal/shared/topic"
	"spritely/pkg/broker"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	broker       *broker.Broker
	Widget       *widget.Widget
	currentColor color.Color
	isCanvas     bool
}

func NewCanvas(b *broker.Broker, offset geom.Offset, elementSize geom.Size) *Sprite {
	c := New(b, offset, elementSize)
	c.isCanvas = true
	return c
}

// as *actor.ActorSystem, spriteSheetOffset geom.Offset, spriteSize widget.Size
func New(b *broker.Broker, offset geom.Offset, elementSize geom.Size) *Sprite {
	sprite := Sprite{
		broker: b,
	}

	sprite.Widget = widget.NewWithColorElements(sprite.initWidget(8), elementSize, offset)
	go sprite.mailbox()
	return &sprite
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

func (s *Sprite) Render(dst *ebiten.Image) {
	s.Widget.Render(dst)
}
func (s *Sprite) handleClick(coord geom.Coordinate) {
	if !s.isCanvas {
		return
	}
	if !s.Widget.IsWithinBounds(coord) {
		return
	}
	local := s.Widget.ToLocalCoordinate(coord)
	s.SetPixel(local)
	s.Widget.SelectElement(local)
}

func (s *Sprite) handleRightClick(coord geom.Coordinate) {
	if !s.isCanvas {
		return
	}
	if !s.Widget.IsWithinBounds(coord) {
		return
	}

	if !s.isCanvas {
		return
	}
	local := s.Widget.ToLocalCoordinate(coord)
	s.broker.Publish(message.Message{
		Topic:   topic.SET_CURRENT_COLOR,
		Payload: s.Widget.Elements[local.Y][local.X].Graphic,
	})
}

func (s *Sprite) SetPixel(coord geom.Coordinate) {
	s.Widget.Elements[coord.Y][coord.X].SetGraphic(s.currentColor)
	s.broker.Publish(message.Message{
		Topic:   topic.SET_PIXEL,
		Payload: coord,
	})
}
