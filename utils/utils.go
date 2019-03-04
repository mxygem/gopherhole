package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FilledCount counts how many positions are
// filled in the current board
func FilledCount(c int, b [][]string) (int, bool) {
	var f int

	for _, r := range b {
		for _, s := range r {
			if s != " " {
				f++
			}
		}
	}

	if f <= c-2 || f >= c+2 {
		return f, false
	}

	return f, true
}

// GophersExist returns true if a gopher is in
// the current board
func GophersExist(b [][]string) bool {
	for _, r := range b {
		for _, s := range r {
			if s == "g" {
				return true
			}
		}
	}

	return false
}

// FillBoardWith is a test helper that returns a
// board full of a particular item/character
func FillBoardWith(item string, b [][]string) [][]string {
	if item == "" {
		return [][]string{}
	}

	for i, r := range b {
		for ii := range r {
			b[i][ii] = item
		}
	}

	return b
}

// CheckTestError provides a clean/reusable way
// to validate that errors are either nil or what
// is expected
func CheckTestError(t testing.TB, tcErr, err error) {
	if tcErr != nil {
		assert.EqualError(t, err, tcErr.Error())
	} else {
		assert.NoError(t, err)
	}
}
