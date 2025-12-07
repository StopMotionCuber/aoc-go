package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
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
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	manifold := make([][]int64, len(lines))
	amountSplits := 0
	for i, line := range lines {
		manifold[i] = make([]int64, len(line))
		if i == 0 {
			for j := range line {
				if line[j] == 'S' {
					manifold[i][j] = 1
				}
			}
		} else {
			for j := range line {
				if manifold[i-1][j] != 0 {
					// Above something is coming down
					// Check whether we need to split
					if line[j] == '^' {
						manifold[i][j+1] += manifold[i-1][j]
						manifold[i][j-1] += manifold[i-1][j]
						amountSplits++
					} else {
						manifold[i][j] += manifold[i-1][j]
					}
				}
			}
		}
	}
	if part2 {
		sum := int64(0)
		for _, v := range manifold[len(manifold)-1] {
			sum += v
		}
		return sum
	}
	return amountSplits
}
