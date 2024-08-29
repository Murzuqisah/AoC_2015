package day2tests_test

import (
	"testing"

	"advent/days/day_2/part2"
)

type Cases struct {
	name      string
	Dimension []int
	expected  int
}

var Case = []Cases{
	{
		name:      "Case1",
		Dimension: []int{2, 3, 4},
		expected:  34,
	},
	{
		name:      "Case2",
		Dimension: []int{1, 1, 10},
		expected:  14,
	},
}

func TestLength(t *testing.T) {
	for _, tt := range Case {
		t.Run(tt.name, func(t *testing.T) {
			actual := part2.Length(tt.Dimension)
			if actual != tt.expected {
				t.Errorf("expected %d, but got %d", tt.expected, actual)
			}
		})
	}
}
