package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	MaxLineLength = 1024
	NumDigits     = 12
)

func maxJoltage(line string) int64 {
	if len(line) < NumDigits {
		return 0
	}

	result := make([]byte, 0, NumDigits)
	start := 0

	for i := 0; i < NumDigits; i++ {
		maxPos := len(line) - (NumDigits - i)

		bestPos := start
		bestDigit := line[start]

		for p := start; p <= maxPos; p++ {
			if line[p] > bestDigit {
				bestDigit = line[p]
				bestPos = p
			}
		}

		result = append(result, bestDigit)
		start = bestPos + 1
	}

	joltage, err := strconv.ParseInt(string(result), 10, 64)
	if err != nil {
		return 0
	}
	return joltage
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		log.Fatalf("Failed to open list.txt: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := int64(0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) >= NumDigits {
			joltage := maxJoltage(line)
			total += joltage
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	fmt.Printf("Total output joltage: %d\n", total)
}
