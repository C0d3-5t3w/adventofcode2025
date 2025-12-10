package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	MaxCounters = 16
	MaxButtons  = 20
)

var (
	gNumCounters  int
	gNumButtons   int
	gTargets      [MaxCounters]int
	gButtons      [MaxButtons][MaxCounters]int
	gButtonCounts [MaxButtons]int
	gCoeff        [MaxCounters][MaxButtons]int
)

func parseButton(s string) []int {
	var indices []int
	s = strings.Trim(s, "()")
	parts := strings.Split(s, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		var num int
		fmt.Sscanf(p, "%d", &num)
		indices = append(indices, num)
	}
	return indices
}

func parseTargets(s string) []int {
	var targets []int
	s = strings.Trim(s, "{}")
	parts := strings.Split(s, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		var num int
		fmt.Sscanf(p, "%d", &num)
		targets = append(targets, num)
	}
	return targets
}

func checkSolution(presses []int) bool {
	for c := 0; c < gNumCounters; c++ {
		sum := 0
		for b := 0; b < gNumButtons; b++ {
			sum += gCoeff[c][b] * presses[b]
		}
		if sum != gTargets[c] {
			return false
		}
	}
	return true
}

func solveRecursive(presses []int, buttonIdx int, currentSum int64, bestSoFar int64) int64 {
	if currentSum >= bestSoFar {
		return math.MaxInt64
	}

	if buttonIdx == gNumButtons {
		if checkSolution(presses) {
			return currentSum
		}
		return math.MaxInt64
	}

	maxPresses := math.MaxInt32
	for i := 0; i < gButtonCounts[buttonIdx]; i++ {
		counter := gButtons[buttonIdx][i]
		if counter < gNumCounters && gTargets[counter] < maxPresses {
			maxPresses = gTargets[counter]
		}
	}

	if bestSoFar != math.MaxInt64 && int64(bestSoFar)-currentSum < int64(maxPresses) {
		maxPresses = int(bestSoFar - currentSum)
	}

	result := bestSoFar

	for p := 0; p <= maxPresses; p++ {
		presses[buttonIdx] = p
		subResult := solveRecursive(presses, buttonIdx+1, currentSum+int64(p), result)
		if subResult < result {
			result = subResult
		}
	}

	presses[buttonIdx] = 0
	return result
}

func solveGaussWithOptimization() int64 {
	matrix := make([][]float64, gNumCounters)
	for i := 0; i < gNumCounters; i++ {
		matrix[i] = make([]float64, gNumButtons+1)
		for j := 0; j < gNumButtons; j++ {
			matrix[i][j] = float64(gCoeff[i][j])
		}
		matrix[i][gNumButtons] = float64(gTargets[i])
	}

	pivotRow := 0
	pivotCol := make([]int, MaxCounters)
	numPivots := 0
	isPivotCol := make([]bool, MaxButtons)

	for col := 0; col < gNumButtons && pivotRow < gNumCounters; col++ {
		maxRow := pivotRow
		for row := pivotRow + 1; row < gNumCounters; row++ {
			if math.Abs(matrix[row][col]) > math.Abs(matrix[maxRow][col]) {
				maxRow = row
			}
		}

		if math.Abs(matrix[maxRow][col]) < 1e-9 {
			continue
		}

		matrix[pivotRow], matrix[maxRow] = matrix[maxRow], matrix[pivotRow]

		pivotVal := matrix[pivotRow][col]
		for j := 0; j <= gNumButtons; j++ {
			matrix[pivotRow][j] /= pivotVal
		}

		for row := 0; row < gNumCounters; row++ {
			if row != pivotRow && math.Abs(matrix[row][col]) > 1e-9 {
				factor := matrix[row][col]
				for j := 0; j <= gNumButtons; j++ {
					matrix[row][j] -= factor * matrix[pivotRow][j]
				}
			}
		}

		pivotCol[numPivots] = col
		numPivots++
		isPivotCol[col] = true
		pivotRow++
	}

	var freeVars []int
	for col := 0; col < gNumButtons; col++ {
		if !isPivotCol[col] {
			freeVars = append(freeVars, col)
		}
	}
	numFree := len(freeVars)

	if numFree == 0 {
		solution := make([]float64, MaxButtons)
		for i := numPivots - 1; i >= 0; i-- {
			col := pivotCol[i]
			solution[col] = matrix[i][gNumButtons]
			for j := col + 1; j < gNumButtons; j++ {
				solution[col] -= matrix[i][j] * solution[j]
			}
		}

		var total int64
		intSol := make([]int, MaxButtons)
		for i := 0; i < gNumButtons; i++ {
			intSol[i] = int(math.Round(solution[i]))
			if intSol[i] < 0 {
				return math.MaxInt64
			}
			total += int64(intSol[i])
		}

		if checkSolution(intSol) {
			return total
		}
		return math.MaxInt64
	}

	maxFreeVal := 0
	for i := 0; i < gNumCounters; i++ {
		if gTargets[i] > maxFreeVal {
			maxFreeVal = gTargets[i]
		}
	}

	var best int64 = math.MaxInt64
	freeVals := make([]int, MaxButtons)

	for {
		solution := make([]float64, MaxButtons)
		for i := 0; i < numFree; i++ {
			solution[freeVars[i]] = float64(freeVals[i])
		}

		valid := true
		for i := 0; i < numPivots; i++ {
			col := pivotCol[i]
			val := matrix[i][gNumButtons]
			for j := col + 1; j < gNumButtons; j++ {
				val -= matrix[i][j] * solution[j]
			}
			solution[col] = val

			if val < -0.5 {
				valid = false
				break
			}
		}

		if valid {
			intSol := make([]int, MaxButtons)
			var total int64
			allNonneg := true
			for i := 0; i < gNumButtons; i++ {
				intSol[i] = int(math.Round(solution[i]))
				if intSol[i] < 0 {
					allNonneg = false
					break
				}
				total += int64(intSol[i])
			}

			if allNonneg && total < best && checkSolution(intSol) {
				best = total
			}
		}

		idx := 0
		for idx < numFree {
			freeVals[idx]++
			if freeVals[idx] <= maxFreeVal {
				break
			}
			freeVals[idx] = 0
			idx++
		}
		if idx == numFree {
			break
		}
	}

	return best
}

