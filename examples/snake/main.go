package main

import (
	"math/rand"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/window"
)

//// Game Sate /////////////////////////////////////////////////////////////////

const (
	width  = 18
	height = 18
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
	border
	fruit
	up
	right
	down
	left
	tail
)

//// Main //////////////////////////////////////////////////////////////////////

func main() {
	defer cozely.Recover()

	pixel.SetResolution(pixel.XY{180, 180})
	cozely.Configure(cozely.UpdateStep(.25))
	window.Events.Resize = func() {
		origin = pixel.Resolution().Minus(pixel.XY{width, height}.TimesXY(cellsize)).Slash(2)
	}

	err := cozely.Run(menu{})
	if err != nil {
		panic(err)
	}
}

//// Menu Loop /////////////////////////////////////////////////////////////////

type menu struct{}

func (menu) Enter() {
	setupGrid()
	score = 0
}

func (menu) Leave() {}

func (menu) React() {
	if input.MenuBack.Pressed() {
		cozely.Stop(nil)
	}

	if input.MenuSelect.Pressed() {
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

//// Game Loop /////////////////////////////////////////////////////////////////

type loop struct{}

func (loop) Enter() {
	input.ShowMouse(false)
}

func (loop) Leave() {
	input.ShowMouse(true)
}

func (loop) React() {
	if input.MenuBack.Pressed() {
		cozely.Stop(nil)
	}

	if input.MenuUp.Pressed() && direction != down {
		next = up
	}
	if input.MenuRight.Pressed() && direction != left {
		next = right
	}
	if input.MenuDown.Pressed() && direction != up {
		next = down
	}
	if input.MenuLeft.Pressed() && direction != right {
		next = left
	}
}

func (loop) Update() {
	direction = next
	switch direction {
	case up:
		snake.Y--
	case right:
		snake.X++
	case down:
		snake.Y++
	case left:
		snake.X--
	}
	advance()
}

func (loop) Render() {
	pixel.Clear(1)
	drawGrid()
}

//// Game Over Loop ////////////////////////////////////////////////////////////

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
	if input.MenuBack.Pressed() {
		cozely.Stop(nil)
	}

	if input.MenuSelect.Pressed() {
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
	pixel.Clear(8)
	if counter%2 == 0 {
		c := pixel.Cursor{
			Position: pixel.XY{40, 16},
		}
		c.Print("*** GAME OVER ***")
	}
	drawGrid()
}

//// Game Logic /////////////////////////////////////////////////////////////////

func setupGrid() {
	grid = [width][height]cell{}
	for x := 0; x < width; x++ {
		grid[x][0] = border
		grid[x][height-1] = border
	}
	for y := 0; y < height; y++ {
		grid[0][y] = border
		grid[width-1][y] = border
	}
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

//// Game Rendering ////////////////////////////////////////////////////////////

var (
	cellsize = pixel.XY{8, 8}
	origin   pixel.XY
)

func drawGrid() {
	pixel.Box(
		11, 3, 0, 0,
		origin.Plus(cellsize).MinusS(1),
		origin.Plus(pixel.XY{width, height}.TimesXY(cellsize)).Minus(cellsize).PlusS(1),
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
				pixel.Box(
					2, 15, 0, 2,
					p,
					p.Plus(cellsize),
				)
			}
			if x == snake.X && y == snake.Y {
				pixel.Box(
					2, 14, 0, 2,
					p,
					p.Plus(cellsize),
				)
			}
		}
	}
	c := pixel.Cursor{
		Position: pixel.XY{25, pixel.Resolution().Y - 10},
	}
	if score > 0 {
		c.Printf("Score: %2d", score)
	}
	if bestscore > 0 {
		c.Position = pixel.XY{109, pixel.Resolution().Y - 10}
		c.Printf("Best: %2d", bestscore)
	}
}
