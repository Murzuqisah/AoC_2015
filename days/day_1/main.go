package main

import (
	"fmt"
	"os"

	"advent/days/day_1/part1"
	"advent/days/day_1/part2"
)

func main() {
	file := os.Args[1]
	// fmt.Println(file)

	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file", err)
	}

	fmt.Println(len(input))
	if len(input) < 1 {
		fmt.Println("File is empty")
	}

	floor := part1.FloorCount(string(input))
	fmt.Println("Final floor: ", floor)

	position := part2.Position(string(input), -1)
	fmt.Println(position)
}
