package main

import (
	"image/color"
	"time"

	"tinygo.org/x/tinydraw"
)

const (
	NO_ROTATION = iota
	ROTATE_90
	ROTATE_180
	ROTATE_270

	PANEL_SIZE = 31
)

type coord struct {
	x int16
	y int16
	p uint8
	d int16
}

var pacman coord
var direction int16
var directionMod int16
var walls [1024]bool
var pills [6][1024]bool
var complete [6]bool
var pillColor = color.RGBA{6, 6, 6, 255}
var superPillColor = color.RGBA{255, 255, 255, 255}
var wallColor = color.RGBA{0, 0, 155, 255}
var pillsEaten uint32

func pacmanGame() {

	walls = [1024]bool{
		true, true, true, true, true, true, true, true, true, true, false, false, false, true, false, true, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, true, false, false, false, false, true, true, false, true, true, true, false, true, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, true, true, false, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, true, true, false, false, false, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, true, true, true, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, true, true, true, false, true, true, true, true, true, true, false, true, true, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, true, true, false, true, true, true, false, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, false, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, true, true, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, false, false, false, false, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, false, false, false, false, true, true, false, true, true, false, true, false, false, false, false, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, false, false, false, true, false, false, false, false, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, false, false, true, true, true, false, true, true, true, true, true, false, true, false, false, false, false, true, false, true, true, false, false, false, false, true, true, false, false, false, false, false, true, false, true, true, true, false, true, true, true, true, true, false, true, false, false, false, false, true, false, true, true, true, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, false, true, true, true, true, true, false, true, false, false, false, false, true, false, true, true, true, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, false, true, true, false, false, false, false, true, false, false, false, false, true, false, true, true, false, false, false, false, true, true, false, false, false, false, true, true, false, true, true, true, false, true, true, false, true, true, false, true, false, false, false, false, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, false, false, false, false, true, true, false, true, true, false, true, false, false, false, false, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, false, false, false, false, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, true, true, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, true, true, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, true, true, false, true, true, false, true, true, true, false, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, false, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, true, true, true, true, false, true, true, true, true, true, false, true, true, true, true, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, true, true, true, true, true, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, true, true, false, false, false, false, true, true, false, true, true, false, true, true, true, false, true, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, true, true, false, true, true, false, true, true, false, true, true, false, false, false, false, false, false, false, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, true, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, false, false, false, true, false, true, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true,
	}

	createPills()

	directionMod = NO_ROTATION

	pacman = coord{6, 15, 0, 1}

	var x, y int16
	t := time.Now()
	for {
		//display.FillScreen(color.RGBA{0, 0, 0, 255})

		drawWalls()
		drawPills()

		if time.Since(t).Milliseconds() > 60 {
			t = time.Now()
			getInput()
		}

		x, y = pacmanCoords(pacman.x, pacman.y, pacman.p)
		tinydraw.FilledRectangle(display, x, y, 1, 1, color.RGBA{255, 255, 0, 255})

		display.Display()

	}
}

func pacmanCoords(x, y int16, panel uint8) (int16, int16) {
	if panel == 0 {
		return (PANEL_SIZE + 1) + x, y
	} else if panel == 1 {
		return 64 + x, y
	} else if panel == 2 {
		return 96 + x, y
	} else if panel == 3 {
		return 128 + x, y
	} else if panel == 4 { // TOP
		return 160 + x, y
	} else if panel == 5 { // BOTTOM
		return x, y
	}
	return 0, 0
}

func getInput() {
	/*if bt.Buffered() > 0 || uart.Buffered() > 0 {
		var d int16
		var data byte
		for bt.Buffered() > 0 {
			data, _ = bt.ReadByte()
		}
		for uart.Buffered() > 0 {
			data, _ = uart.ReadByte()
		}
		switch data {
		case 119:
			d = 0 + directionMod
			//y--
			break
		case 97:
			d = 3 + directionMod
			//x--
			break
		case 115:
			d = 2 + directionMod
			//y++
			break
		case 100:
			d = 1 + directionMod
			//x++
			break
		default:
		}

		if d >= 4 {
			d -= 4
		}
		if movePacman(d) {
			return
		}
	}*/
	movePacman(pacman.d)
}

