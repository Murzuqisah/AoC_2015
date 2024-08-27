package day1

// func Position(input string, position int) int {
// 	floorNum := map[byte]int{
// 		'(': 1,
// 		')': -1,
// 	}

// 	floor := 0
// 	count := 0
// 	for i, num := range floorNum {
// 		if level, step := floorNum[byte(num)]; step {
// 			floor += level
// 			count++
// 			if floor == position {
// 				position = int(i)
// 			}

// 		}
// 	}
// 	return count
// }

func Position(input string, target int) int {
	floor := 0

	for i, level := range input {
		if level == '(' {
			floor++
		} else if level == ')' {
			floor--
		}

		if floor == target {
			return i + 1
		}
	}
	return 0
}
