package file

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/dfirebaugh/spritely-go/internal/palette"
	"github.com/dfirebaugh/spritely-go/internal/sprite"
	"github.com/dfirebaugh/spritely-go/pkg/broker"
	"github.com/dfirebaugh/spritely-go/pkg/geom"
	"github.com/dfirebaugh/spritely-go/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	log "github.com/sirupsen/logrus"
)

func Save(dst *ebiten.Image, sprites [][]*sprite.Sprite, offset geom.Offset) {
	copyToImage(dst, sprites, offset)

	f, err := os.Create("spritely.spritesheet.png")
	if err != nil {
		log.Errorf("error saving sprite sheet: %s", err)
		return
	}
	defer f.Close()

	os.WriteFile(
		"spritely.spt",
		[]byte(fmt.Sprintf("#spritely-sprite\n==\n%s\n", Encode(sprites))),
		0o644,
	)
	os.WriteFile(
		"spritely.plt",
		[]byte(fmt.Sprintf("#spritely-palette\n==\n%s\n", encodePalette(palette.DefaultColors))),
		0o644,
	)

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, dst)
	if err != nil {
		log.Errorf("error encoding sprite sheet: %s", err)
	}
}

func copyToImage(dst *ebiten.Image, sprites [][]*sprite.Sprite, offset geom.Offset) {
	for _, row := range sprites {
		for _, sprite := range row {
			for _, sRow := range sprite.Widget.Elements {
				for _, element := range sRow {
					ebitenutil.DrawRect(
						dst, float64(element.Offset.X-offset.X),
						float64(element.Offset.Y-offset.Y),
						float64(element.Size.Width),
						float64(element.Size.Height), element.Graphic.(color.Color))
				}
			}
		}
	}
}

// Encode the sprite sheet into a hex representation
//
//	given that we have a defined palette of (16) colors
//	we can represent a sprite by a set of hex digits
func Encode(sprites [][]*sprite.Sprite) string {
	var result string
	r1 := make(map[int]string, 32)
	for _, spriteRow := range sprites {
		for _, sprite := range spriteRow {
			// result = result + sprite.Encode() + "\n"
			hex := strings.Split(sprite.Encode(), "\n")
			r1[0] = fmt.Sprintf("%s %s", r1[0], hex[0])
			r1[1] = fmt.Sprintf("%s %s", r1[1], hex[1])
			r1[2] = fmt.Sprintf("%s %s", r1[2], hex[2])
			r1[3] = fmt.Sprintf("%s %s", r1[3], hex[3])
			r1[4] = fmt.Sprintf("%s %s", r1[4], hex[4])
			r1[5] = fmt.Sprintf("%s %s", r1[5], hex[5])
			r1[6] = fmt.Sprintf("%s %s", r1[6], hex[6])
			r1[7] = fmt.Sprintf("%s %s", r1[7], hex[7])
		}
	}
	println()
	result = strings.Join([]string{r1[0], r1[1], r1[2], r1[3], r1[4], r1[5], r1[6], r1[7]}, "\n")

	return result
}

func Decode(hex string) [][]color.Color {
	var result [][]color.Color
	for _, row := range strings.Split(hex, "\n") {
		var resultRow []color.Color
		for _, c := range row {
			i, _ := strconv.Atoi(string(c))
			resultRow = append(resultRow, palette.DefaultColors[i])
		}
		result = append(result, resultRow)
	}
	return result
}

func parseSprite(encoded string, offset geom.Offset, size geom.Size) *sprite.Sprite {
	s := sprite.New(broker.NewBroker(), offset, size)
	var elements [][]*widget.Element
	var elementRow []*widget.Element
	x := 0
	y := 0
	for _, digit := range strings.Split(encoded, "") {
		if y == 8 {
			y = 0
		}
		if x == 8 {
			x = 0
			y++
			elements = append(elements, elementRow)
			elementRow = []*widget.Element{}
		}
		num, _ := strconv.ParseInt(digit, 16, 64)
		elementRow = append(elementRow, &widget.Element{
			Graphic: palette.DefaultColors[num],
		})
		x++
	}

	s.Widget.SetElements(elements)
	return s
}

// encode palette into something that we could parse
func encodePalette(p []color.Color) string {
	var s string

	for _, c := range p {
		r, g, b, _ := c.RGBA()
		s = s + fmt.Sprintf("%x,%x,%x\n", r, g, b)
	}
	return s
}
