package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min(n ...int) int {
	if len(n) == 0 {
		return 0
	}

	sort.Slice(n, func(i, j int) bool {
		return n[i] < n[j]
	})
	return n[0]
}

func minTwo(n ...int) (int, int) {
	if len(n) < 2 {
		return 0, 0
	}

	sort.Slice(n, func(i, j int) bool {
		return n[i] < n[j]
	})
	return n[0], n[1]
}

func playOne(input []string) {
	total := 0
	for _, dim := range input {
		dimSlice := strings.Split(dim, "x")

		l, _ := strconv.Atoi(dimSlice[0])
		w, _ := strconv.Atoi(dimSlice[1])
		h, _ := strconv.Atoi(dimSlice[2])

		a1 := l * w
		a2 := l * h
		a3 := h * w

		total += (2*a1 + 2*a2 + 2*a3 + min(a1, a2, a3))
	}
	fmt.Printf("Total Wrapping Paper: %d\n", total)
}

func playTwo(input []string) {
	total := 0
	for _, dim := range input {
		dimSlice := strings.Split(dim, "x")

		l, _ := strconv.Atoi(dimSlice[0])
		w, _ := strconv.Atoi(dimSlice[1])
		h, _ := strconv.Atoi(dimSlice[2])

		m1, m2 := minTwo(l, w, h)
		total += (m1*2 + m2*2 + l*w*h)
	}
	fmt.Printf("Total Ribbon: %d\n", total)
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
		playOne(input)
	case 2:
		playTwo(input)
	}
}
