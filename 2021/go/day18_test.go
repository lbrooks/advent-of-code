package main

import (
	"fmt"
	"testing"
)

func TestLineToSnailfish(t *testing.T) {
	line := "[1,2]"
	fish := NewSnailfish(line)

	if line != fish.String() {
		t.Fatalf("'%s' was not deserialized-serialized right: '%v'", line, fish)
	}
}

type explodeTest struct {
	depthExplode int
	timesExplode int
	input        string
	expected     string
}

func TestExplode(t *testing.T) {
	cases := []explodeTest{
		{1, 1, "[[1,2],3]", "[0,5]"},
		{1, 1, "[1,[2,3]]", "[3,0]"},
		{1, 1, "[1,[[2,3],4]]", "[3,[0,7]]"},
		{1, 1, "[[1,[2,3]],4]", "[[3,0],7]"},

		{1, 2, "[[1,2],3]", "[0,5]"},
		{1, 2, "[1,[2,3]]", "[3,0]"},
		{1, 2, "[1,[[2,3],4]]", "[3,0]"},
		{1, 2, "[[1,[2,3]],4]", "[0,7]"},

		{4, 1, "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}

	for _, c := range cases {
		fish := NewSnailfish(c.input)

		fmt.Print(fish)
		for i := 0; i < c.timesExplode; i++ {
			Explode(fish, c.depthExplode)
			fmt.Print(" -> ")
			fmt.Print(fish)
		}
		fmt.Print("\n")
		if fish.String() != c.expected {
			t.Fatalf("'%s' was supposed to explode as '%s' but exploded as '%v'", c.input, c.expected, fish)
		}
	}
}

type splitTest struct {
	splitSize int
	input     string
	expected  string
}

func TestSplit(t *testing.T) {
	cases := []splitTest{
		{5, "[1,2]", "[1,2]"},
		{5, "[1,5]", "[1,[2,3]]"},
		{5, "[5,1]", "[[2,3],1]"},
		{5, "[1,6]", "[1,[3,3]]"},
		{5, "[6,1]", "[[3,3],1]"},
		{5, "[5,6]", "[[2,3],6]"},
		{7, "[[[[0,3],2],[9,[0,8]]],[1,1]]", "[[[[0,3],2],[[4,5],[0,8]]],[1,1]]"},
	}

	for _, c := range cases {
		fish := NewSnailfish(c.input)

		fmt.Print(fish)
		Split(fish, c.splitSize)
		fmt.Print(" -> ")
		fmt.Print(fish)
		fmt.Print("\n")

		if fish.String() != c.expected {
			t.Fatalf("'%s' was supposed to split as '%s' but split as '%v'", c.input, c.expected, fish)
		}
	}
}

type magnitudeTest struct {
	input    string
	expected int
}

func TestMagnitude(t *testing.T) {
	cases := []magnitudeTest{
		{"[9,1]", (3 * 9) + (2 * 1)},
		{"[1,9]", (3 * 1) + (2 * 9)},
		{"[[9,1],[1,9]]", (3 * 29) + (2 * 21)},
	}

	for _, c := range cases {
		actual := Magnitude(NewSnailfish(c.input))
		if actual != c.expected {
			t.Fatalf("Magnitude of '%s' was supposed to be '%v' but was '%v'", c.input, c.expected, actual)
		}
	}
}
