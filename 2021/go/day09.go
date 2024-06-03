package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

const (
	InfoColor    = "\033[1;34m%d\033[0m"
	NoticeColor  = "\033[1;36m%d\033[0m"
	WarningColor = "\033[1;33m%d\033[0m"
	ErrorColor   = "\033[1;31m%d\033[0m"
	DebugColor   = "\033[0;36m%d\033[0m"
)

type coord struct {
	r, c int
}

func newCoord(r, c int) coord {
	return coord{r, c}
}

type grid struct {
	grid [][]int
	lows map[coord]struct{}

	seen [][]bool
}

func newGrid(board [][]int) *grid {
	s := make([][]bool, len(board))
	for i := range s {
		s[i] = make([]bool, len(board[0]))
	}
	return &grid{
		grid: board,
		lows: make(map[coord]struct{}),
		seen: s,
	}
}

func (g *grid) String() string {
	var output strings.Builder
	for r, row := range g.grid {
		for c, cell := range row {
			if _, has := g.lows[newCoord(r, c)]; has {
				output.WriteString(fmt.Sprintf(InfoColor, cell))
			} else {
				output.WriteString(fmt.Sprintf("%d", cell))
			}
		}
		output.WriteString("\n")
	}
	return output.String()
}

func (g *grid) BasinString() string {
	var output strings.Builder
	for _, row := range g.grid {
		for _, cell := range row {
			if cell == 9 {
				output.WriteString(fmt.Sprintf("%d", cell))
			} else {
				output.WriteString(fmt.Sprintf(InfoColor, cell))
			}
		}
		output.WriteString("\n")
	}
	return output.String()
}

func (g *grid) markLows() int {
	g.markCorners()
	g.markEdges()
	g.markCenter()

	return g.sumLows()
}

func (g *grid) sumLows() (sum int) {
	for coord := range g.lows {
		sum += g.grid[coord.r][coord.c] + 1
	}
	return
}

func (g *grid) markCorners() {
	{
		r, c := 0, 0
		num := g.grid[r][c]
		if num < g.grid[r][c+1] && num < g.grid[r+1][c] {
			g.lows[newCoord(r, c)] = struct{}{}
		}
	}
	{
		r, c := len(g.grid)-1, 0
		num := g.grid[r][c]
		if num < g.grid[r][c+1] && num < g.grid[r-1][c] {
			g.lows[newCoord(r, c)] = struct{}{}
		}
	}
	{
		r, c := 0, len(g.grid[0])-1
		num := g.grid[r][c]
		if num < g.grid[r][c-1] && num < g.grid[r+1][c] {
			g.lows[newCoord(r, c)] = struct{}{}
		}
	}
	{
		r, c := len(g.grid)-1, len(g.grid[0])-1
		num := g.grid[r][c]
		if num < g.grid[r][c-1] && num < g.grid[r-1][c] {
			g.lows[newCoord(r, c)] = struct{}{}
		}
	}
}

func (g *grid) markEdges() {
	{
		r := 0
		for c := 1; c < len(g.grid[r])-1; c++ {
			num := g.grid[r][c]
			if num < g.grid[r][c-1] && num < g.grid[r][c+1] && num < g.grid[r+1][c] {
				g.lows[newCoord(r, c)] = struct{}{}
			}
		}
	}
	{
		c := 0
		for r := 1; r < len(g.grid)-1; r++ {
			num := g.grid[r][c]
			if num < g.grid[r-1][c] && num < g.grid[r+1][c] && num < g.grid[r][c+1] {
				g.lows[newCoord(r, c)] = struct{}{}
			}
		}
	}
	{
		r := len(g.grid) - 1
		for c := 1; c < len(g.grid[r])-1; c++ {
			num := g.grid[r][c]
			if num < g.grid[r][c-1] && num < g.grid[r][c+1] && num < g.grid[r-1][c] {
				g.lows[newCoord(r, c)] = struct{}{}
			}
		}
	}
	{
		c := len(g.grid[0]) - 1
		for r := 1; r < len(g.grid)-1; r++ {
			num := g.grid[r][c]
			if num < g.grid[r-1][c] && num < g.grid[r+1][c] && num < g.grid[r][c-1] {
				g.lows[newCoord(r, c)] = struct{}{}
			}
		}
	}
}

func (g *grid) markCenter() {
	for r := 1; r < len(g.grid)-1; r++ {
		for c := 1; c < len(g.grid[r])-1; c++ {
			num := g.grid[r][c]

			if num < g.grid[r-1][c] && num < g.grid[r+1][c] && num < g.grid[r][c-1] && num < g.grid[r][c+1] {
				g.lows[newCoord(r, c)] = struct{}{}
			}
		}
	}
}

func (g *grid) countBasins() {
	for r := 1; r < len(g.grid)-1; r++ {
		for c := 1; c < len(g.grid[r])-1; c++ {
			num := g.grid[r][c]

			if num < g.grid[r-1][c] && num < g.grid[r+1][c] && num < g.grid[r][c-1] && num < g.grid[r][c+1] {
				g.lows[newCoord(r, c)] = struct{}{}
			}
		}
	}
}

func createGrid(input []string) *grid {
	board := make([][]int, len(input))
	for rowIdx, row := range input {
		rowSlice := make([]int, len(row))
		for cellIdx, cell := range strings.Split(row, "") {
			num, _ := strconv.Atoi(cell)
			rowSlice[cellIdx] = num
		}
		board[rowIdx] = rowSlice
	}

	return newGrid(board)
}

func playOne(input []string) int {
	g := createGrid(input)
	return g.markLows()
}

func (g *grid) basinArea(r, c int) int {
	if r < 0 || r >= len(g.grid) {
		return 0
	}
	if c < 0 || c >= len(g.grid[0]) {
		return 0
	}
	if g.grid[r][c] == 9 {
		return 0
	}
	if g.seen[r][c] {
		return 0
	}

	g.seen[r][c] = true
	return 1 + g.basinArea(r+1, c) + g.basinArea(r-1, c) + g.basinArea(r, c-1) + g.basinArea(r, c+1)
}

func playTwo(input []string) int {
	g := createGrid(input)

	basins := make([]int, 0)
	for r, row := range g.grid {
		for c := range row {
			size := g.basinArea(r, c)
			if size > 0 {
				basins = append(basins, size)
			}
		}
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	ans := 1
	for i := 0; i < 3 && i < len(basins); i++ {
		ans *= basins[i]
	}

	return ans
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", playOne(input))
	fmt.Printf("Part 2: %d\n", playTwo(input))
}
