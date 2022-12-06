package main

import (
	"fmt"
	"unicode"

	"github.com/ravitri/aoc/2022/helper"
)

func part1(in []string) int {
	var found rune
	var s int
	total := 0

	for _, line := range in {

		str := []rune(line)
		first := []rune(str)[:len(line)/2]
		second := []rune(str)[len(line)/2:]
		found = 0

		for _, f := range first {
			for _, s := range second {
				if f == s {
					found = f
					break
				}
				continue
			}

			if found != 0 {
				break
			} else {
				continue
			}
		}

		if found != 0 {
			s = score(found)
			total += s
		}
	}
	return total
}

func common2(in []string) rune {

	var found rune

	for _, i := range in[0] {
		for _, j := range in[1] {
			if i == j {
				for _, k := range in[2] {
					if j == k {
						found = k
					}
				}
			}
		}
	}

	return found
}

func part2(in []string) int {
	total := 0

	length := len(in)

	for i := 1; i <= length/3; i++ {
		temp := []string{}
		for j := ((3 * i) - 2); j <= (3 * i); j++ {
			temp = append(temp, in[j-1])
		}
		c := common2(temp)
		s := score(c)
		total += s
	}

	return total
}

func score(in rune) int {
	// a-z = 1-26
	// A-Z = 27-52

	if unicode.IsLower(in) {
		// ASCII 'a' = 97 and we need 1 as value
		return int(in - 96)
	}

	// ASCII 'A' is 65 and we need 27 as value
	return int(in - 38)
}

func main() {

	file, err := helper.ReadFile("input.txt")

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
	}

	score1 := part1(file)
	score2 := part2(file)

	fmt.Println("Part 1 -> Sum =", score1)
	fmt.Println("Part 2 -> Sum =", score2)
}
