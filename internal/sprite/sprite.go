package sprite

import (
	"image/color"
	"spritely/internal/shared/message"
	"spritely/internal/shared/topic"
	"spritely/internal/tool"
	"spritely/pkg/broker"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	broker       *broker.Broker
	Widget       *widget.Widget
	currentColor color.Color
	currentTool  tool.Tool
	isCanvas     bool
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
	if !s.Widget.IsWithinBounds(coord) {
		return
	}
	s.executeCanvasOp(coord)
}

func (s *Sprite) handleRightClick(coord geom.Coordinate) {
	if !s.Widget.IsWithinBounds(coord) {
		return
	}
	if !s.isCanvas {
		return
	}
	local := s.Widget.ToLocalCoordinate(coord)
	s.setCurrentColor(s.Widget.GetElement(local).Graphic.(color.Color))
}

func (s *Sprite) setPixel(coord geom.Coordinate) {
	s.Widget.Elements[coord.Y][coord.X].SetGraphic(s.currentColor)
	s.broker.Publish(message.Message{
		Topic:   topic.SET_PIXEL,
		Payload: coord,
	})
}
