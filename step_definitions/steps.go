package stepdefinitions

import (
	"fmt"

	"github.com/jaysonesmith/gopherhole/board"
	"github.com/jaysonesmith/gopherhole/support"
)

func (sc *ScenarioContext) ANewGameIsRequestedWithNoBoardSizeSet() error {
	sc.Board = board.New(0, 0, 0)

	return nil
}

func (sc *ScenarioContext) ANewGameIsRequestedWithAXBoardSize(x, y int) error {
	sc.Board = board.New(x, y, 0)

	return nil
}

func (sc *ScenarioContext) AXBoardMustBeReturned(x, y int) error {
	err := support.CheckBoardDimensions(x, y, sc.Board)
	if err != nil {
		return fmt.Errorf("board dimensions incorrect: %s", err.Error())
	}

	return nil
}
