package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

var extract_instructions = regexp.MustCompile(`^move ([0-9]+) from ([0-9]+) to ([0-9]+)`)

type Room struct {
	warehouse [][]string
}

func (r *Room) MoveStack(count, from, to int) {
	val := append([]string{}, r.warehouse[from][0:count]...)
	r.warehouse[from] = r.warehouse[from][count:]
	r.warehouse[to] = append(val, r.warehouse[to]...)
}

func (r *Room) StackTops() string {
	tops := ""
	for _, stack := range r.warehouse {
		if len(stack) > 0 {
			tops += stack[0]
		} else {
			tops += " "
		}
	}
	return tops
}

func (r *Room) Print() {
	for _, row := range r.warehouse {
		fmt.Println(strings.Join(row, ","))
	}
}

func parseRoom(input []string) (Room, int) {
	data := Room{
		warehouse: make([][]string, 9),
	}
	for i := range data.warehouse {
		data.warehouse[i] = make([]string, 0)
	}

	re := regexp.MustCompile(`^(?:\s+([0-9]+))+`)
	startingRow := 0
	for rowIdx, row := range input {
		if re.MatchString(row) {
			startingRow = rowIdx + 2
			break
		}
		for charIdx, char := range row {
			if char == '[' || char == ']' || char == ' ' {
				continue
			}
			dataIdx := (charIdx - 1) / 4
			data.warehouse[dataIdx] = append(data.warehouse[dataIdx], string(char))
		}
	}

	return data, startingRow
}

func parseInstructions(input string) (count, from, to int) {
	ins := extract_instructions.FindStringSubmatch(input)

	count, _ = strconv.Atoi(ins[1])

	from, _ = strconv.Atoi(ins[2])
	from -= 1

	to, _ = strconv.Atoi(ins[3])
	to -= 1

	return
}

func part1(input []string) string {
	room, startingRow := parseRoom(input)

	for i := startingRow; i < len(input); i++ {
		moveCount, from, to := parseInstructions(input[i])

		for m := 0; m < moveCount; m++ {
			room.MoveStack(1, from, to)
		}
	}

	return room.StackTops()
}

func part2(input []string) string {
	room, startingRow := parseRoom(input)

	for i := startingRow; i < len(input); i++ {
		room.MoveStack(parseInstructions(input[i]))
	}

	return room.StackTops()
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}
