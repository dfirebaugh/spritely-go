package spritesheet

import (
	"fmt"
	"image/color"
	"strings"
	"time"

	"github.com/dfirebaugh/spritely-go/internal/message"
	"github.com/dfirebaugh/spritely-go/internal/topic"
	"github.com/dfirebaugh/spritely-go/pkg/geom"
	"github.com/dfirebaugh/spritely-go/pkg/widget"
)

var isLeftJustPressed = false

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
		case topic.LEFT_CLICK_JUST_PRESSED.String():
			isLeftJustPressed = true
			go func() {
				time.Sleep(200 * time.Millisecond)
				isLeftJustPressed = false
			}()

		case topic.SAVE.String():
			if !isLeftJustPressed {
				continue
			}

			println("Selected Sprite Encoding:")
			selectedEncoded := ss.sprites[ss.selected.Y][ss.selected.X].Encode()
			var selectedOutput string
			for _, line := range splitLines(selectedEncoded) {
				selectedOutput += fmt.Sprintf("\"%s\"\n", line)
			}
			fmt.Println(selectedOutput)

			var fullSpriteSheet string
			spriteHeight := len(ss.sprites[0][0].Widget.Elements)
			for y := 0; y < 4; y++ {
				for i := 0; i < spriteHeight; i++ {
					var line string
					for x := 0; x < 8; x++ {
						encodedSprite := splitLines(ss.sprites[y][x].Encode())
						line += encodedSprite[i]
					}
					fullSpriteSheet += fmt.Sprintf("\"%s\"\n", line)
				}
				fullSpriteSheet += "\n"
			}

			fmt.Println("Entire Sprite Sheet Encoding in 8x4 grid (with quotes):")
			fmt.Println(fullSpriteSheet)

			isLeftJustPressed = false

		}
	}
}

func splitLines(encoded string) []string {
	return strings.Split(strings.TrimSpace(encoded), "\n")
}

func formatEncodedSprite(encoded string) string {
	var output string
	for _, line := range splitLines(encoded) {
		output += fmt.Sprintf("\"%s\"\n", line)
	}
	return output
}
