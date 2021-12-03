package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func mostCommonBit(in []string, pos int) rune {
	one, zero := 0, 0
	for _, bin := range in {
		digit := bin[pos]

		switch digit {
		case '0':
			zero++

		case '1':
			one++
		}
	}

	if one > zero {
		return '1'
	}
	return '0'
}

func LeastCommonBit(in []string, pos int) rune {
	one, zero := 0, 0
	for _, bin := range in {
		digit := bin[pos]

		switch digit {
		case '0':
			zero++

		case '1':
			one++
		}
	}

	if one < zero {
		return '1'
	}
	return '0'
}

func gammaA(input []string) int64 {

	var output []rune
	for i := 0; i < len(input[0]); i++ {
		output = append(output, mostCommonBit(input, i))
	}
	ans, _ := strconv.ParseInt(string(output), 2, 64)
	return ans
}

func epsilonA(input []string) int64 {

	var output []rune
	for i := 0; i < len(input[0]); i++ {
		output = append(output, LeastCommonBit(input, i))
	}
	ans, _ := strconv.ParseInt(string(output), 2, 64)
	return ans
}

func powerA(input []string) int64 {

	g := gammaA(input)
	e := epsilonA(input)
	return g * e
}

func main() {

	fileName := "input.txt"
	list := readFile(fileName)

	ans := powerA(list)
	fmt.Println("Power is ", ans)
}
