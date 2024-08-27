package main

import (
	"os"
	"fmt"
)

func main() {
	argument := os.Args[0:]
	if len(argument) < 2 {
		return
	}
	file := argument[0]

	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file", err)
	}

	if len(input) < 1 {
		fmt.Println("File is empty")
	}

	floor := FloorCount((input))
	fmt.Println("Final floor: ", floor)

}

func FloorCount(s []byte) int{
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

// func FloorCount(floors []byte) int {
// 	floor := map[byte]int {
// 		'(' : 1,
// 		')'	: -1,
// 	}

// 	count := 0
// 	for _, char := range floors {
// 		if key, value := floor[char]; value {
// 			count += key
// 		}
// 	}
// 	return count
// }