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
	result := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		if part2 {
			curResult := getJoltsTwelve(line)
			println(curResult)
			result += int(getJoltsTwelve(line))
		} else {
			result += int(getJolts(line))
		}
	}
	return result
}

func getJoltsTwelve(input string) int64 {
	maxNum := int32(0)
	startIdx := 0
	sum := int64(0)

	for i := 11; i >= 0; i-- {
		intermediateIdx := 0
		for idx, char := range input[startIdx : len(input)-i] {
			if (char - '0') > maxNum {
				maxNum = char - '0'
				intermediateIdx = idx
			}
		}
		startIdx += intermediateIdx + 1
		sum *= 10
		sum += int64(maxNum)
		maxNum = 0
	}
	return sum
}

func getJolts(input string) int32 {
	maxNum := int32(0)
	maxIdx := -1
	// Find largest
	for idx, char := range input[:len(input)-1] {
		if (char - '0') > maxNum {
			maxNum = char - '0'
			maxIdx = idx
		}
	}
	sum := 10 * maxNum
	maxNum = 0
	// Find second largest
	for _, char := range input[maxIdx+1:] {
		if (char - '0') > maxNum {
			maxNum = char - '0'

		}
	}
	return sum + maxNum
}
