package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne(scanner *bufio.Scanner) {
	var depth, horizontal int

	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")

		distance, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		switch instruction[0] {
		case "up":
			depth -= distance
		case "down":
			depth += distance
		case "forward":
			horizontal += distance
		}
	}

	fmt.Println("Depth:", depth, "Distance:", horizontal, "Product:", (depth * horizontal))
}

func partTwo(scanner *bufio.Scanner) {
	var depth, horizontal, aim int

	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")

		distance, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		switch instruction[0] {
		case "up":
			aim -= distance
		case "down":
			aim += distance
		case "forward":
			horizontal += distance
			depth += (aim * distance)
		}
	}

	fmt.Println("Depth:", depth, "Distance:", horizontal, "Product:", (depth * horizontal))
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
		partOne(bufio.NewScanner(os.Stdin))
	case 2:
		partTwo(bufio.NewScanner(os.Stdin))
	}
}
