package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var mappings map[string]string
var entries map[string]int

func round() {
	newRecords := make(map[string]int)
	for pair, count := range entries {
		if _, has := mappings[pair]; !has {
			newRecords[pair] += count
			continue
		}
		parts := strings.Split(pair, "")
		newRecords[parts[0]+mappings[pair]] += count
		newRecords[mappings[pair]+parts[1]] += count
	}
	entries = newRecords
}

func populateEntries(input []string) {
	entries = make(map[string]int)
	asSlice := strings.Split(input[0], "")
	for i, j := 0, 1; j < len(asSlice); i, j = i+1, j+1 {
		entries[asSlice[i]+asSlice[j]]++
	}
}

func generateMappings(input []string) {
	mappings = make(map[string]string)
	for i := 2; i < len(input); i++ {
		line := strings.Split(input[i], "")

		mappings[line[0]+line[1]] = line[6]
	}
}

func play(input []string, iterations int) {
	populateEntries(input)
	generateMappings(input)

	for i := 1; i <= iterations; i++ {
		round()
	}

	counts := make(map[string]int)
	for k, c := range entries {
		letters := strings.Split(k, "")
		counts[letters[0]] += c
		counts[letters[1]] += c
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

	log.Printf("Max: %d\tMin: %d\tDiff: %d\n", max, min, (max-min)/2)
}

func main() {
	buffer := 1
	var err error
	if len(os.Args) > 1 {
		if buffer, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(("Could not convert arg to number: " + os.Args[1]))
		}
	}

	input := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	switch buffer {
	case 1:
		play(input, 10)
	case 2:
		play(input, 40)
	}
}
