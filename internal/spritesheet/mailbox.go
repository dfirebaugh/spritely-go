package spritesheet

import (
	"image/color"
	"spritely/internal/message"
	"spritely/internal/topic"
	"spritely/pkg/geom"
	"spritely/pkg/widget"
)

func (ss *SpriteSheet) mailbox() {
	msg := ss.broker.Subscribe()
	for {
		m := <-msg
		switch m.GetTopic() {
		case topic.LEFT_CLICK.String():
			ss.handleClick(m.GetPayload().(geom.Coordinate))
		case topic.SET_CURRENT_COLOR.String():
			ss.currentColor = m.GetPayload().(color.Color)
		case topic.SET_PIXEL.String():
			coord := m.GetPayload().(geom.Coordinate)
			ss.sprites[ss.selected.Y][ss.selected.X].Widget.Elements[coord.Y][coord.X].SetGraphic(ss.currentColor)
		case topic.COPY.String():
			ss.broker.Publish(message.Message{
				Topic:   topic.PUSH_TO_CLIPBOARD,
				Payload: ss.sprites[ss.selected.Y][ss.selected.X].Widget.Elements,
			})
		case topic.UPDATE_CANVAS.String():
			ss.sprites[ss.selected.Y][ss.selected.X].Widget.SetElements(m.GetPayload().([][]*widget.Element))
		case topic.LEFT.String():
			ss.selectFromLocal(geom.Coordinate{
				X: ss.selected.X - 1,
				Y: ss.selected.Y,
			})
		case topic.RIGHT.String():
			ss.selectFromLocal(geom.Coordinate{
				X: ss.selected.X + 1,
				Y: ss.selected.Y,
			})
		case topic.UP.String():
			ss.selectFromLocal(geom.Coordinate{
				X: ss.selected.X,
				Y: ss.selected.Y - 1,
			})
		case topic.DOWN.String():
			ss.selectFromLocal(geom.Coordinate{
				X: ss.selected.X,
				Y: ss.selected.Y + 1,
			})
		case topic.SAVE.String():
			// ss.Encode()
			ss.save()

		}
	}
}
