package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MemoKey struct {
	node       string
	visitedDac bool
	visitedFft bool
}

var (
	graph     = make(map[string][]string)
	memo      = make(map[MemoKey]int64)
	onPath    = make(map[string]bool)
	callCount int64
)

func countPaths(node string, visitedDac, visitedFft bool) int64 {
	callCount++
	if callCount%10000000 == 0 {
		fmt.Printf("[Progress] Calls: %d\n", callCount)
	}

	if node == "dac" {
		visitedDac = true
	}
	if node == "fft" {
		visitedFft = true
	}

	if node == "out" {
		if visitedDac && visitedFft {
			return 1
		}
		return 0
	}

	if onPath[node] {
		return 0
	}

	key := MemoKey{node, visitedDac, visitedFft}
	if val, ok := memo[key]; ok {
		return val
	}

	outputs := graph[node]
	if len(outputs) == 0 {
		return 0
	}

	onPath[node] = true

	var total int64 = 0
	for _, next := range outputs {
		total += countPaths(next, visitedDac, visitedFft)
	}

	onPath[node] = false

	memo[key] = total

	return total
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		name := strings.TrimSpace(parts[0])
		outputsStr := strings.TrimSpace(parts[1])
		outputs := strings.Fields(outputsStr)

		graph[name] = outputs
	}

	fmt.Printf("Total devices: %d\n", len(graph))

	total := countPaths("svr", false, false)

	fmt.Printf("Total DFS calls: %d\n", callCount)
	fmt.Printf("Answer: %d\n", total)
}
