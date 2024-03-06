package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type snailfish struct {
	parent *snailfish

	left  *snailfish
	right *snailfish

	val int
}

func Add(left, right *snailfish) *snailfish {
	fish := &snailfish{left: left, right: right}
	left.parent = fish
	right.parent = fish
	return fish
}

func AddAll(input []string) *snailfish {
	var s *snailfish
	for _, line := range input {
		parsed := ParseSnailfish(line)

		if s == nil {
			s = parsed
		} else {
			s = s.add(parsed)
		}
	}

	return s
}

func (s *snailfish) add(other *snailfish) *snailfish {
	result := &snailfish{
		left:  s,
		right: other,
	}

	// Don't forget to re-parent the child nodes
	s.parent = result
	other.parent = result

	result.Reduce()
	return result
}

// reduce simplifies a snailfish number by repeatedly calling explode then split until both perform no operations
func (s *snailfish) Reduce() {
	for {
		// First, try to explode a number
		if s.explode() {
			// If we did, start the reduction over
			continue
		}

		// Then, try splitting a number
		if s.split() {
			// If we did, start the reduction over
			continue
		}

		// If neither simplification could be performed, we're done reducing
		return
	}
}

// explode "balances" the snailfish number by exploding the pair nested four levels deep. It does this by:
//
// * The left number of the pair is added to the rightmost child of the pair's parent's left child, if any
// * The right number of the pair is added to the leftmost child of the pair's parent's right child, if any
// * The entire pair is replaced with a literal 0
//
// Return true iff a pair explodes. At most one explosion per call.
func (s *snailfish) explode() bool {
	target := findExplodeTarget(0, s)
	if target == nil {
		return false
	}

	// The left target is the rightmost child of the parent's left child
	leftTarget := target
	// Go "up" until we can go left
	for leftTarget.parent != nil {
		old := leftTarget
		leftTarget = leftTarget.parent

		if leftTarget.left != old {
			// Now that we can finally go left, do so
			leftTarget = leftTarget.left
			break
		}
	}

	if leftTarget != s {
		// Go "right" until we hit a leaf
		for leftTarget.right != nil {
			leftTarget = leftTarget.right
		}
	}

	// The right target is the leftmost child of the parent's right child
	rightTarget := target
	// Go "up" until we can go right
	for rightTarget.parent != nil {
		old := rightTarget
		rightTarget = rightTarget.parent

		if rightTarget.right != old {
			// Now that we can finally go right, do so
			rightTarget = rightTarget.right
			break
		}
	}

	if rightTarget != s {
		// Go "left" until we hit a leaf
		for rightTarget.left != nil {
			rightTarget = rightTarget.left
		}
	}

	if leftTarget != nil {
		// If we found a left target, add the target's left leaf to it
		leftTarget.val += target.left.val
	}

	if rightTarget != nil {
		// If we found a right target, add the target's right leaf to it
		rightTarget.val += target.right.val
	}

	// Replace the target with a literal 0
	target.left = nil
	target.right = nil
	target.val = 0

	return true
}

// findExplodeTarget finds the first node whose descendents are literals and is nested 4 levels deep
func findExplodeTarget(n int, s *snailfish) *snailfish {
	if s.left == nil && s.right == nil {
		// Hit a leaf, go back up
		return nil
	}

	if n == 4 && s.left.left == nil && s.right.right == nil {
		return s
	}

	if target := findExplodeTarget(n+1, s.left); target != nil {
		return target
	}

	return findExplodeTarget(n+1, s.right)
}

// split finds the first snailfish literal >= 10 and turns it into a pair where the left element is the floor of the
// literal divided by 2, and the right element is the ceil of the literal divided by 2.
//
// Returns true iff a split was performed. At most one split per call.
func (s *snailfish) split() bool {
	target := findSplitTarget(s)
	if target == nil {
		return false
	}

	left := int(math.Floor(float64(target.val) / 2.0))
	right := int(math.Ceil(float64(target.val) / 2.0))

	target.val = 0
	target.left = &snailfish{parent: target, val: left}
	target.right = &snailfish{parent: target, val: right}

	return true
}

func findSplitTarget(s *snailfish) *snailfish {
	if s.left == nil && s.right == nil {
		if s.val >= 10 {
			return s
		}

		return nil
	}

	if target := findSplitTarget(s.left); target != nil {
		return target
	}

	return findSplitTarget(s.right)
}

// magnitude returns the magnitude of the number, which is 3 times its left child plus 2 times its right child.
//
// The magnitude of a literal is the literal value itself.
func (s *snailfish) magnitude() int {
	if s.left == nil && s.right == nil {
		return s.val
	}

	return 3*s.left.magnitude() + 2*s.right.magnitude()
}

func (s *snailfish) String() string {
	if s.left != nil && s.right != nil {
		return fmt.Sprintf("[%s,%s]", s.left, s.right)
	}
	return strconv.Itoa(s.val)
}

func (s *snailfish) getDepth() (depth int) {
	for w := s; w != nil; w = w.parent {
		depth++
	}
	depth -= 1
	return
}

func ParseSnailfish(input string) *snailfish {
	var root *snailfish
	var working *snailfish

	writeLeft := true
	for _, c := range input {
		if c == '[' {
			nf := &snailfish{}
			if working != nil {
				nf.parent = working
				if writeLeft {
					working.left = nf
				} else {
					working.right = nf
				}
			} else {
				root = nf
			}
			working = nf
			writeLeft = true
		} else if c == ']' {
			working = working.parent
		} else if c == ',' {
			writeLeft = false
		} else {
			diff := int(c) - int('0')
			c := &snailfish{val: diff, parent: working}
			if writeLeft {
				working.left = c
			} else {
				working.right = c
			}
		}
	}

	return root
}

func playOne(input []string) {
	fish := AddAll(input)
	fmt.Printf("Fish %v has magnitude %d\n", fish, fish.magnitude())
}

func playTwo(input []string) {
	var mag int
	for i, line1 := range input {
		for j, line2 := range input {
			if i == j {
				continue
			}
			m := AddAll([]string{line1, line2}).magnitude()
			if m > mag {
				mag = m
			}
		}
	}
	fmt.Printf("Max Magintude %d\n", mag)
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
