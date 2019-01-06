package board

import "math"

type Board [][]string

// New returns a new board to the user based on
// an input dimension number. The board size
// returned defaults to 4 for anything under that
// size.
func New(x, y, d int) Board {
	x = defaultDim(x)
	y = defaultDim(y)

	b := make(Board, x)
	for i := range b {
		b[i] = make([]string, y)
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
// Also allows for a seed to be specified for testing
// purposes
// func fillBoard(p, s int, b *Board) {
// 	if p == 0 {
// 		return
// 	}

// }
