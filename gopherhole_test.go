package main

import (
	"github.com/DATA-DOG/godog"
	steps "github.com/jaysonesmith/gopherhole/step_definitions"
)

func FeatureContext(s *godog.Suite) {
	var sc steps.ScenarioContext

	sc.Steps(s)
}
