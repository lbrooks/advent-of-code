package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lbrooks/advent-of-code/utils"
)

type bingo struct {
	marked        []int
	numberToIndex map[int]int
	indexToNumber map[int]int
}

func newBingo(rows []string) *bingo {
	b := bingo{
		marked:        make([]int, 25),
		numberToIndex: make(map[int]int),
		indexToNumber: make(map[int]int),
	}

	idx := 0
	for _, r := range rows {
		for _, c := range strings.Split(r, " ") {
			num := strings.TrimSpace(c)
			if num == "" {
				continue
			}
			v, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
			b.numberToIndex[v] = idx
			b.indexToNumber[idx] = v
			idx++
		}
	}

	return &b
}

func (b *bingo) printNum(idx int) string {
	if b.marked[idx] == 1 {
		return fmt.Sprintf("(%d)", b.indexToNumber[idx])
	}
	return fmt.Sprintf("%d", b.indexToNumber[idx])
}

func (b *bingo) printBoard() {
	fmt.Printf(`		%4s %4s %4s %4s %s
		%4s %4s %4s %4s %s
		%4s %4s %4s %4s %s
		%4s %4s %4s %4s %s
		%4s %4s %4s %4s %s
`,
		b.printNum(0), b.printNum(1), b.printNum(2), b.printNum(3), b.printNum(4),
		b.printNum(5), b.printNum(6), b.printNum(7), b.printNum(8), b.printNum(9),
		b.printNum(10), b.printNum(11), b.printNum(12), b.printNum(13), b.printNum(14),
		b.printNum(15), b.printNum(16), b.printNum(17), b.printNum(18), b.printNum(19),
		b.printNum(20), b.printNum(21), b.printNum(22), b.printNum(23), b.printNum(24),
	)
}

func (b *bingo) playRound(num int) bool {
	if idx, has := b.numberToIndex[num]; has {
		b.marked[idx] = 1

		if (b.marked[0] + b.marked[1] + b.marked[2] + b.marked[3] + b.marked[4]) == 5 {
			return true
		}
		if (b.marked[5] + b.marked[6] + b.marked[7] + b.marked[8] + b.marked[9]) == 5 {
			return true
		}
		if (b.marked[10] + b.marked[11] + b.marked[12] + b.marked[13] + b.marked[14]) == 5 {
			return true
		}
		if (b.marked[15] + b.marked[16] + b.marked[17] + b.marked[18] + b.marked[19]) == 5 {
			return true
		}
		if (b.marked[20] + b.marked[21] + b.marked[22] + b.marked[23] + b.marked[24]) == 5 {
			return true
		}

		if (b.marked[0] + b.marked[5] + b.marked[10] + b.marked[15] + b.marked[20]) == 5 {
			return true
		}
		if (b.marked[1] + b.marked[6] + b.marked[11] + b.marked[16] + b.marked[21]) == 5 {
			return true
		}
		if (b.marked[2] + b.marked[7] + b.marked[12] + b.marked[17] + b.marked[22]) == 5 {
			return true
		}
		if (b.marked[3] + b.marked[8] + b.marked[13] + b.marked[18] + b.marked[23]) == 5 {
			return true
		}
		if (b.marked[4] + b.marked[9] + b.marked[14] + b.marked[19] + b.marked[24]) == 5 {
			return true
		}
	}
	return false
}

func (b *bingo) sumUnmarked() int {
	sum := 0
	for i, v := range b.marked {
		if v == 0 {
			sum += b.indexToNumber[i]
		}
	}
	return sum
}

func (b *bingo) getUnmarked() string {
	vals := make([]string, 0)
	for i, v := range b.marked {
		if v == 0 {
			vals = append(vals, strconv.Itoa(b.indexToNumber[i]))
		}
	}
	return strings.Join(vals, " + ")
}

func partOne(input []string) int {
	numbers := make([]int, 0)

	idx := 0
	for _, v := range strings.Split(input[0], ",") {
		num := strings.TrimSpace(v)
		if num != "" {
			i, _ := strconv.Atoi(num)
			numbers = append(numbers, i)
		}
	}
	idx++

	roundWin := -1
	value := -1
	for idx < len(input) {
		board := make([]string, 5)

		idx++
		board[0] = input[idx]
		idx++
		board[1] = input[idx]
		idx++
		board[2] = input[idx]
		idx++
		board[3] = input[idx]
		idx++
		board[4] = input[idx]
		idx++

		b := newBingo(board)
		for i, n := range numbers {
			if b.playRound(n) {
				if roundWin < 0 || i < roundWin {
					roundWin = i
					value = b.sumUnmarked() * n
				}
				break
			}
		}
	}

	return value
}

func partTwo(input []string) int {
	numbers := make([]int, 0)

	idx := 0
	for _, v := range strings.Split(input[idx], ",") {
		i, _ := strconv.Atoi(v)
		numbers = append(numbers, i)
	}
	idx++

	roundWin := -1
	value := -1
	for idx < len(input) {
		board := make([]string, 5)

		idx++
		board[0] = input[idx]
		idx++
		board[1] = input[idx]
		idx++
		board[2] = input[idx]
		idx++
		board[3] = input[idx]
		idx++
		board[4] = input[idx]
		idx++

		b := newBingo(board)
		for i, n := range numbers {
			if b.playRound(n) {
				if i > roundWin {
					roundWin = i
					value = b.sumUnmarked() * n
				}
				break
			}
		}
	}

	return value
}

func main() {
	input := utils.ReadPiped()

	fmt.Printf("Part 1: %d\n", partOne(input))
	fmt.Printf("Part 2: %d\n", partTwo(input))
}
