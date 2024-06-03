package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

func playRound(fish map[int]int) map[int]int {
	nextRound := make(map[int]int)
	for days, numFish := range fish {
		if days == 0 {
			nextRound[8] = nextRound[8] + numFish
			nextRound[6] = nextRound[6] + numFish
		} else {
			nextRound[days-1] = nextRound[days-1] + numFish
		}
	}
	return nextRound
}

func spawn(input string, forRounds int) int {
	fish := make(map[int]int, 0)

	for _, v := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(v)
		fish[num] = fish[num] + 1
	}

	for i := 0; i < forRounds; i++ {
		fish = playRound(fish)
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}

	return sum
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", spawn(input[0], 80))
	fmt.Printf("Part 2: %d\n", spawn(input[0], 256))
}
