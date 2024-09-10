package part1

import (
	"fmt"
	"strconv"
)

const maxvalues = 1000

// get the number of lit bulbs in a rectangle defined by the given values
func LightBulbs(values []string) int {
	var grid [maxvalues][maxvalues]int
	var currentGrid [maxvalues][maxvalues]int

	var startX, startY, endX, endY int
	var action string
	count := 0

	if len(values) == 7 {
		action = values[0] + " " + values[1]
		startX, _ = strconv.Atoi(values[2])
		startY, _ = strconv.Atoi(values[3])
		endX, _ = strconv.Atoi(values[5])
		endY, _ = strconv.Atoi(values[6])
		fmt.Println(startX, startX, endX, endY)
	} else if len(values) == 6 {
		action = values[0]
		startX, _ = strconv.Atoi(values[1])
		startY, _ = strconv.Atoi(values[2])
		endX, _ = strconv.Atoi(values[4])
		endY, _ = strconv.Atoi(values[5])
	}

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			switch action {
			case "turn on":
				grid[x][y] = 1
			case "turn off":
				grid[x][y] = 0
			case "toggle":
				grid[x][y] ^= 1
			}
		}
	}
	currentGrid = grid
	for i := 0; i < maxvalues; i++ {
		for j := 0; j < maxvalues; j++ {
			if currentGrid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

/*
A way to store or organize related data, such as values or light bulb positions, you could use a map in Go.
To keep track of the positions of light bulbs, you could use a map of values to boolean values:
The lightBulbs map is used to store the positions of light bulbs within the rectangle.
The keys of the map are formatted strings representing the values, and the values are boolean true.
The length of the map is then returned, which represents the total number of light bulbs.


func LightBulbs(values [][]int) int {
    startX, startY := values[0][0], values[0][1]
    endX, endY := values[1][0], values[1][1]

    lightBulbs := make(map[string]bool)
    for x := startX; x <= endX; x++ {
        for y := startY; y <= endY; y++ {
            lightBulbs[fmt.Sprintf("%d,%d", x, y)] = true

        }
    }

    return len(lightBulbs)
}
*/