func solveMinPresses(targets []int, numCounters int, buttons [][]int, buttonCounts []int, numButtons int) int64 {
	gNumCounters = numCounters
	gNumButtons = numButtons

	for i := 0; i < numCounters; i++ {
		gTargets[i] = targets[i]
	}

	for b := 0; b < numButtons; b++ {
		gButtonCounts[b] = buttonCounts[b]
		for i := 0; i < buttonCounts[b]; i++ {
			gButtons[b][i] = buttons[b][i]
		}
	}

	for i := range gCoeff {
		for j := range gCoeff[i] {
			gCoeff[i][j] = 0
		}
	}

	for b := 0; b < numButtons; b++ {
		for i := 0; i < buttonCounts[b]; i++ {
			counter := buttons[b][i]
			if counter < numCounters {
				gCoeff[counter][b] = 1
			}
		}
	}

	result := solveGaussWithOptimization()

	if result == math.MaxInt64 {
		if numButtons <= 8 {
			presses := make([]int, MaxButtons)
			result = solveRecursive(presses, 0, 0, math.MaxInt64)
		}
	}

	if result == math.MaxInt64 {
		return -1
	}
	return result
}

func main() {
	fp, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open list.txt: %v\n", err)
		os.Exit(1)
	}
	defer fp.Close()

	var total int64
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		braceIdx := strings.Index(line, "{")
		if braceIdx == -1 {
			continue
		}

		closeIdx := strings.Index(line, "}")
		if closeIdx == -1 {
			continue
		}
		targetsStr := line[braceIdx : closeIdx+1]
		targets := parseTargets(targetsStr)
		numCounters := len(targets)

		var buttons [][]int
		var buttonCounts []int

		p := 0
		for p < len(line) {
			for p < len(line) && line[p] != '(' && line[p] != '{' {
				p++
			}
			if p >= len(line) || line[p] == '{' {
				break
			}

			closeP := strings.Index(line[p:], ")")
			if closeP == -1 {
				break
			}
			buttonStr := line[p : p+closeP+1]
			indices := parseButton(buttonStr)
			buttons = append(buttons, indices)
			buttonCounts = append(buttonCounts, len(indices))
			p = p + closeP + 1
		}
		numButtons := len(buttons)

		minPresses := solveMinPresses(targets, numCounters, buttons, buttonCounts, numButtons)
		if minPresses >= 0 {
			total += minPresses
		}
	}

	fmt.Println(total)
}
