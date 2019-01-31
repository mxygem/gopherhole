package utils

import "fmt"

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

// FullBoardOf is a test helper that returns a
// board full of a particular item/character
func FullBoardOf(item string) [][]string {
	if item == "" {
		return [][]string{}
	}

	b := [][]string{
		[]string{"", "", ""},
		[]string{"", "", ""},
		[]string{"", "", ""},
	}

	for i, r := range b {
		for _, s := range r {
			s = item
			fmt.Println(s)
		}
		fmt.Println(b)
	}

	return b
}
