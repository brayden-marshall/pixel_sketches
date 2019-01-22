package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

const (
	DIRECTION_RIGHT = iota
	DIRECTION_UP
	DIRECTION_LEFT
	DIRECTION_DOWN
)

const windowSize float64 = 600

type Ant struct {
	Row       int
	Col       int
	Direction int
}

func drawWorld(world [][]bool, ant Ant, imd *imdraw.IMDraw) {
	var interval = windowSize / float64(len(world))
	var min, max pixel.Vec
	for row := range world {
		min = pixel.V(float64(row)*interval, 0)
		max = min.Add(pixel.V(interval, interval))
		for col := range world[row] {
			if ant.Row == row && ant.Col == col {
				imd.Color = colornames.Red
			} else if world[row][col] {
				imd.Color = colornames.Black
			} else {
				imd.Color = colornames.White
			}
			imd.Push(min)
			imd.Push(max)
			imd.Rectangle(0)
			min = min.Add(pixel.V(0, interval))
			max = max.Add(pixel.V(0, interval))
		}
	}
}

// using a separate modulo function to properly deal with negatives
func modulo(x, y int) int {
	return (x%y + y) % y
}

func moveAnt(world *[][]bool, ant *Ant) {
	antSquare := (*world)[ant.Row][ant.Col]
	(*world)[ant.Row][ant.Col] = !antSquare
	if antSquare {
		ant.Direction = modulo(ant.Direction-1, 4)
	} else {
		ant.Direction = modulo(ant.Direction+1, 4)
	}

	switch ant.Direction {
	// direction changes are calculated with modulo for edge wrapping
	case DIRECTION_RIGHT:
		ant.Col = modulo(ant.Col+1, len(*world))
	case DIRECTION_UP:
		ant.Row = modulo(ant.Row-1, len(*world))
	case DIRECTION_LEFT:
		ant.Col = modulo(ant.Col-1, len(*world))
	case DIRECTION_DOWN:
		ant.Row = modulo(ant.Row+1, len(*world))
	default:
		fmt.Printf("Error in moveAnt(): no matching direction.")
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Langton's Ant",
		Bounds: pixel.R(0, 0, windowSize, windowSize),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	const worldSize int = 20
	world := make([][]bool, worldSize)
	for i := range world {
		world[i] = make([]bool, worldSize)
	}

	var ant Ant = Ant{
		Row:       worldSize / 2,
		Col:       worldSize / 2,
		Direction: DIRECTION_LEFT,
	}

	ticker := time.NewTicker(time.Millisecond)
	for !win.Closed() {
		win.Clear(colornames.Snow)
		imd.Clear()
		drawWorld(world, ant, imd)
		imd.Draw(win)
		win.Update()
		select {
		case <-ticker.C:
			moveAnt(&world, &ant)
		}
	}
	ticker.Stop()
}

func main() {
	pixelgl.Run(run)
}
