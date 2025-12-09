package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const MaxPoints = 1000

type Point struct {
	x, y int
}

type HSegment struct {
	y, xMin, xMax int
}

type VSegment struct {
	x, yMin, yMax int
}

var (
	hSegments []HSegment
	vSegments []VSegment
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isInsideOrOnBoundary(x, y int) bool {
	crossings := 0

	for _, seg := range vSegments {
		if seg.x >= x && y >= seg.yMin && y <= seg.yMax {
			if seg.x == x {
				return true
			}
			if y > seg.yMin && y <= seg.yMax {
				crossings++
			}
		}
	}

	for _, seg := range hSegments {
		if seg.y == y && x >= seg.xMin && x <= seg.xMax {
			return true
		}
	}

	return crossings%2 == 1
}

func isRectangleValid(x1, y1, x2, y2 int) bool {
	minX, maxX := x1, x2
	if x1 > x2 {
		minX, maxX = x2, x1
	}
	minY, maxY := y1, y2
	if y1 > y2 {
		minY, maxY = y2, y1
	}

	if !isInsideOrOnBoundary(minX, minY) {
		return false
	}
	if !isInsideOrOnBoundary(minX, maxY) {
		return false
	}
	if !isInsideOrOnBoundary(maxX, minY) {
		return false
	}
	if !isInsideOrOnBoundary(maxX, maxY) {
		return false
	}

	for x := minX; x <= maxX; x++ {
		if !isInsideOrOnBoundary(x, maxY) {
			return false
		}
	}
	for x := minX; x <= maxX; x++ {
		if !isInsideOrOnBoundary(x, minY) {
			return false
		}
	}

	for y := minY; y <= maxY; y++ {
		if !isInsideOrOnBoundary(minX, y) {
			return false
		}
	}
	for y := minY; y <= maxY; y++ {
		if !isInsideOrOnBoundary(maxX, y) {
			return false
		}
	}

	return true
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
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			continue
		}
		points = append(points, Point{x, y})
		if len(points) >= MaxPoints {
			fmt.Fprintln(os.Stderr, "Too many points")
			os.Exit(1)
		}
	}

	count := len(points)
	fmt.Printf("Read %d red tile coordinates\n", count)

	for i := 0; i < count; i++ {
		next := (i + 1) % count
		x1, y1 := points[i].x, points[i].y
		x2, y2 := points[next].x, points[next].y

		if y1 == y2 {

			xMin, xMax := x1, x2
			if x1 > x2 {
				xMin, xMax = x2, x1
			}
			hSegments = append(hSegments, HSegment{y: y1, xMin: xMin, xMax: xMax})
		} else if x1 == x2 {

			yMin, yMax := y1, y2
			if y1 > y2 {
				yMin, yMax = y2, y1
			}
			vSegments = append(vSegments, VSegment{x: x1, yMin: yMin, yMax: yMax})
		} else {
			fmt.Fprintf(os.Stderr, "Warning: non-axis-aligned segment between (%d,%d) and (%d,%d)\n",
				x1, y1, x2, y2)
		}
	}

	fmt.Printf("Built %d horizontal and %d vertical segments\n", len(hSegments), len(vSegments))

	sort.Slice(hSegments, func(i, j int) bool {
		if hSegments[i].y != hSegments[j].y {
			return hSegments[i].y < hSegments[j].y
		}
		return hSegments[i].xMin < hSegments[j].xMin
	})
	sort.Slice(vSegments, func(i, j int) bool {
		if vSegments[i].x != vSegments[j].x {
			return vSegments[i].x < vSegments[j].x
		}
		return vSegments[i].yMin < vSegments[j].yMin
	})

	var maxArea int64 = 0
	bestI, bestJ := -1, -1

	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			width := int64(abs(points[j].x-points[i].x)) + 1
			height := int64(abs(points[j].y-points[i].y)) + 1
			area := width * height

			if area > maxArea {
				if isRectangleValid(points[i].x, points[i].y, points[j].x, points[j].y) {
					maxArea = area
					bestI = i
					bestJ = j
				}
			}
		}
	}

	if bestI >= 0 {
		fmt.Printf("Best rectangle: (%d,%d) to (%d,%d)\n",
			points[bestI].x, points[bestI].y, points[bestJ].x, points[bestJ].y)
	}
	fmt.Printf("Largest valid rectangle area: %d\n", maxArea)
}
