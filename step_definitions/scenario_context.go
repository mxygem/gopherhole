package stepdefinitions

import (
	"github.com/jaysonesmith/gopherhole/board"
)

type ScenarioContext struct {
	Board board.Board
	Item  string
	X     int
	Y     int
}
