package day4tests_test

import (
	"testing"

	"advent/days/day_4/part2"
)

type Cases struct {
	name     string
	input    string
	value    string
	hash     string
	newHash  string
	expected string
}

var Case = []Cases{
	{
		name:     "case 1",
		input:    "abcdef",
		value:    "abcdef6742839",
		hash:     "e80b5017098950fc58aad83c8c14978e",
		newHash:  "000000072a1e4320d13deee9d934ae29",
		expected: "6742839",
	},
	{
		name:     "case 2",
		input:    "pqrstuv",
		value:    "pqrstuv5714438",
		hash:     "70717273747576d41d8cd98f00b204e9800998ecf8427e",
		newHash:  "00000094434e1914548b3a1af245fb27",
		expected: "5714438",
	},
}

func TestGetIdeal(t *testing.T) {
	for _, tc := range Case {
		t.Run(tc.name, func(t *testing.T) {
			actual := part2.GetIdeal(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
