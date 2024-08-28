package day2

func Wrapping(dimensions []int) int {
	if len(dimensions) != 3 {
		return 0
	}

	l := dimensions[0]
	w := dimensions[1]
	h := dimensions[2]

	// length, width, height := value[0], value[1], value[2]
	surfaceArea := 2*(l*w+w*h+h*l) + (l * w)

	return surfaceArea
}
