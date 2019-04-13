package stepdefinitions

import (
	"github.com/jaysonesmith/gopherhole/board"
)

type ScenarioContext struct {
	Board  board.Board
	Char   string
	X      int
	Y      int
	Errors Errors
}

type Errors struct {
	PlacementError string
}
