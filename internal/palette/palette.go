package palette

import "image/color"

const PaletteSize = 4

type Palette []color.Color

// because color.Black is actually a different type
var Black = color.RGBA{0, 0, 0, 255}

var (
	DefaultColors = Palette{
		color.RGBA{127, 36, 84, 255},
		color.RGBA{28, 43, 83, 255},
		color.RGBA{0, 135, 81, 255},
		color.RGBA{171, 82, 54, 255},
		color.RGBA{96, 88, 79, 255},
		color.RGBA{195, 195, 198, 255},
		color.RGBA{255, 241, 233, 255},
		color.RGBA{237, 27, 81, 255},
		color.RGBA{250, 162, 27, 255},
		color.RGBA{247, 236, 47, 255},
		color.RGBA{93, 187, 77, 255},
		color.RGBA{81, 166, 220, 255},
		color.RGBA{131, 118, 156, 255},
		color.RGBA{241, 118, 166, 255},
		color.RGBA{252, 204, 171, 255},
		Black,
	}
)

func (Palette) To2D() [][]color.Color {
	var palette [][]color.Color
	i := 0
	for y := 0; y < PaletteSize; y++ {
		var paletteRow []color.Color
		for x := 0; x < PaletteSize; x++ {
			paletteRow = append(paletteRow, DefaultColors[i])
			i++
		}
		palette = append(palette, paletteRow)
	}
	return palette
}
