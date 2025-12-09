package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	accessibleCount := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if c >= len(grid[r]) || grid[r][c] != '@' {
				continue
			}

			adjacentRolls := 0
			for d := 0; d < 8; d++ {
				nr := r + dx[d]
				nc := c + dy[d]

				if nr >= 0 && nr < rows && nc >= 0 && nc < len(grid[nr]) {
					if grid[nr][nc] == '@' {
						adjacentRolls++
					}
				}
			}

			if adjacentRolls < 4 {
				accessibleCount++
			}
		}
	}

	fmt.Printf("Accessible paper rolls: %d\n", accessibleCount)
}
