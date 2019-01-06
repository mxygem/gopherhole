package board

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

	fillBoard(d, &b)

	return b
}

func defaultDim(i int) int {
	if i < 4 {
		return 4
	}

	return i
}

func fillBoard(d int, b *Board) {
	if d == 0 {
		return
	}

}

// difficulty determines how many gopher/hole
// pairs should be used to fill in the current board
// levels:
// 0 - empty
// 1 - easy - 25% of space
// 2 - medium - 50% of space
// 3 - hard - 75% of space
func difficulty(x, y, d int) int {
	switch d {
	case 1:
		return 3
	default:
		return 0
	}
}
