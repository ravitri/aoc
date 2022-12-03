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

func round1(theirs string, mine string) int {

	points := 0

	criteria := map[string]map[string]int{}
	criteria["A"] = map[string]int{}
	criteria["B"] = map[string]int{}
	criteria["C"] = map[string]int{}
	criteria["A"]["Y"] = Win
	criteria["A"]["Z"] = Loss
	criteria["A"]["X"] = Draw
	criteria["B"]["X"] = Loss
	criteria["B"]["Y"] = Draw
	criteria["B"]["Z"] = Win
	criteria["C"]["X"] = Win
	criteria["C"]["Y"] = Loss
	criteria["C"]["Z"] = Draw

	score := map[string]int{}

	score["X"] = 1
	score["Y"] = 2
	score["Z"] = 3

	points = criteria[theirs][mine] + score[mine]

	return points
}

func round2(theirs string, mine string) int {

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

	criteria := map[string]map[string]int{}
	criteria["A"] = map[string]int{}
	criteria["B"] = map[string]int{}
	criteria["C"] = map[string]int{}
	criteria["A"]["Y"] = Win
	criteria["A"]["Z"] = Loss
	criteria["A"]["X"] = Draw
	criteria["B"]["X"] = Loss
	criteria["B"]["Y"] = Draw
	criteria["B"]["Z"] = Win
	criteria["C"]["X"] = Win
	criteria["C"]["Y"] = Loss
	criteria["C"]["Z"] = Draw

	score := map[string]int{}

	score["X"] = 1
	score["Y"] = 2
	score["Z"] = 3

	points = criteria[theirs][replace[theirs][mine]] + score[replace[theirs][mine]]

	return points
}

func rps(input []string) (int, int) {

	score := 0
	score1 := 0

	for _, play := range input {
		p := strings.Split(play, " ")
		s := round1(p[0], p[1])
		s1 := round2(p[0], p[1])
		score += s
		score1 += s1
	}

	return score, score1
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
