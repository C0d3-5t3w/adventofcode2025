package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	MaxBoxes    = 1000
	Connections = 1000
)

type Point struct {
	x, y, z int
}

type Edge struct {
	box1, box2 int
	distance   float64
}

var (
	parent  [MaxBoxes]int
	rankArr [MaxBoxes]int
)

func initUnionFind(n int) {
	for i := 0; i < n; i++ {
		parent[i] = i
		rankArr[i] = 0
	}
}

func find(x int) int {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func unionSets(x, y int) bool {
	px := find(x)
	py := find(y)

	if px == py {
		return false
	}

	if rankArr[px] < rankArr[py] {
		parent[px] = py
	} else if rankArr[px] > rankArr[py] {
		parent[py] = px
	} else {
		parent[py] = px
		rankArr[px]++
	}
	return true
}

func distance(a, b Point) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening list.txt\n")
		os.Exit(1)
	}
	defer file.Close()

	var boxes []Point
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		boxes = append(boxes, Point{x, y, z})

		if len(boxes) >= MaxBoxes {
			break
		}
	}

	numBoxes := len(boxes)
	fmt.Printf("Read %d junction boxes\n", numBoxes)

	var edges []Edge
	for i := 0; i < numBoxes; i++ {
		for j := i + 1; j < numBoxes; j++ {
			edges = append(edges, Edge{
				box1:     i,
				box2:     j,
				distance: distance(boxes[i], boxes[j]),
			})
		}
	}

	fmt.Printf("Calculated %d distances\n", len(edges))

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	initUnionFind(numBoxes)

	actualConnections := 0
	limit := Connections
	if limit > len(edges) {
		limit = len(edges)
	}

	for i := 0; i < limit; i++ {
		if unionSets(edges[i].box1, edges[i].box2) {
			actualConnections++
		}
	}

	fmt.Printf("Processed %d pairs, made %d actual connections\n", limit, actualConnections)

	circuitSize := [MaxBoxes]int{}
	for i := 0; i < numBoxes; i++ {
		root := find(i)
		circuitSize[root]++
	}

	sizes := circuitSize[:]
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	numCircuits := 0
	for i := 0; i < numBoxes; i++ {
		if circuitSize[i] > 0 {
			numCircuits++
		}
	}

	fmt.Printf("Number of circuits: %d\n", numCircuits)
	fmt.Printf("Three largest circuits: %d, %d, %d\n", sizes[0], sizes[1], sizes[2])

	answer := int64(sizes[0]) * int64(sizes[1]) * int64(sizes[2])
	fmt.Printf("\nAnswer: %d\n", answer)
}
