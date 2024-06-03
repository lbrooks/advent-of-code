package main

import (
	"fmt"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type node struct {
	id    string
	paths []*node
}

var (
	allNodes map[string]*node
	visited  map[string]int
	path     []string
)

var (
	shortestPath      string
	shortestPathNodes int
)

var (
	longestPath      string
	longestPathNodes int
)

func countPathsToExit(room *node) int {
	if room.id == "end" {
		if shortestPathNodes == 0 || shortestPathNodes > (len(path)+1) {
			shortestPathNodes = len(path) + 1
			shortestPath = strings.Join(path, " -> ") + " -> end"
		}
		if longestPathNodes < (len(path) + 1) {
			longestPathNodes = len(path) + 1
			longestPath = strings.Join(path, " -> ") + " -> end"
		}
		return 1
	}
	if strings.ToLower(room.id) == room.id && visited[room.id] > 0 {
		return 0
	}
	visited[room.id] += 1
	path = append(path, room.id)
	pathsToExit := 0
	for _, next := range room.paths {
		pathsToExit += countPathsToExit(next)
	}
	visited[room.id] -= 1
	path = path[:len(path)-1]
	return pathsToExit
}

func countPathsToExitWithDouble(room *node, allowedDouble string) int {
	ad := allowedDouble

	if room.id == "end" {
		if shortestPathNodes == 0 || shortestPathNodes > (len(path)+1) {
			shortestPathNodes = len(path) + 1
			shortestPath = strings.Join(path, " -> ") + " -> end"
		}
		if longestPathNodes < (len(path) + 1) {
			longestPathNodes = len(path) + 1
			longestPath = strings.Join(path, " -> ") + " -> end"
		}
		return 1
	}

	if strings.ToLower(room.id) == room.id {
		if visited[room.id] > 1 {
			return 0
		}
		if visited[room.id] > 0 {
			if room.id == "start" {
				return 0
			}

			if ad == "" || ad == room.id {
				ad = room.id
			} else {
				return 0
			}
		}
	}

	visited[room.id] += 1
	path = append(path, room.id)
	pathsToExit := 0
	for _, next := range room.paths {
		pathsToExit += countPathsToExitWithDouble(next, ad)
	}
	visited[room.id] -= 1
	path = path[:len(path)-1]
	return pathsToExit
}

func playOne(input []string) {
	allNodes = make(map[string]*node)
	visited = make(map[string]int)
	path = make([]string, 0)

	for _, in := range input {
		path := strings.Split(in, "-")
		if _, has := allNodes[path[0]]; !has {
			allNodes[path[0]] = &node{id: path[0], paths: make([]*node, 0)}
		}
		if _, has := allNodes[path[1]]; !has {
			allNodes[path[1]] = &node{id: path[1], paths: make([]*node, 0)}
		}
		allNodes[path[0]].paths = append(allNodes[path[0]].paths, allNodes[path[1]])
		allNodes[path[1]].paths = append(allNodes[path[1]].paths, allNodes[path[0]])
	}

	pathCount := countPathsToExit(allNodes["start"])

	fmt.Printf("Number of paths: %d\nShortest Path: %d\tPath: %s\nLongest Path: %d\tPath: %s\n",
		pathCount, shortestPathNodes, shortestPath, longestPathNodes, longestPath)
}

func playTwo(input []string) {
	allNodes = make(map[string]*node)
	visited = make(map[string]int)
	path = make([]string, 0)

	for _, in := range input {
		path := strings.Split(in, "-")
		if _, has := allNodes[path[0]]; !has {
			allNodes[path[0]] = &node{id: path[0], paths: make([]*node, 0)}
		}
		if _, has := allNodes[path[1]]; !has {
			allNodes[path[1]] = &node{id: path[1], paths: make([]*node, 0)}
		}
		allNodes[path[0]].paths = append(allNodes[path[0]].paths, allNodes[path[1]])
		allNodes[path[1]].paths = append(allNodes[path[1]].paths, allNodes[path[0]])
	}

	pathCount := countPathsToExitWithDouble(allNodes["start"], "")

	fmt.Printf("Number of paths: %d\nShortest Path: %d\tPath: %s\nLongest Path: %d\tPath: %s",
		pathCount, shortestPathNodes, shortestPath, longestPathNodes, longestPath)
}

func main() {
	input := utils.ReadPiped()

	playOne(input)
	playTwo(input)
}
