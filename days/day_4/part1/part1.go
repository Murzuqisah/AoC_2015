package part1

import (
	"crypto/md5"
	"strconv"
	"strings"
)

func CalculateChecksum(input string) (result string) {
	value := md5.Sum([]byte(input))

	var hash, newHash string
	var zeroCount int

	prefix := strings.Repeat("0", zeroCount)
	if !strings.HasPrefix(string(value), prefix) {
		num := 1
		for {
			hash = input + strconv.Itoa(num)

			newHash = string(md5.Sum([]byte(hash)))
			if strings.HasPrefix(newHash, zeroCount) {
				result = strconv.Itoa(num)
			}
			num++
		}
	}
	return
}

// func CalculateChecksum(input string) string {
// 	// Split the input string into individual characters
// 	characters := strings.Split(input, "")

// 	// Count occurrences of each character
// 	counts := make(map[string]int)
// 	for _, char := range characters {
// 		counts[char]++
// 	}

// 	// Sort the characters by their count in descending order
// 	sortedCharacters := make([]string, 0, len(counts))
// 	for char, count := range counts {
// 		sortedCharacters = append(sortedCharacters, char+strconv.Itoa(count))
// 	}
// 	sort.Slice(sortedCharacters, func(i, j int) bool {
// 		firstCount, _ := strconv.Atoi(strings.Split(sortedCharacters[i], "")[1])
// 		secondCount, _ := strconv.Atoi(strings.Split(sortedCharacters[j], "")[1])
// 		return firstCount > secondCount
// 	})

// 	// Create the checksum by concatenating the first two characters
// 	checksum := sortedCharacters[0] + sortedCharacters[1]

// 	// Calculate the md5 checksum
// 	hasher := md5.New()
// 	hasher.Write([]byte(checksum))
// 	checksumMD5 := fmt.Sprintf("%x", hasher.Sum(nil))

// 	return checksumMD5
// }

// func main() {
// 	secretKey := "pqrstuv"
// 	number := 0

// 	for {
// 		input := secretKey + strconv.Itoa(number)
// 		hasher := md5.New()
// 		hasher.Write([]byte(input))
// 		hash := fmt.Sprintf("%x", hasher.Sum(nil))

// 		if hash[:5] == "00000" {
// 			fmt.Println("Lowest number:", number)
// 			break
// 		}

// 		number++
// 	}
// }
