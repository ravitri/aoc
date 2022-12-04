package main

import (
	"fmt"
	"strings"

	"github.com/ravitri/aoc/2022/helper"
)

const (
	Loss int = 0
	Draw int = 3
	Win  int = 6
)

var (
	criteria = map[string]map[string]int{
		"A": {
			"X": Draw,
			"Y": Win,
			"Z": Loss,
		},
		"B": {
			"X": Loss,
			"Y": Draw,
			"Z": Win,
		},
		"C": {
			"X": Win,
			"Y": Loss,
			"Z": Draw,
		},
	}

	score1 = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
)

func part1(theirs string, mine string) int {
	points := 0
	points = criteria[theirs][mine] + score1[mine]
	return points
}

func part2(theirs string, mine string) int {
	points := 0
	replace := map[string]map[string]string{}
	replace["A"] = map[string]string{}
	replace["B"] = map[string]string{}
	replace["C"] = map[string]string{}
	replace["A"]["Y"] = "X"
	replace["A"]["Z"] = "Y"
	replace["A"]["X"] = "Z"
	replace["B"]["X"] = "X"
	replace["B"]["Y"] = "Y"
	replace["B"]["Z"] = "Z"
	replace["C"]["X"] = "Y"
	replace["C"]["Y"] = "Z"
	replace["C"]["Z"] = "X"

	points = criteria[theirs][replace[theirs][mine]] + score1[replace[theirs][mine]]

	return points
}

func rps(input []string) (int, int) {

	score1 := 0
	score2 := 0

	for _, play := range input {
		p := strings.Split(play, " ")
		s1 := part1(p[0], p[1])
		s2 := part2(p[0], p[1])
		score1 += s1
		score2 += s2
	}

	return score1, score2
}

func main() {

	file, err := helper.ReadFile("input.txt")

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
	}

	points, points1 := rps(file)

	fmt.Println("Part 1 -> Points =", points)
	fmt.Println("Part 2 -> Points =", points1)
}
