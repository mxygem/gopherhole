package main

import (
	"github.com/DATA-DOG/godog"
	steps "github.com/jaysonesmith/gopherhole/step_definitions"
)

func FeatureContext(s *godog.Suite) {
	var sc steps.ScenarioContext
	s.Step(`^a new game is requested with no board size set$`, sc.ANewGameIsRequestedWithNoBoardSizeSet)
	s.Step(`^a (\d+)x(\d+) board must be returned$`, sc.AXBoardMustBeReturned)
	s.Step(`^a new game is requested with a (\d+)x(\d+) board size$`, sc.ANewGameIsRequestedWithAXBoardSize)
}
