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
	rawMaze := strings.Split(input, "\n")
	// Delete all empty lines from maze
	var maze [][]byte
	for _, line := range rawMaze {
		if strings.TrimSpace(line) != "" {
			maze = append(maze, []byte(line))
		}
	}
	if !part2 {
		return getAccessibleCount(maze, part2)
	}
	lastChange := -1
	sum := 0
	for lastChange != 0 {
		lastChange = getAccessibleCount(maze, part2)
		sum += lastChange
	}
	return sum
}

func getAccessibleCount(maze [][]byte, part2 bool) int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	full_count := 0
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] != '@' {
				continue
			}
			count := 0
			for _, dir := range directions {
				x := i + dir[0]
				y := j + dir[1]
				if x >= 0 && x < len(maze) && y >= 0 && y < len(maze[i]) {
					if maze[x][y] == '@' {
						count++
					}
				}
			}
			if count < 4 {
				full_count++
				if part2 {
					maze[i][j] = '.'
				}
			}
		}
	}
	return full_count
}
