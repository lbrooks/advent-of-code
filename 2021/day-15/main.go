package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var distanceInto [][]int

func getDistanceTo(r, c int) int {
	if r < 0 || r >= len(distanceInto) {
		return -1
	}
	if c < 0 || c >= len(distanceInto[r]) {
		return -1
	}
	return distanceInto[r][c]
}

func populateDistances(input []string, multiplier int) {
	tileRows := len(input)
	distanceInto = make([][]int, tileRows*multiplier)
	for r, line := range input {
		tileColumns := len(line)

		for mr := 0; mr < multiplier; mr++ {
			distanceInto[r+(tileRows*mr)] = make([]int, tileColumns*multiplier)
		}

		for c, val := range line {
			num, _ := strconv.Atoi(string(val))
			for mr := 0; mr < multiplier; mr++ {
				for mc := 0; mc < multiplier; mc++ {
					newVal := (num + mr + mc) % 9
					if newVal == 0 {
						newVal = 9
					}
					distanceInto[r+(tileRows*mr)][c+(tileColumns*mc)] = newVal
				}
			}
		}
	}
}

func findSmallestRemaining(distanceTable map[string]int, visited map[string]bool) (r, c int) {
	minDist := -1
	var coord string
	for pos, dist := range distanceTable {
		if visited[pos] {
			continue
		}
		if minDist < 0 || dist < minDist {
			minDist = dist
			coord = pos
		}
	}
	if coord == "" {
		return -1, -1
	}

	co := strings.Split(coord, "-")
	r, _ = strconv.Atoi(co[0])
	c, _ = strconv.Atoi(co[1])
	return
}

func findLightestPath(startR, startC, endR, endC int) int {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Println("Time taken for calculation:", duration.Nanoseconds()/1_000_000, "ms")
	}()

	endID := fmt.Sprintf("%d-%d", endR, endC)

	log.Println("Created Distance Table")

	distanceQueue := make(PriorityQueue, 0)
	itemTable := make(map[string]*Item)
	for r := 0; r < len(distanceInto); r++ {
		for c := 0; c < len(distanceInto[0]); c++ {
			cellId := fmt.Sprintf("%d-%d", r, c)
			item := &Item{r: r, c: c, id: cellId, priority: math.MaxInt}
			if startR == r && startC == c {
				item.priority = 0
			}
			heap.Push(&distanceQueue, item)
			itemTable[cellId] = item
		}
	}

	visited := make(map[string]bool)

	log.Println("Started processing")

	// for r, c := findSmallestRemaining(distanceTable, visited); r >= 0 && c >= 0; r, c = findSmallestRemaining(distanceTable, visited) {
	for i := heap.Pop(&distanceQueue); distanceQueue.Len() > 0; i = heap.Pop(&distanceQueue) {
		selfId := i.(*Item).id

		if visited[selfId] {
			continue
		}

		r, c := i.(*Item).r, i.(*Item).c

		visited[selfId] = true

		upId := fmt.Sprintf("%d-%d", r-1, c)
		if !visited[upId] {
			upDist := getDistanceTo(r-1, c)
			if upDist >= 0 {
				upDist += i.(*Item).priority
				upCell := itemTable[upId]
				if upDist < upCell.priority {
					distanceQueue.update(upCell, upDist)
				}
			}
		}

		downId := fmt.Sprintf("%d-%d", r+1, c)
		if !visited[downId] {
			downDist := getDistanceTo(r+1, c)
			if downDist >= 0 {
				downDist += i.(*Item).priority
				downCell := itemTable[downId]
				if downDist < downCell.priority {
					distanceQueue.update(downCell, downDist)
				}
			}
		}

		leftId := fmt.Sprintf("%d-%d", r, c-1)
		if !visited[leftId] {
			leftDist := getDistanceTo(r, c-1)
			if leftDist >= 0 {
				leftDist += i.(*Item).priority
				leftCell := itemTable[leftId]
				if leftDist < leftCell.priority {
					distanceQueue.update(leftCell, leftDist)
				}
			}
		}

		rightId := fmt.Sprintf("%d-%d", r, c+1)
		if !visited[rightId] {
			rightDist := getDistanceTo(r, c+1)
			if rightDist >= 0 {
				rightDist += i.(*Item).priority
				rightCell := itemTable[rightId]
				if rightDist < rightCell.priority {
					distanceQueue.update(rightCell, rightDist)
				}
			}
		}

		if r == endR && c == endC {
			break
		}
	}

	return itemTable[endID].priority
}

func playOne(input []string) {
	populateDistances(input, 1)

	length := findLightestPath(0, 0, len(distanceInto)-1, len(distanceInto[0])-1)

	log.Printf("Lowest Weight: %d\n", length)
}

func playTwo(input []string) {
	log.Println("Start")
	populateDistances(input, 5)

	length := findLightestPath(0, 0, len(distanceInto)-1, len(distanceInto[0])-1)

	log.Printf("Lowest Weight: %d\n", length)
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
