package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func readFile(f *os.File) []Line {
	input := bufio.NewScanner(f)
	lines := []Line{}
	for input.Scan() {
		coords := strings.Split(input.Text(), " -> ")
		lines = append(lines, Line{
			start: parseCoordinates(coords[0]),
			end:   parseCoordinates(coords[1]),
		})
	}
	return lines
}

func (l *Line) lineType() string {
	if l.start.x == l.end.x {
		return "y"
	} else if l.start.y == l.end.y {
		return "x"
	} else if (l.start.x < l.end.x && l.start.y < l.end.y) || (l.start.x > l.end.x && l.start.y > l.end.y) {
		return "dd"
	} else {
		return "di"
	}
}

func (l *Line) getPoints(lineType string) []Point {
	var start, finish Point
	points := []Point{}
	if lineType == "x" {
		if l.start.x < l.end.x {
			start = l.start
			finish = l.end
		} else {
			start = l.end
			finish = l.start
		}
	} else if lineType == "y" {
		if l.start.y < l.end.y {
			start = l.start
			finish = l.end
		} else {
			start = l.end
			finish = l.start
		}
	} else if lineType == "dd" {
		if l.start.y < l.end.y {
			start = l.start
			finish = l.end
		} else {
			start = l.end
			finish = l.start
		}
	} else if lineType == "di" {
		if l.start.y < l.end.y {
			start = l.start
			finish = l.end
		} else {
			start = l.end
			finish = l.start
		}
	}
	newPoint := start
	for newPoint != finish {
		points = append(points, newPoint)
		switch lineType {
		case "x":
			newPoint = Point{newPoint.x + 1, newPoint.y}
		case "y":
			newPoint = Point{newPoint.x, newPoint.y + 1}
		case "dd":
			newPoint = Point{newPoint.x + 1, newPoint.y + 1}
		case "di":
			newPoint = Point{newPoint.x - 1, newPoint.y + 1}
		}
	}
	points = append(points, newPoint)
	return points
}

func parseCoordinates(s string) Point {
	coords := strings.Split(s, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parseCoordinates: x parsing %v\n", err)
		os.Exit(1)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parseCoordinates: y parsing %v\n", err)
		os.Exit(1)
	}
	return Point{x, y}
}

func intersectionCount(file *os.File) int {
	field := make(map[Point]int)

	coords := readFile(file)
	for _, line := range coords {
		for _, point := range line.getPoints(line.lineType()) {
			field[point]++
		}
	}

	counter := 0
	for _, v := range field {
		if v >= 2 {
			counter++
		}
	}
	fmt.Println(counter)
	file.Close()

	return counter
}

func main() {

	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "hydrothermal: %v", err)
		os.Exit(1)
	}

	// Part A = 6007 - Yet to generalize

	ans := intersectionCount(f)
	fmt.Println("Part B: Intersection of the horizontal, vertical and diagnonal lines is", ans)

}
