package spritesheet

import (
	"image/color"
	"spritely/assets/icons"
	"spritely/internal/message"
	"spritely/internal/sprite"
	"spritely/internal/topic"
	"spritely/pkg/broker"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	_ "image/png"

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
	page          int
	tabWidget     *widget.Widget
	tabs          [][]*ebiten.Image
}

func New(broker *broker.Broker, spriteSheetOffset geom.Offset, spriteSize geom.Size) *SpriteSheet {
	ss := SpriteSheet{
		broker:        broker,
		pixelSize:     2,
		spriteRowSize: 8,
		columnSize:    4,
		rowSize:       8,
		page:          0,
	}

	// var tabs [][]*ebiten.Image
	// var tabRow []*ebiten.Image
	// for i := 0; i < 1; i++ {
	// 	tabRow = append(tabRow, ebiten.NewImageFromImage(icons.Tabs.SubImage(image.Rectangle{
	// 		image.Point{
	// 			X: 0,
	// 			Y: 0,
	// 		},
	// 		image.Point{
	// 			X: 8,
	// 			Y: 8,
	// 		},
	// 	})))
	// }
	// tabs = append(tabs, tabRow)

	var selectionRects [][]color.Color
	for y := 0; y < ss.columnSize; y++ {
		var spriteRow []*sprite.Sprite
		var selectionRow []color.Color
		for x := 0; x < ss.rowSize; x++ {
			spriteRow = append(spriteRow, sprite.New(broker, geom.Offset{
				X: float64(x*ss.pixelSize*ss.spriteRowSize) + spriteSheetOffset.X,
				Y: float64(y*ss.pixelSize*ss.spriteRowSize) + spriteSheetOffset.Y + float64(ss.pixelSize*spriteSize.Height),
			}, geom.Size{
				Width:  ss.pixelSize,
				Height: ss.pixelSize,
			}))
			selectionRow = append(selectionRow, color.Transparent)
		}
		ss.sprites = append(ss.sprites, spriteRow)
		selectionRects = append(selectionRects, selectionRow)
	}
	ss.Widget = widget.NewSelectableColors(
		selectionRects,
		geom.Size{
			Width:  8 * ss.pixelSize,
			Height: 8 * ss.pixelSize,
		},
		geom.Offset{
			X: spriteSheetOffset.X,
			Y: spriteSheetOffset.Y + float64(ss.pixelSize*spriteSize.Height),
		})

	ss.tabWidget = widget.NewSelectableImages(
		[][]*ebiten.Image{{icons.Tabs}},
		geom.Size{
			Height: 8 * ss.pixelSize,
			Width:  8 * ss.pixelSize,
		},
		spriteSheetOffset)

	go ss.mailbox()
	return &ss
}

func (ss SpriteSheet) Render(dst *ebiten.Image) {
	for _, row := range ss.sprites {
		for _, s := range row {
			s.Render(dst)
		}
	}

	ss.tabWidget.Render(dst)
	ss.Widget.Render(dst)
}

func (ss *SpriteSheet) handleClick(coordinate geom.Coordinate) {
	// if ss.tabWidget.IsWithinBounds(coordinate) {
	// 	local := ss.tabWidget.ToLocalCoordinate(coordinate)
	// 	ss.tabWidget.SelectElement(local)
	// }
	if !ss.Widget.IsWithinBounds(coordinate) {
		return
	}
	ss.selectSprite(coordinate)
}

func (ss SpriteSheet) GetSprite(local geom.Coordinate) *sprite.Sprite {
	return ss.sprites[ss.selected.Y][ss.selected.X]
}

func (ss *SpriteSheet) selectSprite(coordinate geom.Coordinate) {
	ss.selectFromLocal(ss.Widget.ToLocalCoordinate(coordinate))
}

func (ss *SpriteSheet) selectFromLocal(local geom.Coordinate) {
	if !ss.Widget.IsWithinLocalBounds(local) {
		return
	}
	ss.selected = local
	ss.Widget.SelectElement(local)
	ss.broker.Publish(message.Message{
		Topic:   topic.UPDATE_CANVAS,
		Payload: ss.GetSprite(ss.selected).Widget.Elements,
	})
}
