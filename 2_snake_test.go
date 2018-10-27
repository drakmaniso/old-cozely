package cozely_test

import (
	"math/rand"

	"github.com/cozely/cozely"
	"github.com/cozely/cozely/color"
	"github.com/cozely/cozely/color/pico8"
	"github.com/cozely/cozely/input"
	"github.com/cozely/cozely/pixel"
	"github.com/cozely/cozely/resource"
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

	err := resource.Path("testdata/")
	if err != nil {
		panic(err)
	}
	color.Load(&pico8.Palette)
	pixel.SetResolution(pixel.XY{180, 180})
	cozely.Configure(cozely.UpdateStep(.25))

	err = cozely.Run(menu{})
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
	pixel.Clear(pico8.DarkBlue)
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
	pixel.Clear(pico8.DarkBlue)
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
	pixel.Clear(pico8.Red)
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
	for x := range grid {
		for y := range grid[x] {
			if x == 0 || x == len(grid)-1 || y == 0 || y == len(grid[x])-1 {
				grid[x][y] = border
			} else {
				grid[x][y] = empty
			}
		}
	}
	head.X, head.Y = int16(len(grid)/2), int16(len(grid[0])/2)
	grid[head.X][head.Y] = up
	grid[head.X][head.Y+1] = up
	grid[head.X][head.Y+2] = up
	grid[head.X][head.Y+3] = up
	grid[head.X][head.Y+4] = up
	grid[head.X][head.Y+5] = tail
	direction, next = up, up
	addFruit()

}

// addFruit randomly places a fruit in an empty cell of the grid.
func addFruit() {
	for {
		x, y := rand.Intn(len(grid)), rand.Intn(len(grid[0]))
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
	var c cell
	c, grid[head.X][head.Y] = grid[head.X][head.Y], direction

	if c == fruit {
		// Eat and grow
		score++
		addFruit()
		return
	}

	if c != empty {
		// Hit border or snake body
		cozely.Goto(gameover{})
		return
	}

	// Remove last section of the tail
	s := head
	for i := 0; i < len(grid)*len(grid[0]); i++ {
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
	g := pixel.XY{int16(len(grid)), int16(len(grid[0]))} // grid size
	s := pixel.Picture("body").Size().Minuss(1)          // cell size
	o := pixel.Resolution().Minus(s.Times(g)).Slashs(2)  // grid origin

	// Draw grid background
	pixel.Box("playground").Paint(
		o.Plus(s),
		s.Times(g.Minuss(2)).Pluss(1),
		0,
		0,
	)

	// Draw grid content
	for x := range grid {
		for y := range grid[x] {
			i := pixel.XY{int16(x), int16(y)}
			p := o.Plus(i.Times(s))
			switch grid[i.X][i.Y] {
			case fruit:
				pixel.Picture("fruit").Paint(p, 0, 0)
			case up, right, down, left, tail:
				if i == head {
					pixel.Picture("head").Paint(p, 0, 0)
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
				} else {
					pixel.Picture("body").Paint(p, 0, 0)
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
