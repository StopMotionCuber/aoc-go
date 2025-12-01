package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strconv"
	"strings"
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

	currentNum := 50
	result := 0

	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "L") {
			// get int after L
			numStr := strings.TrimPrefix(line, "L")
			num, _ := strconv.Atoi(numStr)
			currentNum -= num
		}
		if strings.HasPrefix(line, "R") {
			numStr := strings.TrimPrefix(line, "R")
			num, _ := strconv.Atoi(numStr)
			currentNum += num
		}
		if part2 {
			for currentNum < 0 {
				currentNum += 100
				result += 1
			}
			for currentNum >= 100 {
				currentNum -= 100
				result += 1
			}
			//println("Landed on ", currentNum, " result=", result)
		} else {
			currentNum = (currentNum + 100) % 100
			if currentNum == 0 {
				result += 1
			}
		}
	}
	// solve part 1 here
	return result
}
