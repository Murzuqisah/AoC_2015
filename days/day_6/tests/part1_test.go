package tests_test

import (
	"testing"

	part1 "advent/days/day_6/part_1"
)

func TestLightBulbs_SingleLightbulb(t *testing.T) {
	instructions := []string{"turn on 500,500 through 500,500"}
	expected := 1
	actual := part1.LightBulbs(instructions)
	if actual != expected {
		t.Errorf("Expected %d lit bulbs, got %d", expected, actual)
	}
}

func TestLightBulbs_ZeroCoordinates(t *testing.T) {
	instructions := []string{"turn on 0,0 through 0,0"}
	expected := 1

	actual := part1.LightBulbs(instructions)
	if actual != expected {
		t.Errorf("Expected %d lit bulbs, got %d", expected, actual)
	}
}

func TestLightBulbs_MultipleInstructionsSameAction(t *testing.T) {
	instructions := []string{
		"turn on 0,0 through 999,999",
		"turn on 0,0 through 999,999",
	}

	expected := 1000000
	actual := part1.LightBulbs(instructions)

	if actual != expected {
		t.Errorf("Expected %d lit bulbs, got %d", expected, actual)
	}
}

func TestLightBulbs_MultipleInstructionsNonOverlappingRanges(t *testing.T) {
	instructions := []string{
		"turn on 0,0 through 1,1",
		"turn on 3,3 through 4,4",
	}

	expected := 8
	actual := part1.LightBulbs(instructions)

	if actual != expected {
		t.Errorf("Expected %d lit bulbs, but got %d", expected, actual)
	}
}

func TestLightBulbs_MultipleInstructionsWithOverlappingRanges(t *testing.T) {
	instructions := []string{
		"turn on 0,0 through 2,2",
		"turn off 1,1 through 3,3",
		"toggle 2,2 through 4,4",
	}

	expected := 14
	actual := part1.LightBulbs(instructions)

	if actual != expected {
		t.Errorf("Expected %d lit bulbs, but got %d", expected, actual)
	}
}

func TestLightBulbs_StartCoordinatesGreaterThanEndCoordinates(t *testing.T) {
	instructions := []string{"turn on 500,500 through 499,499"}
	expected := 0

	actual := part1.LightBulbs(instructions)
	if actual != expected {
		t.Errorf("Expected %d lit bulbs, got %d", expected, actual)
	}
}
