package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type Point struct {
	x int
	y int
}

func (p *Point) Translate(x, y int) {
	p.x += x
	p.y += y
}

func (p *Point) Equal(other *Point) bool {
	return p.x == other.x && p.y == other.y
}

func (p *Point) IsAdjacent(other *Point) bool {
	if p.Equal(other) {
		return true
	}

	deltaX := p.x - other.x
	closeX := deltaX >= -1 && deltaX <= 1

	deltaY := p.y - other.y
	closeY := deltaY >= -1 && deltaY <= 1

	return closeX && closeY
}

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type TrackingPoint struct {
	point *Point
	track map[string]bool
}

func NewTrackingPoint() *TrackingPoint {
	tp := &TrackingPoint{
		point: &Point{x: 0, y: 0},
		track: make(map[string]bool),
	}
	tp.track[tp.point.String()] = true
	return tp
}

func (p *TrackingPoint) String() string {
	return fmt.Sprintf("(%v) %d", p.point, len(p.track))
}

func (p *TrackingPoint) Translate(x, y int) {
	p.point.Translate(x, y)
	p.track[p.point.String()] = true
}

func PrintPoints(head, tail *TrackingPoint) {
	lowestX := 0
	if head.point.x < tail.point.x {
		lowestX = head.point.x - 1
	} else {
		lowestX = tail.point.x - 1
	}

	highestY := 0
	if head.point.y < tail.point.y {
		highestY = tail.point.y + 1
	} else {
		highestY = head.point.y + 1
	}

	fmt.Printf("Head: %v | Tail: %v\n", head, tail)
	fmt.Print("\n")
	for y := highestY; y > highestY-6; y-- {
		for x := lowestX; x < lowestX+6; x++ {
			spot := &Point{x, y}

			symbol := "."
			color := White

			isHead := spot.Equal(head.point)
			isHeadAdj := head.point.IsAdjacent(spot)
			isTail := spot.Equal(tail.point)
			isTailAdj := tail.point.IsAdjacent(spot)

			if (isHead || isHeadAdj) && (isTail || isTailAdj) {
				color = Green
			} else if isHead || isHeadAdj {
				color = Blue
			} else if isTail || isTailAdj {
				color = Yellow
			}

			if isHead && isTail {
				symbol = "B"
			} else if isHead {
				symbol = "H"
			} else if isTail {
				symbol = "T"
			} else if isHeadAdj || isTailAdj {
				symbol = "#"
			}

			fmt.Printf("%s%s%s", color, symbol, Reset)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

}

func (p *TrackingPoint) MoveAdjacentTo(other *TrackingPoint) {
	if p.point.IsAdjacent(other.point) {
		return
	}

	tX, tY := 0, 0
	if p.point.x > other.point.x {
		tX = -1
	} else if p.point.x < other.point.x {
		tX = 1
	}
	if p.point.y > other.point.y {
		tY = -1
	} else if p.point.y < other.point.y {
		tY = 1
	}
	p.Translate(tX, tY)
}

func Part1(input []string) int {
	head := NewTrackingPoint()
	tail := NewTrackingPoint()

	for _, ins := range input {
		parts := strings.Split(ins, " ")
		direction := parts[0]
		distance, _ := strconv.Atoi(parts[1])

		for i := 0; i < distance; i++ {
			// PrintPoints(head, tail)

			if direction == "U" {
				head.Translate(0, 1)
			} else if direction == "D" {
				head.Translate(0, -1)
			} else if direction == "L" {
				head.Translate(-1, 0)
			} else if direction == "R" {
				head.Translate(1, 0)
			}

			tail.MoveAdjacentTo(head)
		}
	}

	return len(tail.track)
}

func Part2(input []string) int {
	knots := make([]*TrackingPoint, 0)
	for i := 0; i < 10; i++ {
		knots = append(knots, NewTrackingPoint())
	}

	for _, ins := range input {
		parts := strings.Split(ins, " ")
		direction := parts[0]
		distance, _ := strconv.Atoi(parts[1])

		for i := 0; i < distance; i++ {
			if direction == "U" {
				knots[0].Translate(0, 1)
			} else if direction == "D" {
				knots[0].Translate(0, -1)
			} else if direction == "L" {
				knots[0].Translate(-1, 0)
			} else if direction == "R" {
				knots[0].Translate(1, 0)
			}

			for k := 1; k < len(knots); k++ {
				knots[k].MoveAdjacentTo(knots[k-1])
			}
		}
	}

	return len(knots[len(knots)-1].track)
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", Part1(input))
	fmt.Printf("Part 2: %d\n", Part2(input))
}
