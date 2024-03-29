package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type coord struct {
	col, row int
}

func (c coord) fold(col, row int) coord {
	if col >= 0 {
		if c.col > col {
			return coord{
				col: col - (c.col - col),
				row: c.row,
			}
		}
		return c
	}
	if row >= 0 {
		if c.row > row {
			return coord{
				col: c.col,
				row: row - (c.row - row),
			}
		}
		return c
	}
	return coord{}
}

func printBoard(points map[coord]struct{}, foldCol, foldRow int) {
	var cellCount, rowCount int
	for c := range points {
		if c.col > cellCount {
			cellCount = c.col
		}
		if c.row > rowCount {
			rowCount = c.row
		}
	}

	board := make([][]string, rowCount+1)
	for rowIdx := range board {
		board[rowIdx] = make([]string, cellCount+1)
		for colIdx := range board[rowIdx] {
			board[rowIdx][colIdx] = "."

			if foldRow == rowIdx {
				board[rowIdx][colIdx] = "<"
			}
			if foldCol == colIdx {
				board[rowIdx][colIdx] = "^"
			}
		}
	}
	for c := range points {
		board[c.row][c.col] = "#"
	}

	for _, r := range board {
		fmt.Printf("|%s|\n", strings.Join(r, ""))
	}
	fmt.Println("")
}

func playOne(input []string) {
	points := make(map[coord]struct{})

	for _, v := range input {
		if strings.Contains(v, ",") {
			c := strings.Split(v, ",")
			col, _ := strconv.Atoi(c[0])
			row, _ := strconv.Atoi(c[1])

			points[coord{col: col, row: row}] = struct{}{}
		} else if strings.Contains(v, "=") {
			col, row := -1, -1
			foldOn, _ := strconv.Atoi(v[13:])
			switch v[11] {
			case 'x':
				col = foldOn
			case 'y':
				row = foldOn
			}

			for c := range points {
				delete(points, c)
				points[c.fold(col, row)] = struct{}{}
			}

			break
		}
	}

	fmt.Printf("Number of points remaining: %d\n", len(points))
}

func playTwo(input []string) {
	points := make(map[coord]struct{})

	for _, v := range input {
		if strings.Contains(v, ",") {
			c := strings.Split(v, ",")
			col, _ := strconv.Atoi(c[0])
			row, _ := strconv.Atoi(c[1])

			points[coord{col: col, row: row}] = struct{}{}
		} else if strings.Contains(v, "=") {
			col, row := -1, -1
			foldOn, _ := strconv.Atoi(v[13:])
			switch v[11] {
			case 'x':
				col = foldOn
			case 'y':
				row = foldOn
			}

			for c := range points {
				delete(points, c)
				points[c.fold(col, row)] = struct{}{}
			}
		}
	}

	printBoard(points, -1, -1)
}

func main() {
	input := utils.ReadPiped()

	playOne(input)
	playTwo(input)
}
