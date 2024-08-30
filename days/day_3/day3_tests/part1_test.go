package day3tests_test

import (
	"testing"

	"advent/days/day_3/part1"
)

type Cases struct {
	name       string
	directions string
	expected   int
}

var Case = []Cases{
	{
		name:       "Case1",
		directions: ">",
		expected:   2,
	},
	{
		name:       "Case2",
		directions: "^>v<",
		expected:   4,
	},
	{
		name:       "Case3",
		directions: "^v^v^v^v^v",
		expected:   2,
	},
}

func TestHouses(t *testing.T) {
	for _, tt := range Case {
		t.Run(tt.name, func(t *testing.T) {
			actual := part1.Houses(tt.directions)
			if actual != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, actual)
			}
		})
	}
}
