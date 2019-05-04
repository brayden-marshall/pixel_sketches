package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
	"time"
)

const (
	windowHeight = 900
	windowWidth  = 600
)

func dot(imd *imdraw.IMDraw, x, y float64) {
	imd.Push(pixel.V(x, y))
	imd.Circle(0.4, 0.4)
}

func fern(win *pixelgl.Window, imd *imdraw.IMDraw) {
	// random generator setup
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// setting origin
	x1 := 0.0
	y1 := 0.0
	// x2 and y2 used to simplify calculation of next dot location
	var x2 float64
	var y2 float64

	const scaleFactor = 80.0
	const iterations = 300000

	win.Clear(colornames.Black)
	dot(imd, x1*scaleFactor, y1*scaleFactor)
	for i := 0; i < iterations && !win.Closed(); i++ {
		// only draw to screen every 1000 iterations
		if i%1000 == 0 {
			imd.Draw(win)
			win.Update()
			imd.Clear()
		}

		r := random.Float64()
		if r < 0.1 {
			x2 = 0
			y2 = 0.16 * y1
		} else if r < 0.86 {
			x2 = (0.85 * x1) + (0.04 * y1)
			y2 = (-0.04 * x1) + (0.85 * y1) + 1.6
		} else if r < 0.93 {
			x2 = (0.2 * x1) - (0.26 * y1)
			y2 = (0.23 * x1) + (0.22 * y1) + 1.6
		} else {
			x2 = (-0.15 * x1) + (0.28 * y1)
			y2 = (0.26 * x1) + (0.24 * y1) + 0.44
		}
		x1 = x2
		y1 = y2
		dot(imd, x1*scaleFactor, y1*scaleFactor)
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Barnsley Fern",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// set up imdraw parameters
	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	m := pixel.IM.Moved(pixel.V(windowWidth/2, 10))
	imd.SetMatrix(m)

	fern(win, imd)
	fmt.Println("done")

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
