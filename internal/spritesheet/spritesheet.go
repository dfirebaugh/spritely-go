package spritesheet

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strings"

	"github.com/dfirebaugh/spritely-go/internal/message"
	"github.com/dfirebaugh/spritely-go/internal/palette"
	"github.com/dfirebaugh/spritely-go/internal/sprite"
	"github.com/dfirebaugh/spritely-go/internal/topic"
	"github.com/dfirebaugh/spritely-go/pkg/broker"
	"github.com/dfirebaugh/spritely-go/pkg/geom"
	"github.com/dfirebaugh/spritely-go/pkg/widget"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	broker        *broker.Broker
	Widget        *widget.Widget
	sprites       [][]*sprite.Sprite
	currentColor  color.Color
	pixelSize     int
	spriteRowSize int
	selected      geom.Coordinate
	rowSize       int
	columnSize    int
}

func New(broker *broker.Broker, spriteSheetOffset geom.Offset, spriteSize geom.Size) *SpriteSheet {
	ss := SpriteSheet{
		broker:        broker,
		pixelSize:     2,
		spriteRowSize: 8,
		columnSize:    4,
		rowSize:       8,
	}

	var selectionRects [][]color.Color
	for y := 0; y < ss.columnSize; y++ {
		var spriteRow []*sprite.Sprite
		var selectionRow []color.Color
		for x := 0; x < ss.rowSize; x++ {
			spriteRow = append(spriteRow, sprite.New(broker, geom.Offset{
				X: float64(x*ss.pixelSize*ss.spriteRowSize) + spriteSheetOffset.X,
				Y: float64(y*ss.pixelSize*ss.spriteRowSize) + spriteSheetOffset.Y,
			}, geom.Size{
				Width:  ss.pixelSize,
				Height: ss.pixelSize,
			}))
			selectionRow = append(selectionRow, color.Transparent)
		}
		ss.sprites = append(ss.sprites, spriteRow)
		selectionRects = append(selectionRects, selectionRow)
	}
	ss.Widget = widget.NewSelectableColors(selectionRects, geom.Size{
		Width:  8 * ss.pixelSize,
		Height: 8 * ss.pixelSize,
	}, spriteSheetOffset)

	// Try to load from an existing sprite_sheet file
	if _, err := os.Stat("./sprite_sheet"); err == nil {
		fmt.Println("Loading sprite sheet from file...")
		err = ss.loadSpriteSheetFromFile("./sprite_sheet")
		if err != nil {
			fmt.Println("Error loading sprite sheet:", err)
		} else {
			fmt.Println("Sprite sheet loaded successfully.")
		}
	}

	go ss.mailbox()
	return &ss
}

func (ss *SpriteSheet) loadSpriteSheetFromFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}
		for x, char := range line {
			colorIdx := charToIndex(char)
			ss.sprites[y/ss.spriteRowSize][x/ss.spriteRowSize].Widget.Elements[y%ss.spriteRowSize][x%ss.spriteRowSize].SetGraphic(palette.DefaultColors[colorIdx])
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (ss SpriteSheet) Render(dst *ebiten.Image) {
	for _, row := range ss.sprites {
		for _, s := range row {
			s.Render(dst)
		}
	}
	ss.Widget.Render(dst)
}

func (ss *SpriteSheet) handleClick(coordinate geom.Coordinate) {
	if !ss.Widget.IsWithinBounds(coordinate) {
		return
	}
	ss.selectSprite(coordinate)
}

func (ss SpriteSheet) GetSprite(local geom.Coordinate) *sprite.Sprite {
	return ss.sprites[ss.selected.Y][ss.selected.X]
}

func (ss *SpriteSheet) selectSprite(coordinate geom.Coordinate) {
	ss.selectFromLocal(ss.Widget.ToLocalCoordinate(coordinate))
}

func (ss *SpriteSheet) selectFromLocal(local geom.Coordinate) {
	if !ss.Widget.IsWithinLocalBounds(local) {
		return
	}
	ss.selected = local
	ss.Widget.SelectElement(local)
	ss.broker.Publish(message.Message{
		Topic:   topic.UPDATE_CANVAS,
		Payload: ss.GetSprite(ss.selected).Widget.Elements,
	})
}
