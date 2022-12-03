package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func playRound(fish map[int]int) map[int]int {
	nextRound := make(map[int]int)
	for days, numFish := range fish {
		if days == 0 {
			nextRound[8] = nextRound[8] + numFish
			nextRound[6] = nextRound[6] + numFish
		} else {
			nextRound[days-1] = nextRound[days-1] + numFish
		}
	}
	return nextRound
}

func spawn(scanner *bufio.Scanner, forRounds int) {
	fish := make(map[int]int, 0)

	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(v)
		fish[num] = fish[num] + 1
	}

	for i := 0; i < forRounds; i++ {
		fish = playRound(fish)
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}

	fmt.Println("Fish Count: ", sum)
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
		spawn(bufio.NewScanner(os.Stdin), 80)
	case 2:
		spawn(bufio.NewScanner(os.Stdin), 256)
	}
}
