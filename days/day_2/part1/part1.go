package part1

import (
	"fmt"
	"sort"
)

func Wrapping(dimensions []int) int {
	if len(dimensions) != 3 {
		return 0
	}

	l := dimensions[0]
	w := dimensions[1]
	h := dimensions[2]

	value1, value2 := Min(dimensions)
	fmt.Println(value1, value2)

	surfaceArea := 2*(l*w+l*h+h*w) + (value1 * value2)

	return surfaceArea
}

func Min(values []int) (int, int) {
	sort.Ints(values)

	val1, val2 := values[0], values[1]

	return val1, val2
}
