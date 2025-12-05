package main

import (
	"strconv"
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

	// Gather ranges
	lines := strings.Split(input, "\n")
	i := 0
	var ranges [][2]int64
	for ; lines[i] != ""; i++ {
		// parse range line
		splitted := strings.Split(lines[i], "-")
		start, _ := strconv.ParseInt(splitted[0], 10, 64)
		end, _ := strconv.ParseInt(splitted[1], 10, 64)
		ranges = append(ranges, [2]int64{start, end})
	}
	if part2 {
		processedRanges := make([][2]int64, 0)
		for _, rng := range ranges {
			mergeRanges(&processedRanges, rng[0], rng[1])
		}
		result := 0
		for _, rng := range processedRanges {
			result += int(rng[1] - rng[0] + 1)
		}
		return result
	}
	i++
	freshItems := 0
	for ; i < len(lines) && lines[i] != ""; i++ {
		current, _ := strconv.ParseInt(lines[i], 10, 64)
		for _, rng := range ranges {
			if current >= rng[0] && current <= rng[1] {
				freshItems++
				break
			}
		}
	}

	// solve part 1 here
	return freshItems
}

func mergeRanges(processedRanges *[][2]int64, start int64, end int64) {
	var overlapsToRemove []int
	for i, rng := range *processedRanges {
		if start > rng[1] || end < rng[0] {
			// No overlap
			continue
		}
		// There is an overlap between the ranges
		// Create a new range that is the intersection of the two
		start = min(rng[0], start)
		end = max(rng[1], end)
		overlapsToRemove = append(overlapsToRemove, i)
	}
	for j := len(overlapsToRemove) - 1; j >= 0; j-- {
		idx := overlapsToRemove[j]
		*processedRanges = append((*processedRanges)[:idx], (*processedRanges)[idx+1:]...)
	}
	*processedRanges = append(*processedRanges, [2]int64{start, end})
}
