package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	A uint8 = 1 << iota
	B
	C
	D
	E
	F
	G
)

type codeInput struct {
	signals []string
	outputs []string
}

func mapSignal(signal string) uint8 {
	var result uint8
	for _, l := range strings.Split(signal, "") {
		switch l {
		case "A", "a":
			result = result | A
		case "B", "b":
			result = result | B
		case "C", "c":
			result = result | C
		case "D", "d":
			result = result | D
		case "E", "e":
			result = result | E
		case "F", "f":
			result = result | F
		case "G", "g":
			result = result | G
		}
	}
	return result
}

type display struct {
	wiring map[uint8]int
}

func determineNumber(signal string) []int {
	switch len(signal) {
	case 2:
		return []int{1}
	case 3:
		return []int{7}
	case 4:
		return []int{4}
	case 5:
		return []int{2, 3, 5}
	case 6:
		return []int{6, 9, 0}
	case 7:
		return []int{8}
	}
	return []int{}
}

func playOne(in []codeInput) int {
	var occ int
	for _, i := range in {
		for _, o := range i.outputs {
			switch len(o) {
			case 2, 4, 3, 7:
				occ++
			}
		}
	}
	return occ
}

func NumOfSetBits(n uint8) uint8 {
	var count uint8
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func playTwo(allInputs []codeInput) int {
	sum := 0

	for _, log := range allInputs {
		numbersToBits := make(map[int]uint8)
		knownNumbers := make(map[uint8]int)

		unknownEntries := make([]string, 0)

		for _, o := range log.signals {
			wires := mapSignal(o)
			numbers := determineNumber(o)
			if len(numbers) == 1 {
				knownNumbers[wires] = numbers[0]
				numbersToBits[numbers[0]] = wires
			} else if len(numbers) > 1 {
				unknownEntries = append(unknownEntries, o)
			}
		}

		for _, o := range unknownEntries {
			asBits := mapSignal(o)
			num := -1

			switch len(o) {
			case 5:
				if (asBits | numbersToBits[1]) == asBits {
					num = 3
				} else if asBits|numbersToBits[4] == numbersToBits[8] {
					num = 2
				} else {
					num = 5
				}
			case 6:
				if (asBits | numbersToBits[1]) != asBits {
					num = 6
				} else if (asBits | numbersToBits[4]) == asBits {
					num = 9
				} else {
					num = 0
				}
			}

			knownNumbers[asBits] = num
			numbersToBits[num] = asBits
		}

		val := 0
		for _, o := range log.outputs {
			val = val * 10
			val += knownNumbers[mapSignal(o)]
		}
		sum += val
	}

	return sum
}

func readInput() []codeInput {
	entries := make([]codeInput, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		record := strings.Split(scanner.Text(), " | ")
		entries = append(entries, codeInput{
			signals: strings.Split(record[0], " "),
			outputs: strings.Split(record[1], " "),
		})
	}
	return entries
}

func main() {
	allEntries := readInput()

	fmt.Printf("Part 1: %d\n", playOne(allEntries))
	fmt.Printf("Part 2: %d\n", playTwo(allEntries))
}
