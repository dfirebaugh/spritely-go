package icons

import (
	"bytes"
	_ "embed"
	"image"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sirupsen/logrus"
)

//go:embed tabs.png
var tabsPng []byte

//go:embed icons.png
var iconsRaw []byte

//go:embed palette.png
var palletRaw []byte

var Tabs *ebiten.Image
var ToolBar *ebiten.Image
var Palette *ebiten.Image

func init() {
	Tabs, _ = loadPNG(tabsPng)
	ToolBar, _ = loadPNG(iconsRaw)
	Palette, _ = loadPNG(palletRaw)
}

func loadPNG(b []byte) (*ebiten.Image, error) {
	var err error
	imgDecoded, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	return ebiten.NewImageFromImage(imgDecoded), nil
}

func ToImage(name string) *ebiten.Image {
	switch name {
	case "palette":
		return Palette
	case "toolbar":
		return ToolBar
	}

	return &ebiten.Image{}
}
