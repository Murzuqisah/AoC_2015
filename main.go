package main

import (
	"fmt"
	"os"

	day1 "advent/day_1"
)

func main() {
	file := os.Args[1]
	fmt.Println(file)

	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file", err)
	}

	fmt.Println(len(input))
	if len(input) < 1 {
		fmt.Println("File is empty")
	}

	floor := day1.FloorCount(string(input))
	fmt.Println("Final floor: ", floor)
}
