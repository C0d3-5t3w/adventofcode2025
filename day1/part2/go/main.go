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

	dial := 50
	zeroCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}

		direction := rune(line[0])
		var distance int
		_, err := fmt.Sscanf(line[1:], "%d", &distance)
		if err != nil {
			continue
		}

		if direction == 'L' {
			firstZero := dial
			if dial == 0 {
				firstZero = 100
			}

			crosses := 0
			if distance >= firstZero {
				crosses = 1 + (distance-firstZero)/100
			}
			zeroCount += crosses

			dial = (dial - distance) % 100
			if dial < 0 {
				dial += 100
			}
		} else if direction == 'R' {
			firstZero := 100 - dial
			if dial == 0 {
				firstZero = 100
			}

			crosses := 0
			if distance >= firstZero {
				crosses = 1 + (distance-firstZero)/100
			}
			zeroCount += crosses

			dial = (dial + distance) % 100
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Answer: %d\n", zeroCount)
}
