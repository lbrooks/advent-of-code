package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lbrooks/advent-of-code/utils"
)

type BoundingBox struct {
	xmin, xmax, ymin, ymax int
}

func (b BoundingBox) InTarget(p probe) int {
	if p.x < b.xmin {
		if p.velocityX == 0 {
			return 1
		}
		return -1
	}
	if p.x > b.xmax {
		return 1
	}

	if p.y > b.ymax {
		return -1
	}
	if p.y < b.ymin {
		return 1
	}

	return 0
}

type probe struct {
	x, y                 int
	velocityX, velocityY int
}

func (p probe) move() probe {
	resVelocityX := p.velocityX
	if resVelocityX > 0 {
		resVelocityX--
	} else if resVelocityX < 0 {
		resVelocityX++
	}
	return probe{
		x:         p.x + p.velocityX,
		y:         p.y + p.velocityY,
		velocityX: resVelocityX,
		velocityY: p.velocityY - 1,
	}
}

func playOne(input BoundingBox) {
	top := input.ymin
	if top < 0 {
		top *= -1
	}
	top -= 1

	res := 0
	for ; top > 0; top-- {
		res += top
	}

	fmt.Println("Max Height", res)
}

func playTwo(input BoundingBox) {
	count := 0
	for vX := input.xmax; vX > 0; vX-- {
		for vY := 10_000; vY > -10_000; vY-- {
			p := probe{velocityX: vX, velocityY: vY}
			tgt := input.InTarget(p)
			for tgt < 0 {
				p = p.move()
				tgt = input.InTarget(p)
			}
			if tgt == 0 {
				count++
			}
		}
	}
	fmt.Println("Number Of Entries:", count)
}

func main() {
	buffer := 1
	var err error
	if len(os.Args) > 1 {
		if buffer, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(("Could not convert arg to number: " + os.Args[1]))
		}
	}

	input := utils.ReadPiped()

	var xmin, xmax, ymin, ymax int
	count, err := fmt.Sscanf(input[0], "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	if count != 4 || err != nil {
		panic(fmt.Sprint("expected 4 values, found:", count, "\nError: \n", err))
	}

	target := BoundingBox{xmin, xmax, ymin, ymax}

	switch buffer {
	case 1:
		playOne(target)
	case 2:
		playTwo(target)
	}
}
