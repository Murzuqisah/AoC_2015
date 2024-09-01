package main

import (
	"crypto/md5"
	"fmt"
	"os"

	"advent/days/day_4/part1"
	"advent/days/day_4/part2"
)

func main() {
	input := os.Args[1]
	value := fmt.Sprintf("%x", md5.New().Sum([]byte(input)))

	fmt.Println(value)
	fmt.Println()

	checksum := part1.CalculateChecksum(input)
	fmt.Println("Checksum five 0's (MD5):", checksum)

	fmt.Println()

	idealValue := part2.GetIdeal(input)
	hash := fmt.Sprintf("%x", md5.New().Sum([]byte(input+idealValue)))

	fmt.Println(hash)
	fmt.Println("Ideal six 0's (MD5):", idealValue)
}
