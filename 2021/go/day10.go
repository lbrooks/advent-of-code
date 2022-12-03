package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/lbrooks/advent-of-code/utils"
)

type node struct {
	prev  *node
	value rune
}

func playOne(input []string) {
	sumFailPoints := 0
	for _, in := range input {
		var queue *node
		for _, b := range in {
			if b == '(' {
				queue = &node{value: ')', prev: queue}
			} else if b == '[' {
				queue = &node{value: ']', prev: queue}
			} else if b == '{' {
				queue = &node{value: '}', prev: queue}
			} else if b == '<' {
				queue = &node{value: '>', prev: queue}
			} else {
				if queue == nil || queue.value != b {
					points := 0

					switch b {
					case ')':
						points += 3
					case ']':
						points += 57
					case '}':
						points += 1197
					case '>':
						points += 25137
					}
					sumFailPoints += points

					break
				} else {
					queue = queue.prev
				}
			}
		}
	}
	fmt.Printf("Points: %d\n", sumFailPoints)
}

func playTwo(input []string) {
	points := make([]int, 0)

	for _, in := range input {
		isFailure := false
		var queue *node
		for _, b := range in {
			if b == '(' {
				queue = &node{value: ')', prev: queue}
			} else if b == '[' {
				queue = &node{value: ']', prev: queue}
			} else if b == '{' {
				queue = &node{value: '}', prev: queue}
			} else if b == '<' {
				queue = &node{value: '>', prev: queue}
			} else {
				if queue == nil || queue.value != b {
					isFailure = true
					break
				} else {
					queue = queue.prev
				}
			}
		}

		if isFailure {
			continue
		}
		if queue == nil {
			continue
		}
		score := 0
		for q := queue; q != nil; q = q.prev {
			score *= 5
			switch q.value {
			case ')':
				score += 1
			case ']':
				score += 2
			case '}':
				score += 3
			case '>':
				score += 4
			}
		}
		if score != 0 {
			points = append(points, score)
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i] < points[j]
	})

	fmt.Printf("Points: %d\n", points[len(points)/2])
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
		playOne(input)
	case 2:
		playTwo(input)
	}
}
