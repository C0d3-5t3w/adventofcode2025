package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		log.Fatalf("Failed to open list.txt: %v", err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	rows := len(grid)
	if rows == 0 {
		log.Fatal("Empty grid")
	}

	cols := len(grid[0])

	// Find starting position 'S'
	startCol := -1
	for c := 0; c < cols; c++ {
		if grid[0][c] == 'S' {
			startCol = c
			break
		}
	}

	if startCol == -1 {
		log.Fatal("Could not find starting position 'S'")
	}

	// Initialize timelines
	timelines := make([]int64, cols)
	nextTimelines := make([]int64, cols)

	timelines[startCol] = 1

	// Process each row
	for r := 1; r < rows; r++ {
		// Reset next_timelines
		for i := range nextTimelines {
			nextTimelines[i] = 0
		}

		for c := 0; c < cols; c++ {
			if timelines[c] > 0 {
				if grid[r][c] == '^' {
					// Split to left and right
					if c-1 >= 0 {
						nextTimelines[c-1] += timelines[c]
					}
					if c+1 < cols {
						nextTimelines[c+1] += timelines[c]
					}
				} else {
					// Continue straight
					nextTimelines[c] += timelines[c]
				}
			}
		}

		// Swap slices
		timelines, nextTimelines = nextTimelines, timelines
	}

	// Calculate total timelines
	var totalTimelines int64
	for c := 0; c < cols; c++ {
		totalTimelines += timelines[c]
	}

	fmt.Printf("Total timelines: %d\n", totalTimelines)
}
