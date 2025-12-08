package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Point struct {
	x int64
	y int64
	z int64
}

type Tuple struct {
	a int
	b int
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	// solve part 1 here
	var allPoints []Point

	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ",")
		if len(split) != 3 {
			continue
		}
		x, _ := strconv.ParseInt(split[0], 10, 64)
		y, _ := strconv.ParseInt(split[1], 10, 64)
		z, _ := strconv.ParseInt(split[2], 10, 64)
		allPoints = append(allPoints, Point{
			x: x,
			y: y,
			z: z,
		})
	}

	squaredDistances := make(map[int]map[int]int64)
	distances := []int64{}
	distancesToTuple := make(map[int64][]Tuple)
	for i := 0; i < len(allPoints); i++ {
		squaredDistances[i] = make(map[int]int64)
		for j := 0; j < i; j++ {
			squaredDistances[i][j] = (allPoints[i].x-allPoints[j].x)*(allPoints[i].x-allPoints[j].x) +
				(allPoints[i].y-allPoints[j].y)*(allPoints[i].y-allPoints[j].y) +
				(allPoints[i].z-allPoints[j].z)*(allPoints[i].z-allPoints[j].z)
			distances = append(distances, squaredDistances[i][j])
			if distancesToTuple[squaredDistances[i][j]] == nil {
				distancesToTuple[squaredDistances[i][j]] = []Tuple{}
			}
			distancesToTuple[squaredDistances[i][j]] = append(distancesToTuple[squaredDistances[i][j]], Tuple{a: i, b: j})
		}
	}

	iterationCount := 10
	slices.Sort(distances)
	if len(allPoints) == 1000 {
		iterationCount = 1000
	}
	if part2 {
		iterationCount = len(distances)
	}

	connections := make(map[int]map[int]bool)
	// connections is a set per point
	for i := 0; i < len(allPoints); i++ {
		connections[i] = make(map[int]bool)
		connections[i][i] = true
	}

	lastDistance := int64(-1)
	innerIdx := 0
	for i := 0; i < iterationCount; i++ {
		currentDistance := distances[i]
		if currentDistance == lastDistance {
			innerIdx++
		} else {
			innerIdx = 0
		}
		tuple := distancesToTuple[currentDistance][innerIdx]
		// Get all connections for both points
		allConns := make(map[int]bool)
		for k := range connections[tuple.a] {
			allConns[k] = true
		}
		for k := range connections[tuple.b] {
			allConns[k] = true
		}
		// Update all connections to point to the same instance
		for k := range allConns {
			connections[k] = allConns
		}
		if part2 {
			if len(allConns) == len(allPoints) {
				// All points are connected
				// Return both x multiplied together
				return allPoints[tuple.a].x * allPoints[tuple.b].x
			}
		}
	}

	setSizes := getSetSizes(connections)
	slices.Sort(setSizes)
	// Return the size of the largest 3 sets multiplied together
	return setSizes[len(setSizes)-1] * setSizes[len(setSizes)-2] * setSizes[len(setSizes)-3]
}

func getSetSizes(connections map[int]map[int]bool) []int {
	seen := make(map[int]bool)
	var sizes []int
	for _, connSet := range connections {
		size := 0
		for k := range connSet {
			if !seen[k] {
				size++
				seen[k] = true
			}
		}
		if size > 0 {
			sizes = append(sizes, size)
		}
	}
	return sizes
}
