package spritesheet

import (
	"image/color"
	"image/png"
	"os"
	"spritely/internal/shared/topic"
	"spritely/internal/sprite"
	"spritely/internal/widgetmediator"
	"spritely/pkg/actor"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
)

type SpriteSheet struct {
	actorSystem     *actor.ActorSystem
	mediator        actor.Address
	address         actor.Address
	widgetAddress   actor.Address
	spriteAddresses []actor.Address
	sprites         [][]color.Color // used for selection only
	pixelSize       int
	spriteRowSize   int
	selected        geom.Coordinate
	rowSize         int
	columnSize      int
}

func New(as *actor.ActorSystem, mediator actor.Address, spriteSheetOffset geom.Offset, spriteSize widget.Size) actor.Address {
	ss := SpriteSheet{
		actorSystem:   as,
		mediator:      mediator,
		pixelSize:     2,
		spriteRowSize: 8,
		columnSize:    4,
		rowSize:       8,
	}

	for y := 0; y < ss.columnSize; y++ {
		var spriteRow []color.Color
		for x := 0; x < ss.rowSize; x++ {
			spriteRow = append(spriteRow, color.Transparent)
			s := sprite.New(as, mediator, geom.Offset{
				X: float64(x*ss.pixelSize*ss.spriteRowSize) + spriteSheetOffset.X,
				Y: float64(y*ss.pixelSize*ss.spriteRowSize) + spriteSheetOffset.Y,
			}, widget.Size{
				Width:  ss.pixelSize,
				Height: ss.pixelSize,
			})
			ss.spriteAddresses = append(ss.spriteAddresses, s)
		}
		ss.sprites = append(ss.sprites, spriteRow)
	}
	ss.widgetAddress = widgetmediator.NewSelectableColors(as, ss.sprites, widget.Size{ // for selection
		Width:  8 * ss.pixelSize,
		Height: 8 * ss.pixelSize,
	})
	as.Lookup(ss.widgetAddress).Message(actor.Message{
		Topic:   topic.SET_OFFSET,
		Payload: spriteSheetOffset,
	})
	// ss.actorSystem.Lookup(ss.spriteAddresses[0]).Message(actor.Message{
	// 	Topic:     topic.GET_PIXELS,
	// 	Requestor: address.CANVAS.Address(),
	// })
	return as.Register(&ss)
}

func (ss *SpriteSheet) update() {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		return
	}

	x, y := ebiten.CursorPosition()

	ss.actorSystem.Lookup(ss.widgetAddress).Message(actor.Message{
		Topic:     topic.HANDLE_CLICK,
		Requestor: ss.address,
		Payload: geom.Coordinate{
			X: x,
			Y: y,
		},
	})
}

func (ss SpriteSheet) render(dst *ebiten.Image) {
	renderMsg := actor.Message{
		Topic:   topic.RENDER,
		Payload: dst,
	}
	for _, s := range ss.spriteAddresses {
		ss.actorSystem.Lookup(s).Message(renderMsg)
	}
	ss.actorSystem.Lookup(ss.widgetAddress).Message(renderMsg)
}

func (ss SpriteSheet) coordToIndex(coord geom.Coordinate) int {
	return (coord.Y * ss.spriteRowSize) + coord.X
}

func (ss *SpriteSheet) handleClick(coordinate geom.Coordinate) {
	ss.selected = coordinate

	ss.actorSystem.Lookup(ss.mediator).Message(actor.Message{
		Topic:     topic.SET_CURRENT_SPRITE,
		Requestor: ss.address,
		Payload:   coordinate,
	})
}

func (ss SpriteSheet) save() {
	spriteSheetImg := ebiten.NewImage(
		ss.spriteRowSize*ss.pixelSize*ss.rowSize,
		ss.spriteRowSize*ss.pixelSize*ss.columnSize,
	)

	ss.copyToImage(spriteSheetImg)

	f, err := os.Create("spritely.spritesheet.png")
	if err != nil {
		log.Errorf("error saving sprite sheet: %s", err)
		return
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, spriteSheetImg)
	if err != nil {
		log.Errorf("error encoding sprite sheet: %s", err)
	}
}

func (ss SpriteSheet) copyToImage(spriteSheetImg *ebiten.Image) {
	for _, sprite := range ss.spriteAddresses {
		ss.actorSystem.Lookup(sprite).Message(actor.Message{
			Topic:   topic.DRAW,
			Payload: spriteSheetImg,
		})
	}
}
