package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type coord struct {
	x float64
	y float64
}

func (c coord) String() string {
	return fmt.Sprintf("(%4.0f, %4.0f)", c.x, c.y)
}

type line struct {
	start coord
	end   coord
}

func (l *line) String() string {
	return fmt.Sprint(l.start, " - ", l.end)
}

func newLine(input string) *line {
	coords := strings.Split(input, " -> ")

	start := strings.Split(coords[0], ",")
	sx, _ := strconv.ParseFloat(start[0], 64)
	sy, _ := strconv.ParseFloat(start[1], 64)

	end := strings.Split(coords[1], ",")
	ex, _ := strconv.ParseFloat(end[0], 64)
	ey, _ := strconv.ParseFloat(end[1], 64)

	return &line{
		start: coord{x: sx, y: sy},
		end:   coord{x: ex, y: ey},
	}
}

func (l *line) insideX(x float64) bool {
	if l.start.x <= l.end.x {
		return l.start.x <= x && x <= l.end.x
	}
	return l.end.x <= x && x <= l.start.x
}

func (l *line) insideY(y float64) bool {
	if l.start.y <= l.end.y {
		return l.start.y <= y && y <= l.end.y
	}
	return l.end.y <= y && y <= l.start.y
}

func (l *line) hasPoint(x, y float64) bool {
	if !l.insideX(x) || !l.insideY(y) {
		return false
	}

	if l.isHorizontal() {
		return y == l.start.y
	}

	if l.isVertical() {
		return x == l.start.x
	}

	slope := (l.start.y - l.end.y) / (l.start.x - l.end.x)
	b := l.end.y - (slope * l.end.x)

	if (slope*x + b) == y {
		return true
	}
	return false
}

func (l *line) isHorizontal() bool {
	return l.start.y == l.end.y
}

func (l *line) isVertical() bool {
	return l.start.x == l.end.x
}

func (l *line) isStraight() bool {
	return l.isHorizontal() || l.isVertical()
}

func getLines(input []string, onlyStraight bool) (lines []*line, maxX, maxY float64) {
	lines = make([]*line, 0)
	for _, line := range input {
		l := newLine(line)

		if onlyStraight && !l.isStraight() {
			continue
		}

		lines = append(lines, l)

		if l.start.y > maxY {
			maxY = l.start.y
		}
		if l.end.y > maxY {
			maxY = l.end.y
		}

		if l.start.x > maxX {
			maxX = l.start.x
		}
		if l.end.x > maxX {
			maxX = l.end.x
		}
	}
	return
}

func partOne(input []string) int {
	lines, maxX, maxY := getLines(input, true)

	traffic := 0
	for x := 0.0; x <= maxX; x++ {
		for y := 0.0; y <= maxY; y++ {
			c := 0
			for _, l := range lines {
				if l.hasPoint(x, y) {
					c++
				}
			}
			if c > 1 {
				traffic++
			}
		}
	}

	return traffic
}

func partTwo(input []string) int {
	lines, maxX, maxY := getLines(input, false)

	traffic := 0
	for x := 0.0; x <= maxX; x++ {
		for y := 0.0; y <= maxY; y++ {
			c := 0
			for _, l := range lines {
				if l.hasPoint(x, y) {
					c++
				}
			}
			if c > 1 {
				traffic++
			}
		}
	}

	return traffic
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", partOne(input))
	fmt.Printf("Part 2: %d\n", partTwo(input))
}
