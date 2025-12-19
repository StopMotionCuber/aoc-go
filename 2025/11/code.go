package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type node struct {
	name         string
	successors   []*node
	predecessors []*node
	visited      bool
	pathCount    int
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

	graph := buildGraph(input)

	if !part2 {
		return getPathCount(graph, "you", "out")
	} else {
		return getPathCount(graph, "svr", "fft")*
			getPathCount(graph, "fft", "dac")*
			getPathCount(graph, "dac", "out") +
			getPathCount(graph, "svr", "dac")*
				getPathCount(graph, "dac", "fft")*
				getPathCount(graph, "fft", "out")
	}

	// calculate how many paths from start to end
}

func resetGraph(graph map[string]*node) {
	for _, n := range graph {
		n.visited = false
		n.pathCount = 0
	}
}

func getPathCount(graph map[string]*node, startNodeName string, endNodeName string) int {
	// reset graph
	resetGraph(graph)
	graph[startNodeName].pathCount = 1
	allNodes := []*node{}
	for _, n := range graph {
		allNodes = append(allNodes, n)
	}

	for len(allNodes) > 0 {
		// find nodes with no unvisited predecessors
		for i := 0; i < len(allNodes); i++ {
			n := allNodes[i]
			canVisit := true
			for _, pred := range n.predecessors {
				if !pred.visited {
					canVisit = false
					break
				}
			}
			if canVisit {
				// visit node
				for _, succ := range n.successors {
					succ.pathCount += n.pathCount
				}
				n.visited = true
				// remove from allNodes
				allNodes = append(allNodes[:i], allNodes[i+1:]...)
				i--
			}
		}
	}

	// return path count to "out"
	return graph[endNodeName].pathCount
}

func buildGraph(input string) map[string]*node {
	graph := map[string]*node{}
	for _, line := range strings.Split(input, "\n") {
		// First pass, only create nodes
		parts := strings.Split(line, ":")
		name := strings.TrimSpace(parts[0])

		graph[name] = &node{name: name}
	}
	graph["out"] = &node{name: "out"}
	// Second pass, create edges
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		name := strings.TrimSpace(parts[0])
		successors := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, succName := range successors {
			succName = strings.TrimSpace(succName)
			if succNode, exists := graph[succName]; exists {
				graph[name].successors = append(graph[name].successors, succNode)
				succNode.predecessors = append(succNode.predecessors, graph[name])
			}
		}
	}
	return graph
}
