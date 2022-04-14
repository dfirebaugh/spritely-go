package main

import (
	"image/color"
	"log"
	"spritely/internal/mediator"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Width    int
	Height   int
	mediator mediator.Mediator
}

const (
	ScreenWidth  = 1024
	ScreenHeight = 1024
)

func (g *Game) Update() error {
	g.mediator.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	g.mediator.Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 256, 256
}

func (g *Game) Run() {
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("spritely")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func main() {
	println("running spritely...")

	game := &Game{
		Width:    ScreenWidth,
		Height:   ScreenHeight,
		mediator: mediator.New(),
	}

	game.Run()
}
