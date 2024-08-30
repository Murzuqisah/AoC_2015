package day1tests_test

import (
	"testing"

	"advent/days/day_1/part1"
)

type Cases struct {
	argument string
	input    string
	expected int
}

var Case = []Cases{
	{
		argument: "case1",
		input:    "(())",
		expected: 0,
	},
	{
		argument: "case2",
		input:    "()()",
		expected: 0,
	},
	{
		argument: "case3",
		input:    "(((",
		expected: 3,
	},
	{
		argument: "case4",
		input:    "(()(()(",
		expected: 3,
	},
	{
		argument: "case5",
		input:    "))(((((",
		expected: 3,
	},
	{
		argument: "case6",
		input:    "())",
		expected: -1,
	},
	{
		argument: "case7",
		input:    "))(",
		expected: -1,
	},
	{
		argument: "case8",
		input:    ")))",
		expected: -3,
	},
	{
		argument: "case9",
		input:    ")())())",
		expected: -3,
	},
}

func TestFloorCount(t *testing.T) {
	for _, tc := range Case {
		t.Run(tc.argument, func(t *testing.T) {
			actual := part1.FloorCount(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %d but got %d", tc.expected, actual)
			}
		})
	}
}
