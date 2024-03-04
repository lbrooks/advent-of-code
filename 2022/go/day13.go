package main

import (
	"encoding/json"
	"fmt"

	"github.com/lbrooks/advent-of-code/utils"
)

type Pair struct {
	index int
	left  []interface{}
	right []interface{}
}

func (p Pair) String() string {
	return fmt.Sprintf("I(%d) - L(%v) v R(%v)", p.index, p.left, p.right)
}

type Problem struct {
	pairs []Pair
}

func (p Problem) String() string {
	return fmt.Sprintf("Pairs: %v", p.pairs)
}

func ParseInput(input []string) (*Problem, error) {
	pairs := make([]Pair, 0)

	for i, index := 0, 1; i <= len(input); i, index = i+3, index+1 {
		var left, right []interface{}
		err := json.Unmarshal([]byte(input[i]), &left)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(input[i+1]), &right)
		if err != nil {
			return nil, err
		}

		pairs = append(pairs, Pair{index, left, right})
	}

	return &Problem{pairs}, nil
}

func Part1(prob *Problem) int {
	fmt.Println(prob)
	return 0
}

func Part2(prob *Problem) int {
	return 0
}

func main() {
	input := utils.ReadPiped()

	p1, err := ParseInput(input)
	if err != nil {
		return
	}
	p2, err := ParseInput(input)
	if err != nil {
		return
	}

	fmt.Printf("Part 1: %d\n", Part1(p1))
	fmt.Printf("Part 2: %d\n", Part2(p2))
}
