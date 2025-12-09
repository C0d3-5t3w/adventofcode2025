package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		log.Fatal("Failed to open list.txt: ", err)
	}
	defer file.Close()

	var ranges []Range
	scanner := bufio.NewScanner(file)
	parsingRanges := true
	freshCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				start, errStart := strconv.ParseInt(parts[0], 10, 64)
				end, errEnd := strconv.ParseInt(parts[1], 10, 64)
				if errStart == nil && errEnd == nil {
					ranges = append(ranges, Range{start: start, end: end})
				}
			}
		} else {
			id, err := strconv.ParseInt(line, 10, 64)
			if err == nil {
				isFresh := false
				for _, r := range ranges {
					if id >= r.start && id <= r.end {
						isFresh = true
						break
					}
				}
				if isFresh {
					freshCount++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file: ", err)
	}

	fmt.Printf("Number of fresh ingredient IDs: %d\n", freshCount)
}
