package day3tests_test

import (
	"testing"

	"advent/days/day_3/part2"
)

type Tests struct {
	Name       string
	directions string
	expected   int
}

var Test = []Tests{
	{
		Name:       "Case1",
		directions: "^v",
		expected:   3,
	},
	{
		Name:       "Case2",
		directions: "^v^v^v^v^v",
		expected:   11,
	},
	{
		Name:       "Case",
		directions: ">> >> <^ ^^ vv v>",
		expected:   6,
	},
}

func TestPresents(t *testing.T) {
	for _, tc := range Test {
		t.Run(tc.Name, func(t *testing.T) {
			actual := part2.Presents(tc.directions)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
