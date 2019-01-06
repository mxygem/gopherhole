package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		xl   int
		yl   int
	}{
		{name: "0 defaults to 4x4", x: 0, xl: 4, yl: 4},
		// {name: "1 defaults to 4x4", x: 1, xl: 4, yl: 4},
		// {name: "2 defaults to 4x4", x: 2, xl: 4, yl: 4},
		// {name: "3 defaults to 4x4", x: 3, xl: 4, yl: 4},
		// {name: "4 returns 4x4", x: 4, xl: 4, yl: 4},
		// {name: "10 returns 10x10", x: 10, xl: 10, yl: 10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := New(tc.x)

			assert.Equal(tt, tc.xl, len(actual))
			assert.Equal(tt, tc.yl, len(actual[0]))
		})
	}
}
