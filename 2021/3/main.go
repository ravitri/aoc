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

func mostAndLeastCommonBit(input []string, pos int) (rune, rune) {
	one, zero := 0, 0
	for _, bin := range input {
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
	} else if zero > one {
		return '0', '1'
	}

	return '2', '2'
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

func calcRating(input []string, pos int, mode bool) string {
	if len(input) == 1 {
		fmt.Println("Returning ", input[0])
		fmt.Println("Position ", pos)
		return input[0]
	}
	most, least := mostAndLeastCommonBit(input, pos)
	rating := make([]string, 0)

	check := least

	if mode {
		check = most

		if most == least {
			check = '1'
		}
	} else {
		check = least

		if most == least {
			check = '0'
		}
	}

	for _, l := range input {
		if rune(l[pos]) == check {
			rating = append(rating, l)
		}
	}
	fmt.Println("Rating is ", rating)
	fmt.Printf("most=%s least=%s pos=%d\n", string(most), string(least), pos)
	return calcRating(rating, pos+1, mode)
}

func lifeSupport(input []string) int64 {

	oxygen := calcRating(input, 0, true)
	fmt.Println("----------------------")
	carbondioxide := calcRating(input, 0, false)
	fmt.Printf("Oxygen is %d and CO2 is %d\n", binaryToDecimal(oxygen), binaryToDecimal(carbondioxide))
	return (binaryToDecimal(oxygen) * binaryToDecimal(carbondioxide))

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

	ans = lifeSupport(list)
	fmt.Println("Life Support rating is ", ans)
}
