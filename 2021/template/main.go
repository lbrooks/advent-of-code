package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func playOne(input []string) {
	fmt.Println("Not Yet Implemented")
}

func playTwo(intput []string) {
	fmt.Println("Not Yet Implemented")
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
		positions = append(positions, scanner.Text())
	}

	switch buffer {
	case 1:
		playOne(positions)
	case 2:
		playTwo(positions)
	}
}
