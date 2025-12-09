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

const MAX_BOXES = 1000

type Point struct {
	x, y, z int
}

type Edge struct {
	box1, box2 int
	distance   float64
}

var parent [MAX_BOXES]int
var rankArr [MAX_BOXES]int

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
	dx := int64(a.x) - int64(b.x)
	dy := int64(a.y) - int64(b.y)
	dz := int64(a.z) - int64(b.z)
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func main() {
	file, err := os.Open("../list.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening list.txt\n")
		os.Exit(1)
	}
	defer file.Close()

	var boxes [MAX_BOXES]Point
	numBoxes := 0

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

		boxes[numBoxes] = Point{x, y, z}
		numBoxes++
		if numBoxes >= MAX_BOXES {
			break
		}
	}

	fmt.Printf("Read %d junction boxes\n", numBoxes)

	numEdges := (numBoxes * (numBoxes - 1)) / 2
	edges := make([]Edge, 0, numEdges)

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

	connectionsNeeded := numBoxes - 1
	connectionsMade := 0
	var lastBox1, lastBox2 int

	for i := 0; i < len(edges) && connectionsMade < connectionsNeeded; i++ {
		if unionSets(edges[i].box1, edges[i].box2) {
			connectionsMade++
			lastBox1 = edges[i].box1
			lastBox2 = edges[i].box2

			if connectionsMade == connectionsNeeded {
				fmt.Printf("Last connection (#%d): boxes %d and %d\n", connectionsMade,
					lastBox1, lastBox2)
				fmt.Printf("Coordinates: (%d,%d,%d) and (%d,%d,%d)\n", boxes[lastBox1].x,
					boxes[lastBox1].y, boxes[lastBox1].z, boxes[lastBox2].x,
					boxes[lastBox2].y, boxes[lastBox2].z)
			}
		}
	}

	fmt.Printf("Made %d connections to form single circuit\n", connectionsMade)

	answer := int64(boxes[lastBox1].x) * int64(boxes[lastBox2].x)
	fmt.Printf("\nAnswer: %d\n", answer)
}
