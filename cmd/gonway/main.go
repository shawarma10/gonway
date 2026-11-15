package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width  = 160
	height = 120
)

type Game struct {
	canvas *image.RGBA
	cells [160][120]bool
}

func NewGame() *Game {
	g := &Game{canvas: image.NewRGBA(image.Rect(0, 0, width, height))}

	for x := range width {
		for y := range height/2 {
			g.cells[x][y] = true
		}
	}
	
	return g
}

func (g *Game) Update() error {
	g.canvas = image.NewRGBA(image.Rect(0, 0, width, height))
	for i, x := range g.cells {
		for j, y := range x {
			if y {
				g.canvas.Set(i, j, color.White)
			} else {
				g.canvas.Set(i, j, color.Black)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.canvas.Pix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(width*4, height*4)
	ebiten.SetWindowTitle("Gonway")

	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
