package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type octopus struct {
	grid [][]int
	lit  int

	boomed [][]bool
}

func newOctopus(input []string) *octopus {
	grid := make([][]int, len(input))
	boomed := make([][]bool, len(input))

	for r, v := range input {
		boomed[r] = make([]bool, len(v))
		grid[r] = make([]int, len(v))
		for c, n := range v {
			l := int(n - '0')
			grid[r][c] = l
		}
	}

	return &octopus{
		grid:   grid,
		lit:    0,
		boomed: boomed,
	}
}

func (g *octopus) resetBoom() bool {
	countEntries := 0
	countBoomed := 0
	for r, row := range g.boomed {
		for c := range row {
			countEntries++
			if g.boomed[r][c] {
				countBoomed++
			}
			g.boomed[r][c] = false
		}
	}
	return countEntries == countBoomed
}
func (g *octopus) addOne() {
	for r, row := range g.grid {
		for c, val := range row {
			g.grid[r][c] = val + 1
		}
	}
}
func (g *octopus) boom(r, c int, addOne bool) int {
	if r < 0 || r >= len(g.grid) {
		return 0
	}
	if c < 0 || c >= len(g.grid[r]) {
		return 0
	}
	if addOne {
		g.grid[r][c] += 1
	}
	if g.grid[r][c] < 10 {
		return 0
	}
	if g.boomed[r][c] {
		return 0
	}
	g.boomed[r][c] = true
	return 1 + g.boom(r-1, c-1, true) + g.boom(r-1, c, true) + g.boom(r-1, c+1, true) +
		g.boom(r, c-1, true) + g.boom(r, c+1, true) +
		g.boom(r+1, c-1, true) + g.boom(r+1, c, true) + g.boom(r+1, c+1, true)
}

func (g *octopus) removeEnergy() {
	for r, row := range g.grid {
		for c, cell := range row {
			if cell > 9 {
				g.grid[r][c] = 0
			}
		}
	}
}

func (g *octopus) playRound() bool {
	g.addOne()

	for r, row := range g.grid {
		for c := range row {
			g.lit += g.boom(r, c, false)
		}
	}

	g.removeEnergy()

	return g.resetBoom()
}

func (g *octopus) String() string {
	var out strings.Builder
	for _, row := range g.grid {
		for _, cell := range row {
			out.WriteString(fmt.Sprintf("%d", cell))
		}
		out.WriteString("\n")
	}
	return out.String()
}

func playOne(input []string) {
	game := newOctopus(input)

	for i := 0; i < 100; i++ {
		game.playRound()
	}

	fmt.Printf("Boomed: %d\n", game.lit)
}

func playTwo(input []string) {
	game := newOctopus(input)

	round := 0
	allFlash := false
	for !allFlash {
		round++
		allFlash = game.playRound()
	}

	fmt.Printf("Round: %d\n", round)
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
