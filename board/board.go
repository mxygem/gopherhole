package board

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var r *rand.Rand

// Board is the representation of the game's
// playing field
type Board [][]string

type positions struct {
	gx int
	gy int
	hx int
	hy int
}

// New returns a new board to the user based on
// an input dimension number. The board size
// returned defaults to 4 for anything under that
// size.
func New(x, y int) Board {
	x = defaultDim(x)
	y = defaultDim(y)

	b := make(Board, x)
	for i := range b {
		b[i] = make([]string, y)
		for j := range b[i] {
			b[i][j] = " "
		}
	}

	return b
}

func defaultDim(i int) int {
	if i < 4 {
		return 4
	}

	return i
}

// Fill pseudorandomly fills the board with
// the desired number of gopher/hole pairs. It
// also allows for a seed to be specified for
// deterministic generation
func (b Board) Fill(d, s int) {
	// empty difficulty
	if d == 0 {
		return
	}

	setRand(s)
	dc := difficulty(len(b), len(b[0]), d)

	// find spots for all the desired gopher/hole pairs
	for i := 0; i < dc; i++ {
		pos := b.findPositions()
		if pos.gx == -1 {
			continue
		}

		b[pos.gx][pos.gy] = "g"
		b[pos.hx][pos.hy] = "o"

		// b.Print()
	}
}

// Start removes the gophers from a board
func (b Board) Start() {
	for i, r := range b {
		for j, c := range r {
			if c == "g" {
				b[i][j] = " "
			}
		}
	}
}

// setRand sets the intended seed otherwise uses
// a time based seed
func setRand(s int) {
	if s > 0 {
		r = rand.New(rand.NewSource(int64(s)))
		return
	}

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

var diffLevels = map[int]float64{0: 0.0, 1: 0.3, 2: 0.5}

// difficulty determines how many gopher/hole
// pairs should be used to fill in the current
// board for the desired difficulty
// levels:
// 0 - empty
// 1 - medium - up to 30% of space
// 2 - hard - up to 50% of space
func difficulty(x, y, d int) int {
	return int(math.Ceil((float64(x*y) * diffLevels[d]) / 2))
}

// findPositions looks for a valid position to
// place a gopher and then finds a valid place to
// place a hole based on that gopher if one
// exists
func (b Board) findPositions() positions {
	gx, gy := b.gopherArea()
	if gx == -1 {
		return positions{-1, -1, -1, -1}
	}

	hx, hy := b.holeArea(gx, gy)
	if hx == -1 {
		return b.findPositions()
	}

	return positions{gx, gy, hx, hy}
}

func (b Board) gopherArea() (int, int) {
	// pick a gopher position and then make sure
	// it's empty AND its surrounding 8 squares
	// don't already contain a gopher otherwise,
	// look elsewhere.
	// Not an optimal solution.
	var x, y int

	xl := len(b)
	yl := len(b[0])

	for i := 0; i < xl*yl; i++ {
		x = r.Intn(xl)
		y = r.Intn(yl)

		if b[x][y] == " " && !b.surroundingGopher(x, y) {
			return x, y
		}
	}

	return -1, -1
}

// surroundingGopher checks for whether or not
// the 8 positions surrounding the passed in
// coordinates contain a gopher
func (b Board) surroundingGopher(x, y int) bool {
	pos := [][]int{
		// top row
		[]int{-1, -1},
		[]int{0, -1},
		[]int{+1, -1},
		// middle row
		[]int{-1, 0},
		[]int{+1, 0},
		// bottom row
		[]int{-1, +1},
		[]int{0, +1},
		[]int{+1, +1},
	}

	for _, p := range pos {
		nx := p[0] + x
		ny := p[1] + y

		if b.withinBounds(nx, ny) && b[nx][ny] == "g" {
			return true
		}
	}

	return false
}

// Similar to gopherArea, except it checks for in
// bounds before checking for empty.
// Returns -1, -1 if no suitable space was found
func (b Board) holeArea(x, y int) (int, int) {
	// determining hole position position is
	// "random" while taking into account the edge
	// of the board
	// 0 = up
	// 1 = right
	// 2 = down
	// 3 = left
	di := shuffleDirections()

	for _, i := range di {
		pos := [][]int{
			[]int{x, y + 1},
			[]int{x + 1, y},
			[]int{x, y - 1},
			[]int{x - 1, y},
		}

		px := pos[i-1][0]
		py := pos[i-1][1]

		ok := b.canPlace(px, py)
		if !ok {
			continue
		}

		return px, py
	}

	return -1, -1
}

func shuffleDirections() []int {
	d := []int{1, 2, 3, 4}

	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

	return d
}

func (b Board) canPlace(x, y int) bool {
	if inBounds := b.withinBounds(x, y); !inBounds {
		return false
	}

	// check occupancy
	if b[x][y] == " " {
		return true
	}

	return false
}

func (b Board) withinBounds(x, y int) bool {
	xl := len(b)
	yl := len(b[0])

	// check within bounds
	if x < 0 || x >= xl || y < 0 || y >= yl {
		return false
	}

	return true
}

func (b Board) Print() {
	for _, r := range b {
		fmt.Println(r)
	}
	fmt.Println("----------")
}

// WriteChar is the method used to add a
// character to a board
func (b Board) WriteChar(input string, x, y int) error {
	b[x][y] = input
	return nil
}
