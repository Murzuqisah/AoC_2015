package main

import (
	"fmt"
	"os"
	"strings"

	part1 "advent/days/day_6/part_1"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	input := os.Args[1]
	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Split the file content into lines and clean up the data
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	// Call LightBulbs function with all instructions
	count := part1.LightBulbs(lines)

	fmt.Printf("Lit bulbs count: %d\n", count)
}
