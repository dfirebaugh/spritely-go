package input

import (
	"spritely/internal/clipboard"
	"spritely/internal/shared/topic"
	"spritely/pkg/actor"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputController struct {
	actorSystem *actor.ActorSystem
	mediator    actor.Address
	address     actor.Address
	clipboard   actor.Address
}

func New(as *actor.ActorSystem, mediator actor.Address) actor.Address {
	return as.Register(&InputController{actorSystem: as, mediator: mediator,
		clipboard: clipboard.New(as),
	})
}

func (i InputController) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyS) {
		println("save")
		// i.actorSystem.Lookup(i.mediator).Message(actor.Message{
		// 	Topic: topic.SAVE,
		// })
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyC) {
		println("copy")
		i.actorSystem.Lookup(i.mediator).Message(actor.Message{
			Topic:     topic.GET_SELECTED_SPRITE,
			Requestor: i.clipboard,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyV) {
		println("paste")
		i.actorSystem.Lookup(i.clipboard).Message(actor.Message{
			Topic:     topic.PASTE,
			Requestor: i.mediator,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyZ) {
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyY) {
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyShift) && inpututil.IsKeyJustPressed(ebiten.KeyZ) {
	}
}
