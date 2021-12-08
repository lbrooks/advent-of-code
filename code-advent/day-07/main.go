package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func playOne(positions []int) {
	counts := make(map[int]int)
	for _, v := range positions {
		counts[v] = counts[v] + 1
	}

	atPos := 0
	minFuel := -1

	for endPos := range counts {
		var fuel int
		for pos, occ := range counts {
			if endPos > pos {
				fuel += (endPos - pos) * occ
			} else {
				fuel += (pos - endPos) * occ
			}
		}
		if minFuel < 0 || fuel < minFuel {
			minFuel = fuel
			atPos = endPos
		}
	}

	fmt.Printf("Fuel Used: %d to come to Position: %d\n", minFuel, atPos)
}

var sumCache map[int]int

func getFuelSum(distance int) int {
	d := distance
	if distance < 0 {
		d *= -1
	}

	if used, has := sumCache[d]; has {
		return used
	}
	if d == 0 {
		return 0
	}
	sumCache[d] = d + getFuelSum(d-1)
	return sumCache[d]
}

func playTwo(positions []int) {
	maxVal := 0
	counts := make(map[int]int)
	for _, v := range positions {
		counts[v] = counts[v] + 1
		if v > maxVal {
			maxVal = v
		}
	}

	atPos := 0
	minFuel := -1

	for endPos := 0; endPos <= maxVal; endPos++ {
		var fuel int
		for pos, occ := range counts {
			fuel += getFuelSum(endPos-pos) * occ
		}
		if minFuel < 0 || fuel < minFuel {
			atPos = endPos
			minFuel = fuel
		}
	}

	fmt.Printf("Fuel Used: %d to come to Position: %d\n", minFuel, atPos)
}

func main() {
	sumCache = make(map[int]int)

	buffer := 1
	var err error
	if len(os.Args) > 1 {
		if buffer, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(("Could not convert arg to number: " + os.Args[1]))
		}
	}

	positions := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(v)
		positions = append(positions, num)
	}

	switch buffer {
	case 1:
		playOne(positions)
	case 2:
		playTwo(positions)
	}
}
