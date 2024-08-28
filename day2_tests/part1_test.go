package day2tests_test

import (
	"testing"

	day2 "advent/day_2"
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
			actual := day2.Wrapping(tc.value)
			if actual != tc.expected {
				t.Errorf("expected %d, but got %d", tc.expected, actual)
			}
		})
	}
}
