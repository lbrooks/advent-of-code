package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

func envelops(line string) bool {
	ranges := strings.Split(line, ",")
	first := strings.Split(ranges[0], "-")
	second := strings.Split(ranges[1], "-")

	fs, _ := strconv.Atoi(first[0])
	fe, _ := strconv.Atoi(first[1])

	ss, _ := strconv.Atoi(second[0])
	se, _ := strconv.Atoi(second[1])

	if fs <= ss && fe >= se {
		return true
	}
	if ss <= fs && se >= fe {
		return true
	}
	return false
}

func overlap(line string) bool {
	ranges := strings.Split(line, ",")
	first := strings.Split(ranges[0], "-")
	second := strings.Split(ranges[1], "-")

	fs, _ := strconv.Atoi(first[0])
	fe, _ := strconv.Atoi(first[1])

	ss, _ := strconv.Atoi(second[0])
	se, _ := strconv.Atoi(second[1])

	if fe < ss {
		return false
	}
	if fs > se {
		return false
	}
	return true
}

func part1(input []string) int {
	count := 0
	for _, l := range input {
		if envelops(l) {
			count++
		}
	}
	return count
}

func part2(input []string) int {
	count := 0
	for _, l := range input {
		if overlap(l) {
			count++
		}
	}
	return count
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}
