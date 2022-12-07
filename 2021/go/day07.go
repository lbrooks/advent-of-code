package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

func playOne(positions []int) int {
	counts := make(map[int]int)
	for _, v := range positions {
		counts[v] = counts[v] + 1
	}

	// atPos := 0
	minFuel := -1

	for endPos := range counts {
		var fuel int
		for pos, occ := range counts {
			if endPos > pos {
				fuel += (endPos - pos) * occ
			} else {
				fuel += (pos - endPos) * occ
			}
		}
		if minFuel < 0 || fuel < minFuel {
			minFuel = fuel
			// atPos = endPos
		}
	}

	return minFuel
}

var sumCache map[int]int

func getFuelSum(distance int) int {
	d := distance
	if distance < 0 {
		d *= -1
	}

	if used, has := sumCache[d]; has {
		return used
	}
	if d == 0 {
		return 0
	}
	sumCache[d] = d + getFuelSum(d-1)
	return sumCache[d]
}

func playTwo(positions []int) int {
	maxVal := 0
	counts := make(map[int]int)
	for _, v := range positions {
		counts[v] = counts[v] + 1
		if v > maxVal {
			maxVal = v
		}
	}

	// atPos := 0
	minFuel := -1

	for endPos := 0; endPos <= maxVal; endPos++ {
		var fuel int
		for pos, occ := range counts {
			fuel += getFuelSum(endPos-pos) * occ
		}
		if minFuel < 0 || fuel < minFuel {
			// atPos = endPos
			minFuel = fuel
		}
	}

	return minFuel
}

func main() {
	sumCache = make(map[int]int)

	positions := make([]int, 0)
	for _, v := range strings.Split(utils.ReadPiped()[0], ",") {
		num, _ := strconv.Atoi(v)
		positions = append(positions, num)
	}

	fmt.Printf("Part 1: %d\n", playOne(positions))
	fmt.Printf("Part 2: %d\n", playTwo(positions))
}
