package support

import (
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			err := CheckBoardDimensions(tc.x, tc.y, tc.b)

			assert.Equal(tt, tc.err, err)
		})
	}
}
