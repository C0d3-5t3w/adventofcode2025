package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	MaxLine = 16384
	MaxRows = 10
)

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	var maxLen int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() && len(lines) < MaxRows {
		line := scanner.Text()
		if len(line) > maxLen {
			maxLen = len(line)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	numLines := len(lines)
	opRow := numLines - 1
	numNumberRows := opRow

	for i := 0; i < numLines; i++ {
		if len(lines[i]) < maxLen {
			lines[i] = lines[i] + strings.Repeat(" ", maxLen-len(lines[i]))
		}
	}

	var grandTotal int64 = 0
	col := 0

	for col < maxLen {

		allSpace := true
		for r := 0; r < numNumberRows; r++ {
			if col < len(lines[r]) && lines[r][col] != ' ' {
				allSpace = false
				break
			}
		}

		if allSpace {
			col++
			continue
		}

		startCol := col

		for col < maxLen {
			allSpace := true
			for r := 0; r < numNumberRows; r++ {
				if col < len(lines[r]) && lines[r][col] != ' ' {
					allSpace = false
					break
				}
			}
			if allSpace {
				break
			}
			col++
		}
		endCol := col

		var numbers []int64
		op := '+'

		c := endCol - 1
		for c >= startCol {

			hasDigit := false
			for r := 0; r < numNumberRows; r++ {
				if c < len(lines[r]) && lines[r][c] >= '0' && lines[r][c] <= '9' {
					hasDigit = true
					break
				}
			}

			if !hasDigit {
				c--
				continue
			}

			var num int64 = 0
			for r := 0; r < numNumberRows; r++ {
				if c < len(lines[r]) && lines[r][c] >= '0' && lines[r][c] <= '9' {
					num = num*10 + int64(lines[r][c]-'0')
				}
			}
			numbers = append(numbers, num)
			c--
		}

		for c := startCol; c < endCol; c++ {
			if c < len(lines[opRow]) && (lines[opRow][c] == '+' || lines[opRow][c] == '*') {
				op = rune(lines[opRow][c])
				break
			}
		}

		if len(numbers) > 0 {
			var result int64
			if op == '+' {
				result = 0
				for _, num := range numbers {
					result += num
				}
			} else {
				result = 1
				for _, num := range numbers {
					result *= num
				}
			}
			grandTotal += result
		}
	}

	fmt.Printf("Grand Total: %d\n", grandTotal)
}
