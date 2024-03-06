package main

import (
	"testing"
)

type depthTest struct {
	fish  string
	depth int
}

func TestDepth(t *testing.T) {
	allTests := []depthTest{
		{"[1,3]", 1},
		{"[[2,3],3]", 2},
	}

	for _, s := range allTests {
		left := ParseSnailfish(s.fish)
		for ; left.left != nil; left = left.left {
		}

		d := left.getDepth()
		if d != s.depth {
			t.Errorf("Fish(%v) : Depth = %d; want %d", s.fish, d, s.depth)
		}
	}
}

func TestParse(t *testing.T) {
	allTests := []string{
		"[[4,[[1,2],1]],1]",
		"[1,9]",
		"[[[[9,8],[5,3]],5],[[6,9],[9,[6,8]]]]",
	}

	for _, s := range allTests {
		parsed := ParseSnailfish(s).String()
		if parsed != s {
			t.Errorf("Parse Failure : Got = %d; want %d", parsed, s)
		}
	}
}

type reduceTest struct {
	fish   string
	result string
}

func TestReduce(t *testing.T) {
	allTests := []reduceTest{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
	}

	for _, s := range allTests {
		p := ParseSnailfish(s.fish)
		p.Reduce()
		parsed := p.String()
		if parsed != s.result {
			t.Errorf("Reduce(%s) : Got = %s; want %s", s.fish, parsed, s.result)
		}
	}
}
