package board

type Board [][]string

func New(x, y int) Board {
	return make(Board, x)
}
