package main

            /*
const (
	NO_ROTATION = iota
	ROTATE_90
	ROTATE_180
	ROTATE_270
)

type coord struct {
	x int16
	y int16
	p uint8
}

var snake [10]coord
var direction int16
var directionMod int16

func snakeGame() {

	directionMod = NO_ROTATION

	snake[0] = coord{0, 0, 0}
	snake[1] = coord{0, 1, 0}
	snake[2] = coord{0, 2, 0}
	snake[3] = coord{0, 3, 0}
	snake[4] = coord{0, 4, 0}
	snake[5] = coord{0, 5, 0}
	snake[6] = coord{0, 6, 0}
	snake[7] = coord{0, 7, 0}
	snake[8] = coord{1, 7, 0}
	snake[9] = coord{2, 7, 0}

	var x, y int16
	var i int16
	for {
		display.FillScreen(color.RGBA{0, 0, 0, 255})

		getInput()

		for i = 0; i < 10; i++ {
			x, y = snakeCoords(snake[i].x, snake[i].y, snake[i].p)
			tinydraw.FilledRectangle(display, x, y, 4, 4, color.RGBA{0, 255, 0, 255})
		}
		//x, y = snakeCoords(snake[0].x, snake[0].y, snake[0].p)
		//tinydraw.FilledRectangle(display, x, y, 4, 4, color.RGBA{0, 255, 0, 255})

		display.Display()

		//time.Sleep(100 * time.Millisecond)
	}
}

func snakeCoords(x, y int16, panel uint8) (int16, int16) {
	if panel == 0 {
		return 32 + x*4, y * 4
	} else if panel == 1 {
		return 64 + x*4, y * 4
	} else if panel == 2 {
		return 96 + x*4, y * 4
	} else if panel == 3 {
		return 128 + x*4, y * 4
	} else if panel == 4 { // TOP
		return 160 + x*4, y * 4
	} else if panel == 5 { // BOTTOM
		return x * 4, y * 4
	}
	return 0, 0
}

func getInput() {
	var d int16
	if uart.Buffered() > 0 {
		data, _ := uart.ReadByte()
		x := snake[0].x
		y := snake[0].y
		p := snake[0].p
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
				x = 7
				p = 3
			} else if p == 1 || p == 2 || p == 3 {
				x = 7
				p--
			} else if p == 4 {
				x = y
				y = 0
				p = 3
			} else if p == 5 {
				x = 7 - y
				y = 7
				p = 3
			}
		}

		if x == 8 {
			if p == 0 || p == 1 || p == 2 {
				x = 0
				p++
			} else if p == 3 {
				x = 0
				p = 0
			} else if p == 4 {
				x = 7 - y
				y = 0
				p = 1
			} else if p == 5 {
				x = y
				y = 7
				p = 1
			}
		}

		if y < 0 {
			if p == 0 {
				y = 7
				p = 4
			} else if p == 1 {
				y = 7 - x
				x = 7
				p = 4
			} else if p == 2 {
				y = 0
				x = 7 - x
				p = 4
			} else if p == 3 {
				y = x
				x = 0
				p = 4
			} else if p == 4 {
				x = 7 - x
				y = 0
				p = 2
			} else if p == 5 {
				y = 7
				p = 0
			}
		}

		if y == 8 {
			if p == 0 {
				y = 0
				p = 5
			} else if p == 1 {
				y = x
				x = 7
				p = 5
			} else if p == 2 {
				x = 7 - x
				y = 7
				p = 5
			} else if p == 3 {
				y = 7 - x
				x = 0
				p = 5
			} else if p == 4 {
				y = 0
				p = 0
			} else if p == 5 {
				x = 7 - x
				y = 7
				p = 2
			}
		}

		changeDirection(snake[0].p, p)

		for i := 9; i > 0; i-- {
			snake[i].x = snake[i-1].x
			snake[i].y = snake[i-1].y
			snake[i].p = snake[i-1].p
		}
		snake[0].x = x
		snake[0].y = y
		snake[0].p = p
		println("SNAKE", x, y, p, directionMod)
	}
}

func changeDirection(oldP, p uint8) {
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
}

// only works between {-128,-128} & {255,255}
// We flatten the cube in the following way:
// +---+---+---+---+
// | a | b | c | d | horizontal faces of the cube
// | v | > | ʌ | < | bottom face, rotated
// | ɔ | p | ɐ | q | first row, rotated π rads
// | ↑ | → | ↓ | ← | top face, rotated
// +---+---+---+---+
func cubeCoords(x, y int16) (int16, int16) {
	if x < 0 {
		x += 128
	}
	if x > 127 {
		x -= 128
	}
	if y < 0 {
		y += 128
	}
	if y > 127 {
		y -= 128
	}

	if y < 32 { // horizontal faces
		return 32 + x, y
	} else if y < 64 { // bottom faces
		y -= 32
		if x >= 96 {
			x, y = y, 127-x
		} else if x >= 64 {
			x = 95 - x
			y = 31 - y
		} else if x >= 32 {
			x, y = 31-y, x-32
		}
		return x, y
	} else if y < 96 { // horizontal rotated
		y -= 64
		if x < 32 {
			x = 95 - x
		} else if x < 64 {
			x = 159 - x
		} else if x < 96 {
			x = 95 - x
		} else {
			x = 159 - x
		}
		return 32 + x, 31 - y
	} else if y < 128 { // top face
		y -= 96
		if x >= 96 {
			x, y = 31-y, x-96
		} else if x >= 64 {
			x = 95 - x
			y = 31 - y
		} else if x >= 32 {
			x, y = y, 63-x
		}
		return 160 + x, y
	}
	return 0, 0
}
       */