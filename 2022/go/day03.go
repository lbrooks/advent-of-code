package main

import (
	"fmt"

	"github.com/lbrooks/advent-of-code/utils"
)

var priorities = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func part1_duplicate_items(input []string) int {
	score := 0
	for _, line := range input {
		counts := make(map[rune]uint8)
		for _, f := range line[0 : len(line)/2] {
			counts[f] = 1
		}

		for _, s := range line[len(line)/2:] {
			if _, has := counts[s]; has {
				score += priorities[s]
				break
			}
		}
	}
	return score
}

func part2_group_badges(input []string) int {
	score := 0

	for i := 0; i < len(input); {
		counts := make(map[rune]uint8)

		for _, f := range input[i] {
			counts[f] = counts[f] | 1
		}
		i++
		for _, f := range input[i] {
			counts[f] = counts[f] | 2
		}
		i++
		for _, f := range input[i] {
			counts[f] = counts[f] | 4
		}
		i++

		for l, c := range counts {
			if c&7 == 7 {
				score += priorities[l]
				break
			}
		}
	}
	return score
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", part1_duplicate_items(input))
	fmt.Printf("Part 2: %d\n", part2_group_badges(input))
}
