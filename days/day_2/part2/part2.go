package part2

import (
	"sort"
)

// shortest distance around its sides
// or the smallest perimeter of any face
// + bow: cubic feet volume of the present -> volume of the cube/cuboid
func Length(Dimension []int) int {
	if len(Dimension) != 3 {
		return 0
	}

	var perimeter []int

	l := Dimension[0]
	w := Dimension[1]
	h := Dimension[2]

	fullLength := 0
	perim1 := 2 * (l + w)
	perim2 := 2 * (w + h)
	perim3 := 2 * (l + h)

	perimeter = []int{perim1, perim2, perim3}
	sort.Ints(perimeter)
	// fmt.Println(perimeter)

	// 2*small[0] + 2*small[1]

	bowLength := l * w * h

	fullLength = perimeter[0] + bowLength

	return fullLength
}
