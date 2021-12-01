package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open file input.txt")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var list []int
	for scanner.Scan() {

		num, _ := strconv.Atoi(scanner.Text())
		list = append(list, num)
	}

	file.Close()

	sonarSweep(list)

}

func sonarSweep(input []int) {

	i := input[0]
	increased := 0

	for _, num := range input[1:] {
		if num > i {
			increased++
		}
		i = num
	}

	fmt.Println("Increased count is ", increased)
}

func main() {
	readFile()
}
