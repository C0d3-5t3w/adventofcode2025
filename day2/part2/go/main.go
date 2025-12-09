package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func buildRepeated(pattern uint64, patternLen int, repeats int) uint64 {
	result := uint64(0)
	multiplier := uint64(1)

	for i := 0; i < repeats; i++ {
		result = result*multiplier + pattern
		if i == 0 {
			for j := 0; j < patternLen; j++ {
				multiplier *= 10
			}
		}
	}

	return result
}

const maxCandidates = 1000

var candidates [maxCandidates]uint64
var numCandidates int

func addCandidate(val uint64) {
	if numCandidates < maxCandidates {
		candidates[numCandidates] = val
		numCandidates++
	}
}

func nextInvalidID(start uint64) uint64 {
	numCandidates = 0

	str := fmt.Sprintf("%d", start)
	startLen := len(str)

	for totalLen := startLen; totalLen <= startLen+2 && totalLen <= 20; totalLen++ {
		for patternLen := 1; patternLen <= totalLen/2; patternLen++ {
			if totalLen%patternLen != 0 {
				continue
			}

			repeats := totalLen / patternLen
			if repeats < 2 {
				continue
			}

			minPattern := uint64(1)
			for i := 1; i < patternLen; i++ {
				minPattern *= 10
			}
			maxPattern := minPattern*10 - 1

			lo := minPattern
			hi := maxPattern
			best := uint64(0)

			for lo <= hi {
				mid := lo + (hi-lo)/2
				num := buildRepeated(mid, patternLen, repeats)

				if num >= start {
					best = num
					hi = mid - 1
				} else {
					lo = mid + 1
				}
			}

			if best > 0 {
				addCandidate(best)
			}
		}
	}

	result := uint64(0)
	for i := 0; i < numCandidates; i++ {
		if result == 0 || candidates[i] < result {
			result = candidates[i]
		}
	}

	return result
}

func sumInvalidIDsInRange(start uint64, end uint64) uint64 {
	sum := uint64(0)

	candidate := nextInvalidID(start)

	for candidate != 0 && candidate <= end {
		sum += candidate
		candidate = nextInvalidID(candidate + 1)
	}

	return sum
}

func main() {
	content, err := os.ReadFile("../list.txt")
	if err != nil {
		fmt.Printf("Failed to open list.txt: %v\n", err)
		os.Exit(1)
	}

	contentStr := string(content)
	totalSum := uint64(0)

	// Parse ranges from content (format: "number-number")
	re := regexp.MustCompile(`(\d+)-(\d+)`)
	matches := re.FindAllStringSubmatch(contentStr, -1)

	for _, match := range matches {
		rangeStart, _ := strconv.ParseUint(match[1], 10, 64)
		rangeEnd, _ := strconv.ParseUint(match[2], 10, 64)

		if rangeEnd > 0 {
			rangeSum := sumInvalidIDsInRange(rangeStart, rangeEnd)
			totalSum += rangeSum
		}
	}

	fmt.Printf("Sum of all invalid IDs: %d\n", totalSum)
}
