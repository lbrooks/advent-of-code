package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func partOne(scanner *bufio.Scanner) {
	numbers := make([]int, 0)

	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), ",") {
		num := strings.TrimSpace(v)
		if num != "" {
			i, _ := strconv.Atoi(num)
			numbers = append(numbers, i)
		}
	}

	roundWin := -1
	value := -1

	for scanner.Scan() {
		board := make([]string, 5)

		scanner.Scan()
		board[0] = scanner.Text()
		scanner.Scan()
		board[1] = scanner.Text()
		scanner.Scan()
		board[2] = scanner.Text()
		scanner.Scan()
		board[3] = scanner.Text()
		scanner.Scan()
		board[4] = scanner.Text()

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

	fmt.Println("Value:", value)
}

func partTwo(scanner *bufio.Scanner) {
	numbers := make([]int, 0)

	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), ",") {
		i, _ := strconv.Atoi(v)
		numbers = append(numbers, i)
	}

	roundWin := -1
	value := -1
	for scanner.Scan() {
		board := make([]string, 5)

		scanner.Scan()
		board[0] = scanner.Text()
		scanner.Scan()
		board[1] = scanner.Text()
		scanner.Scan()
		board[2] = scanner.Text()
		scanner.Scan()
		board[3] = scanner.Text()
		scanner.Scan()
		board[4] = scanner.Text()

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

	fmt.Println("Value:", value)
}

func main() {
	buffer := 1
	var err error
	if len(os.Args) > 1 {
		if buffer, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatal(("Could not convert arg to number: " + os.Args[1]))
		}
	}
	switch buffer {
	case 1:
		partOne(bufio.NewScanner(os.Stdin))
	case 2:
		partTwo(bufio.NewScanner(os.Stdin))
	}
}
