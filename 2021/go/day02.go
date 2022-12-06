package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

func partOne(input []string) {
	var depth, horizontal int

	for _, line := range input {
		instruction := strings.Split(line, " ")

		distance, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		switch instruction[0] {
		case "up":
			depth -= distance
		case "down":
			depth += distance
		case "forward":
			horizontal += distance
		}
	}

	fmt.Printf("Part 1: %d\n", (depth * horizontal))
}

func partTwo(input []string) {
	var depth, horizontal, aim int

	for _, line := range input {
		instruction := strings.Split(line, " ")

		distance, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		switch instruction[0] {
		case "up":
			aim -= distance
		case "down":
			aim += distance
		case "forward":
			horizontal += distance
			depth += (aim * distance)
		}
	}

	fmt.Printf("Part 2: %d\n", (depth * horizontal))
}

func main() {
	input := utils.ReadPiped()

	partOne(input)
	partTwo(input)
}
