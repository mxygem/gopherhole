package stepdefinitions

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/jaysonesmith/gopherhole/board"
	"github.com/jaysonesmith/gopherhole/support"
	"github.com/jaysonesmith/gopherhole/utils"
	"github.com/pkg/errors"
)

var difficulty = map[string]int{"medium": 1, "hard": 2}
var characters = map[string]string{
	"gopher": "g", "gophers": "g",
	"hole": "h", "holes": "h",
	"space": " ", "spaces": " ",
	"earth": "e",
}

// Steps is a mapping of gherkin regex to step
// defintion methods
func (sc *ScenarioContext) Steps(s *godog.Suite) {
	s.Step(`^a new game is requested with no board size set$`, sc.ANewGameIsRequestedWithNoBoardSizeSet)
	s.Step(`^a (\d+)x(\d+) board must be returned$`, sc.AXBoardMustBeReturned)
	s.Step(`^a new game is requested with a (\d+)x(\d+) board size$`, sc.ANewGameIsRequestedWithAXBoardSize)
	s.Step(`^a (\d+)x(\d+) board is filled at (\w+) difficulty$`, sc.AXBoardIsFilledAtDifficulty)
	s.Step(`^approximately (\d+) spaces will be filled$`, sc.ApproximatelySpacesWillBeFilled)
	s.Step(`^a medium (\d+)x(\d+) board$`, sc.AMediumXBoard)
	s.Step(`^a new game is started$`, sc.ANewGameIsStarted)
	s.Step(`^no gophers should be returned to the player$`, sc.NoGophersShouldBeReturnedToThePlayer)

	s.Step(`^a (\d+)x(\d+) sized board full of (\w+)$`, sc.ABoardFullOf)
	s.Step(`^a (\w+) is entered into position \((\d+), (\d+)\)$`, sc.IsEnteredToPosition)
	s.Step(`^that position must contain the expected character$`, sc.ThatPositionMustContainTheExpectedCharacter)
	s.Step(`^that position must contain a (\w+) character$`, sc.ThatPositionMustContainACharacter)
	s.Step(`^a placement error of "([^"]*)" must be returned$`, sc.APlacementErrorOfMustBeReturned)
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
	sc.Board.Fill(difficulty[d], 1)

	return nil
}

func (sc *ScenarioContext) ApproximatelySpacesWillBeFilled(c int) error {
	if f, ok := utils.FilledCount(c, sc.Board); !ok {
		return fmt.Errorf("expected amount of spaces were not filled. expected: %d found: %d", c, f)
	}

	return nil
}

func (sc *ScenarioContext) AMediumXBoard(x, y int) error {
	sc.Board = board.New(x, y)
	sc.Board.Fill(1, 1)

	return nil
}

func (sc *ScenarioContext) ANewGameIsStarted() error {
	sc.Board.RemoveGophers()

	return nil
}

func (sc *ScenarioContext) NoGophersShouldBeReturnedToThePlayer() error {
	if utils.GophersExist(sc.Board) {
		sc.Board.Print()
		return fmt.Errorf("Unexpected gophers found in board")
	}

	return nil
}

func (sc *ScenarioContext) ABoardFullOf(x, y int, item string) error {
	sc.Board = board.New(x, y)
	utils.FillBoardWith(characters[item], sc.Board)

	return nil
}

func (sc *ScenarioContext) IsEnteredToPosition(char string, x, y int) error {
	c := characters[char]
	sc.Char = c
	sc.X = x
	sc.Y = y
	if err := sc.Board.WriteChar(c, x, y); err != nil {
		sc.Errors.PlacementError = err.Error()
	}

	return nil
}

func (sc *ScenarioContext) ThatPositionMustContainTheExpectedCharacter() error {
	foundChar, err := sc.Board.CharAt(sc.X, sc.Y)
	if err != nil {
		return nil
	}

	if sc.Char != foundChar {
		return errors.Errorf("%s not found at (%d, %d)", sc.Char, sc.X, sc.Y)
	}

	return nil
}

func (sc *ScenarioContext) ThatPositionMustContainACharacter(char string) error {
	foundChar, err := sc.Board.CharAt(sc.X, sc.Y)
	if err != nil {
		return nil
	}

	if characters[char] != foundChar {
		return errors.Errorf("%s not found at (%d, %d)", sc.Char, sc.X, sc.Y)
	}

	return nil
}

func (sc *ScenarioContext) APlacementErrorOfMustBeReturned(err string) error {
	if sc.Errors.PlacementError != err {
		return errors.Errorf(`expected error of "%s" but found "%s"`, err, sc.Errors.PlacementError)
	}

	return nil
}
