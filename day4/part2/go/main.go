package main

import (
	"bufio"
	"fmt"
	"os"
)

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func countAdjacentRolls(grid []string, rows, cols, r, c int) int {
	count := 0
	for d := 0; d < 8; d++ {
		nr := r + dx[d]
		nc := c + dy[d]
		if nr >= 0 && nr < rows && nc >= 0 && nc < len(grid[nr]) {
			if grid[nr][nc] == '@' {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open list.txt: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var grid []string
	var rows, cols int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
		rows++
		if len(line) > cols {
			cols = len(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Grid size: %d rows x %d cols\n", rows, cols)

	totalRemoved := 0
	removedAny := true

	for removedAny {
		removedAny = false
		removedThisRound := 0

		toRemove := make([][]bool, rows)
		for i := range toRemove {
			toRemove[i] = make([]bool, cols)
		}

		for r := 0; r < rows; r++ {
			for c := 0; c < len(grid[r]); c++ {
				if grid[r][c] != '@' {
					continue
				}

				adjacentRolls := countAdjacentRolls(grid, rows, cols, r, c)

				if adjacentRolls < 4 {
					toRemove[r][c] = true
					removedThisRound++
					removedAny = true
				}
			}
		}

		for r := 0; r < rows; r++ {
			for c := 0; c < len(grid[r]); c++ {
				if toRemove[r][c] {
					gridRow := []rune(grid[r])
					gridRow[c] = '.'
					grid[r] = string(gridRow)
				}
			}
		}

		if removedThisRound > 0 {
			fmt.Printf("Removed %d rolls this round\n", removedThisRound)
		}
		totalRemoved += removedThisRound
	}

	fmt.Printf("Total paper rolls removed: %d\n", totalRemoved)
}
