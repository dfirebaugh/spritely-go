package sprite

import (
	"fmt"
	"image/color"
	"spritely/internal/message"
	"spritely/internal/palette"
	"spritely/internal/tool"
	"spritely/internal/topic"
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

// Encode outputs the sprite data as a hex string.
//   this might be convenient if we want to package the spritesheet in a cart
func (s *Sprite) Encode() string {
	var result string
	for _, row := range s.Widget.Elements {
		for _, p := range row {
			result = fmt.Sprintf("%s%x", result, getIndex(p.Graphic.(color.Color), palette.DefaultColors))
		}
		result = fmt.Sprintf("%s\n", result)
	}
	return result
}

func getIndex(value color.Color, slice []color.Color) int {
	if value == color.Black {
		return 0
	}
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
