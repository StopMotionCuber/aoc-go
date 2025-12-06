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
	// solve part 1 here
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	operators := make([]string, len(lines))
	_ = operators

	numbers := make([][]int, len(lines)-1)
	for i, line := range lines {
		if i == len(lines)-1 {
			operators = strings.Fields(line)
		} else {
			numStrs := strings.Fields(line)
			numbers[i] = make([]int, len(numStrs))
			for j, numStr := range numStrs {
				numbers[i][j], _ = strconv.Atoi(numStr)
			}
		}
	}

	sum := 0

	if part2 {
		currentResult := 0
		number := 0
		currentOperator := lines[len(lines)-1][0]
		// Recalculate the numbers
		for i := 0; i < len(lines[0]); i++ {
			for j := 0; j < len(lines)-1; j++ {
				if lines[j][i] != ' ' {
					number = number*10 + int(lines[j][i]-'0')
				}
			}
			if currentResult == 0 {
				currentResult = number
			} else {
				if currentOperator == '+' {
					currentResult += number
				} else if number != 0 {
					currentResult *= number
				}
			}
			if (i >= len(lines[0])-1) || lines[len(lines)-1][i+1] != ' ' {
				// New operator
				if i != len(lines[0])-1 {
					currentOperator = lines[len(lines)-1][i+1]
				}
				sum += currentResult
				currentResult = 0
			}
			number = 0
		}
		return sum
	}

	for i, op := range operators {
		partResult := numbers[0][i]
		for j := 1; j < len(numbers); j++ {
			switch op {
			case "+":
				partResult += numbers[j][i]
			case "*":
				partResult *= numbers[j][i]
			}
		}
		sum += partResult
	}

	return sum
}
