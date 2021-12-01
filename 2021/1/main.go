package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(f string) *os.File {

	file, err := os.Open(f)

	if err != nil {
		log.Fatalf("failed to open file %s", f)
	}

	return file
}

func convertStringToInt(f *os.File) []int {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var l []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		l = append(l, num)
	}

	f.Close()

	return l
}

func sonarSweepA(input []int) int {

	i := input[0]
	increased := 0

	for _, num := range input[1:] {
		if num > i {
			increased++
		}
		i = num
	}

	return increased
}

func slidingWindow(n []int, i int) int {
	return n[i] + n[i+1] + n[i+2]
}

func sonarSweepB(input []int) int {

	sum := slidingWindow(input, 0)
	increased := 0

	for i := 1; i <= len(input)-3; i++ {
		if slidingWindow(input, i) > sum {
			increased++
		}
		sum = slidingWindow(input, i)
	}
	return increased
}

func main() {

	fileName := "input.txt"
	file := readFile(fileName)

	var list []int
	list = convertStringToInt(file)

	result := sonarSweepA(list)
	fmt.Println("Sonar Sweep A increased count is:", result)

	result = sonarSweepB(list)
	fmt.Println("Sonar Sweep B increased count is:", result)
}
