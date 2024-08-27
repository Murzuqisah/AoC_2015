package day1

func FloorCount(s string) int {
	/*
		( -> 1 floor up
		) -> 1 floor down
	*/
	floor := 0
	for _, char := range s {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}
	return floor
}

/*   ALTERNATIVE TWO    */

// func FloorCount(floors string) int {
// 	floor := map[byte]int{
// 		'(': 1,
// 		')': -1,
// 	}

// 	count := 0
// 	for _, char := range floors {
// 		if key, value := floor[byte(char)]; value {
// 			count += key
// 		}
// 	}
// 	return count
// }
