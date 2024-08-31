package main

import (
	"fmt"

	"advent/days/day_4/part1"
)

func main() {
	input := "abcdef"
	checksum := part1.CalculateChecksum(input)
	fmt.Println("Checksum (MD5):", checksum)
}
