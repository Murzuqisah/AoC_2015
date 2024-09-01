package day5_test

import (
	"testing"

	"advent/days/day_5/part_1"
)

type TestData struct {
	name     string
	input    string
	expected bool
}

var TestDatas = []TestData{
	{
		name:     "test 1",
		input:    "ugknbfddgicrmopn",
		expected: true,
	},
	{
		name:     "test 2",
		input:    "aaa",
		expected: true,
	},
	{
		name:     "test 3",
		input:    "jchzalrnumimnmhp",
		expected: false,
	},
	{
		name:     "test 4",
		input:    "haegwjzuvuyypxyu",
		expected: false,
	},
	{
		name:     "test 5",
		input:    "dvszwmarrgswjxmb",
		expected: false,
	},
}

func TestNaughtyAndNice(t *testing.T) {
	for _, tt := range TestDatas {
		t.Run(tt.name, func(t *testing.T) {
			actual := part_1.NaugtyAndNice(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
