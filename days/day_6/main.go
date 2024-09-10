package main

import (
	"fmt"
	"os"
	"strings"

	part1 "advent/days/day_6/part_1"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go <input_file>")
	}

	input := args[0]
	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	word := strings.ReplaceAll(string(file), ",", " ")
	lines := strings.Split(word, "\n")
	i := 0
	count := 0
	for i < len(lines)-1 {
		words := strings.Fields(lines[i])
		fmt.Println(part1.LightBulbs(words))
		i++
		fmt.Println()
		fmt.Printf("count:%v", count+part1.LightBulbs(words))
		fmt.Println()
	}
	fmt.Println(count)
}
