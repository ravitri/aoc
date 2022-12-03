package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/ravitri/aoc/2022/helper"
)

func gatherCalories(in []string) ([]int, error) {

	calories := make([]int, 0)
	cal := 0

	for _, line := range in {
		if line == "" {
			calories = append(calories, cal)
			cal = 0
			continue
		}
		c, _ := strconv.Atoi(line)
		cal += c
	}

	calories = append(calories, cal)

	sort.Ints(calories)

	return calories, nil
}

func main() {

	file, _ := helper.ReadFile("input.txt")

	calories, err := gatherCalories(file)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
	}

	fmt.Println("Part 1 -> Calories =", calories[len(calories)-1])
	fmt.Println("Part 2 -> Calories =", calories[len(calories)-1]+calories[len(calories)-2]+calories[len(calories)-3])
}