func movePacman(d int16) bool {
	x := pacman.x
	y := pacman.y
	p := pacman.p
	switch d {
	case 0:
		y--
		break
	case 1:
		x++
		break
	case 2:
		y++
		break
	case 3:
		x--
		break
	}

	if x < 0 {
		if p == 0 {
			x = PANEL_SIZE
			p = 3
		} else if p == 1 || p == 2 || p == 3 {
			x = PANEL_SIZE
			p--
		} else if p == 4 {
			x = y
			y = 0
			p = 3
		} else if p == 5 {
			x = PANEL_SIZE - y
			y = PANEL_SIZE
			p = 3
		}
	}

	if x == (PANEL_SIZE + 1) {
		if p == 0 || p == 1 || p == 2 {
			x = 0
			p++
		} else if p == 3 {
			x = 0
			p = 0
		} else if p == 4 {
			x = PANEL_SIZE - y
			y = 0
			p = 1
		} else if p == 5 {
			x = y
			y = PANEL_SIZE
			p = 1
		}
	}

	if y < 0 {
		if p == 0 {
			y = PANEL_SIZE
			p = 4
		} else if p == 1 {
			y = PANEL_SIZE - x
			x = PANEL_SIZE
			p = 4
		} else if p == 2 {
			y = 0
			x = PANEL_SIZE - x
			p = 4
		} else if p == 3 {
			y = x
			x = 0
			p = 4
		} else if p == 4 {
			x = PANEL_SIZE - x
			y = 0
			p = 2
		} else if p == 5 {
			y = PANEL_SIZE
			p = 0
		}
	}

	if y == (PANEL_SIZE + 1) {
		if p == 0 {
			y = 0
			p = 5
		} else if p == 1 {
			y = x
			x = PANEL_SIZE
			p = 5
		} else if p == 2 {
			x = PANEL_SIZE - x
			y = PANEL_SIZE
			p = 5
		} else if p == 3 {
			y = PANEL_SIZE - x
			x = 0
			p = 5
		} else if p == 4 {
			y = 0
			p = 0
		} else if p == 5 {
			x = PANEL_SIZE - x
			y = PANEL_SIZE
			p = 2
		}
	}

	if checkCollision(x, y, p) {
		return false
	}
	if o := changeDirection(pacman.p, p); o > 0 {
		println("CHANGE DIRECTION", pacman.d, d, o, directionMod)
		d -= o
		if d < 0 {
			d += 4
		}
	}

	pacman.x = x
	pacman.y = y
	pacman.p = p
	pacman.d = d

	if pills[p][x*32+y] {
		pillsEaten++
		pills[p][x*32+y] = false
	}
	println("PACMAN", x, y, p, d, directionMod)
	return true
}

func changeDirection(oldP, p uint8) int16 {
	if oldP == p {
		return 0
	}
	o := directionMod
	switch oldP {
	case 0:
		break
	case 1:
		if p == 4 {
			directionMod += 3
		} else if p == 5 {
			directionMod += 1
		}
		break
	case 2:
		if p == 4 {
			directionMod += 2
		} else if p == 5 {
			directionMod += 2
		}
		break
	case 3:
		if p == 4 {
			directionMod += 1
		} else if p == 5 {
			directionMod += 3
		}
		break
	case 4:
		if p == 1 {
			directionMod += 1
		} else if p == 2 {
			directionMod += 2
		} else if p == 3 {
			directionMod += 3
		}
		break
	case 5:
		if p == 1 {
			directionMod += 3
		} else if p == 2 {
			directionMod += 2
		} else if p == 3 {
			directionMod += 1
		}
		break
	}

	if directionMod >= 4 {
		directionMod -= 4
	}
	o -= directionMod
	if o < 0 {
		o += 4
	}
	if o > 3 {
		o -= 4
	}
	return o
}

func drawWalls() {
	for i := int16(0); i < 32; i++ {
		for j := int16(0); j < 32; j++ {
			if walls[i*32+j] {
				// BOTTOM
				display.SetPixel(i, j, wallColor)
				// FRONT
				display.SetPixel(32+i, j, wallColor)
				// RIGHT
				display.SetPixel(64+i, j, wallColor)
				// BACK
				display.SetPixel(96+(31-i), j, wallColor)
				// LEFT
				display.SetPixel(128+(31-i), j, wallColor)
				// TOP
				display.SetPixel(160+i, 31-j, wallColor)
			}
		}
	}
}

func drawPills() {
	for i := int16(0); i < 32; i++ {
		for j := int16(0); j < 32; j++ {
			for p := int16(0); p < 6; p++ {
				if pills[p][i*32+j] {
					if (i == 1 && j == 1) ||
						(i == 1 && j == 30) ||
						(i == 30 && j == 1) ||
						(i == 30 && j == 30) {
						if p == 5 {
							display.SetPixel(i, j, superPillColor)
						} else {
							display.SetPixel(32+p*32+i, j, superPillColor)
						}
					} else {
						if p == 5 {
							display.SetPixel(i, j, pillColor)
						} else {
							display.SetPixel(32+p*32+i, j, pillColor)
						}
					}
				}
			}
		}
	}
}

func createPills() {
	var c int16
	for i := int16(0); i < 32; i++ {
		for j := int16(0); j < 32; j++ {
			if ((i < 6 || i >= 27) &&
				((j >= 10 && j <= 12) || (j >= 16 && j <= 18))) ||
				(i >= 11 && i <= 20 && j >= 13 && j <= 16) {
				continue
			}
			c = i*32 + j
			if !walls[c] {
				// BOTTOM
				pills[5][c] = true
				// FRONT
				pills[0][c] = true
				// RIGHT
				pills[1][c] = true
				// BACK
				pills[2][(31-i)*32+j] = true
				// LEFT
				pills[3][(31-i)*32+j] = true
				// TOP
				pills[4][i*32+(31-j)] = true
			}
		}
	}
}

func checkCollision(x, y int16, panel uint8) bool {
	switch panel {
	case 0:
		return walls[x*32+y]
		break
	case 1:
		return walls[x*32+y]
		break
	case 2:
		return walls[(31-x)*32+y]
		break
	case 3:
		return walls[(31-x)*32+y]
		break
	case 4:
		return walls[x*32+(31-y)]
		break
	case 5:
		return walls[x*32+y]
		break
	}
	return true
}
