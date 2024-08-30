package day1tests_test

import (
	"testing"

	"advent/days/day_1/part2"
)

type Tests struct {
	name     string
	Input    string
	Position int
	expected int
}

var Test = []Tests{
	{
		name:     "Case1",
		Input:    "(())",
		Position: 0,
		expected: 4,
	},
	{
		name:     "Case2",
		Input:    ")))((()))",
		Position: -1,
		expected: 1,
	},
}

func TestPosition(t *testing.T) {
	for _, tt := range Test {
		t.Run(tt.name, func(t *testing.T) {
			actual := part2.Position(tt.Input, tt.Position)
			if actual != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}
