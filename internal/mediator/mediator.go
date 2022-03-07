package mediator

import (
	"spritely/internal/colorpicker"
	"spritely/internal/input"
	"spritely/internal/shared/topic"
	"spritely/internal/sprite"
	"spritely/internal/spritesheet"
	"spritely/internal/toolbar"
	"spritely/pkg/actor"
	"spritely/pkg/geom"
	"spritely/pkg/idempotency"
	"spritely/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mediator struct {
	actorSystem        *actor.ActorSystem
	address            actor.Address
	canvas             actor.Address
	inputController    actor.Address
	colorPicker        actor.Address
	toolBar            actor.Address
	spriteSheet        actor.Address
	idempotencyManager idempotency.IdempotencyManager
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
	spriteSheetSize   = widget.Size{
		Width:  spriteSize,
		Height: spriteSize,
	}
	toolBarSize = widget.Size{
		Width:  toolBarElementSize,
		Height: toolBarElementSize,
	}
	canvasSize = widget.Size{
		Width:  canvasPixelSize,
		Height: canvasPixelSize,
	}
)

func New() Mediator {
	actorSystem := actor.New()
	mediator := Mediator{
		actorSystem: actorSystem,
	}
	mediatorAddress := actorSystem.Register(&mediator)
	mediator.inputController = input.New(actorSystem, mediatorAddress)
	mediator.canvas = sprite.New(actorSystem, mediatorAddress, canvasOffset, canvasSize)
	mediator.colorPicker = colorpicker.New(actorSystem, mediatorAddress, colorPickerOffset, colorPickerPixelSize)
	mediator.toolBar = toolbar.New(actorSystem, mediatorAddress, toolbarOffset, toolBarSize)
	mediator.spriteSheet = spritesheet.New(actorSystem, mediatorAddress, spriteSheetOffset, spriteSheetSize)
	return mediator
}
func (m *Mediator) Update() {
	updateMsg := actor.Message{
		Topic: topic.UPDATE,
	}
	go m.actorSystem.Lookup(m.canvas).Message(updateMsg)
	go m.actorSystem.Lookup(m.inputController).Message(updateMsg)
	go m.actorSystem.Lookup(m.colorPicker).Message(updateMsg)
	go m.actorSystem.Lookup(m.toolBar).Message(updateMsg)
	go m.actorSystem.Lookup(m.spriteSheet).Message(updateMsg)
}
func (m *Mediator) Render(dst *ebiten.Image) {
	renderMsg := actor.Message{
		Topic:   topic.RENDER,
		Payload: dst,
	}
	m.actorSystem.Lookup(m.canvas).Message(renderMsg)
	m.actorSystem.Lookup(m.inputController).Message(renderMsg)
	m.actorSystem.Lookup(m.colorPicker).Message(renderMsg)
	m.actorSystem.Lookup(m.toolBar).Message(renderMsg)
	m.actorSystem.Lookup(m.spriteSheet).Message(renderMsg)
}
