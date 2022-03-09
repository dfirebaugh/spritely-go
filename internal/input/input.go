package input

import (
	"spritely/internal/clipboard"
	"spritely/internal/shared/message"
	"spritely/internal/shared/topic"
	"spritely/pkg/broker"
	"spritely/pkg/geom"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Controller struct {
	broker    *broker.Broker
	clipboard *clipboard.Controller
}

func New(b *broker.Broker) *Controller {
	i := Controller{broker: b, clipboard: clipboard.New()}
	go i.mailbox()
	return &i
}

func (i Controller) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		coordinate := geom.Coordinate{
			X: x,
			Y: y,
		}
		i.broker.Publish(message.Message{
			Topic:   topic.LEFT_CLICK,
			Payload: coordinate,
		})
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		x, y := ebiten.CursorPosition()
		coordinate := geom.Coordinate{
			X: x,
			Y: y,
		}
		i.broker.Publish(message.Message{
			Topic:   topic.RIGHT_CLICK,
			Payload: coordinate,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyS) {
		i.broker.Publish(message.Message{
			Topic: topic.SAVE,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyC) {
		i.broker.Publish(message.Message{
			Topic: topic.COPY,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyV) {
		i.broker.Publish(message.Message{
			Topic:   topic.PASTE,
			Payload: i.clipboard.Pixels,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		i.broker.Publish(message.Message{
			Topic: topic.UNDO,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyY) {
		i.broker.Publish(message.Message{
			Topic: topic.REDO,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyShift) && inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		i.broker.Publish(message.Message{
			Topic: topic.REDO,
		})
	}
}
