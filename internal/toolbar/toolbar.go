package toolbar

import (
	_ "image/png"
	"log"
	"time"

	"github.com/dfirebaugh/spritely-go/assets"
	"github.com/dfirebaugh/spritely-go/internal/message"
	"github.com/dfirebaugh/spritely-go/internal/tool"
	"github.com/dfirebaugh/spritely-go/internal/topic"
	"github.com/dfirebaugh/spritely-go/pkg/broker"
	"github.com/dfirebaugh/spritely-go/pkg/geom"
	"github.com/dfirebaugh/spritely-go/pkg/widget"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	penImg  *ebiten.Image
	fillImg *ebiten.Image
	dragImg *ebiten.Image
	undoImg *ebiten.Image
	redoImg *ebiten.Image
	loadImg *ebiten.Image
	saveImg *ebiten.Image
	infoImg *ebiten.Image
)

func init() {
	var err error
	penImg, _, err = ebitenutil.NewImageFromFileSystem(assets.AssetFS, "icons/Pen.png")
	if err != nil {
		log.Fatal(err)
	}
	fillImg, _, err = ebitenutil.NewImageFromFileSystem(assets.AssetFS, "icons/Fill.png")
	if err != nil {
		log.Fatal(err)
	}
	dragImg, _, err = ebitenutil.NewImageFromFileSystem(assets.AssetFS, "icons/Drag.png")
	if err != nil {
		log.Fatal(err)
	}
	undoImg, _, err = ebitenutil.NewImageFromFileSystem(assets.AssetFS, "icons/Undo.png")
	if err != nil {
		log.Fatal(err)
	}
	redoImg, _, err = ebitenutil.NewImageFromFileSystem(assets.AssetFS, "icons/Redo.png")
	if err != nil {
		log.Fatal(err)
	}
	loadImg, _, err = ebitenutil.NewImageFromFileSystem(assets.AssetFS, "icons/Load.png")
	if err != nil {
		log.Fatal(err)
	}
	saveImg, _, err = ebitenutil.NewImageFromFileSystem(assets.AssetFS, "icons/Save.png")
	if err != nil {
		log.Fatal(err)
	}
	infoImg, _, err = ebitenutil.NewImageFromFileSystem(assets.AssetFS, "icons/Info.png")
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
		tb.broker.Publish(message.Message{
			Topic:   topic.SET_CURRENT_TOOL,
			Payload: tool.Pencil,
		})
	case tool.Fill:
		tb.currentTool = tool.Fill
		tb.broker.Publish(message.Message{
			Topic:   topic.SET_CURRENT_TOOL,
			Payload: tool.Fill,
		})
	case tool.Drag:
		tb.currentTool = tool.Drag
		tb.broker.Publish(message.Message{
			Topic:   topic.SET_CURRENT_TOOL,
			Payload: tool.Drag,
		})
	case tool.Undo:
		go tb.delaySwitch()
	case tool.Redo:
		go tb.delaySwitch()
	case tool.Open:
		go tb.delaySwitch()
	case tool.Save:
		tb.broker.Publish(message.Message{
			Topic: topic.SAVE,
		})
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
