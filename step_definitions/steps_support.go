package stepdefinitions

import "github.com/jaysonesmith/gopherhole/board"

func filledCount(c int, b board.Board) (int, bool) {
	var f int

	for _, r := range b {
		for _, s := range r {
			if s != " " {
				f++
			}
		}
	}

	if f <= c-2 && f >= c+2 {
		return f, false
	}

	return f, true
}
