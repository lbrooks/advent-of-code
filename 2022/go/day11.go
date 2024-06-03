package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

var parseID = regexp.MustCompile(`Monkey ([0-9]+):`)
var parseItems = regexp.MustCompile(`Starting items: ([0-9 ,]+)`)
var parseOperation = regexp.MustCompile(`Operation: new = old (.+)`)
var parseTest = regexp.MustCompile(`Test: divisible by ([0-9]+)`)
var parseTrue = regexp.MustCompile(`If true: throw to monkey ([0-9]+)`)
var parseFalse = regexp.MustCompile(`If false: throw to monkey ([0-9]+)`)

type MonkeyOperation func(old int) int

type Monkey struct {
	id    int
	items []int

	operation MonkeyOperation

	divisibleBy    int
	recipientTrue  int
	recipientFalse int
	inspectedCount int
}

func (m *Monkey) String() string {
	var output strings.Builder
	output.WriteString("Monkey ")
	output.WriteString(strconv.Itoa(m.id))
	output.WriteString(": ")
	for _, v := range m.items {
		output.WriteString(strconv.Itoa(v))
		output.WriteString(", ")
	}
	return output.String()
}

func ParseMonkeyOperation(input string, reg *regexp.Regexp) MonkeyOperation {
	op := reg.FindStringSubmatch(input)[1]

	parts := strings.Split(op, " ")
	mathOp := parts[0]
	against := parts[1]

	if against == "old" {
		switch mathOp {
		case "+":
			return func(old int) int {
				return old + old
			}
		case "-":
			return func(old int) int {
				return 0
			}
		case "*":
			return func(old int) int {
				return old * old
			}
		case "/":
			return func(old int) int {
				return 1
			}
		}
	}

	val, _ := strconv.Atoi(against)
	switch mathOp {
	case "+":
		return func(old int) int {
			return old + val
		}
	case "-":
		return func(old int) int {
			return old - val
		}
	case "*":
		return func(old int) int {
			return old * val
		}
	case "/":
		return func(old int) int {
			return old / val
		}
	}

	return func(old int) int {
		return 0
	}
}

func RegexToNumber(input string, reg *regexp.Regexp) int {
	match := reg.FindStringSubmatch(input)[1]
	num, err := strconv.Atoi(match)
	if err != nil {
		fmt.Printf("Problem Parsing, input [%s], match [%s]\n", input, match)
		return 1
	}
	return num
}

func RegexToNumberSlice(input string, reg *regexp.Regexp) (nums []int) {
	match := reg.FindStringSubmatch(input)[1]
	stringItems := strings.Split(match, ", ")

	nums = make([]int, 0, len(stringItems))
	for _, si := range stringItems {
		v, _ := strconv.Atoi(si)
		nums = append(nums, v)
	}
	return
}

func ParseMonkey(input []string) (*Monkey, error) {
	if len(input) != 6 {
		return nil, fmt.Errorf("Unable to parse from input: %v", strings.Join(input, " | "))
	}

	return &Monkey{
		id:             RegexToNumber(input[0], parseID),
		items:          RegexToNumberSlice(input[1], parseItems),
		operation:      ParseMonkeyOperation(input[2], parseOperation),
		divisibleBy:    RegexToNumber(input[3], parseTest),
		recipientTrue:  RegexToNumber(input[4], parseTrue),
		recipientFalse: RegexToNumber(input[5], parseFalse),
		inspectedCount: 0,
	}, nil
}

func CalculateScore(monkeys []*Monkey) int {
	counts := make([]int, 0, len(monkeys))
	for _, m := range monkeys {
		counts = append(counts, m.inspectedCount)
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})
	return counts[0] * counts[1]
}

func Part1(input []*Monkey) int {
	for i := 0; i < 20; i++ {
		MonkeyRound(input, func(worry int) int {
			return worry / 3
		})
	}

	return CalculateScore(input)
}

func Part2(input []*Monkey) int {
	bigMod := 1
	for _, m := range input {
		bigMod *= m.divisibleBy
	}

	for i := 0; i < 10000; i++ {
		MonkeyRound(input, func(worry int) int {
			return worry % bigMod
		})
	}

	return CalculateScore(input)
}

func MonkeyRound(allMonkeys []*Monkey, worryReducer func(worry int) int) {
	for _, monkey := range allMonkeys {
		MonkeyInspect(monkey, allMonkeys, worryReducer)
	}
}

func MonkeyInspect(monkey *Monkey, allMonkeys []*Monkey, worryReducer func(worry int) int) {
	for _, worry := range monkey.items {
		monkey.inspectedCount += 1
		newWorry := worryReducer(monkey.operation(worry))
		recipient := 0
		if (newWorry % monkey.divisibleBy) == 0 {
			recipient = monkey.recipientTrue
		} else {
			recipient = monkey.recipientFalse
		}
		allMonkeys[recipient].items = append(allMonkeys[recipient].items, newWorry)
	}
	monkey.items = make([]int, 0)
}

func main() {
	input := utils.ReadPiped()
	p1Monkey := make([]*Monkey, 0)
	p2Monkey := make([]*Monkey, 0)
	for i := 0; i < len(input); {
		m1, err := ParseMonkey(input[i : i+6])
		if err != nil {
			fmt.Println(err)
			return
		}
		m2, _ := ParseMonkey(input[i : i+6])

		p1Monkey = append(p1Monkey, m1)
		p2Monkey = append(p2Monkey, m2)
		i += 7
	}

	fmt.Printf("Part 1: %d\n", Part1(p1Monkey))
	fmt.Printf("Part 2: %d\n", Part2(p2Monkey))
}
