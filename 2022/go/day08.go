package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type Tree struct {
	height  int
	visited bool
	visible bool
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func PrintForest(grid [][]*Tree) {
	for _, r := range grid {
		for _, c := range r {
			color := Red
			if c.visible {
				color = Green
			}
			fmt.Printf("%s%d%s", color, c.height, Reset)
		}
		fmt.Print("\n")
	}
}

func Part1(forest [][]*Tree) int {
	// Drip From Sides
	for r := 0; r < len(forest); r++ {
		runningHeight := forest[r][0].height
		for c := 1; c < len(forest); c++ {
			if forest[r][c].height > runningHeight {
				forest[r][c].visible = true
				runningHeight = forest[r][c].height
			}
		}

		runningHeight = forest[r][len(forest)-1].height
		for c := len(forest) - 2; c > 0; c-- {
			if forest[r][c].height > runningHeight {
				forest[r][c].visible = true
				runningHeight = forest[r][c].height
			}
		}
	}

	// Drip From Top.Bottom
	for c := 0; c < len(forest); c++ {
		runningHeight := forest[0][c].height
		for r := 1; r < len(forest); r++ {
			if forest[r][c].height > runningHeight {
				forest[r][c].visible = true
				runningHeight = forest[r][c].height
			}
		}

		runningHeight = forest[len(forest)-1][c].height
		for r := len(forest) - 2; r > 0; r-- {
			if forest[r][c].height > runningHeight {
				forest[r][c].visible = true
				runningHeight = forest[r][c].height
			}
		}
	}

	numVisible := 0
	for _, row := range forest {
		for _, tree := range row {
			if tree.visible {
				numVisible++
			}
		}
	}

	return numVisible
}

func VisScore(forest [][]*Tree, r, c int) int {
	northVis := 0
	for y := r - 1; y >= 0; y-- {
		northVis++
		if forest[r][c].height <= forest[y][c].height {
			break
		}
	}

	southVis := 0
	for y := r + 1; y < len(forest); y++ {
		southVis++
		if forest[r][c].height <= forest[y][c].height {
			break
		}
	}

	westVis := 0
	for y := c - 1; y >= 0; y-- {
		westVis++
		if forest[r][c].height <= forest[r][y].height {
			break
		}
	}

	eastVis := 0
	for y := c + 1; y < len(forest); y++ {
		eastVis++
		if forest[r][c].height <= forest[r][y].height {
			break
		}
	}

	return northVis * southVis * westVis * eastVis
}

func Part2(forest [][]*Tree) int {
	largestView := 0
	for r := 0; r < len(forest); r++ {
		for c := 0; c < len(forest); c++ {
			vis := VisScore(forest, r, c)
			if largestView < vis {
				largestView = vis
			}
		}
	}

	return largestView
}

func main() {
	input := utils.ReadPiped()

	forest_1 := make([][]*Tree, 0, len(input))
	forest_2 := make([][]*Tree, 0, len(input))
	for r, line := range input {
		rowVisible := r == 0 || r == len(input)-1

		heights := strings.Split(line, "")
		row_1 := make([]*Tree, 0, len(heights))
		row_2 := make([]*Tree, 0, len(heights))
		for c, h := range heights {
			colVisible := c == 0 || c == len(line)-1

			v, _ := strconv.Atoi(string(h))
			row_1 = append(row_1, &Tree{height: v, visited: rowVisible || colVisible, visible: rowVisible || colVisible})
			row_2 = append(row_2, &Tree{height: v, visited: rowVisible || colVisible, visible: rowVisible || colVisible})
		}
		forest_1 = append(forest_1, row_1)
		forest_2 = append(forest_2, row_2)
	}

	fmt.Printf("Part 1: %d\n", Part1(forest_1))
	fmt.Printf("Part 2: %d\n", Part2(forest_2))

}
