package main

import (
	"fmt"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

func partOne(records [][]string) {
	counts := make([]int, 12)
	rows := 0

	for _, value := range records {
		rows += 1

		for i, c := range value {
			if c == "1" {
				counts[i] += 1
			}
		}
	}

	gamma, epsilon := 0, 0
	for _, v := range counts {
		gamma = (gamma << 1)
		epsilon = (epsilon << 1)

		if v > (rows - v) {
			gamma += 1
		} else {
			epsilon += 1
		}
	}

	fmt.Printf("Part 1: %d\n", (gamma * epsilon))
}

func partTwo(records [][]string) {
	entryLength := len(records[0])

	values := records
	oxy := 0
	for i := 0; i < entryLength; i++ {
		leadZero, leadOne := make([][]string, 0), make([][]string, 0)
		for _, v := range values {
			if v[0] == "0" {
				leadZero = append(leadZero, v[1:])
			} else {
				leadOne = append(leadOne, v[1:])
			}
		}

		oxy = oxy << 1
		if len(leadZero) > len(leadOne) {
			values = leadZero
		} else {
			oxy += 1
			values = leadOne
		}
	}

	values = records
	co2 := 0
	for i := 0; i < entryLength; i++ {
		leadZero, leadOne := make([][]string, 0), make([][]string, 0)
		for _, v := range values {
			if v[0] == "0" {
				leadZero = append(leadZero, v[1:])
			} else {
				leadOne = append(leadOne, v[1:])
			}
		}

		co2 = co2 << 1
		if len(leadZero) == 0 {
			co2 += 1
			values = leadOne
		} else if len(leadOne) == 0 {
			values = leadZero
		} else if len(leadZero) <= len(leadOne) {
			values = leadZero
		} else {
			co2 += 1
			values = leadOne
		}
	}

	fmt.Printf("Part 2: %d\n", oxy*co2)
}

func main() {
	records := make([][]string, 0)
	input := utils.ReadPiped()
	for _, l := range input {
		records = append(records, strings.Split(l, ""))
	}

	partOne(records)
	partTwo(records)
}
