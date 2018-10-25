package cozely_test

import (
	"math/rand"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/pico8"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
)

//// Constants /////////////////////////////////////////////////////////////////

const (
	gridwidth, gridheight = 18, 18
)

var (
	resolution = pixel.XY{180, 180}
	cellsize   = pixel.XY{8, 8}
	origin     = resolution.Minus(pixel.XY{gridwidth, gridheight}.TimesXY(cellsize)).Slash(2)
)

//// Game Sate /////////////////////////////////////////////////////////////////

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

var (
	grid             [18][18]cell
	head             struct{ X, Y int16 }
	direction, next  cell
	score, bestscore int
)

//// Main //////////////////////////////////////////////////////////////////////

func Example_snake() {
	defer cozely.Recover()

	color.Load(&pico8.Palette)
	pixel.SetResolution(resolution)
	cozely.Configure(cozely.UpdateStep(.25))

	err := cozely.Run(menu{})
	if err != nil {
		panic(err)
	}
	//Output:
}

//// Menu Loop /////////////////////////////////////////////////////////////////

type menu struct{}

func (menu) Enter() {
	setupGrid()
	score = 0
}

func (menu) Leave() {}

func (menu) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}

	if input.Select.Pressed() {
		cozely.Goto(game{})
	}
	return
}

func (menu) Update() {}

func (menu) Render() {
	pixel.Clear(1)
	cur := pixel.Cursor{
		Color:    pico8.White,
		Position: pixel.XY{25, 16},
	}
	cur.Print("Press [space] to start")
	drawGrid()
}

//// Game Loop /////////////////////////////////////////////////////////////////

type game struct{}

func (game) Enter() {
	input.ShowMouse(false)
}

func (game) Leave() {
	input.ShowMouse(true)
}

func (game) React() {
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}

	if input.Up.Pressed() && direction != down {
		next = up
	}
	if input.Right.Pressed() && direction != left {
		next = right
	}
	if input.Down.Pressed() && direction != up {
		next = down
	}
	if input.Left.Pressed() && direction != right {
		next = left
	}
}

func (game) Update() {
	direction = next
	step()
}

func (game) Render() {
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
	if input.Close.Pressed() {
		cozely.Stop(nil)
	}

	if input.Select.Pressed() {
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
		cur := pixel.Cursor{
			Color:    pico8.White,
			Position: pixel.XY{40, 16},
		}
		cur.Print("*** GAME OVER ***")
	}
	drawGrid()
}

//// Game Logic /////////////////////////////////////////////////////////////////

// setupGrid prepares the grid for a new game.
func setupGrid() {
	grid = [gridwidth][gridheight]cell{}
	for x := 0; x < gridwidth; x++ {
		grid[x][0] = border
		grid[x][gridheight-1] = border
	}
	for y := 0; y < gridheight; y++ {
		grid[0][y] = border
		grid[gridwidth-1][y] = border
	}
	head.X, head.Y = gridwidth/2, gridheight/2
	grid[head.X][head.Y] = up
	grid[head.X][head.Y+1] = up
	grid[head.X][head.Y+2] = up
	grid[head.X][head.Y+3] = up
	grid[head.X][head.Y+4] = up
	grid[head.X][head.Y+5] = tail
	direction, next = up, up
	addFruit()

}

// addFruit randaomly places a fruit in an empty cell of the grid.
func addFruit() {
	for {
		x, y := rand.Intn(gridwidth), rand.Intn(gridheight)
		if grid[x][y] == empty {
			grid[x][y] = fruit
			return
		}
	}
}

// step advances the snake in the current direction.
func step() {
	switch direction {
	case up:
		head.Y--
	case right:
		head.X++
	case down:
		head.Y++
	case left:
		head.X--
	}

	if grid[head.X][head.Y] == fruit {
		// Eat and grow
		score++
		grid[head.X][head.Y] = direction
		addFruit()
		return
	}

	if grid[head.X][head.Y] != empty {
		// Hit border or snake body
		cozely.Goto(gameover{})
		return
	}

	// Remove last section of the tail
	grid[head.X][head.Y] = direction
	s := head
	for i := 0; i < gridwidth*gridheight; i++ {
		ns := s
		switch grid[s.X][s.Y] {
		case up:
			ns.Y++
		case right:
			ns.X--
		case down:
			ns.Y--
		case left:
			ns.X++
		}
		if grid[ns.X][ns.Y] == tail {
			grid[s.X][s.Y] = tail
			grid[ns.X][ns.Y] = empty
			return
		}
		s = ns
	}
}

//// Game Rendering ////////////////////////////////////////////////////////////

// drawGrid draws the current game state.
func drawGrid() {
	rectangle := pixel.Box("builtins/rectangle")
	fill := pixel.Box("builtins/fill")
	{
		// Draw background
		o := origin.Plus(cellsize).MinusS(1)
		s := pixel.XY{gridwidth - 2, gridheight - 2}.TimesXY(cellsize).PlusS(3)
		fill.Paint(o, s, 0, pico8.DarkGreen)
		rectangle.Paint(o, s, 0, pico8.Green)
	}

	// Draw grid content
	var s struct{ X, Y int16 }
	for s.X = int16(0); s.X < gridwidth; s.X++ {
		for s.Y = int16(0); s.Y < gridheight; s.Y++ {
			p := pixel.XY(s).TimesXY(cellsize)
			p = origin.Plus(p)
			switch grid[s.X][s.Y] {
			case fruit:
				fill.Paint(p.PlusS(2), cellsize.MinusS(3), 0, pico8.Red)
			case up, right, down, left, tail:
				fill.Paint(p, cellsize.PlusS(1), 0, pico8.Peach)
				rectangle.Paint(p, cellsize.PlusS(1), 0, pico8.DarkPurple)
			}
			if s.X == head.X && s.Y == head.Y {
				fill.Paint(p, cellsize.PlusS(1), 0, pico8.Pink)
				rectangle.Paint(p, cellsize.PlusS(1), 0, pico8.DarkPurple)
				switch next {
				case up:
					pixel.Point(p.Plus(pixel.XY{3, 2}), 0, pico8.DarkBlue)
					pixel.Point(p.Plus(pixel.XY{5, 2}), 0, pico8.DarkBlue)
				case down:
					pixel.Point(p.Plus(pixel.XY{3, 8 - 2}), 0, pico8.DarkBlue)
					pixel.Point(p.Plus(pixel.XY{5, 8 - 2}), 0, pico8.DarkBlue)
				case left:
					pixel.Point(p.Plus(pixel.XY{2, 3}), 0, pico8.DarkBlue)
					pixel.Point(p.Plus(pixel.XY{2, 5}), 0, pico8.DarkBlue)
				case right:
					pixel.Point(p.Plus(pixel.XY{8 - 2, 3}), 0, pico8.DarkBlue)
					pixel.Point(p.Plus(pixel.XY{8 - 2, 5}), 0, pico8.DarkBlue)
				}
			}
		}
	}

	// Display score
	cur := pixel.Cursor{
		Position: pixel.XY{25, pixel.Resolution().Y - 10},
		Color:    pico8.White,
	}
	if score > 0 {
		cur.Printf("Score: %2d", score)
	}
	if bestscore > 0 {
		cur.Position = pixel.XY{109, pixel.Resolution().Y - 10}
		cur.Printf("Best: %2d", bestscore)
	}
}
