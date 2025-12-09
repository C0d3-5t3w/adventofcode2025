package main

import (
	"fmt"
	"os"
	"strconv"
)

func nextInvalidID(start uint64) uint64 {
	str := strconv.FormatUint(start, 10)
	len := len(str)

	targetLen := len
	if len%2 != 0 {
		targetLen = len + 1
	}

	for targetLen <= 20 {
		half := targetLen / 2

		minHalf := uint64(1)
		for i := 1; i < half; i++ {
			minHalf *= 10
		}

		maxHalf := minHalf*10 - 1

		startHalf := minHalf

		if targetLen == len || (len%2 != 0 && targetLen == len+1) {
			multiplier := uint64(1)
			for i := 0; i < half; i++ {
				multiplier *= 10
			}
			multiplier += 1

			needed := (start + multiplier - 1) / multiplier
			if needed > startHalf {
				startHalf = needed
			}
		}

		if startHalf <= maxHalf {
			multiplier := uint64(1)
			for i := 0; i < half; i++ {
				multiplier *= 10
			}

			return startHalf*multiplier + startHalf
		}

		targetLen += 2
	}

	return 0
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
		return
	}

	contentStr := string(content)
	totalSum := uint64(0)

	i := 0
	for i < len(contentStr) {

		for i < len(contentStr) && (contentStr[i] < '0' || contentStr[i] > '9') {
			i++
		}
		if i >= len(contentStr) {
			break
		}

		rangeStart := uint64(0)
		for i < len(contentStr) && contentStr[i] >= '0' && contentStr[i] <= '9' {
			rangeStart = rangeStart*10 + uint64(contentStr[i]-'0')
			i++
		}

		for i < len(contentStr) && contentStr[i] != '-' && (contentStr[i] < '0' || contentStr[i] > '9') {
			i++
		}
		if i < len(contentStr) && contentStr[i] == '-' {
			i++
		}

		rangeEnd := uint64(0)
		for i < len(contentStr) && contentStr[i] >= '0' && contentStr[i] <= '9' {
			rangeEnd = rangeEnd*10 + uint64(contentStr[i]-'0')
			i++
		}

		if rangeEnd > 0 {
			rangeSum := sumInvalidIDsInRange(rangeStart, rangeEnd)
			totalSum += rangeSum
		}
	}

	fmt.Printf("Sum of all invalid IDs: %d\n", totalSum)
}
