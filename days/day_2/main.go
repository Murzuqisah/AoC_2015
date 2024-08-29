package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"advent/days/day_2/part1"
	"advent/days/day_2/part2"
)

func main() {
	inputs := os.Args[1]

	values, err := os.ReadFile(inputs)
	if err != nil {
		fmt.Println("error reading file", err)
	}

	if len(values) < 1 {
		fmt.Println("File is empty")
	}

	sum := 0
	size := 0

	scanner := bufio.NewScanner(strings.NewReader(string(values)))
	for scanner.Scan() {
		line := scanner.Text()
		dimensions := strings.Split(line, "x")

		if len(dimensions) != 3 {
			fmt.Println("Invalid format:", line)
			continue
		}

		l, err1 := strconv.Atoi(dimensions[0])
		w, err2 := strconv.Atoi(dimensions[1])
		h, err3 := strconv.Atoi(dimensions[2])

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("error parsing dimensions", line)
			continue
		}

		surfaceArea := part1.Wrapping([]int{l, w, h})
		sum += surfaceArea

		perimeter := part2.Length([]int{l, w, h})
		size += perimeter

	}

	fmt.Println(sum)
	fmt.Println(size)
}
