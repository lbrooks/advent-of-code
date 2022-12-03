package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/lbrooks/advent-of-code/utils"
)

func calculate(input []string, topN int) (max, total int) {
	maxCalories := make([]int, 0)
	workingCalories := 0
	for _, line := range input {
		if len(line) == 0 {
			maxCalories = append(maxCalories, workingCalories)
			sort.Slice(maxCalories, func(i, j int) bool {
				return maxCalories[i] > maxCalories[j]
			})
			if len(maxCalories) > topN {
				maxCalories = maxCalories[0:topN]
			}
			workingCalories = 0
			continue
		}
		cals, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("ERROR: " + line)
			break
		}
		workingCalories += cals
	}

	maxCalories = append(maxCalories, workingCalories)
	sort.Slice(maxCalories, func(i, j int) bool {
		return maxCalories[i] > maxCalories[j]
	})
	if len(maxCalories) > topN {
		maxCalories = maxCalories[0:topN]
	}

	for _, v := range maxCalories {
		if v > max {
			max = v
		}
		total += v
	}
	return
}

func main() {
	input := utils.ReadPiped()

	m1, t1 := calculate(input, 1)
	fmt.Printf("Part 1: Max: %d; Total: %d\n", m1, t1)

	m2, t2 := calculate(input, 3)
	fmt.Printf("Part 2: Max: %d; Total: %d\n", m2, t2)
}
