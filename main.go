package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/aykevl/ledsgo"
)

const (
	size = 32
)

/*
var (
	uart = machine.UART1
	bt   = machine.UART2
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
)*/

func main() {

	//uart.Configure(machine.UARTConfig{TX: machine.USBCDC_DM_PIN, RX: machine.USBCDC_DP_PIN})
	//bt.Configure(machine.UARTConfig{TX: tx, RX: rx, BaudRate: 9600})
	//uart.Write([]byte("Echo console enabled. Type something then press enter:\r\n"))

	//pacmanGame()

	//fullRefreshes := uint(0)
	previousSecond := int64(0)
	//demo := colorCoordinateAt
	//demo := noiseAt
	//demo := fireAt
	//demo := radiance
	//demo := hyperspace
	demoSeconds := 0
	demoF := 0
	var funcs [6]func(x, y, z int, t time.Time) color.RGBA
	funcs[0] = noiseAt
	funcs[1] = fireAt
	funcs[2] = radiance
	funcs[3] = hyperspace
	funcs[4] = colorCoordinateAt
	funcs[5] = noiseAt
	demo := funcs[demoF]
	resetGame()
	for {
		start := time.Now()
		if demoF != 5 {
			drawPixels(start, demo)
			display.Display()
		} else {
			drawPills()
			drawWalls()

			d := pacman.d - 1 + int16(rand.Int31n(3))
			if d < 0 {
				d += 4
			}
			d = d % 4
			movePacman(d)

			x, y := pacmanCoords(pacman.x, pacman.y, pacman.p)
			display.SetPixel(x, y, color.RGBA{255, 255, 0, 255})

			display.Display()

		}

		second := (start.UnixNano() / int64(time.Second))
		if second != previousSecond {
			previousSecond = second

			//newFullRefreshes := getFullRefreshes()
			//animationTime := time.Since(start)
			//animationFPS := int64(10 * time.Second / animationTime)
			//print("#", second, " screen=", newFullRefreshes-fullRefreshes, "fps animation=", animationTime.String(), "/", (animationFPS / 10), ".", animationFPS%10, "fps\r\n")
			//fullRefreshes = newFullRefreshes
			demoSeconds++
			if demoSeconds >= 30 {
				demoSeconds = 0
				demoF++
				if demoF >= 6 {
					demoF = 0
				} else if demoF == 5 {
					resetGame()
				}
				demo = funcs[demoF]
			}
		}
	}
}

// drawPixels updates every pixel on the cube by calling getColor for each pixel
// and drawing it to the screen. It maps virtual (3D) pixels to physical (2D)
// pixels in the process.
func drawPixels(t time.Time, getColor func(x, y, z int, t time.Time) color.RGBA) {
	// Somewhat arbitrarily picking the top left of the topmost panel as the (0,
	// 0, 31) of the 3D cube.
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			display.SetPixel(int16(x+size*5), int16(y), getColor(x+1, y+1, 0, t))
			display.SetPixel(int16(x+size*4), int16(y), getColor(0, x+1, y+1, t))
			display.SetPixel(int16(x+size*3), int16(y), getColor(size-x, 0, y+1, t))
			display.SetPixel(int16(x+size*2), int16(y), getColor(size+1, size-x, y+1, t))
			display.SetPixel(int16(x+size*1), int16(y), getColor(x+1, size+1, y+1, t))
			display.SetPixel(int16(x+size*0), int16(y), getColor(x+1, size-y, size+1, t))
		}
	}
}

// noiseAt returns noise at the specified location.
func noiseAt(x, y, z int, t time.Time) color.RGBA {
	const (
		spread = 4096 / size // higher means the noise gets more detailed
		speed  = 20          // higher means slower
	)
	hue := uint16(ledsgo.Noise4(uint32(t.UnixNano()>>speed), uint32(x*spread), uint32(y*spread), uint32(z*spread))) * 2
	return ledsgo.Color{hue, 0xff, 0xff}.Spectrum()
}

// fireAt returns fire at the specified location.
func fireAt(x, y, z int, t time.Time) color.RGBA {
	const pointsPerCircle = 12  // how many LEDs there are per turn of the torch
	const cooling = 1792 / size // higher means faster cooling
	const detail = 12800 / size // higher means more detailed flames
	const speed = 12            // higher means faster
	const screenHeight = size + 1
	if z == 0 {
		return color.RGBA{}
	}
	heat := ledsgo.Noise3(uint32((31-z)*detail)-uint32((t.UnixNano()>>20)*speed), uint32(x*detail), uint32(y*detail))/32 + (128 * 8)
	heat -= uint16(screenHeight-z) * cooling
	if heat < 0 {
		heat = 0
	}
	return heatMap(heat)
}

// heatMap maps a color in the range 0..2047 to a color in a heat index. Useful
// for making flames.
func heatMap(index uint16) color.RGBA {
	if index < 128*8 {
		// red only
		return color.RGBA{uint8(index / 4), 0, 0, 255}
	}
	if index < 224*8 {
		// red-yellow
		return color.RGBA{255, uint8(uint32(index-128*8) / 3), 0, 255}
	}
	// yellow-white
	return color.RGBA{255, 255, (uint8(index - 224*8)), 255}
}

// colorCoordinateAt returns a color based on the 3 coordinates given. Useful
// for getting the virtual->physical pixel mapping right.
func colorCoordinateAt(x, y, z int, t time.Time) color.RGBA {
	// X represents red (more red to the right)
	// Y represents green (more green to the bottom)
	// Z represents blue (more blue to the bottom)
	return color.RGBA{uint8(x * 255 / (size + 1)), uint8(y * 255 / (size + 1)), uint8(z * 255 / (size + 1)), 0xff}
}

// radiance shows colors radiating out of the center.
func radiance(x, y, z int, now time.Time) color.RGBA {
	const circleX = 33 / 2 * 256
	const circleY = 33 / 2 * 256
	px := (x * (8192 / size)) - 4224         // .8
	py := (y * (8192 / size)) - 4224         // .8
	distance := ledsgo.Sqrt((px*px + py*py)) // .8
	hue := uint16(ledsgo.Noise1(uint32(distance>>0)-uint32(now.UnixNano()>>18))) + 0x8000
	return ledsgo.Color{hue, 0xff, 0xff}.Spectrum()
}

// hyperspace is a demo that should look a little bit like a hyperspace scene in
// a sci-fi movie.
func hyperspace(x, y, z int, now time.Time) color.RGBA {
	const circleX = (size/2 + 0.5) * 256
	const circleY = (size/2 + 0.5) * 256
	const cylinderRadius = 50 * 256 // higher number means more complexity
	// Calculate distance from the circle center.
	px := (x << 8) - circleX                 // .8
	py := (y << 8) - circleY                 // .8
	distance := ledsgo.Sqrt((px*px + py*py)) // .8

	// Normalize this distance.
	px = px * cylinderRadius / distance
	py = py * cylinderRadius / distance

	// Now the tricky part. Imagine this is a cyliner with px and py on
	// the outer circle of the cylinder. The cylinder moves through 3D
	// space in the direction of one of the flat sides of the cylinder
	// (the third coordinate).
	alpha := int(ledsgo.Noise3(uint32(px), uint32(py), uint32(distance/4)-uint32(now.UnixNano()>>16)))
	alpha -= 10000
	if alpha < 0 {
		alpha = 0
	}
	c := color.RGBA{0xaa, 0xaa, 0xff, 0xff}
	c = ledsgo.ApplyAlpha(c, uint8(alpha/256))
	return c
}
