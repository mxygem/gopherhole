package board

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Board [][]string

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

// fillBoard pseudorandomly fills the board with
// the desired number of gopher/hole pairs it
// Also allows for a seed to be specified for
// deterministic generation
func fillBoard(d, s int, bo *Board) {
	b := *bo
	if d == 0 {
		return
	}

	setRand(s)
	xl := len(b)
	yl := len(b[0])
	dc := difficulty(xl, yl, d)

	for i := 0; i < dc; i++ {
		// determine hole position
		x := rand.Intn(xl)
		y := rand.Intn(yl)
		fmt.Printf("hole at x: %d y: %d\n", x, y)

		// determine gopher position
		// position is "random" while taking into
		// account the edge of the board
		// 0 = up
		// 1 = right
		// 2 = down
		// 3 = left
		di := rand.Intn(4)
		fmt.Println(di)

		// TODO: Gopher placement direction
		// 1. Boundary checks - make sure position is
		//		within board limits
		// 2. Occupation checks - make sure position
		//		doesn't already have data
		// 3. If unable to place gopher, will need to
		// 		find a new spot for both

		// place gopher
		switch di {
		case 0:
			fmt.Printf("gopher up from x: %d y: %d\n", x, y)
			// ok := canPlaceUp()
			// b[x-1][y] = "g"
		case 1:
			fmt.Printf("gopher right from x: %d y: %d\n", x, y)
			// ok := canPlaceRight()
			// b[x][y+1] = "g"
		case 2:
			fmt.Printf("gopher down from x: %d y: %d\n", x, y)
			// ok := canPlaceDown()
			// b[x+1][y] = "g"
		case 3:
			fmt.Printf("gopher left from x: %d y: %d\n", x, y)
			// ok := canPlaceLeft()
			// b[x][y-1] = "g"
		}

		// place hole
		b[x][y] = "o"

		printBoard(&b)
	}
}

func canPlaceUp(x, y, d int, b Board) bool {
	xl := len(b)
	yl := len(b[0])

	ok := withinBounds(xl, yl, x, y, d)
	if !ok {
		return ok
	}

	return spaceOpen(x-1, y, &b)
}

func withinBounds(xl, yl, x, y, d int) bool {
	switch d {
	case 0: // up
		if x == 0 {
			return false
		}
	case 1: // right
		if y >= (yl - 1) {
			return false
		}
	case 2: // lower
		if x <= (xl - 1) {
			return false
		}
	case 3: // left
		if y == 0 {
			return false
		}
	}

	return true
}

func spaceOpen(x, y int, b *Board) bool {
	return true
}

func printBoard(b *Board) {
	for _, r := range *b {
		fmt.Println(r)
	}
}

// setRand sets the intended seed otherwise uses
// a time based seed
func setRand(s int) {
	if s > 0 {
		rand.NewSource(int64(s))
		return
	}

	rand.NewSource(time.Now().UnixNano())
}
