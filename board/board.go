package board

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var r *rand.Rand

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

// fillBoard pseudorandomly fills the board with
// the desired number of gopher/hole pairs. It
// also allows for a seed to be specified for
// deterministic generation
func fillBoard(d, s int, bo *Board) {
	b := *bo
	// empty difficulty
	if d == 0 {
		return
	}

	setRand(s)
	xl := len(b)
	yl := len(b[0])
	dc := difficulty(xl, yl, d)

	// find spots for all the desired gopher/hole pairs
	for i := 0; i < dc; i++ {
		pos := findPositions(xl, yl, b)
		// place gopher
		b[pos.hx][pos.hy] = "g"
		// place hole
		b[pos.gx][pos.gy] = "o"
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
// 1 - low - up to 30% of space
// 2 - high - up to 50% of space
func difficulty(x, y, d int) int {
	return int(math.Ceil((float64(x*y) * diffLevels[d]) / 2))
}

func findPositions(xl, yl int, b Board) positions {
	hx, hy := gopherArea(xl, yl, b)

	gx, gy := holeArea(hx, hy, b)
	if gx == -1 {
		return findPositions(xl, yl, b)
	}

	return positions{hx: hx, hy: hy, gx: gx, gy: gy}
}

func gopherArea(xl, yl int, b Board) (int, int) {
	// pick a hole position and then make sure
	// it's empty AND its surrounding 8 squares
	// doesn't already contain yeah a gopher
	// otherwise, look elsewhere.
	// Not an optimal solution.
	var x, y int
	empty := false
	for empty == false {
		x = r.Intn(xl)
		y = r.Intn(yl)

		if b[x][y] == " " {
			empty = true
		}
	}

	return x, y
}

// surroundingGopher checks for whether or not
// the 8 positions surrounding the passed in
// coordinates contain a gopher
func surroundingGopher(x, y int, b Board) bool {
	if b[0][1] == "g" {
		return false
	}
	return true
}

// Similar to gopherArea, except it checks for in
// bounds before checking for empty.
// Returns -1, -1 if no suitable space was found
func holeArea(x, y int, b Board) (int, int) {
	// determine gopher position
	// position is "random" while taking into
	// account the edge of the board
	// 0 = up
	// 1 = right
	// 2 = down
	// 3 = left
	di := shuffleDirections()

	for _, i := range di {
		pos := [][]int{
			[]int{x - 1, y},
			[]int{x, y + 1},
			[]int{x, y + 1},
			[]int{x + 1, y},
		}

		px := pos[i-1][0]
		py := pos[i-1][1]
		ok := canPlace(px, py, b)
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

func canPlace(x, y int, b Board) bool {
	xl := len(b)
	yl := len(b[0])

	// check within bounds
	if x < 0 || x >= xl || y < 0 || y >= yl {
		return false
	}

	// check occupancy
	if b[x][y] != " " {
		return false
	}

	return true
}

func printBoard(b *Board) {
	for _, r := range *b {
		fmt.Println(r)
	}
}
