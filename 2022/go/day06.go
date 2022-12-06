package main

import (
	"fmt"

	"github.com/lbrooks/advent-of-code/utils"
)

func firstSetUniqueCharacters(input string, length int) int {
	for i := length - 1; i < len(input); i++ {
		data := make(map[byte]bool)
		for j := i - (length - 1); j <= i; j++ {
			data[input[j]] = true
		}

		if len(data) == length {
			return i + 1
		}
	}
	return -1
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", firstSetUniqueCharacters(input[0], 4))
	fmt.Printf("Part 2: %d\n", firstSetUniqueCharacters(input[0], 14))
}
