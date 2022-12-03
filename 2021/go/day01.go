package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lbrooks/advent-of-code/utils"
)

func countIncreaseOccurancesBuffer(input []string, numToCount int) {
	buffer := make([]int, numToCount)
	idxToReplace := 0
	var numIncreases, rowsVisited, prev int

	for _, line := range input {
		rowsVisited++

		current, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		if rowsVisited > numToCount {
			if (prev + current - buffer[idxToReplace]) > prev {
				numIncreases++
			}
		}

		prev -= buffer[idxToReplace]
		prev += current
		buffer[idxToReplace] = current
		idxToReplace = (idxToReplace + 1) % numToCount
	}

	fmt.Printf("Total Increase Count: %d\n", numIncreases)
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
		countIncreaseOccurancesBuffer(input, 1)
	case 2:
		countIncreaseOccurancesBuffer(input, 3)
	}
}
