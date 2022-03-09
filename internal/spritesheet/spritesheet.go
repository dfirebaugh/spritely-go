package spritesheet

import (
	"image/color"
	"spritely/internal/shared/message"
	"spritely/internal/shared/topic"
	"spritely/internal/sprite"
	"spritely/pkg/broker"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	broker        *broker.Broker
	Widget        *widget.Widget
	sprites       [][]*sprite.Sprite
	currentColor  color.Color
	pixelSize     int
	spriteRowSize int
	selected      geom.Coordinate
	rowSize       int
	columnSize    int
}

func New(broker *broker.Broker, spriteSheetOffset geom.Offset, spriteSize geom.Size) *SpriteSheet {
	ss := SpriteSheet{
		broker:        broker,
		pixelSize:     2,
		spriteRowSize: 8,
		columnSize:    4,
		rowSize:       8,
	}

	var selectionRects [][]color.Color
	for y := 0; y < ss.columnSize; y++ {
		var spriteRow []*sprite.Sprite
		var selectionRow []color.Color
		for x := 0; x < ss.rowSize; x++ {
			// spriteRow = append(spriteRow, color.Transparent)
			spriteRow = append(spriteRow, sprite.New(broker, geom.Offset{
				X: float64(x*ss.pixelSize*ss.spriteRowSize) + spriteSheetOffset.X,
				Y: float64(y*ss.pixelSize*ss.spriteRowSize) + spriteSheetOffset.Y,
			}, geom.Size{
				Width:  ss.pixelSize,
				Height: ss.pixelSize,
			}))
			selectionRow = append(selectionRow, color.Transparent)
		}
		ss.sprites = append(ss.sprites, spriteRow)
		selectionRects = append(selectionRects, selectionRow)
	}
	ss.Widget = widget.NewSelectableColors(selectionRects, geom.Size{
		Width:  8 * ss.pixelSize,
		Height: 8 * ss.pixelSize,
	}, spriteSheetOffset)

	go ss.mailbox()
	return &ss
}

func (ss SpriteSheet) Render(dst *ebiten.Image) {
	for _, row := range ss.sprites {
		for _, s := range row {
			s.Render(dst)
		}
	}
	ss.Widget.Render(dst)
}

func (ss *SpriteSheet) handleClick(coordinate geom.Coordinate) {
	if !ss.Widget.IsWithinBounds(coordinate) {
		return
	}
	ss.selected = ss.Widget.ToLocalCoordinate(coordinate)
	ss.Widget.SelectElement(ss.selected)
	ss.broker.Publish(message.Message{
		Topic:   topic.SET_CANVAS,
		Payload: ss.GetSprite(ss.selected).Widget.Elements,
	})
}

func (ss SpriteSheet) GetSprite(local geom.Coordinate) *sprite.Sprite {
	return ss.sprites[ss.selected.Y][ss.selected.X]
}
