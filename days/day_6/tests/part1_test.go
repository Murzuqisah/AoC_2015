package tests_test

import (
	"testing"

	part1 "advent/days/day_6/part_1"
)

type Tests struct {
	name       string
	coordinate []string
	expected   int
}

var tests = []Tests{
	{
		name:       "Single light bulb turn on",
		coordinate: []string{"turn on 500,500 through 500,500"},
		expected:   1,
	},
	{
		name:       "Single light bulb turn off",
		coordinate: []string{"turn off 500,500 through 500,500"},
		expected:   0,
	},
	{
		name:       "Single light bulb toggle",
		coordinate: []string{"toggle 500,500 through 500,500"},
		expected:   2,
	},
	{
		name:       "Multiple light bulbs turn on",
		coordinate: []string{"turn on 0,0 through 999,999"},
		expected:   1000000,
	},
}

func TestLightBulbs(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := [1000][1000]int{}
			for _, coordinate := range tt.coordinate {
				steps := "turn on"
				startX, startY := 500, 500
				endX, endY := 500, 500
				part1.LightBulbs(&grid, steps, startX, startY, endX, endY)
			}
			actual := 0
			for i := 0; i < 1000; i++ {
				for j := 0; j < 1000; j++ {
					actual += grid[i][j]
				}
			}
			if actual != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, actual)
			}
		})
	}
}

/*
package part1_test

import "testing"

func TestLightBulbs_OverlappingRanges(t *testing.T) {
	grid := [maxCoordinates][maxCoordinates]int{}
	steps := []string{"turn on", "turn off", "toggle"}
	coordinates := [][]int{{0, 0}, {2, 2}}

	for _, step := range steps {
		LightBulbs(&grid, step, coordinates[0][0], coordinates[0][1], coordinates[1][0], coordinates[1][1])
	}

	expected := 9
	actual := 0
	for i := 0; i < maxCoordinates; i++ {
		for j := 0; j < maxCoordinates; j++ {
			actual += grid[i][j]
		}
	}

	if actual != expected {
		t.Errorf("Expected %d lit bulbs, but got %d", expected, actual)
	}
}

func TestLightBulbs_SingleLightbulbTurnOff(t *testing.T) {
	var grid [maxCoordinates][maxCoordinates]int
	startX, startY := 500, 500
	endX, endY := 500, 500
	steps := "turn off"

	expected := 0
	part1.LightBulbs(&grid, steps, startX, startY, endX, endY)
	if grid[startX][startY] != expected {
		t.Errorf("LightBulbs() for single lightbulb with 'turn off' command = %v, want %v", grid[startX][startY], expected)
	}
}

func TestLightBulbs_SingleLightbulbTurnOn(t *testing.T) {
	grid := [maxCoordinates][maxCoordinates]int{}
	steps := "turn on"
	startX, startY := 500, 500
	endX, endY := 500, 500

	expected := 1
	part1.LightBulbs(&grid, steps, startX, startY, endX, endY)
	count := 0
	for i := 0; i < maxCoordinates; i++ {
		for j := 0; j < maxCoordinates; j++ {
			count += grid[i][j]
		}
	}
	if count != expected {
		t.Errorf("Expected %d lit bulbs, got %d", expected, count)
	}
}

func TestLightBulbs_LargeRangeTurnOn(t *testing.T) {
	grid := [maxCoordinates][maxCoordinates]int{}
	steps := "turn on"
	startX, startY := 0, 0
	endX, endY := maxCoordinates-1, maxCoordinates-1

	expected := maxCoordinates * maxCoordinates
	part1.LightBulbs(&grid, steps, startX, startY, endX, endY)
	count := 0
	for i := 0; i < maxCoordinates; i++ {
		for j := 0; j < maxCoordinates; j++ {
			count += grid[i][j]
		}
	}
	if count != expected {
		t.Errorf("Expected %d lit bulbs, got %d", expected, count)
	}
}

func TestLightBulbs_LargeRangeTurnOff(t *testing.T) {
	grid := [maxCoordinates][maxCoordinates]int{}
	steps := "turn off"
	startX, startY := 0, 0
	endX, endY := maxCoordinates-1, maxCoordinates-1

	expected := 0
	part1.LightBulbs(&grid, steps, startX, startY, endX, endY)
	count := 0
	for i := 0; i < maxCoordinates; i++ {
		for j := 0; j < maxCoordinates; j++ {
			count += grid[i][j]
		}
	}
	if count != expected {
		t.Errorf("Expected %d lit bulbs, got %d", expected, count)
	}
}

func TestLightBulbs_LargeRangeToggle(t *testing.T) {
	grid := [maxCoordinates][maxCoordinates]int{}
	steps := "toggle"
	startX, startY := 0, 0
	endX, endY := maxCoordinates-1, maxCoordinates-1

	expected := maxCoordinates * maxCoordinates / 2
	part1.LightBulbs(&grid, steps, startX, startY, endX, endY)
	count := 0
	for i := 0; i < maxCoordinates; i++ {
		for j := 0; j < maxCoordinates; j++ {
			count += grid[i][j]
		}
	}
	if count != expected {
		t.Errorf("Expected %d lit bulbs, got %d", expected, count)
	}
}

*/
