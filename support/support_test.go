package support

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jaysonesmith/gopherhole/board"
)

func TestCheckBoardDimensions(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		y    int
		b    board.Board
		err  error
	}{
		{
			name: "Correct dimensions",
			x:    2,
			y:    2,
			b:    [][]string{[]string{"", ""}, []string{"", ""}},
			err:  nil,
		},
		{
			name: "X dimension is 0",
			x:    2,
			y:    2,
			b:    [][]string{},
			err:  fmt.Errorf("X dimension is incorrect. expected 2 found 0, Y dimension not checked as X is 0"),
		},
		{
			name: "Incorrect X dimension",
			x:    2,
			y:    2,
			b:    [][]string{[]string{"", ""}},
			err:  fmt.Errorf("X dimension is incorrect. expected 2 found 1"),
		},
		{
			name: "Correct X but incorrect y dimension",
			x:    2,
			y:    2,
			b:    [][]string{[]string{"", "", ""}, []string{"", "", ""}},
			err:  fmt.Errorf("Y dimension is incorrect. expected 2 found 3"),
		},
		{
			name: "Both dimensions incorrect",
			x:    2,
			y:    2,
			b:    [][]string{[]string{"", "", ""}, []string{"", "", ""}, []string{"", "", ""}},
			err:  fmt.Errorf("X dimension is incorrect. expected 2 found 3, Y dimension is incorrect. expected 2 found 3"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			err := CheckBoardDimensions(tc.x, tc.y, tc.b)

			assert.Equal(tt, tc.err, err)
		})
	}
}
