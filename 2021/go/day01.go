package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func countIncreaseOccurancesBuffer(scanner *bufio.Scanner, numToCount int) {
	buffer := make([]int, numToCount)
	idxToReplace := 0
	var numIncreases, rowsVisited, prev int

	for scanner.Scan() {
		line := scanner.Text()
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

	switch buffer {
	case 1:
		countIncreaseOccurancesBuffer(bufio.NewScanner(os.Stdin), 1)
	case 2:
		countIncreaseOccurancesBuffer(bufio.NewScanner(os.Stdin), 3)
	}
}
