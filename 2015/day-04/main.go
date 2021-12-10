package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func playOne(input string) {
	firstOccurance := -1
	for i := 1; i < 999999; i++ {
		in := input + strconv.Itoa(i)
		inBytes := []byte(in)
		hash := fmt.Sprintf("%x", md5.Sum(inBytes))
		if strings.HasPrefix(hash, "00000") {
			firstOccurance = i
			break
		}
	}
	fmt.Printf("At Number: %d\n", firstOccurance)
}

func playTwo(input string) {
	firstOccurance := -1
	for i := 1; i < 99999999; i++ {
		in := input + strconv.Itoa(i)
		inBytes := []byte(in)
		hash := fmt.Sprintf("%x", md5.Sum(inBytes))
		if strings.HasPrefix(hash, "000000") {
			firstOccurance = i
			break
		}
	}
	fmt.Printf("At Number: %d\n", firstOccurance)
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
