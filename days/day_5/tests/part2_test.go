package day5_test

import (
	part2 "advent/days/day_5/part_2"
	"testing"
)

type Cases struct {
	name     string
	input    string
	expected bool
}

var Case = []Cases{
	{
		name:     "case 1",
		input:    "qjhvhtzxzqqjkmpb",
		expected: true,
	},
	{
		name:     "case 2",
		input:    "xxyxx",
		expected: true,
	},
	{
		name:     "case 3",
		input:    "uurcxstgmygtbstg",
		expected: false,
	},
	{
		name:     "case 4",
		input:    "ieodomkazucvgmuy",
		expected: false,
	},
}

func TestIsNice(t *testing.T) {
	for _, tc := range Case {
		t.Run(tc.name, func(t *testing.T) {
			actual := part2.IsNice(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
