package mediator

import (
	"spritely/internal/colorpicker"
	"spritely/internal/input"
	"spritely/internal/sprite"
	"spritely/internal/spritesheet"
	"spritely/internal/toolbar"
	"spritely/internal/widgets"
	"spritely/pkg/broker"
	"spritely/pkg/geom"
	"spritely/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mediator struct {
	lastMsg         string
	broker          *broker.Broker
	colorPicker     *colorpicker.ColorPicker
	toolBar         *toolbar.ToolBar
	spriteSheet     *spritesheet.SpriteSheet
	inputController *input.Controller
	canvas          *sprite.Sprite
	widgets         []*widget.Widget
}

const (
	toolBarElementSize   = 16
	spriteSize           = 8
	canvasNPixels        = 8
	canvasPixelSize      = 20
	colorPickerPixelSize = 16
)

var (
	canvasOffset      = geom.Offset{X: 0, Y: 0}
	toolbarOffset     = geom.Offset{X: 0, Y: float64(canvasNPixels * canvasPixelSize)}
	colorPickerOffset = geom.Offset{X: float64(canvasNPixels * canvasPixelSize), Y: 0}
	spriteSheetOffset = geom.Offset{X: 0, Y: toolbarOffset.Y + float64(toolBarElementSize)}
	spriteSheetSize   = geom.Size{
		Width:  spriteSize,
		Height: spriteSize,
	}
	toolBarSize = geom.Size{
		Width:  toolBarElementSize,
		Height: toolBarElementSize,
	}
	canvasSize = geom.Size{
		Width:  canvasPixelSize,
		Height: canvasPixelSize,
	}
)

func New() Mediator {
	b := broker.NewBroker()
	mediator := Mediator{
		broker: b,
		// inputController: input.New(b),
		// colorPicker:     colorpicker.New(b, colorPickerOffset, colorPickerPixelSize),
		// toolBar:         toolbar.New(b, toolbarOffset, toolBarSize),
		// spriteSheet:     spritesheet.New(b, spriteSheetOffset, spriteSheetSize),
		// canvas:          sprite.NewCanvas(b, canvasOffset, canvasSize),
		widgets: widgets.LoadWidgets(),
	}

	go mediator.broker.Start()

	return mediator
}
func (m *Mediator) Update() {
	// m.inputController.Update()
}
func (m *Mediator) Render(dst *ebiten.Image) {
	// m.colorPicker.Widget.Render(dst)
	// m.toolBar.Widget.Render(dst)
	// m.spriteSheet.Render(dst)
	// m.canvas.Render(dst)
	for _, w := range m.widgets {
		w.Render(dst)
	}
}
