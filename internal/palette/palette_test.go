package palette

import "testing"

func Test1DTo2D(t *testing.T) {
	dd := DefaultColors.To2D()
	i := 0

	for _, row := range dd {
		for _, c := range row {
			if c != DefaultColors[i] {
				t.Errorf("palette should successfully convert to a 2D slice of colors")
			}
			i++
		}
	}
}
