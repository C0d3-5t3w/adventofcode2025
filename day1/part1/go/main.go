package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open list.txt: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	dial := 50
	zeroCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		direction := rune(line[0])
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}

		if direction == 'L' {
			dial = (dial - distance) % 100
			if dial < 0 {
				dial += 100
			}
		} else if direction == 'R' {
			dial = (dial + distance) % 100
		}

		if dial == 0 {
			zeroCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Answer: %d\n", zeroCount)
}
