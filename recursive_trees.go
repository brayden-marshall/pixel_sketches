package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math"
)

const (
	windowSize float64 = 800

	// tilt is how much (in degrees) we change the angle of each branch
	tilt float64 = 30.0
	// shrink is how big each new branch is relative to the previous
	shrink float64 = 0.56
)

func recursiveTree(imd *imdraw.IMDraw, size int, start pixel.Vec, length, angle float64) {
	if size < 1 {
		return
	}

	angleRad := degreesToRadians(angle)
	dx := length * math.Sin(angleRad)
	dy := length * math.Cos(angleRad)
	end := start.Add(pixel.V(dx, dy))

	drawBranch(imd, start, end)

	recursiveTree(imd, size-1, end, length*shrink, angle-tilt)
	recursiveTree(imd, size-1, end, length*shrink, angle+tilt)
}

func degreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

// convenience function to simplify drawing lines with a start and end vector
func drawBranch(imd *imdraw.IMDraw, start, end pixel.Vec) {
	imd.Push(start)
	imd.Push(end)
	imd.Line(2)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Recursive Trees",
		Bounds: pixel.R(0, 0, windowSize, windowSize),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.Black

	recursiveTree(imd, 8, pixel.V(windowSize/2, 0), windowSize/3, 0.0)

	for !win.Closed() {
		win.Clear(colornames.Snow)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
