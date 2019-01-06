package board

type Board [][]string

// New returns a new board to the user based on
// an input dimension number. The board size
// returned defaults to 4 for anything under that
// size.
func New(x, y int) Board {
	x = defaultDim(x)
	y = defaultDim(y)

	b := make([][]string, x)
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
