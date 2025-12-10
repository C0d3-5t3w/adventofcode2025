package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func parseLights(s string) (target int, numLights int) {
	target = 0
	numLights = 0
	for i := 1; i < len(s) && s[i] != ']'; i++ {
		if s[i] == '#' {
			target |= 1 << numLights
		}
		numLights++
	}
	return target, numLights
}

func parseButton(s string) (mask int, end int) {
	mask = 0
	i := 1
	for i < len(s) && s[i] != ')' {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			mask |= 1 << num
		} else {
			i++
		}
	}
	if i < len(s) && s[i] == ')' {
		i++
	}
	return mask, i
}

func minPresses(target int, buttons []int, numLights int) int {
	maxState := 1 << numLights
	dist := make([]int, maxState)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	queue := make([]int, 0, maxState)
	dist[0] = 0
	queue = append(queue, 0)

	for front := 0; front < len(queue); front++ {
		state := queue[front]
		if state == target {
			return dist[target]
		}

		for _, btn := range buttons {
			nextState := state ^ btn
			if dist[nextState] == math.MaxInt32 {
				dist[nextState] = dist[state] + 1
				queue = append(queue, nextState)
			}
		}
	}

	if dist[target] == math.MaxInt32 {
		return -1
	}
	return dist[target]
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open list.txt:", err)
		os.Exit(1)
	}
	defer file.Close()

	var total int64 = 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		bracketIdx := strings.Index(line, "[")
		if bracketIdx == -1 {
			continue
		}

		target, numLights := parseLights(line[bracketIdx:])

		var buttons []int

		closeBracketIdx := strings.Index(line[bracketIdx:], "]")
		if closeBracketIdx == -1 {
			continue
		}
		p := bracketIdx + closeBracketIdx + 1

		for p < len(line) {

			for p < len(line) && line[p] != '(' && line[p] != '{' {
				p++
			}
			if p >= len(line) || line[p] == '{' {
				break
			}

			mask, end := parseButton(line[p:])
			buttons = append(buttons, mask)
			p += end
		}

		minVal := minPresses(target, buttons, numLights)
		if minVal >= 0 {
			total += int64(minVal)
		}
	}

	fmt.Println(total)
}
