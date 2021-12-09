package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(f string) []string {

	file, err := os.Open(f)

	if err != nil {
		log.Fatalf("failed to open file %s", f)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []string
	for scanner.Scan() {
		input = strings.Split(scanner.Text(), ",")
	}

	file.Close()

	return input
}

func minimum(list []int) int {
	min := math.MaxInt32
	for i := range list {
		if i < min {
			min = i
		}
	}
	return min
}

func maximum(list []int) int {
	max := 0
	for i := range list {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {

	filename := "input.txt"
	list := readFile(filename)

	fuel := make([]int, 0)
	for _, s := range list {
		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "crab submarines: loading numbers %v\n", err)
			os.Exit(1)
		}
		fuel = append(fuel, val)
	}

	min := minimum(fuel)
	max := maximum(fuel)

	distance := math.MaxInt32

	for i := min; i <= max; i++ {
		score := 0
		dist := 0
		for _, j := range fuel {
			difference := int(math.Abs(float64(j - i)))
			dist = difference * (difference + 1) / 2
			score += dist
		}
		if score < distance {
			distance = score
		}
	}

	fmt.Printf("Lowest distance is %d\n", distance)
}
