package spritesheet

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"time"

	"github.com/dfirebaugh/spritely-go/internal/message"
	"github.com/dfirebaugh/spritely-go/internal/palette"
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

			var fullSpriteSheet string
			spriteHeight := len(ss.sprites[0][0].Widget.Elements)
			for y := 0; y < 4; y++ {
				for i := 0; i < spriteHeight; i++ {
					var line string
					for x := 0; x < 8; x++ {
						encodedSprite := splitLines(ss.sprites[y][x].Encode())
						line += encodedSprite[i]
					}
					fullSpriteSheet += fmt.Sprintf("%s\n", line)
				}
				fullSpriteSheet += "\n"
			}

			fmt.Println(selectedOutput)
			err := os.WriteFile("sprite_sheet", []byte(fullSpriteSheet), 0o644)
			if err != nil {
				fmt.Println("Error writing to sprite_sheet.txt:", err)
			}
			fmt.Println("Entire Sprite Sheet Encoding in 8x4 grid (with quotes):")
			fmt.Println(fullSpriteSheet)

			// // Save full sprite sheet to "sprite_sheet.h"
			// err = os.WriteFile("sprite_sheet.h", []byte(fullSpriteSheet), 0o644)
			// if err != nil {
			// 	fmt.Println("Error writing to sprite_sheet.h:", err)
			// }

			fmt.Println("Sprite Sheet as unsigned char array in C syntax:")
			cSyntaxOutput := formatAsCUnsignedCharArray(fullSpriteSheet)
			fmt.Println(cSyntaxOutput)

			err = os.WriteFile("sprite_sheet.h", []byte(`
#ifndef SPRITE_SHEET_H
#define SPRITE_SHEET_H

`+cSyntaxOutput+`

#endif

        `), 0o644)
			if err != nil {
				fmt.Println("Error writing to sprite_sheet.char.h:", err)
			}

			err = ss.saveSpriteSheetAsPNG("sprite_sheet.png", fullSpriteSheet)
			if err != nil {
				fmt.Println("Error saving sprite sheet as PNG:", err)
			}
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

func formatAsCUnsignedCharArray(encoded string) string {
	var output strings.Builder
	output.WriteString("unsigned char spriteSheet[] = {\n")
	for _, line := range splitLines(encoded) {
		for i := 0; i < len(line); i += 2 {
			highNibble := line[i]
			lowNibble := line[i+1]

			high := uint8(0)
			low := uint8(0)

			if highNibble >= 'A' && highNibble <= 'F' {
				high = (highNibble - 'A' + 10)
			} else if highNibble >= 'a' && highNibble <= 'f' {
				high = (highNibble - 'a' + 10)
			} else {
				high = highNibble - '0'
			}

			if lowNibble >= 'A' && lowNibble <= 'F' {
				low = (lowNibble - 'A' + 10)
			} else if lowNibble >= 'a' && lowNibble <= 'f' {
				low = (lowNibble - 'a' + 10)
			} else {
				low = lowNibble - '0'
			}

			packedByte := (high << 4) | low
			output.WriteString(fmt.Sprintf("0x%02X, ", packedByte))
		}
		output.WriteString("\n")
	}
	output.WriteString("};\n")
	return output.String()
}

func (ss *SpriteSheet) saveSpriteSheetAsPNG(filename, encoded string) error {
	width := 8 * len(ss.sprites[0][0].Widget.Elements[0])
	height := 4 * len(ss.sprites[0][0].Widget.Elements)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	lines := splitLines(encoded)
	y := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue // Skip empty lines
		}
		for x, char := range line {
			colorIdx := charToIndex(char)
			img.Set(x, y, palette.DefaultColors[colorIdx])
		}
		y++
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func charToIndex(c rune) int {
	if c >= 'A' && c <= 'F' {
		return int(c - 'A' + 10)
	} else if c >= 'a' && c <= 'f' {
		return int(c - 'a' + 10)
	}
	return int(c - '0')
}
