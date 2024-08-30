package part2

func Presents(s string) int {
	present := make(map[[2]int]bool)

	// two positions for Santa & Robo-Santa
	Santa := [2]int{0, 0}
	RoboSanta := [2]int{0, 0}

	present[Santa] = true

	for i, visit := range s {
		var housePosition *[2]int
		if i%2 == 0 {
			housePosition = &Santa
		} else {
			housePosition = &RoboSanta
		}

		if visit == '>' {
			housePosition[0] += 1
		} else if visit == '<' {
			housePosition[0] -= 1
		} else if visit == '^' {
			housePosition[1] += 1
		} else if visit == 'v' {
			housePosition[1] -= 1
		}
		present[*housePosition] = true
	}
	return len(present)
}
