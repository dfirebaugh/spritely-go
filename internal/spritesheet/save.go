package spritesheet

import (
	"github.com/dfirebaugh/spritely-go/internal/file"

	"github.com/hajimehoshi/ebiten/v2"
)

func (ss *SpriteSheet) save() {
	file.Save(ebiten.NewImage(
		ss.spriteRowSize*ss.pixelSize*ss.rowSize,
		ss.spriteRowSize*ss.pixelSize*ss.columnSize,
	), ss.sprites, ss.Widget.Offset)
}
