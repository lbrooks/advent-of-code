package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lbrooks/advent-of-code/utils"
)

type doubleKeyMap map[int]map[int]int

func newDkm() doubleKeyMap {
	return make(map[int]map[int]int)
}

func (d doubleKeyMap) add(k1, k2, v int) {
	if _, has := d[k1]; !has {
		d[k1] = make(map[int]int)
	}
	d[k1][k2] += v
}

func playOne(input string) {
	presents := newDkm()
	x, y := 0, 0
	for _, d := range input {
		presents.add(x, y, 1)

		switch d {
		case '^':
			y++
		case 'v':
			y--
		case '<':
			x--
		case '>':
			x++
		}
	}
	presents.add(x, y, 1)

	houses := 0
	for _, v := range presents {
		houses += len(v)
	}

	fmt.Printf("Houses Delevered To: %d\n", houses)
}

func playTwo(input string) {
	presents := newDkm()
	isRobo := false
	x, y := 0, 0
	rx, ry := 0, 0
	for _, d := range input {
		if isRobo {
			presents.add(rx, ry, 1)

			switch d {
			case '^':
				ry++
			case 'v':
				ry--
			case '<':
				rx--
			case '>':
				rx++
			}
		} else {
			presents.add(x, y, 1)

			switch d {
			case '^':
				y++
			case 'v':
				y--
			case '<':
				x--
			case '>':
				x++
			}
		}
		isRobo = !isRobo
	}
	presents.add(x, y, 1)
	presents.add(rx, ry, 1)

	houses := 0
	for _, v := range presents {
		houses += len(v)
	}

	fmt.Printf("Houses Delevered To: %d\n", houses)
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

	switch buffer {
	case 1:
		playOne(input[0])
	case 2:
		playTwo(input[0])
	}
}
