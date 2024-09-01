package main

import (
	"fmt"
	"os"
	"strings"

	"advent/days/day_5/part_1"
)

func main() {
	args := os.Args[1]

	file, err := os.ReadFile(args)
	if err != nil {
		fmt.Println("error reading file:", err)
	}

	if len(file) < 1 {
		fmt.Println("file is empty")
	}

	str := string(file)
	words := strings.Fields(str)
	count := 0

	for _, word := range words {
		if part_1.NaugtyAndNice(word) {
			count++
		}
	}
	fmt.Printf("Number of nice words: %d\n", count)
}
