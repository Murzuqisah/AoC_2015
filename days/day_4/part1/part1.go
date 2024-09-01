package part1

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func CalculateChecksum(input string) (result string) {
	var zeroCount int = 5 // Define the number of leading zeroes
	prefix := strings.Repeat("0", zeroCount)

	num := 1
	for {
		hash := input + strconv.Itoa(num)
		value := md5.Sum([]byte(hash))

		if strings.HasPrefix(hex.EncodeToString(value[:]), prefix) {
			result = strconv.Itoa(num)
			break
		}
		num++
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

/*
// Function to compute the MD5 hash of a string
func computeMD5Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// Function to check if a hash starts with a given number of zeroes
func startsWithZeroes(hash string, zeroCount int) bool {
	prefix := strings.Repeat("0", zeroCount)
	return strings.HasPrefix(hash, prefix)
}

// Function to find the lowest number that results in an MD5 hash with at least five leading zeroes
func findLowestNumber(secretKey string, zeroCount int) int {
	num := 1
	for {
		// Create the string to hash
		data := secretKey + strconv.Itoa(num)

		// Compute the MD5 hash
		hash := computeMD5Hash(data)

		// Check if hash starts with the required number of zeroes
		if startsWithZeroes(hash, zeroCount) {
			return num
		}

		// Increment the number
		num++
	}
}

func main() {
	// Example secret keys
	secretKey1 := "abcdef"
	secretKey2 := "pqrstuv"

	// Define the number of leading zeroes required
	zeroCount := 5

	// Find and print results
	fmt.Printf("Lowest number for '%s' with %d leading zeroes: %d\n", secretKey1, zeroCount, findLowestNumber(secretKey1, zeroCount))
	fmt.Printf("Lowest number for '%s' with %d leading zeroes: %d\n", secretKey2, zeroCount, findLowestNumber(secretKey2, zeroCount))
}


*/
