package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func getWinner(play string) string {
	switch play {
	case "rock":
		return "paper"
	case "paper":
		return "scissors"
	case "scissors":
		return "rock"
	}
	return "--"
}

func getLoser(play string) string {
	switch play {
	case "rock":
		return "scissors"
	case "paper":
		return "rock"
	case "scissors":
		return "paper"
	}
	return "--"
}

func tranlate_opp_play(play string) string {
	switch play {
	case "A":
		return "rock"
	case "B":
		return "paper"
	case "C":
		return "scissors"
	}
	return "--"
}

func part1_you_play(play string) string {
	switch play {
	case "X":
		return "rock"
	case "Y":
		return "paper"
	case "Z":
		return "scissors"
	}
	return "--"
}

func part2_you_play(opp_play, play string) string {
	if play == "Y" {
		return opp_play
	}

	switch play {
	case "X":
		return getLoser(opp_play)
	case "Y":
		return opp_play
	case "Z":
		return getWinner(opp_play)
	}
	return "--"
}

func part1_score(opp, you string) int {
	opp_play := tranlate_opp_play(opp)
	you_play := part1_you_play(you)

	return pointsForPlay(you_play) + pointsForGame(opp_play, you_play)
}

func part2_score(opp, you string) int {
	opp_play := tranlate_opp_play(opp)
	you_play := part2_you_play(opp_play, you)

	return pointsForPlay(you_play) + pointsForGame(opp_play, you_play)
}

func pointsForGame(opp, you string) int {
	if opp == you {
		return 3
	}
	if opp == "rock" && you == "paper" {
		return 6
	} else if opp == "paper" && you == "scissors" {
		return 6
	} else if opp == "scissors" && you == "rock" {
		return 6
	}
	return 0
}

func pointsForPlay(p string) int {
	switch p {
	case "rock":
		return 1
	case "paper":
		return 2
	case "scissors":
		return 3
	}
	return 0
}

func main() {
	input := utils.ReadPiped()

	part1_points := 0
	part2_points := 0
	for _, line := range input {
		game := strings.Split(line, " ")

		part1_points += part1_score(game[0], game[1])
		part2_points += part2_score(game[0], game[1])
	}

	fmt.Println("Part 1:", part1_points)
	fmt.Println("Part 2:", part2_points)
}
