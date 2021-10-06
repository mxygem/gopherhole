package game

import "github.com/jaysonesmith/gopherhole/board"

// Game is the foundation object that contains
// the board, its status, and the current error
// if any
type Game struct {
	Board  *board.Board
	Status int
	Error  string
}

// New returns a pointer to a new Game
func New() *Game {
	return &Game{}
}
