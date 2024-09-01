package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

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
