package main

import (
	"fmt"
	"os"

	"advent/days/day_3/part1"
	"advent/days/day_3/part2"
)

func main() {
	file := os.Args[1]

	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	if len(input) < 1 {
		fmt.Println("file is empty")
	}

	visits := part1.Houses(string(input))

	presents := part2.Presents(string(input))
	fmt.Println(visits)
	fmt.Println(presents)
}
