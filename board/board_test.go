package board

import (
	"testing"

	"github.com/jaysonesmith/gopherhole/utils"

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
		{name: "4x4 Easy", x: 4, y: 4, d: 1, expected: 3},
		{name: "4x4 Medium", x: 4, y: 4, d: 2, expected: 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {

			actual := difficulty(tc.x, tc.y, tc.d)

			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func TestFill(t *testing.T) {
	testCases := []struct {
		name string
		dim  int
		diff int
		b    *Board
	}{
		{
			name: "Empty",
			dim:  4,
			diff: 0,
			b: &Board{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
			},
		},
		{
			name: "Medium",
			dim:  4,
			diff: 1,
			b: &Board{
				[]string{" ", " ", " ", "o"},
				[]string{" ", " ", " ", "g"},
				[]string{" ", "g", " ", " "},
				[]string{" ", "o", "o", "g"},
			},
		},
		{
			name: "Hard",
			dim:  4,
			diff: 2,
			b: &Board{
				[]string{" ", "g", " ", "o"},
				[]string{" ", "o", " ", "g"},
				[]string{" ", "g", " ", " "},
				[]string{" ", "o", "o", "g"},
			},
		},
		{
			name: "Hard 5x5",
			dim:  5,
			diff: 2,
			b: &Board{
				[]string{" ", " ", "o", " ", " "},
				[]string{"g", "o", "g", " ", " "},
				[]string{" ", " ", " ", "o", "g"},
				[]string{"g", " ", " ", " ", " "},
				[]string{"o", " ", "o", "g", " "},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			setRand(1)
			b := New(tc.dim, tc.dim)

			b.Fill(tc.diff, 1)

			b.Print()
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

			ok := b.canPlace(tc.x, tc.y)

			assert.Equal(tt, tc.ok, ok)
		})
	}
}

func TestGopherArea(t *testing.T) {
	testCases := []struct {
		name string
		b    Board
		x    int
		y    int
	}{
		{
			name: "Empty Board",
			b: Board{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
			},
			x: 1, y: 3,
		},
		{
			name: "Force placement elsewhere via gopher at initial position",
			b: Board{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", "g"},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
			},
			x: 3, y: 3,
		},
		{
			name: "Force placement elsewhere via gopher next to initial position",
			b: Board{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", "g", " "},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
			},
			x: 3, y: 3,
		},
		{
			name: "Force placement elsewhere via gopher near first two positions",
			b: Board{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", "g", " "},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", "g", " "},
			},
			x: 1, y: 0,
		},
		{
			name: "Force placement elsewhere via hole at initial position",
			b: Board{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", "o"},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
			},
			x: 3, y: 3,
		},
		{
			name: "Check third position",
			b: Board{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", "g"},
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", "g"},
			},
			x: 1, y: 0,
		},
		{
			name: "Full board except one",
			b: Board{
				[]string{"g", "o", "o", "o"},
				[]string{"o", "o", " ", "o"},
				[]string{"g", "o", "o", "o"},
				[]string{"o", "g", "o", "g"},
			},
			x: 1, y: 2,
		},
		{
			name: "Force bottom row",
			b: Board{
				[]string{"o", "o", "o", "o"},
				[]string{"o", "o", "o", "o"},
				[]string{"o", "o", "o", "o"},
				[]string{"o", "o", " ", "o"},
			},
			x: 3, y: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			setRand(1)

			x, y := tc.b.gopherArea()

			assert.Equal(tt, tc.x, x)
			assert.Equal(tt, tc.y, y)
		})
	}
}

func TestHoleArea(t *testing.T) {
	setRand(1)
	testCases := []struct {
		name string
		x    int
		y    int
		hx   int
		hy   int
		b    Board
	}{
		{
			name: "No available spots",
			x:    0, y: 0, hx: -1, hy: -1,
			b: Board{
				[]string{"g", "o", "", ""},
			},
		},
		{
			name: "Can only place to the right",
			x:    0, y: 0, hx: 0, hy: 1,
			b: Board{
				[]string{"g", " ", "o", "o"},
			},
		},
		{
			name: "Can only place downward",
			x:    0, y: 0, hx: 1, hy: 0,
			b: Board{
				[]string{"g", "o", "o", "o"},
				[]string{" ", "o", "o", "o"},
			},
		},
		{
			name: "Can only place upward",
			x:    1, y: 1, hx: 0, hy: 1,
			b: Board{
				[]string{"o", " ", "o", "o"},
				[]string{"o", "g", "o", "o"},
			},
		},
		{
			name: "Can only place to the left",
			x:    3, y: 3, hx: 3, hy: 2,
			b: Board{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", "g"},
				[]string{" ", "g", "o", "o"},
				[]string{" ", " ", " ", "g"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			x, y := tc.b.holeArea(tc.x, tc.y)

			assert.Equal(tt, tc.hx, x)
			assert.Equal(tt, tc.hy, y)
		})
	}
}

func TestSurroundingGopher(t *testing.T) {
	testCases := []struct {
		name     string
		x        int
		y        int
		b        Board
		expected bool
	}{
		{
			name: "Middle of board, all empty",
			x:    1,
			y:    1,
			b: Board{
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
			},
			expected: false,
		},
		{
			name: "Middle of board, one found above",
			x:    1,
			y:    1,
			b: Board{
				[]string{" ", "g", " "},
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
			},
			expected: true,
		},
		{
			name: "Left edge of board, empty",
			x:    0,
			y:    1,
			b: Board{
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
			},
			expected: false,
		},
		{
			name: "Left edge of board, gopher below",
			x:    0,
			y:    1,
			b: Board{
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
				[]string{"g", " ", " "},
			},
			expected: false,
		},
		{
			name: "Bottom-right corner of board, empty",
			x:    2,
			y:    2,
			b: Board{
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
			},
			expected: false,
		},
		{
			name: "Top-right corner of board, two found",
			x:    0,
			y:    2,
			b: Board{
				[]string{" ", "g", " "},
				[]string{" ", " ", "g"},
				[]string{" ", " ", " "},
			},
			expected: true,
		},
		{
			name: "Right-hand side of board, one found",
			x:    1,
			y:    2,
			b: Board{
				[]string{" ", "g", " "},
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
			},
			expected: true,
		},
		{
			name: "Lower-left corner of board, one at diagonal",
			x:    2,
			y:    0,
			b: Board{
				[]string{" ", " ", " "},
				[]string{" ", "g", " "},
				[]string{" ", " ", " "},
			},
			expected: true,
		},
		{
			name: "Lower-left corner of board, one out of range",
			x:    2,
			y:    0,
			b: Board{
				[]string{" ", " ", "g"},
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.b.surroundingGopher(tc.x, tc.y)

			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func TestStart(t *testing.T) {
	b := New(5, 5)
	b.Fill(1, 0)

	b.Start()

	assert.False(t, utils.GophersExist(b))
}

func TestWriteChar(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		x        int
		y        int
		board    Board
		expected Board
		err      error
	}{
		{
			name:     "Write gopher to empty position",
			input:    "g",
			x:        0,
			y:        0,
			board:    Board{[]string{" ", " ", " "}},
			expected: Board{[]string{"g", " ", " "}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			err := tc.board.WriteChar(tc.input, tc.x, tc.y)

			assert.Equal(tt, tc.expected, tc.board)
			assert.Equal(tt, tc.err, err)
		})
	}
}
