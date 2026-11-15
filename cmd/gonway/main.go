package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	width  = 160
	height = 120
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

type Game struct {
	canvas *image.RGBA
	cells [width][height]bool
	ticks int
	speed int
	isPaused bool
}

func NewGame() *Game {
	g := &Game{canvas: image.NewRGBA(image.Rect(0, 0, width, height))}
	g.speed = 6
	return g
}

func (g *Game) GetNeighbors(x, y int) []Point {
	neighbors := []Point{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if x+i >= 0 && x+i < width && y+j >= 0 && y+j < height {
				neighbors = append(neighbors, NewPoint(x+i, y+j))
			}
		}
	}

	return neighbors
}

func (g *Game) GetAliveNeighbors(x, y int) int {
	neighbors := g.GetNeighbors(x, y)
	alive := 0

	for _, n := range neighbors {
		if g.cells[n.X][n.Y] {
			alive++
		}
	}

	return alive
}

func (g *Game) UpdateGeneration() {
	var buf [width][height]bool
	for x, row := range g.cells {
		for y, cell := range row {
			alive := g.GetAliveNeighbors(x, y)

			if cell && (alive == 2 || alive == 3) {
				buf[x][y] = true
			} else if !cell && alive == 3 {
				buf[x][y] = true
			} else {
				buf[x][y] = false
			}
		}
	}

	g.cells = buf
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.cells[x][y] = true
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		x, y := ebiten.CursorPosition()
		g.cells[x][y] = false
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.speed = 1
	}
	
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.speed = 2
	}
	
	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.speed = 3
	}
	
	if inpututil.IsKeyJustPressed(ebiten.Key4) {
		g.speed = 4
	}
	
	if inpututil.IsKeyJustPressed(ebiten.Key5) {
		g.speed = 5
	}

	if inpututil.IsKeyJustPressed(ebiten.Key6) {
		g.speed = 6
	}
	
	if inpututil.IsKeyJustPressed(ebiten.Key7) {
		g.speed = 7
	}
	
	if inpututil.IsKeyJustPressed(ebiten.Key8) {
		g.speed = 8
	}
	
	if inpututil.IsKeyJustPressed(ebiten.Key9) {
		g.speed = 9
	}
	
	if inpututil.IsKeyJustPressed(ebiten.Key0) {
		g.speed = 10
	}
	
	// Reset
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		for i, x := range g.cells {
			for j, _ := range x {
				g.cells[i][j] = false
			}
		}
	}
	
	// Pause/Unpause
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.isPaused = !g.isPaused
		if g.isPaused {
			ebiten.SetWindowTitle("Gonway [Paused]")
		} else {
			ebiten.SetWindowTitle("Gonway")
		}
	}
	
	if !g.isPaused {
		g.ticks++
		if g.ticks%g.speed == 0 {
			g.UpdateGeneration()
		}
	}
	
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
