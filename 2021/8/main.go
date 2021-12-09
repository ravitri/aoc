package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var unique = []int{2, 4, 3, 7}

func readFile(f string) []map[string][]string {

	inputs := []map[string][]string{}
	file, err := os.Open(f)

	if err != nil {
		log.Fatalf("failed to open file %s", f)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		input := map[string][]string{
			"numbers": strings.Split(line[0], " "),
			"result":  strings.Split(line[1], " "),
		}
		inputs = append(inputs, input)
	}

	file.Close()

	return inputs
}

func main() {

	filename := "input.txt"
	list := readFile(filename)

	count := 0
	for _, v := range list {
		for _, numbers := range v["result"] {
			for _, n := range unique {
				if len(numbers) == n {
					count++
				}
			}
		}
	}

	fmt.Printf("Count of unique numbers is %d\n", count)
}
