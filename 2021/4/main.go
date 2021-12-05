package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	board   [5][5]int
	checked [5][5]bool
}

func readFile(f string) []string {

	file, err := os.Open(f)

	if err != nil {
		log.Fatalf("failed to open file %s", f)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var input []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		input = append(input, line)
	}

	file.Close()

	return input
}

func checkBoard(lines []string) *Board {
	b := &Board{}

	for x, line := range lines {
		numbers, err := lineToInt(line, " ")
		if err != nil {
			fmt.Printf("Failed to parse numbers for board: %v\n", err)
		}
		for y, num := range numbers {
			b.board[x][y] = num
		}
	}
	return b
}

func parseBoards(lines []string) []*Board {
	boards := make([]*Board, 0)
	tempLine := make([]string, 0)

	for x, line := range lines {

		if line == "" || x == len(lines)-1 {
			if x == len(lines)-1 {
				tempLine = append(tempLine, line)
			}
			board := checkBoard(tempLine)
			boards = append(boards, board)
			tempLine = make([]string, 0)
			continue
		}

		tempLine = append(tempLine, line)

	}
	return boards
}

func lineToInt(l string, sep string) ([]int, error) {
	numbers := make([]int, 0)

	numString := strings.Split(l, sep)
	for _, n := range numString {

		if n == "" {
			continue
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("Failed to convert string to integer %v\n", err)
			return nil, err
		}
		numbers = append(numbers, i)
	}
	return numbers, nil
}

func (b *Board) check(n int) {
	for i, r := range b.board {
		for j, v := range r {
			if v == n {
				b.checked[i][j] = true
			}
		}
	}
}

func (b *Board) win() bool {
	for r := 0; r < 5; r++ {
		winner := true
		for c := 0; c < 5; c++ {
			if !b.checked[r][c] {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}
	for c := 0; c < 5; c++ {
		winner := true
		for _, r := range b.checked {
			if !r[c] {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}
	return false
}

func (b *Board) score() int {
	sum := 0
	for x, r := range b.board {
		for y, v := range r {
			if !b.checked[x][y] {
				sum += v
			}
		}
	}
	return sum
}

func partA(boards []*Board, numbers []int) int {
	for _, num := range numbers {
		for _, b := range boards {
			b.check(num)
			if b.win() {
				return b.score() * num
			}
		}
	}
	return 0
}

func partB(boards []*Board, numbers []int) int {

	winners := make([]int, 0)

	for _, num := range numbers {
		for i, b := range boards {
			b.check(num)

			if b.win() {
				foundWinner := false
				for _, w := range winners {
					if i == w {
						foundWinner = true
						break
					}
				}
				if !foundWinner {
					winners = append(winners, i)
				}
			}
		}
		if len(winners) == len(boards) {
			return boards[winners[len(winners)-1]].score() * num
		}
	}
	return 0
}

func main() {
	fileName := "input.txt"
	list := readFile(fileName)

	numbers, err := lineToInt(list[0], ",")
	if err != nil {
		fmt.Printf("Failed to parse the string of numbers: %v\n", err)
	}
	boards := parseBoards(list[2:])

	score := partA(boards, numbers)
	fmt.Printf("Score for Part A is %d\n", score)

	score = partB(boards, numbers)
	fmt.Printf("Score for Part B is %d\n", score)

}
