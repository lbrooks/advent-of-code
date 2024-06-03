package main

import (
	"fmt"
	"strconv"

	"github.com/lbrooks/advent-of-code/utils"
)

func toNums(input []string) []int {
	asNums := make([]int, len(input))
	for i, s := range input {
		asNums[i], _ = strconv.Atoi(s)
	}
	return asNums
}

func part1(input []string) (numIncreases int) {
	asNums := toNums(input)

	previous := asNums[0]
	for _, current := range asNums {
		if current > previous {
			numIncreases += 1
		}
		previous = current
	}

	return
}

func part2(input []string) (numIncreases int) {
	asNums := toNums(input)

	previous := asNums[0] + asNums[1] + asNums[2]
	for i := 0; i < len(asNums)-2; i++ {
		current := asNums[i] + asNums[i+1] + asNums[i+2]

		if current > previous {
			numIncreases += 1
		}

		previous = current
	}

	return
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}
