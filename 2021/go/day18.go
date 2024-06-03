package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type Snailfish struct {
	parent *Snailfish

	leftSnail  *Snailfish
	leftNumber *int

	rightSnail  *Snailfish
	rightNumber *int
}

func NewSnailfish(input string) *Snailfish {
	var root *Snailfish = nil
	var workingSnail *Snailfish = nil
	for _, c := range input {
		if c == '[' {
			newSnail := &Snailfish{}
			if workingSnail != nil {
				newSnail.parent = workingSnail

				if workingSnail.HasLeft() {
					workingSnail.rightSnail = newSnail
				} else {
					workingSnail.leftSnail = newSnail
				}
			} else {
				root = newSnail
			}
			workingSnail = newSnail
		} else if c == ']' {
			if workingSnail.parent != nil {
				workingSnail = workingSnail.parent
			}
		} else if c == ',' {
			continue
		} else {
			n, _ := strconv.Atoi(string(c))
			if workingSnail.HasLeft() {
				workingSnail.rightNumber = &n
			} else {
				workingSnail.leftNumber = &n
			}
		}
	}

	return root
}

func (s *Snailfish) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	if s.leftSnail != nil {
		fmt.Fprintf(&sb, "%v", s.leftSnail)
	} else {
		fmt.Fprintf(&sb, "%v", *s.leftNumber)
	}
	sb.WriteString(",")
	if s.rightSnail != nil {
		fmt.Fprintf(&sb, "%v", s.rightSnail)
	} else {
		fmt.Fprintf(&sb, "%v", *s.rightNumber)
	}
	sb.WriteString("]")
	return sb.String()
}

func (s *Snailfish) Depth() int {
	depth := 1
	for w := s; w.parent != nil; w = w.parent {
		depth += 1
	}
	return depth
}

func (s *Snailfish) HasLeft() bool {
	return s.leftSnail != nil || s.leftNumber != nil
}

func (s *Snailfish) HasRight() bool {
	return s.rightSnail != nil || s.rightNumber != nil
}

func Explode(fish *Snailfish, explodePastDepth int) bool {
	if fish == nil {
		return false
	}
	if Explode(fish.leftSnail, explodePastDepth) {
		return true
	}
	if Explode(fish.rightSnail, explodePastDepth) {
		return true
	}
	if fish.leftNumber == nil || fish.rightNumber == nil {
		return false
	}
	if fish.Depth() <= explodePastDepth {
		return false
	}

	closestLeft := fish.parent
	for closestLeft != nil {
		if closestLeft.leftNumber != nil {
			break
		}
		closestLeft = closestLeft.parent
	}
	if closestLeft != nil {
		newNum := *closestLeft.leftNumber + *fish.leftNumber
		closestLeft.leftNumber = &newNum
	}
	if closestLeft != fish.parent {
		newNum := 0
		fish.parent.leftNumber = &newNum
	}
	fish.parent.leftSnail = nil

	closestRight := fish.parent
	for closestRight != nil {
		if closestRight.rightNumber != nil {
			break
		}
		closestRight = closestRight.parent
	}
	if closestRight != nil {
		newNum := *closestRight.rightNumber + *fish.rightNumber
		closestRight.rightNumber = &newNum
	}
	if closestRight != fish.parent {
		newNum := 0
		fish.parent.rightNumber = &newNum
	}
	fish.parent.rightSnail = nil
	return true
}

func Split(fish *Snailfish, splitSize int) bool {
	if fish == nil {
		return false
	}
	if Split(fish.leftSnail, splitSize) {
		return true
	}
	if fish.leftNumber != nil {
		if *fish.leftNumber >= splitSize {
			mid := float64(*fish.leftNumber) / 2.0
			newLeft := int(math.Floor(mid))
			newRight := int(math.Ceil(mid))
			fish.leftSnail = &Snailfish{
				parent:      fish,
				leftNumber:  &newLeft,
				rightNumber: &newRight,
			}
			fish.leftNumber = nil
			return true
		}
	}
	if Split(fish.rightSnail, splitSize) {
		return true
	}
	if fish.rightNumber != nil {
		if *fish.rightNumber >= splitSize {
			mid := float64(*fish.rightNumber) / 2.0
			newLeft := int(math.Floor(mid))
			newRight := int(math.Ceil(mid))
			fish.rightSnail = &Snailfish{
				parent:      fish,
				leftNumber:  &newLeft,
				rightNumber: &newRight,
			}
			fish.rightNumber = nil
			return true
		}
	}
	return false
}

func Magnitude(fish *Snailfish) int {
	leftValue := 0
	if fish.leftSnail != nil {
		leftValue = 3 * Magnitude(fish.leftSnail)
	} else if fish.leftNumber != nil {
		leftValue = 3 * *fish.leftNumber
	}

	rightValue := 0
	if fish.rightSnail != nil {
		rightValue = 2 * Magnitude(fish.rightSnail)
	} else if fish.rightNumber != nil {
		rightValue = 2 * *fish.rightNumber
	}

	return leftValue + rightValue
}

func partOne(inputFish []*Snailfish) *Snailfish {
	var workingFish *Snailfish = nil
	for _, fish := range inputFish {
		if workingFish == nil {
			workingFish = fish
			fmt.Println(workingFish)
		} else {
			newRoot := &Snailfish{
				parent:     nil,
				leftSnail:  workingFish,
				rightSnail: fish,
			}
			workingFish.parent = newRoot
			fish.parent = newRoot

			workingFish = newRoot

			for {
				fmt.Println(workingFish)
				if Explode(workingFish, 4) {
					continue
				}
				if Split(workingFish, 10) {
					continue
				}
				break
			}
		}
	}
	return workingFish
}

func main() {
	input := utils.ReadPiped()

	part1Input := make([]*Snailfish, 0, len(input))
	for _, i := range input {
		part1Input = append(part1Input, NewSnailfish(i))
	}

	part1Fish := partOne(part1Input)

	fmt.Printf("Part 1: Fish %v, Magnitude: %v\n", part1Fish, Magnitude(part1Fish))
}
