package board

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_Dimensions(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		y    int
		xl   int
		yl   int
	}{
		{name: "0 defaults to 4x4", x: 0, y: 0, xl: 4, yl: 4},
		{name: "1 defaults to 4x4", x: 1, y: 1, xl: 4, yl: 4},
		{name: "2 defaults to 4x4", x: 2, y: 2, xl: 4, yl: 4},
		{name: "3 defaults to 4x4", x: 3, y: 3, xl: 4, yl: 4},
		{name: "4 returns 4x4", x: 4, y: 4, xl: 4, yl: 4},
		{name: "10 returns 10x10", x: 10, y: 10, xl: 10, yl: 10},
		{name: "1000 returns 10x10", x: 1000, y: 1000, xl: 1000, yl: 1000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := New(tc.x, tc.y)

			assert.Equal(tt, tc.xl, len(actual))
			assert.Equal(tt, tc.yl, len(actual[0]))
		})
	}
}

func TestDifficulty(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		y        int
		d        int
		expected int
	}{
		{name: "10x10 Empty", x: 10, y: 10, d: 0, expected: 0},
		{name: "10x10 Easy", x: 10, y: 10, d: 1, expected: 15},
		{name: "10x10 Medium", x: 10, y: 10, d: 2, expected: 25},
		{name: "10x10 Hard", x: 10, y: 10, d: 3, expected: 35},
		{name: "4x4 Easy", x: 4, y: 4, d: 1, expected: 3},
		{name: "4x4 Medium", x: 4, y: 4, d: 2, expected: 4},
		{name: "4x4 Hard", x: 4, y: 4, d: 3, expected: 6},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {

			actual := difficulty(tc.x, tc.y, tc.d)

			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func TestFillBoard(t *testing.T) {
	testCases := []struct {
		name string
		d    int
		b    *Board
	}{
		{
			name: "Empty",
			d:    0,
			b: &Board{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
			},
		},
		{
			name: "Easy",
			d:    1,
			b: &Board{
				[]string{" ", " ", " ", "g"},
				[]string{" ", " ", "o", "o"},
				[]string{" ", " ", "g", "g"},
				[]string{" ", " ", " ", "o"},
			},
		},
		{
			name: "Medium",
			d:    2,
			b: &Board{
				[]string{"g", " ", "o", " "},
				[]string{"o", " ", "g", "g"},
				[]string{" ", "o", "g", "o"},
				[]string{" ", " ", " ", " "},
			},
		},
		{
			name: "Hard",
			d:    3,
			b: &Board{
				[]string{"g", " ", "g", " "},
				[]string{"o", " ", "o", " "},
				[]string{"g", "g", "g", "o"},
				[]string{"o", "o", "o", "g"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			b := New(4, 4)
			fmt.Println("board before fill")
			printBoard(&b)

			fillBoard(tc.d, 1, &b)

			fmt.Println("board after fill")
			printBoard(&b)

			assert.Equal(tt, tc.b, &b)
		})
	}
}

func TestCanPlace_Boundaries(t *testing.T) {
	testCases := []struct {
		name string
		bd   int
		x    int
		y    int
		ok   bool
	}{
		{name: "Out of upper bounds", bd: 0, x: -1, y: 0, ok: false},
		{name: "Out of right bounds", bd: 4, x: 0, y: 4, ok: false},
		{name: "Out of lower bounds", bd: 4, x: 4, y: 0, ok: false},
		{name: "Out of left bounds", bd: 4, x: 0, y: -1, ok: false},
		{name: "Inside upper bounds", bd: 4, x: 1, y: 0, ok: true},
		{name: "Inside right bounds", bd: 4, x: 0, y: 0, ok: true},
		{name: "Inside lower bounds", bd: 4, x: 3, y: 0, ok: true},
		{name: "Inside left bounds", bd: 4, x: 0, y: 1, ok: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			b := New(tc.bd, tc.bd)

			ok := canPlace(tc.x, tc.y, b)

			assert.Equal(tt, tc.ok, ok)
		})
	}
}

func TestCanPlaceUp_SpaceOpen(t *testing.T) {
	b := &Board{
		[]string{" ", " ", "g", "o"},
	}

	testCases := []struct {
		name string
		x    int
		y    int
		ok   bool
	}{
		{
			name: "Empty", x: 0, y: 0, ok: true,
		},
		{
			name: "Not empty", x: 0, y: 2, ok: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {

			ok := spaceOpen(tc.x, tc.y, b)

			assert.Equal(tt, tc.ok, ok)
		})
	}
}

func TestGopherArea(t *testing.T) {
	setRand(1)
	testCases := []struct {
		name string
		x    int
		y    int
		gx   int
		gy   int
		b    Board
	}{
		{
			name: "No available spots",
			x:    0, y: 0, gx: -1, gy: -1,
			b: Board{
				[]string{" ", "o", "o", "o"},
			},
		},
		{
			name: "Can only place to the right",
			x:    0, y: 0, gx: 0, gy: 1,
			b: Board{
				[]string{" ", " ", "o", "o"},
			},
		},
		{
			name: "Can only place downward",
			x:    0, y: 0, gx: 1, gy: 0,
			b: Board{
				[]string{" ", "o", "o", "o"},
				[]string{" ", "o", "o", "o"},
			},
		},
		{
			name: "Can only place upward",
			x:    1, y: 1, gx: 0, gy: 1,
			b: Board{
				[]string{"o", " ", "o", "o"},
				[]string{"o", " ", "o", "o"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			x, y := gopherArea(tc.x, tc.y, tc.b)

			fmt.Println("x: ", x)
			fmt.Println("y: ", y)
			assert.Equal(tt, tc.gx, x)
			assert.Equal(tt, tc.gy, y)
		})
	}
}
