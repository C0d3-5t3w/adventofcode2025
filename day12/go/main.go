package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	MaxShapes    = 6
	ShapeSize    = 3
	MaxRotations = 8
	MaxWidth     = 60
	MaxHeight    = 60
)

type Shape struct {
	cells     [ShapeSize][ShapeSize]int
	cellCount int
}

type ShapeVariants struct {
	variants    [MaxRotations]Shape
	numVariants int
	cellCount   int
}

var shapes [MaxShapes]ShapeVariants
var grid [MaxHeight][MaxWidth]int
var width, height int

func rotateShape(s Shape) Shape {
	result := Shape{cellCount: s.cellCount}
	for r := 0; r < ShapeSize; r++ {
		for c := 0; c < ShapeSize; c++ {
			result.cells[c][ShapeSize-1-r] = s.cells[r][c]
		}
	}
	return result
}

func flipShape(s Shape) Shape {
	result := Shape{cellCount: s.cellCount}
	for r := 0; r < ShapeSize; r++ {
		for c := 0; c < ShapeSize; c++ {
			result.cells[r][ShapeSize-1-c] = s.cells[r][c]
		}
	}
	return result
}

func shapesEqual(a, b Shape) bool {
	for r := 0; r < ShapeSize; r++ {
		for c := 0; c < ShapeSize; c++ {
			if a.cells[r][c] != b.cells[r][c] {
				return false
			}
		}
	}
	return true
}

func variantExists(sv *ShapeVariants, s Shape) bool {
	for i := 0; i < sv.numVariants; i++ {
		if shapesEqual(sv.variants[i], s) {
			return true
		}
	}
	return false
}

func generateVariants(base Shape, sv *ShapeVariants) {
	sv.numVariants = 0
	sv.cellCount = base.cellCount

	current := base
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			if !variantExists(sv, current) {
				sv.variants[sv.numVariants] = current
				sv.numVariants++
			}
			current = rotateShape(current)
		}
		current = flipShape(base)
	}
}

func canPlace(s *Shape, row, col int) bool {
	for r := 0; r < ShapeSize; r++ {
		for c := 0; c < ShapeSize; c++ {
			if s.cells[r][c] != 0 {
				nr := row + r
				nc := col + c
				if nr < 0 || nr >= height || nc < 0 || nc >= width {
					return false
				}
				if grid[nr][nc] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func placeShape(s *Shape, row, col, mark int) {
	for r := 0; r < ShapeSize; r++ {
		for c := 0; c < ShapeSize; c++ {
			if s.cells[r][c] != 0 {
				grid[row+r][col+c] = mark
			}
		}
	}
}

func removeShape(s *Shape, row, col int) {
	for r := 0; r < ShapeSize; r++ {
		for c := 0; c < ShapeSize; c++ {
			if s.cells[r][c] != 0 {
				grid[row+r][col+c] = 0
			}
		}
	}
}

func countRemaining(counts []int) int {
	total := 0
	for i := 0; i < MaxShapes; i++ {
		total += counts[i]
	}
	return total
}

func solve(counts []int, placed int) bool {
	if countRemaining(counts) == 0 {
		return true
	}

	shapeIdx := -1
	for i := 0; i < MaxShapes; i++ {
		if counts[i] > 0 {
			shapeIdx = i
			break
		}
	}
	if shapeIdx < 0 {
		return true
	}

	for row := 0; row <= height-1; row++ {
		for col := 0; col <= width-1; col++ {
			for v := 0; v < shapes[shapeIdx].numVariants; v++ {
				s := &shapes[shapeIdx].variants[v]

				if canPlace(s, row, col) {
					placeShape(s, row, col, placed+1)
					counts[shapeIdx]--

					if solve(counts, placed+1) {
						counts[shapeIdx]++
						return true
					}

					removeShape(s, row, col)
					counts[shapeIdx]++
				}
			}
		}
	}

	return false
}

func countTotalCells(counts []int) int {
	total := 0
	for i := 0; i < MaxShapes; i++ {
		total += counts[i] * shapes[i].cellCount
	}
	return total
}

func canFitRegion(w, h int, counts []int) bool {
	totalCells := countTotalCells(counts)
	if totalCells > w*h {
		return false
	}

	width = w
	height = h
	for r := 0; r < MaxHeight; r++ {
		for c := 0; c < MaxWidth; c++ {
			grid[r][c] = 0
		}
	}

	countsCopy := make([]int, MaxShapes)
	copy(countsCopy, counts)

	return solve(countsCopy, 0)
}

func main() {
	f, err := os.Open("../list.txt")
	if err != nil {
		f, err = os.Open("list.txt")
		if err != nil {
			fmt.Println("Cannot open list.txt")
			return
		}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for i := 0; i < MaxShapes; i++ {
		if !scanner.Scan() {
			break
		}

		var base Shape
		base.cellCount = 0

		for r := 0; r < ShapeSize; r++ {
			if !scanner.Scan() {
				break
			}
			line := scanner.Text()
			for c := 0; c < ShapeSize && c < len(line); c++ {
				if line[c] == '#' {
					base.cells[r][c] = 1
					base.cellCount++
				} else {
					base.cells[r][c] = 0
				}
			}
		}

		scanner.Scan()
		generateVariants(base, &shapes[i])
		fmt.Printf("Shape %d: %d cells, %d variants\n", i, shapes[i].cellCount, shapes[i].numVariants)
	}

	fitCount := 0
	regionNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var w, h int
		counts := make([]int, MaxShapes)

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		dims := strings.Split(strings.TrimSpace(parts[0]), "x")
		if len(dims) != 2 {
			continue
		}

		fmt.Sscanf(dims[0], "%d", &w)
		fmt.Sscanf(dims[1], "%d", &h)

		countParts := strings.Fields(strings.TrimSpace(parts[1]))
		if len(countParts) != MaxShapes {
			continue
		}

		for i := 0; i < MaxShapes; i++ {
			fmt.Sscanf(countParts[i], "%d", &counts[i])
		}

		regionNum++

		totalCells := countTotalCells(counts)
		fmt.Printf("Region %d: %dx%d (area=%d, cells=%d) ", regionNum, w, h, w*h, totalCells)

		if canFitRegion(w, h, counts) {
			fitCount++
			fmt.Println("FIT")
		} else {
			fmt.Println("NO")
		}
	}

	fmt.Printf("\nAnswer: %d\n", fitCount)
}
