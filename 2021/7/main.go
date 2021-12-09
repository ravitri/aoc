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

	distance := math.MaxInt32
	for _, i := range fuel {
		score := 0
		for _, j := range fuel {
			if i == j {
				continue
			}
			score += int(math.Abs(float64(j) - float64(i)))
		}
		fmt.Printf("Score for %d is %d\n", i, score)
		if score < distance {
			distance = score
		}
	}

	fmt.Printf("Lowest distance is %d\n", distance)
}
