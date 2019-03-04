package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilledCount(t *testing.T) {
	testCases := []struct {
		name string
		c    int
		b    [][]string
		fc   int
		f    bool
	}{
		{
			name: "Empty",
			f:    true,
		},
		{
			name: "Count is within boundary",
			c:    4,
			b: [][]string{
				[]string{"g", " ", "o", "g"},
				[]string{"o", " ", " ", " "},
			},
			fc: 4,
			f:  true,
		},
		{
			name: "Count is under boundary",
			c:    4,
			b: [][]string{
				[]string{" ", " ", " ", " "},
				[]string{" ", " ", " ", " "},
			},
			fc: 0,
			f:  false,
		},
		{
			name: "Count is over boundary",
			c:    4,
			b: [][]string{
				[]string{"g", "g", "g", "g"},
				[]string{"g", "g", "g", "g"},
			},
			fc: 8,
			f:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			fc, f := FilledCount(tc.c, tc.b)

			assert.Equal(tt, tc.fc, fc)
			assert.Equal(tt, tc.f, f)
		})
	}
}

func TestGophersExist(t *testing.T) {
	testCases := []struct {
		name     string
		b        [][]string
		expected bool
	}{
		{
			name: "Empty",
		},
		{
			name: "Only holes",
			b:    [][]string{[]string{"o"}, []string{"o"}},
		},
		{
			name:     "Single gopher after hole",
			b:        [][]string{[]string{"o"}, []string{"g"}},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := GophersExist(tc.b)

			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func TestFillBoardWith(t *testing.T) {
	testCases := []struct {
		name     string
		item     string
		board    [][]string
		expected [][]string
	}{
		{
			name:     "Empty",
			expected: [][]string{},
		},
		{
			name: "Fill with gophers",
			item: "g",
			board: [][]string{
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
				[]string{" ", " ", " "},
			},
			expected: [][]string{
				[]string{"g", "g", "g"},
				[]string{"g", "g", "g"},
				[]string{"g", "g", "g"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			actual := FillBoardWith(tc.item, tc.board)

			assert.Equal(tt, tc.expected, actual)
		})
	}
}
