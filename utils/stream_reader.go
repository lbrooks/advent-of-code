package utils

import (
	"bufio"
	"os"
)

func ReadPiped() []string {
	input := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
