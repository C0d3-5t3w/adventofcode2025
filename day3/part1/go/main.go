package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func maxJoltage(line string) int {
	max := 0

	for i := 0; i < len(line)-1; i++ {
		for j := i + 1; j < len(line); j++ {
			first := int(line[i] - '0')
			second := int(line[j] - '0')
			value := first*10 + second
			if value > max {
				max = value
			}
		}
	}

	return max
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		log.Fatalf("Failed to open list.txt: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) > 0 {
			joltage := maxJoltage(line)
			total += joltage
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	fmt.Printf("Total output joltage: %d\n", total)
}
