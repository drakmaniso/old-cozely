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
	grid             [width][height]cell
	snake            struct{ X, Y int16 }
	direction, next  cell
	score, bestscore int
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

////////////////////////////////////////////////////////////////////////////////

func main() {
	defer cozely.Recover()

	pixel.SetResolution(pixel.XY{resolution.X, resolution.Y})
	cozely.Configure(cozely.UpdateStep(.25))

	err := cozely.Run(menu{})
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////

type menu struct{}

func (menu) Enter() {
	setupGrid()
	score = 0
}

func (menu) Leave() {}

func (menu) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}

	if input.MenuSelect.Pushed() {
		cozely.Goto(loop{})
	}
	return
}

func (menu) Update() {}

func (menu) Render() {
	pixel.Clear(1)
	c := pixel.Cursor{
		Position: pixel.XY{25, 16},
	}
	c.Print("Press [space] to start")
	drawGrid()
}

////////////////////////////////////////////////////////////////////////////////

type loop struct{}

func (loop) Enter() {
	input.ShowMouse(false)
}

func (loop) Leave() {
	input.ShowMouse(true)
}

func (loop) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
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

func (loop) Update() {
	direction = next

	switch direction {
	case up:
		if snake.Y == 0 {
			cozely.Goto(gameover{})
			return
		}
		snake.Y--
		advance()
	case right:
		if snake.X == width-1 {
			cozely.Goto(gameover{})
			return
		}
		snake.X++
		advance()
	case down:
		if snake.Y == height-1 {
			cozely.Goto(gameover{})
			return
		}
		snake.Y++
		advance()
	case left:
		if snake.X == 0 {
			cozely.Goto(gameover{})
			return
		}
		snake.X--
		advance()
	}
}

////////////////////////////////////////////////////////////////////////////////

func (loop) Render() {
	pixel.Clear(1)
	drawGrid()
}

////////////////////////////////////////////////////////////////////////////////

type gameover struct{}

var counter int

func (gameover) Enter() {
	counter = 0
}

func (gameover) Leave() {
	if score > bestscore {
		bestscore = score
	}
}

func (gameover) React() {
	if input.MenuBack.Pushed() {
		cozely.Stop(nil)
	}

	if input.MenuSelect.Pushed() {
		cozely.Goto(menu{})
	}
	return
}

func (gameover) Update() {
	counter++
	if counter == 16 {
		cozely.Goto(menu{})
	}
}

func (gameover) Render() {
	pixel.Clear(2)
	if counter%2 == 0 {
		c := pixel.Cursor{
			Position: pixel.XY{40, 16},
		}
		c.Print("*** GAME OVER ***")
	}
	drawGrid()
}

////////////////////////////////////////////////////////////////////////////////

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

func advance() {
	if grid[snake.X][snake.Y] == fruit {
		grid[snake.X][snake.Y] = direction
		addFruit()
		score++
		return
	}
	if grid[snake.X][snake.Y] != empty {
		cozely.Goto(gameover{})
		return
	}
	grid[snake.X][snake.Y] = direction
	x, y := snake.X, snake.Y
	for i := 0; i < width*height; i++ {
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
			panic(nil)
		}
		if grid[xx][yy] == tail {
			grid[xx][yy] = empty
			grid[x][y] = tail
			return
		}
		x, y = xx, yy
	}
	panic(nil)
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
	c := pixel.Cursor{
		Position: pixel.XY{25, 170},
	}
	if score > 0 {
		c.Printf("Score: %2d", score)
	}
	if bestscore > 0 {
		c.Position = pixel.XY{109, 170}
		c.Printf("Best: %2d", bestscore)
	}
}
