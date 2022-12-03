package helper

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(f string) ([]string, error) {

	file, err := os.Open(f)
	var lines []string

	if err != nil {
		log.Fatalf("failed to open file %s", f)
	}

	fread := bufio.NewScanner(file)

	for fread.Scan() {
		lines = append(lines, strings.TrimSpace(fread.Text()))
	}

	file.Close()

	return lines, fread.Err()
}
