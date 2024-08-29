package day2tests_test

import (
	"testing"

	"advent/days/day_2/part1"
)

type Tests struct {
	name     string
	value    []int
	expected int
}

var Test = []Tests{
	{
		name:     "Case1",
		value:    []int{2, 3, 4},
		expected: 58,
	},
	{
		name:     "Case2",
		value:    []int{1, 1, 10},
		expected: 43,
	},
}

func TestWrapping(t *testing.T) {
	for _, tc := range Test {
		t.Run(tc.name, func(t *testing.T) {
			actual := part1.Wrapping(tc.value)
			if actual != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, actual)
			}
		})
	}
}
