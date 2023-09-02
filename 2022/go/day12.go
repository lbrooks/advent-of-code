package main

import (
	"fmt"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("{x: %d, y: %d}", p.x, p.y)
}

func (p Point) GetNeighbors(maxX, maxY int) (neighbors []Point) {
	if p.y-1 >= 0 {
		neighbors = append(neighbors, Point{x: p.x, y: p.y - 1})
	}
	if p.x-1 >= 0 {
		neighbors = append(neighbors, Point{x: p.x - 1, y: p.y})
	}
	if p.x+1 < maxX {
		neighbors = append(neighbors, Point{x: p.x + 1, y: p.y})
	}
	if p.y+1 < maxY {
		neighbors = append(neighbors, Point{x: p.x, y: p.y + 1})
	}
	return neighbors
}

type Problem struct {
	maxX    int
	maxY    int
	start   Point
	end     Point
	data    [][]int
	visited [][]*VNode
}

type VNode struct {
	path    []Point
	visited bool
}

func (p Problem) String() string {
	var o strings.Builder
	o.WriteString("Start: ")
	o.WriteString(p.start.String())
	o.WriteString("\nEnd: ")
	o.WriteString(p.end.String())
	o.WriteString("\nGrid:\n")

	for _, row := range p.data {
		for _, cell := range row {
			o.WriteString(fmt.Sprintf("%2d", cell))
			o.WriteString(" ")
		}
		o.WriteString("\n")
	}

	return o.String()
}

func NewProblem(input []string) Problem {
	var start, end Point
	data := make([][]int, 0, len(input))
	visited := make([][]*VNode, 0, len(input))

	var maxX, maxY int

	for x, row := range input {
		v := make([]int, 0, len(row))
		visitedRow := make([]*VNode, 0, len(row))
		maxX = x
		for y, r := range row {
			heightLtr := r
			if r == 'S' {
				start = Point{x, y}
				heightLtr = 'a'
			} else if r == 'E' {
				end = Point{x, y}
				heightLtr = 'z'
			}
			v = append(v, int(heightLtr)-int('a'))
			visitedRow = append(visitedRow, &VNode{path: nil, visited: false})
			maxY = y
		}
		data = append(data, v)
		visited = append(visited, visitedRow)
	}

	return Problem{
		maxX:    maxX + 1,
		maxY:    maxY + 1,
		start:   start,
		end:     end,
		data:    data,
		visited: visited,
	}
}

type Queued struct {
	path []Point
	at   Point
}

func Part1(prob Problem) int {
	queue := make([]Queued, 0)
	queue = append(queue, Queued{path: []Point{prob.start}, at: prob.start})

	for len(queue) > 0 {
		current := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = nil
		}

		if prob.visited[current.at.x][current.at.y].visited {
			if len(prob.visited[current.at.x][current.at.y].path) > len(current.path) {
				prob.visited[current.at.x][current.at.y].path = current.path
			}
			continue
		}

		prob.visited[current.at.x][current.at.y].visited = true
		prob.visited[current.at.x][current.at.y].path = current.path

		if current.at.x == prob.end.x && current.at.y == prob.end.y {
			return len(current.path) - 1
		}

		for _, n := range current.at.GetNeighbors(prob.maxX, prob.maxY) {
			heightDelta := prob.data[n.x][n.y] - prob.data[current.at.x][current.at.y]
			if heightDelta <= 1 {
				queue = append(queue, Queued{path: append(current.path, n), at: n})
			}
		}
	}

	return 0
}

func Part2(prob Problem) int {
	queue := make([]Queued, 0)
	queue = append(queue, Queued{path: []Point{prob.end}, at: prob.end})

	minPathLength := prob.maxX*prob.maxY + 1

	for len(queue) > 0 {
		current := queue[0]
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = nil
		}

		if prob.visited[current.at.x][current.at.y].visited {
			continue
		}

		prob.visited[current.at.x][current.at.y].visited = true
		prob.visited[current.at.x][current.at.y].path = current.path

		if prob.data[current.at.x][current.at.y] == 0 {
			if minPathLength > len(current.path) {
				minPathLength = len(current.path)
			}
		} else {
			for _, n := range current.at.GetNeighbors(prob.maxX, prob.maxY) {
				heightDelta := prob.data[current.at.x][current.at.y] - prob.data[n.x][n.y]
				if heightDelta <= 1 {
					queue = append(queue, Queued{path: append(current.path, n), at: n})
				}
			}
		}
	}

	return minPathLength - 1
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", Part1(NewProblem(input)))
	fmt.Printf("Part 2: %d\n", Part2(NewProblem(input)))
}
