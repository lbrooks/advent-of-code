package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func playOne(input string) {
	floor := 0
	for _, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	fmt.Printf("Floor: %d\n", floor)
}

func playTwo(input string) {
	floor := 0
	basementEntered := 0
	for i, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			basementEntered = i + 1
			break
		}
	}
	fmt.Printf("Into Basement At: %d\n", basementEntered)
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
		playOne(input[0])
	case 2:
		playTwo(input[0])
	}
}
