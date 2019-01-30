package stepdefinitions

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/jaysonesmith/gopherhole/board"
	"github.com/jaysonesmith/gopherhole/support"
)

func (sc *ScenarioContext) Steps(s *godog.Suite) {
	s.Step(`^a new game is requested with no board size set$`, sc.ANewGameIsRequestedWithNoBoardSizeSet)
	s.Step(`^a (\d+)x(\d+) board must be returned$`, sc.AXBoardMustBeReturned)
	s.Step(`^a new game is requested with a (\d+)x(\d+) board size$`, sc.ANewGameIsRequestedWithAXBoardSize)
	s.Step(`^a (\d+)x(\d+) board is filled at (\w+) difficulty$`, sc.AXBoardIsFilledAtDifficulty)
	s.Step(`^approximately (\d+) spaces will be filled$`, sc.ApproximatelySpacesWillBeFilled)
}

func (sc *ScenarioContext) ANewGameIsRequestedWithNoBoardSizeSet() error {
	sc.Board = board.New(0, 0)

	return nil
}

func (sc *ScenarioContext) ANewGameIsRequestedWithAXBoardSize(x, y int) error {
	sc.Board = board.New(x, y)

	return nil
}

func (sc *ScenarioContext) AXBoardMustBeReturned(x, y int) error {
	err := support.CheckBoardDimensions(x, y, sc.Board)
	if err != nil {
		return fmt.Errorf("board dimensions incorrect: %s", err.Error())
	}

	return nil
}

func (sc *ScenarioContext) AXBoardIsFilledAtDifficulty(x, y int, d string) error {
	sc.Board = board.New(x, y)

	difficulty := map[string]int{"medium": 1, "hard": 2}
	sc.Board.Fill(difficulty[d], 1)

	return nil
}

func (sc *ScenarioContext) ApproximatelySpacesWillBeFilled(c int) error {

	if f, ok := filledCount(c, sc.Board); !ok {
		return fmt.Errorf("expected amount of spaces were not filled. expected: %d found: %d", c, f)
	}

	return nil
}
