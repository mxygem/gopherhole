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
	x  int
	y  int
	gx int
	gy int
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
// the desired number of gopher/hole pairs it
// Also allows for a seed to be specified for
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
		// place hole
		b[pos.x][pos.y] = "o"
		// place gopher
		b[pos.gx][pos.gy] = "g"

		printBoard(&b)
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

var diffLevels = map[int]float64{0: 0.0, 1: 0.3, 2: 0.5, 3: 0.7}

// difficulty determines how many gopher/hole
// pairs should be used to fill in the current
// board for the desired difficulty
// levels:
// 0 - empty
// 1 - easy - 30% of space
// 2 - medium - 50% of space
// 3 - hard - 70% of space
func difficulty(x, y, d int) int {
	return int(math.Ceil((float64(x*y) * diffLevels[d]) / 2))
}

func findPositions(xl, yl int, b Board) positions {
	x, y := emptyArea(xl, yl, b)
	fmt.Printf("empty area at x: %d y: %d\n", x, y)

	gx, gy := gopherArea(x, y, b)
	fmt.Printf("gopher area at x: %d y: %d\n", x, y)
	if gx == -1 {
		return findPositions(xl, yl, b)
	}

	return positions{x: x, y: y, gx: gx, gy: gy}
}

func emptyArea(xl, yl int, b Board) (int, int) {
	// determine hole position and make sure it's
	// empty, otherwise, look elsewhere. Not an
	// optimal solution.
	var x, y int
	empty := false
	for empty == false {
		x = rand.Intn(xl)
		y = rand.Intn(yl)
		empty = spaceOpen(x, y, &b)
	}

	return x, y
}

// Similar to emptyArea, except it checks for in
// bounds before checking for empty.
// Returns -1, -1 if no suitable space was found
func gopherArea(x, y int, b Board) (int, int) {
	// determine gopher position
	// position is "random" while taking into
	// account the edge of the board
	// 0 = up
	// 1 = right
	// 2 = down
	// 3 = left
	di := shuffleDirections()

	for _, i := range di {
		// check for open position in direction
		switch i {
		case 0:
			ok := canPlace(x-1, y, b)
			if !ok {
				continue
			}

			return x - 1, y
		case 1:
			ok := canPlace(x, y+1, b)
			if !ok {
				continue
			}

			return x, y + 1
		case 2:
			ok := canPlace(x+1, y, b)
			if !ok {
				continue
			}

			return x + 1, y
		case 3:
			ok := canPlace(x, y-1, b)
			if !ok {
				continue
			}

			return x, y - 1
		}
	}

	return -1, -1
}

func shuffleDirections() []int {
	d := []int{1, 2, 3, 4}

	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

	fmt.Println("shuffled to:", d)
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
		fmt.Printf("b[%d][%d] = %s", x, y, b[x][y])
		return false
	}

	return true
}

func spaceOpen(x, y int, bo *Board) bool {
	fmt.Printf("checking for open space at %d,%d\n", x, y)
	b := *bo

	if b[x][y] == " " {
		fmt.Println("open")
		return true
	}

	fmt.Println("not open")
	return false
}

func printBoard(b *Board) {
	for _, r := range *b {
		fmt.Println(r)
	}
}
