package toolbar

import (
	_ "image/png"
	"log"
	"spritely/internal/shared/topic"
	"spritely/internal/tool"
	"spritely/internal/widgetmediator"
	"spritely/pkg/actor"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var penImg *ebiten.Image
var fillImg *ebiten.Image
var dragImg *ebiten.Image
var undoImg *ebiten.Image
var redoImg *ebiten.Image
var loadImg *ebiten.Image
var saveImg *ebiten.Image
var infoImg *ebiten.Image

func init() {
	var err error
	penImg, _, err = ebitenutil.NewImageFromFile("./assets/icons/Pen.png")
	if err != nil {
		log.Fatal(err)
	}
	fillImg, _, err = ebitenutil.NewImageFromFile("./assets/icons/Fill.png")
	if err != nil {
		log.Fatal(err)
	}
	dragImg, _, err = ebitenutil.NewImageFromFile("./assets/icons/Drag.png")
	if err != nil {
		log.Fatal(err)
	}
	undoImg, _, err = ebitenutil.NewImageFromFile("./assets/icons/Undo.png")
	if err != nil {
		log.Fatal(err)
	}
	redoImg, _, err = ebitenutil.NewImageFromFile("./assets/icons/Redo.png")
	if err != nil {
		log.Fatal(err)
	}
	loadImg, _, err = ebitenutil.NewImageFromFile("./assets/icons/Load.png")
	if err != nil {
		log.Fatal(err)
	}
	saveImg, _, err = ebitenutil.NewImageFromFile("./assets/icons/Save.png")
	if err != nil {
		log.Fatal(err)
	}
	infoImg, _, err = ebitenutil.NewImageFromFile("./assets/icons/Info.png")
	if err != nil {
		log.Fatal(err)
	}
}

type ToolBar struct {
	actorSystem   *actor.ActorSystem
	mediator      actor.Address
	address       actor.Address
	widgetAddress actor.Address
	currentTool   tool.Tool
	offset        geom.Offset
	elementSize   widget.Size
}

func New(as *actor.ActorSystem, mediator actor.Address, offset geom.Offset, elementSize widget.Size) actor.Address {
	elements := [][]*ebiten.Image{
		{
			penImg,
			fillImg,
			dragImg,
			undoImg,
			redoImg,
			loadImg,
			saveImg,
			infoImg,
		},
	}
	toolbar := ToolBar{
		actorSystem: as,
		mediator:    mediator,
		currentTool: tool.Pencil,
		offset:      offset,
		elementSize: elementSize,
	}
	toolbar.widgetAddress = widgetmediator.NewSelectableImages(as, elements, elementSize)
	as.Lookup(toolbar.widgetAddress).Message(actor.Message{
		Topic:   topic.SET_OFFSET,
		Payload: offset,
	})

	return as.Register(&toolbar)
}

func (t *ToolBar) update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		coordinate := geom.Coordinate{
			X: x,
			Y: y,
		}

		t.actorSystem.Lookup(t.widgetAddress).Message(actor.Message{
			Topic:     topic.HANDLE_CLICK,
			Requestor: t.address,
			Payload:   coordinate,
		})
	}
}

func (t *ToolBar) delaySwitch() {
	// time.Sleep(200 * time.Millisecond)
	// t.pick(t.currentTool)
	// t.actorSystem.Lookup(t.widgetAddress).Message(actor.Message{
	// 	Topic:   topic.SET_CURRENT_TOOL,
	// 	Payload: t.currentTool,
	// })
}

func (tb *ToolBar) pick(t tool.Tool) {
	switch t {
	case tool.Pencil:
		tb.currentTool = tool.Pencil
	case tool.Fill:
		tb.currentTool = tool.Fill
	case tool.Drag:
		tb.currentTool = tool.Drag
	case tool.Undo:
		tb.delaySwitch()
	case tool.Redo:
		tb.delaySwitch()
	case tool.Open:
		tb.delaySwitch()
	case tool.Save:
		tb.delaySwitch()
	case tool.Info:
		tb.delaySwitch()
	}
	// tb.actorSystem.Lookup(tb.widgetAddress).Message(actor.Message{
	// 	Topic:   topic.SET_CURRENT_TOOL,
	// 	Payload: t,
	// })
}
