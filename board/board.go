package board

type Board [][]string

// New returns a new board to the user based on
// an input dimension number. The board size
// returned defaults to 4 for anything under that
// size.
func New(x int) Board {
	if x < 4 {
		x = 4
	}

	b := make([][]string, x)
	for i := range b {
		b[i] = make([]string, x)
	}

	return b
}
