package day4tests_test

import (
	"testing"

	"advent/days/day_4/part1"
)

type Tests struct {
	name     string
	input    string
	value    string
	hash     string
	newHash  string
	expected string
}

var tests = []Tests{
	{
		name:     "case 1",
		input:    "abcdef",
		value:    "abcdef609043",
		hash:     "e80b5017098950fc58aad83c8c14978e  -",
		newHash:  "000001dbbfa3a5c83a2d506429c7b00e  -",
		expected: "609043",
	},
	{
		name:     "case 2",
		input:    "pqrstuv",
		value:    "pqrstuv1048970",
		hash:     "d6ce49fc1a64ec3bb8e335b9e7a3e080  -",
		newHash:  "000006136ef2ff3b291c85725f17325c  -",
		expected: "1048970",
	},
}

func TestCalculateChecksum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := part1.CalculateChecksum(tt.input)
			if actual != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, actual)
			}
		})
	}
}
