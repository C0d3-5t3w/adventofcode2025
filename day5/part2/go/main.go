package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int64
	end   int64
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		log.Fatalf("Failed to open list.txt: %v", err)
	}
	defer file.Close()

	var ranges []Range
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			start, err1 := strconv.ParseInt(parts[0], 10, 64)
			end, err2 := strconv.ParseInt(parts[1], 10, 64)

			if err1 == nil && err2 == nil {
				ranges = append(ranges, Range{start: start, end: end})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	var totalFresh int64
	currentStart := ranges[0].start
	currentEnd := ranges[0].end

	for i := 1; i < len(ranges); i++ {
		if ranges[i].start <= currentEnd+1 {
			if ranges[i].end > currentEnd {
				currentEnd = ranges[i].end
			}
		} else {
			totalFresh += currentEnd - currentStart + 1
			currentStart = ranges[i].start
			currentEnd = ranges[i].end
		}
	}

	totalFresh += currentEnd - currentStart + 1

	fmt.Printf("Number of fresh ingredient IDs: %d\n", totalFresh)
}
