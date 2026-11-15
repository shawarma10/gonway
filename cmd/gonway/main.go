package main

import (
	"image"
	"log"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width  = 160
	height = 120
)

type Game struct {
	canvas *image.RGBA
	red bool
	green bool
	blue bool
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyG) {
		g.green = !g.green
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.red = !g.red
	}

	if ebiten.IsKeyPressed(ebiten.KeyB) {
		g.blue = !g.blue
	}
	
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= 0 && x < width && y >= 0 && y < height {
			var redVal, greenVal, blueVal uint8 = 0, 0, 0
			if g.red {
				redVal = 255
			}
			
			if g.green {
				greenVal = 255
			}
			
			if g.blue {
				blueVal = 255
			}
			
			g.canvas.Set(x, y, color.RGBA{redVal, greenVal, blueVal, 255})
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
	ebiten.SetWindowTitle("Hello, Ebitengine")

	g := &Game{canvas: image.NewRGBA(image.Rect(0, 0, width, height))}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
