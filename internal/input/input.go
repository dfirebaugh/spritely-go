package input

import (
	"spritely/internal/clipboard"
	"spritely/internal/message"
	"spritely/internal/tool"
	"spritely/internal/topic"
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
		return
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
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		i.broker.Publish(message.Message{
			Topic:   topic.SET_CURRENT_TOOL,
			Payload: tool.Fill,
		})
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		i.broker.Publish(message.Message{
			Topic:   topic.SET_CURRENT_TOOL,
			Payload: tool.Pencil,
		})
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyS) {
		i.broker.Publish(message.Message{
			Topic: topic.SAVE,
		})
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyC) {
		i.broker.Publish(message.Message{
			Topic: topic.COPY,
		})
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyV) {
		i.broker.Publish(message.Message{
			Topic:   topic.UPDATE_CANVAS,
			Payload: i.clipboard.Pixels,
		})
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		i.broker.Publish(message.Message{
			Topic: topic.UNDO,
		})
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyY) {
		i.broker.Publish(message.Message{
			Topic: topic.REDO,
		})
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && ebiten.IsKeyPressed(ebiten.KeyShift) && inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		i.broker.Publish(message.Message{
			Topic: topic.REDO,
		})
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		i.broker.Publish(message.Message{
			Topic: topic.LEFT,
		})
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		i.broker.Publish(message.Message{
			Topic: topic.RIGHT,
		})
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		i.broker.Publish(message.Message{
			Topic: topic.UP,
		})
		return
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		i.broker.Publish(message.Message{
			Topic: topic.DOWN,
		})
		return
	}
}
