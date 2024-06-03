package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type CPU struct {
	cycle    int
	register int
}

type Command interface {
	CycleCost() int
	RegisterAdjustment() int
}

type NoOp struct{}

func (n NoOp) CycleCost() int {
	return 1
}
func (n NoOp) RegisterAdjustment() int {
	return 0
}
func (n NoOp) String() string {
	return "NoOp"
}

type AddOp struct {
	val int
}

func (n AddOp) CycleCost() int {
	return 2
}
func (n AddOp) RegisterAdjustment() int {
	return n.val
}
func (n AddOp) String() string {
	return fmt.Sprintf("Add %d", n.val)
}

func Part1(input []Command) int {
	val := 0
	cpu := CPU{cycle: 1, register: 1}

	for _, cmd := range input {
		cost := cmd.CycleCost()
		for c := 0; c < cost; c++ {
			if cpu.cycle == 20 || ((cpu.cycle-60)%40) == 0 {
				val += cpu.cycle * cpu.register
			}
			cpu.cycle += 1
		}
		cpu.register += cmd.RegisterAdjustment()
	}

	if cpu.cycle == 20 || ((cpu.cycle-60)%40) == 0 {
		val += cpu.cycle * cpu.register
	}

	return val
}

func Part2(input []Command) string {
	cpu := CPU{cycle: 1, register: 1}
	display := make([]bool, 0, 240)

	for _, cmd := range input {
		cost := cmd.CycleCost()
		for c := 0; c < cost; c++ {
			normalizedCycle := cpu.cycle % 40
			cursorPos := cpu.register + 1
			willDraw := normalizedCycle == cursorPos || normalizedCycle == cursorPos-1 || normalizedCycle == cursorPos+1
			if (cpu.cycle+1)%40 == 0 {
				fmt.Printf("Cycle [%d] Normalized [%d] Register [%d] Cursor [%d] Will Draw [%v]\n", cpu.cycle, normalizedCycle, cpu.register, cursorPos, willDraw)
			}
			display = append(display, willDraw)

			cpu.cycle += 1
		}
		cpu.register += cmd.RegisterAdjustment()
	}

	var sb strings.Builder
	for i, v := range display {
		c := '.'
		if v {
			c = '#'
		}
		sb.WriteRune(c)

		if (i+1)%40 == 0 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func parseToCommand(input string) Command {
	if input == "noop" {
		return NoOp{}
	}
	pieced := strings.Split(input, " ")
	val, _ := strconv.Atoi(pieced[1])
	return AddOp{val}
}

func main() {
	input := utils.ReadPiped()
	asCommands := make([]Command, len(input))
	for i, v := range input {
		asCommands[i] = parseToCommand(v)
	}

	fmt.Printf("Part 1: %d\n", Part1(asCommands))
	fmt.Printf("Part 2:\n%s\n", Part2(asCommands))
}
