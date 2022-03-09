package toolbar

import (
	_ "image/png"
	"log"
	"spritely/internal/tool"
	"spritely/pkg/broker"
	"spritely/pkg/geom"
	"spritely/pkg/widget"
	"time"

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
	broker      *broker.Broker
	currentTool tool.Tool
	offset      geom.Offset
	elementSize geom.Size
	Widget      *widget.Widget
}

func New(broker *broker.Broker, offset geom.Offset, elementSize geom.Size) *ToolBar {
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
		broker:      broker,
		currentTool: tool.Pencil,
		offset:      offset,
		elementSize: elementSize,
	}
	toolbar.Widget = widget.NewSelectableImages(elements, elementSize, offset)

	go toolbar.mailbox()

	return &toolbar
}

func (t *ToolBar) delaySwitch() {
	time.Sleep(100 * time.Millisecond)
	t.pick(t.currentTool)
	t.Widget.SelectElement(geom.Coordinate{
		X: int(t.currentTool),
		Y: 0,
	})
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
		go tb.delaySwitch()
	case tool.Redo:
		go tb.delaySwitch()
	case tool.Open:
		go tb.delaySwitch()
	case tool.Save:
		go tb.delaySwitch()
	case tool.Info:
		go tb.delaySwitch()
	}
}

func (t *ToolBar) handleClick(coord geom.Coordinate) {
	if !t.Widget.IsWithinBounds(coord) {
		return
	}
	local := t.Widget.ToLocalCoordinate(coord)
	t.pick(tool.Tool(local.X))
	t.Widget.SelectElement(local)
}
