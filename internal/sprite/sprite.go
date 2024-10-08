package sprite

import (
	"fmt"
	"image/color"

	"github.com/dfirebaugh/spritely-go/internal/message"
	"github.com/dfirebaugh/spritely-go/internal/palette"
	"github.com/dfirebaugh/spritely-go/internal/tool"
	"github.com/dfirebaugh/spritely-go/internal/topic"
	"github.com/dfirebaugh/spritely-go/pkg/broker"
	"github.com/dfirebaugh/spritely-go/pkg/geom"
	"github.com/dfirebaugh/spritely-go/pkg/widget"

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
			pixelRow = append(pixelRow, palette.Black)
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
//
//	this might be convenient if we want to package the spritesheet in a cart
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

func (s *Sprite) EncodeSelected(start, end geom.Coordinate) string {
	var result string
	for y := start.Y; y <= end.Y && y < len(s.Widget.Elements); y++ {
		for x := start.X; x <= end.X && x < len(s.Widget.Elements[y]); x++ {
			pixel := s.Widget.Elements[y][x]
			result = fmt.Sprintf("%s%x", result, getIndex(pixel.Graphic.(color.Color), palette.DefaultColors))
		}
		result = fmt.Sprintf("%s\n", result)
	}
	return result
}

func getIndex(value color.Color, slice []color.Color) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
