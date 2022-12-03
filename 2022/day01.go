package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	topN := 1
	if len(os.Args) > 1 {
		if v, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(("Could not convert arg to number: " + os.Args[1]))
		} else {
			topN = v
		}
	}

	input := utils.ReadPiped()

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

	fmt.Println("Max Calories:", maxCalories)
	ttlCals := 0
	for _, c := range maxCalories {
		ttlCals += c
	}
	fmt.Println("Total Calories:", ttlCals)
}
