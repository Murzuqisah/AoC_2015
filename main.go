package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	day1 "advent/day_1"
	day2 "advent/day_2"
)

func main() {
	/* ========== DAY 1 ======== */
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

	floor := day1.FloorCount(string(input))
	fmt.Println("Final floor: ", floor)

	position := day1.Position(string(input), -1)
	fmt.Println(position)

	/* ========== DAY 2 ======== */
	inputs := os.Args[1]

	values, err := os.ReadFile(inputs)
	if err != nil {
		fmt.Println("error reading file", err)
	}

	if len(values) < 1 {
		fmt.Println("File is empty")
	}

	sum := 0

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

		surfaceArea := day2.Wrapping([]int{l, w, h})
		sum += surfaceArea
	}

	fmt.Println(sum)
}
