package main

import (
	"bufio"
	"fmt"
	"log"
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
		input = append(input, scanner.Text())
	}

	file.Close()

	return input
}

func compute(input []string) int {

	depth, horizontal := 0, 0
	for _, line := range input {

		kv := strings.Split(line, " ")
		v, _ := strconv.Atoi(kv[1])
		direction, count := kv[0], v

		switch direction {
		case "up":
			depth -= count

		case "down":
			depth += count

		case "forward":
			horizontal += count
		}
	}
	return depth * horizontal
}

func main() {

	fileName := "input.txt"
	list := readFile(fileName)

	ans := compute(list)
	fmt.Println("Answer for part A is:", ans)

}
