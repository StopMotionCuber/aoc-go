package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"math"
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
	// solve part 1 here
	input = strings.Replace(input, "\n", "", -1)

	result := 0

	for _, numberRange := range strings.Split(input, ",") {
		numbers := strings.Split(numberRange, "-")
		if len(numbers) != 2 {
			continue
		}
		first, _ := strconv.Atoi(numbers[0])
		last, _ := strconv.Atoi(numbers[1])
		if part2 {
			for _, id := range getInvalidIDsTaskTwo(first, last) {
				result += id
			}
		} else {
			for _, id := range getInvalidIDsTaskOne(first, last) {
				result += id
			}
		}
	}

	return result
}

func getInvalidIDsTaskOne(start int, end int) []int {
	// Get digitsStart and digitsEnd
	// We need to iterate through all even numbers for digitsStart and digitsEnd
	digitsStart := int(math.Floor(math.Log10(float64(start))) + 1)
	digitsEnd := int(math.Floor(math.Log10(float64(end))) + 1)
	var toReturn []int

	for i := digitsStart; i <= digitsEnd; i++ {
		// We only care for even numbers
		if i%2 == 1 {
			continue
		}
		currentDigits := i / 2
		// We want to start taking a look at 10^i, unless we're at digitsStart
		j := int(math.Pow(10, float64(currentDigits-1)))
		// j should iterate till max numbers
		jMax := j*10 - 1
		if i == digitsEnd {
			jMax = end / j / 10
		}

		if i == digitsStart {
			j = start / j / 10
		}

		for ; j <= jMax; j++ {
			nextNum := j*int(math.Pow(float64(10), float64(currentDigits))) + j
			if nextNum >= start && nextNum <= end {
				toReturn = append(toReturn, nextNum)
			}
		}

	}
	return toReturn
}

func getInvalidIDsTaskTwo(start int, end int) []int {
	// Get digitsStart and digitsEnd
	// We need to iterate through all even numbers for digitsStart and digitsEnd
	digitsStart := int(math.Floor(math.Log10(float64(start))) + 1)
	digitsEnd := int(math.Floor(math.Log10(float64(end))) + 1)
	toReturn := map[int]bool{}

	for k := 2; k <= digitsEnd; k++ {
		// k is the amount of repititions that we want
		for i := digitsStart; i <= digitsEnd; i++ {
			// We only care for numbers dividable by k
			if i%k != 0 {
				continue
			}
			currentDigits := i / k
			// We want to start taking a look at 10^i, unless we're at digitsStart
			j := int(math.Pow(10, float64(currentDigits-1)))
			// j should iterate till max numbers
			jMax := j*10 - 1
			if i == digitsEnd {
				// jMax should be the first currentDigits digits of end
				jMax = end / int(math.Pow(10, float64(digitsEnd-currentDigits)))
			}

			if i == digitsStart {
				// jMax should be the first currentDigits digits of start
				j = start / int(math.Pow(10, float64(digitsStart-currentDigits)))
			}
			print(fmt.Sprintf("Repititions=%d, start=%d, end=%d\n", k, j, jMax))

			for ; j <= jMax; j++ {
				nextNum := j
				for l := 1; l < k; l++ {
					nextNum = nextNum*int(math.Pow(float64(10), float64(currentDigits))) + j
				}
				if nextNum >= start && nextNum <= end {
					toReturn[nextNum] = true
				}
			}

		}

	}
	// return all toReturn keys as a list
	resultSlice := make([]int, 0, len(toReturn))
	for k := range toReturn {
		resultSlice = append(resultSlice, k)
	}
	return resultSlice
}
