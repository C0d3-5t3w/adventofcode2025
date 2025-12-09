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
	maxRows = 10
)

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && len(lines) < maxRows {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	numLines := len(lines)
	if numLines == 0 {
		fmt.Println("No lines read")
		return
	}

	opRow := numLines - 1
	numNumberRows := opRow

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	paddedLines := make([]string, numLines)
	for i, line := range lines {
		paddedLines[i] = line + strings.Repeat(" ", maxLen-len(line))
	}

	grandTotal := int64(0)
	col := 0

	for col < maxLen {

		allSpace := true
		for r := 0; r < numNumberRows; r++ {
			if col < len(paddedLines[r]) && paddedLines[r][col] != ' ' {
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
			allSpace = true
			for r := 0; r < numNumberRows; r++ {
				if col < len(paddedLines[r]) && paddedLines[r][col] != ' ' {
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

		for r := 0; r < numNumberRows; r++ {
			var buf strings.Builder
			for c := startCol; c < endCol && c < len(paddedLines[r]); c++ {
				ch := paddedLines[r][c]
				if ch >= '0' && ch <= '9' {
					buf.WriteByte(ch)
				}
			}

			if buf.Len() > 0 {
				num, _ := strconv.ParseInt(buf.String(), 10, 64)
				numbers = append(numbers, num)
			}
		}

		for c := startCol; c < endCol && c < len(paddedLines[opRow]); c++ {
			if paddedLines[opRow][c] == '+' || paddedLines[opRow][c] == '*' {
				op = rune(paddedLines[opRow][c])
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
