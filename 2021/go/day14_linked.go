package main

import (
	"fmt"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type letter struct {
	val  string
	next *letter
}

func round(head *letter, mappings map[string]map[string]string) *letter {
	if head == nil || head.next == nil {
		return head
	}

	prev, current := head, head.next
	for current != nil {
		if _, has := mappings[prev.val]; has {
			if _, has := mappings[prev.val][current.val]; has {
				prev.next = &letter{val: mappings[prev.val][current.val], next: current}
			}
		}

		prev = current
		current = current.next
	}

	return head
}

func getInput(input []string) *letter {
	var head *letter
	var working *letter
	for _, v := range input[0] {
		next := &letter{val: string(v)}

		if working == nil {
			head = next
		} else {
			working.next = next
		}

		working = next
	}
	return head
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

	// log.Printf("Template:\t%s\n", strings.Join(value, ""))
	for i := 1; i <= iterations; i++ {
		value = round(value, mappings)
		// log.Printf("After Step %d:\t%s\n", i, strings.Join(value, ""))
	}

	counts := make(map[string]int)
	for w := value; w != nil; w = w.next {
		counts[w.val]++
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

	fmt.Printf("Max: %d\tMin: %d\tDiff: %d\n", max, min, max-min)
}

func main() {
	input := utils.ReadPiped()

	play(input, 10)
	play(input, 40)
}
