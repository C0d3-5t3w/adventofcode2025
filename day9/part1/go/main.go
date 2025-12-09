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

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening list.txt")
		os.Exit(1)
	}
	defer file.Close()

	var points []Point
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			continue
		}
		points = append(points, Point{x: x, y: y})
	}

	fmt.Printf("Read %d red tile coordinates\n", len(points))

	var maxArea int64 = 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			width := abs(int64(points[j].x)-int64(points[i].x)) + 1
			height := abs(int64(points[j].y)-int64(points[i].y)) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}

	fmt.Printf("Largest rectangle area: %d\n", maxArea)
}
