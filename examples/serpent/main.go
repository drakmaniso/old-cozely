package main

import (
	"math/rand"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

////////////////////////////////////////////////////////////////////////////////

var (
	resolution = pixel.XY{180, 180}
	cellsize   = pixel.XY{8, 8}
	origin     pixel.XY
)

const (
	width  = 16
	height = 16
)

var (
	grid            [width][height]cell
	snake           struct{ X, Y int16 }
	direction, next cell
)

type cell byte

const (
	empty cell = iota
	fruit
	up
	right
	down
	left
	tail
)

const (
	paused = iota
	playing
	gameover
)

var state int

////////////////////////////////////////////////////////////////////////////////

func main() {
	defer cozely.Recover()

	pixel.SetResolution(pixel.XY{resolution.X, resolution.Y})
	cozely.Configure(cozely.UpdateStep(.25))

	err := cozely.Run(loop{})
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

func (loop) Enter() {
	setupGrid()
	state = paused
}

func setupGrid() {
	grid = [width][height]cell{}
	snake.X, snake.Y = width/2, height/2
	grid[snake.X][snake.Y] = up
	grid[snake.X][snake.Y+1] = up
	grid[snake.X][snake.Y+2] = up
	grid[snake.X][snake.Y+3] = up
	grid[snake.X][snake.Y+4] = up
	grid[snake.X][snake.Y+5] = tail
	direction, next = up, up
	addFruit()

}

func addFruit() {
	for {
		x, y := rand.Intn(width), rand.Intn(height)
		if grid[x][y] == empty {
			grid[x][y] = fruit
			return
		}
	}
}

func (loop) Leave() {
}

////////////////////////////////////////////////////////////////////////////////

func (loop) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}

	if state == paused {
		if input.MenuSelect.Pushed() {
			state = playing
		}
		return
	}

	if state == gameover {
		if input.MenuSelect.Pushed() {
			setupGrid()
			state = paused
		}
		return
	}

	if input.MenuUp.Pushed() && direction != down {
		next = up
	}
	if input.MenuRight.Pushed() && direction != left {
		next = right
	}
	if input.MenuDown.Pushed() && direction != up {
		next = down
	}
	if input.MenuLeft.Pushed() && direction != right {
		next = left
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Update() {
	if state != playing {
		return
	}

	direction = next

	switch direction {
	case up:
		if snake.Y == 0 {
			state = gameover
			return
		}
		snake.Y--
		advance()
	case right:
		if snake.X == width-1 {
			state = gameover
			return
		}
		snake.X++
		advance()
	case down:
		if snake.Y == height-1 {
			state = gameover
			return
		}
		snake.Y++
		advance()
	case left:
		if snake.X == 0 {
			state = gameover
			return
		}
		snake.X--
		advance()
	}
}

func advance() {
	if grid[snake.X][snake.Y] == fruit {
		grid[snake.X][snake.Y] = direction
		addFruit()
		return
	}
	if grid[snake.X][snake.Y] != empty {
		state = gameover
		return
	}
	grid[snake.X][snake.Y] = direction
	x, y := snake.X, snake.Y
	for i := 0; true; i++ {
		xx, yy := x, y
		switch grid[x][y] {
		case up:
			yy++
		case right:
			xx--
		case down:
			yy--
		case left:
			xx++
		default:
			println("OUCH!")
			panic(nil)
		}
		if grid[xx][yy] == tail {
			grid[xx][yy] = empty
			grid[x][y] = tail
			return
		}
		x, y = xx, yy
		if i > 1000 {
			panic(nil)
			println("OOO?")
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() {
	pixel.Clear(1)
	c := pixel.Cursor{}
	if state == paused {
		c.Print("Press [space] to start")
	}
	drawGrid()
	if state == gameover {
		c.Print("*** Game Over ***")
	}
}

func drawGrid() {
	origin = resolution.Minus(pixel.XY{width, height}.TimesXY(cellsize)).Slash(2)
	pixel.Box(
		11, 3, 0, 0,
		origin.MinusS(1),
		origin.Plus(pixel.XY{width, height}.TimesXY(cellsize)).PlusS(1),
	)
	for x := int16(0); x < width; x++ {
		for y := int16(0); y < height; y++ {
			p := pixel.XY{x, y}.TimesXY(cellsize)
			p = origin.Plus(p)
			switch grid[x][y] {
			case fruit:
				pixel.Box(
					8, 8, 0, 5,
					p,
					p.Plus(cellsize),
				)
			case up, right, down, left, tail:
				if x == snake.X && y == snake.Y {
					pixel.Box(
						2, 14, 0, 2,
						p,
						p.Plus(cellsize),
					)
					break
				}
				pixel.Box(
					2, 15, 0, 2,
					p,
					p.Plus(cellsize),
				)
			}
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
