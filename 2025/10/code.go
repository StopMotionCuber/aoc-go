package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
	"gonum.org/v1/gonum/stat/combin"
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
	if part2 {
		return "not implemented"
	}

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		// Split input into 3 parts
		objs := strings.Split(line, " ")
		onOffStr := objs[0][1 : len(objs[0])-1]
		onOff := make([]bool, len(onOffStr))
		for i, char := range onOffStr {
			onOff[i] = char == '#'
		}
		configs := make([][]uint8, 0)
		for i := 1; i < len(objs)-1; i++ {
			splitted := strings.Split(objs[i][1:len(objs[i])-1], ",")
			currentConfig := make([]uint8, 0)
			for _, val := range splitted {
				converted, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				currentConfig = append(currentConfig, uint8(converted))
			}
			configs = append(configs, currentConfig)
		}
		minPresses := getMinPresses(configs, onOff)
		println("Min presses:", minPresses)
		sum += minPresses
	}
	// solve part 1 here
	return sum
}

func getMinPresses(configs [][]uint8, onOff []bool) int {
	// Do DFS to find minimum presses
	onOffCopy := make([]bool, len(onOff))
	for i := 0; i < len(configs); i++ {
		permutationIdx := make([]int, i)
		generator := combin.NewPermutationGenerator(len(configs), i)
		for generator.Next() {
			copy(onOffCopy, onOff)
			presses := generator.Permutation(permutationIdx)
			for _, press := range presses {
				for _, idx := range configs[press] {
					onOffCopy[idx] = !onOffCopy[idx]
				}
			}
			// Check if all are on
			allOff := true
			for _, state := range onOffCopy {
				if state {
					allOff = false
					break
				}
			}
			if allOff {
				println("Found with presses:", i)
				return i
			}
		}
	}
	return -1
}
