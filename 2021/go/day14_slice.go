package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

func round(entry []string, mappings map[string]map[string]string) []string {
	if len(entry) < 2 {
		return entry
	}
	result := make([]string, 0)
	var c1, c2 string
	for i, j := 0, 1; j < len(entry); i, j = i+1, j+1 {
		c1, c2 = entry[i], entry[j]

		result = append(result, c1)
		if _, has := mappings[c1]; !has {
			continue
		}
		if _, has := mappings[c1][c2]; !has {
			continue
		}
		result = append(result, mappings[c1][c2])
	}
	result = append(result, c2)

	return result
}

func getInput(input []string) []string {
	return strings.Split(input[0], "")
}

func generateMappings(input []string) map[string]map[string]string {
	mappings := make(map[string]map[string]string)
	for i := 2; i < len(input); i++ {
		line := strings.Split(input[i], "")
		if _, has := mappings[line[0]]; !has {
			mappings[line[0]] = make(map[string]string)
		}
		mappings[line[0]][line[1]] = line[6]
	}
	return mappings
}

func play(input []string, iterations int) {
	value := getInput(input)
	mappings := generateMappings(input)

	for i := 1; i <= iterations; i++ {
		value = round(value, mappings)
	}

	counts := make(map[string]int)
	for _, v := range value {
		counts[v]++
	}

	min, max := -1, -1
	for _, num := range counts {
		if min < 0 || num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	log.Printf("Max: %d\tMin: %d\tDiff: %d\n", max, min, max-min)
}

func main() {
	buffer := 1
	var err error
	if len(os.Args) > 1 {
		if buffer, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(("Could not convert arg to number: " + os.Args[1]))
		}
	}

	input := utils.ReadPiped()

	switch buffer {
	case 1:
		play(input, 10)
	case 2:
		play(input, 40)
	}
}
