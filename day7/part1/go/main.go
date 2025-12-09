package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MaxRows = 256
	MaxCols = 256
)

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open list.txt: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	grid := make([]string, 0, MaxRows)
	rows := 0
	cols := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if rows == 0 {
			cols = len(line)
		}
		grid = append(grid, line)
		rows++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	startCol := -1
	for c := 0; c < cols; c++ {
		if grid[0][c] == 'S' {
			startCol = c
			break
		}
	}

	if startCol == -1 {
		fmt.Fprintf(os.Stderr, "Could not find starting position 'S'\n")
		os.Exit(1)
	}

	beams := make([]bool, cols)
	nextBeams := make([]bool, cols)

	beams[startCol] = true

	totalSplits := 0

	for r := 1; r < rows; r++ {
		for i := range nextBeams {
			nextBeams[i] = false
		}

		for c := 0; c < cols; c++ {
			if beams[c] {
				if grid[r][c] == '^' {
					totalSplits++

					if c-1 >= 0 {
						nextBeams[c-1] = true
					}
					if c+1 < cols {
						nextBeams[c+1] = true
					}
				} else {
					nextBeams[c] = true
				}
			}
		}

		beams, nextBeams = nextBeams, beams
	}

	fmt.Printf("Total beam splits: %d\n", totalSplits)
}
