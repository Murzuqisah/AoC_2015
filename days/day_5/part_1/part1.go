package part_1

import (
	"strings"
)

func NaugtyAndNice(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if HasVowels(s) && HasConsecutives(s) && !BadPair(s) {
			return true
		} else if BadPair(s) {
			return false
		}
		count++
	}
	return false
}

func HasVowels(s string) bool {
	vowels := "aeiou"

	count := 0
	vowelCount := 0

	for _, char := range s {
		count++
		// checking the vowel count
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		}
		if vowelCount >= 3 {
			return true
		}
	}
	return false
}

func HasConsecutives(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func BadPair(s string) bool {
	badPairs := []string{"ab", "cd", "pq", "xy"}

	// since you are checking for pairs, the iteration should be i < len(s)-1
	for i := 0; i < len(s)-1; i++ {
		for _, pair := range badPairs {
			if s[i:i+2] == pair {
				return true
			}
		}
	}
	return false
}
