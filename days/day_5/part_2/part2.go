package part_2

func IsNice(s string) bool {
	return RepeatingPairs(s) && HasLetterBetween(s)
}

func RepeatingPairs(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		for j := i + 2; j <= len(s)-2; j++ {
			if s[i:i+2] == s[j:j+2] {
				return true
			}
		}
	}
	return false
}

func HasLetterBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

// func CountSubStrings(s, substr string) int {
// 	count := 0
// 	for i := 0; i < len(s); i++ {
// 		if len(s[i:]) < len(substr) {
// 			break
// 		}
// 		if s[i:i+len(substr)] == substr {
// 			count++
// 		}
// 	}
// 	return count
// }
