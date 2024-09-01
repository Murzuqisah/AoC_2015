package part2

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func GetIdeal(input string) (result string) {
	var zeroCount int = 6 // Define the number of leading zeroes
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
