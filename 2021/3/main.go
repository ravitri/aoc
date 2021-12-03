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

func mostAndLeastCommonBit(in []string, pos int) (rune, rune) {
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
		return '1', '0'
	}
	return '0', '1'
}

func powerA(input []string) int64 {

	gamma, epsilon := "", ""
	for i := 0; i < len(input[0]); i++ {
		most, least := mostAndLeastCommonBit(input, i)
		gamma += string(most)
		epsilon += string(least)
	}
	return (binaryToDecimal(gamma) * binaryToDecimal(epsilon))
}

func binaryToDecimal(input string) int64 {
	output, _ := strconv.ParseInt(string(input), 2, 64)
	return output
}

func main() {
	fileName := "input.txt"
	list := readFile(fileName)

	ans := powerA(list)
	fmt.Println("Power Consumption is ", ans)
}
