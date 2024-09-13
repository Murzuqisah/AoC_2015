package part1

import (
	"strconv"
	"strings"
)

const maxvalues = 1000

func LightBulbs(instructions []string) int {
	var grid [maxvalues][maxvalues]int

	for _, instruction := range instructions {
		parts := strings.Fields(instruction)
		var action string
		var startX, startY, endX, endY int

		if parts[0] == "toggle" {
			action = "toggle"
			startX, _ = strconv.Atoi(strings.Split(parts[1], ",")[0])
			startY, _ = strconv.Atoi(strings.Split(parts[1], ",")[1])
			endX, _ = strconv.Atoi(strings.Split(parts[3], ",")[0])
			endY, _ = strconv.Atoi(strings.Split(parts[3], ",")[1])
		} else {
			action = parts[1]
			startX, _ = strconv.Atoi(strings.Split(parts[2], ",")[0])
			startY, _ = strconv.Atoi(strings.Split(parts[2], ",")[1])
			endX, _ = strconv.Atoi(strings.Split(parts[4], ",")[0])
			endY, _ = strconv.Atoi(strings.Split(parts[4], ",")[1])
		}

		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				switch action {
				case "on":
					grid[x][y] = 1
				case "off":
					grid[x][y] = 0
				case "toggle":
					grid[x][y] ^= 1
				}
			}
		}
	}

	count := 0
	for i := 0; i < maxvalues; i++ {
		for j := 0; j < maxvalues; j++ {
			if grid[i][j] == 1 {
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
