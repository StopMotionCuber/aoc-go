package main

import (
	"fmt"
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
}

type Direction int

const (
	N Direction = iota
	E
	S
	W
	NW
	NE
	SE
	SW
)

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
	// Rightmost, lowest point
	rightmostIdx := -1
	rightmostX := -1
	rightmostY := -1
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, ",")
		if len(split) != 2 {
			continue
		}
		x, _ := strconv.ParseInt(split[0], 10, 64)
		y, _ := strconv.ParseInt(split[1], 10, 64)
		allPoints = append(allPoints, Point{
			x: x,
			y: y,
		})
		if int(x) > rightmostX || (int(x) == rightmostX && int(y) > rightmostY) {
			rightmostX = int(x)
			rightmostY = int(y)
			rightmostIdx = len(allPoints) - 1
		}
	}
	println(len(allPoints))

	var openDirections = make(map[int]map[Direction]bool)

	if part2 {
		LineOpenSide := W
		for i := 0; i < len(allPoints); i++ {
			// Shift i so that we start with rightmost

			shiftedIdx := (i + rightmostIdx) % len(allPoints)
			pointBefore := allPoints[(shiftedIdx-1+len(allPoints))%len(allPoints)]
			pointAfter := allPoints[(shiftedIdx+1)%len(allPoints)]
			currentPoint := allPoints[shiftedIdx]
			if i == 0 {
				openDirections[shiftedIdx] = make(map[Direction]bool)
				openDirections[shiftedIdx] = map[Direction]bool{
					NW: true,
					NE: false,
					SE: false,
					SW: false,
				}
				if pointAfter.x == currentPoint.x {
					LineOpenSide = W
				} else {
					LineOpenSide = N
				}
				continue
			}
			if LineOpenSide == W {
				// Determine the next line open side
				// PointAfter y should be same as currentPoint y
				if pointAfter.x < currentPoint.x {
					// Going left
					if pointBefore.y < currentPoint.y {
						// Going left, went down before
						LineOpenSide = N
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: false,
							SE: false,
							NW: true,
							NE: false,
						}
					} else {
						// Going left, went up before
						LineOpenSide = S
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: true,
							SE: false,
							NW: false,
							NE: false,
						}
					}
				} else {
					// Going right
					if pointBefore.y < currentPoint.y {
						// Going right, went down before
						LineOpenSide = S
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: true,
							SE: true,
							NW: true,
							NE: false,
						}
					} else {
						// Going right, went up before
						LineOpenSide = N
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: true,
							SE: false,
							NW: true,
							NE: true,
						}
					}
				}
			} else if LineOpenSide == E {
				if pointAfter.x < currentPoint.x {
					// Going left
					if pointBefore.y < currentPoint.y {
						// Going left, went down before
						LineOpenSide = S
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: true,
							SE: true,
							NW: false,
							NE: true,
						}
					} else {
						// Going left, went up before
						LineOpenSide = N
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: false,
							SE: true,
							NW: true,
							NE: true,
						}
					}
				} else {
					// Going right
					if pointBefore.y < currentPoint.y {
						// Going right, went down before
						LineOpenSide = N
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: false,
							SE: false,
							NW: false,
							NE: true,
						}
					} else {
						// Going right, went up before
						LineOpenSide = S
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: false,
							SE: true,
							NW: false,
							NE: false,
						}
					}
				}
			} else if LineOpenSide == N {
				if pointAfter.y > currentPoint.y {
					// Going down
					if pointBefore.x > currentPoint.x {
						// Going down, went left before
						LineOpenSide = W
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: true,
							SE: false,
							NW: true,
							NE: true,
						}
					} else {
						// Going down, went right before
						LineOpenSide = E
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: false,
							SE: true,
							NW: true,
							NE: true,
						}
					}
				} else {
					// Going up
					if pointBefore.x > currentPoint.x {
						// Going up, went left before
						if len(allPoints) < 50 {
							println("Going up, went left before")
						}
						LineOpenSide = E
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: false,
							SE: false,
							NW: false,
							NE: true,
						}
					} else {
						// Going up, went right before
						LineOpenSide = W
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: false,
							SE: false,
							NW: true,
							NE: false,
						}
					}
				}
			} else if LineOpenSide == S {
				if pointAfter.y > currentPoint.y {
					// Going down
					if pointBefore.x > currentPoint.x {
						// Going down, went left before
						LineOpenSide = E
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: false,
							SE: true,
							NW: false,
							NE: false,
						}
					} else {
						// Going down, went right before
						LineOpenSide = W
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: true,
							SE: false,
							NW: false,
							NE: false,
						}
					}
				} else {
					// Going up
					if pointBefore.x > currentPoint.x {
						// Going up, went left before
						if len(allPoints) < 50 {
							println("Going up, went left before")
						}
						LineOpenSide = W
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: true,
							SE: true,
							NW: true,
							NE: false,
						}
					} else {
						// Going up, went right before
						LineOpenSide = E
						openDirections[shiftedIdx] = map[Direction]bool{
							SW: true,
							SE: true,
							NW: false,
							NE: true,
						}
					}
				}
			}
			if len(allPoints) < 50 {
				println(fmt.Sprintf("Point %d (%d,%d) open sides: NW:%t NE:%t SE:%t SW:%t", shiftedIdx, currentPoint.x, currentPoint.y, openDirections[shiftedIdx][NW], openDirections[shiftedIdx][NE], openDirections[shiftedIdx][SE], openDirections[shiftedIdx][SW]))
			}
		}
	}

	maxArea := int64(0)
	for i := 0; i < len(allPoints); i++ {
	inner:
		for j := i + 1; j < len(allPoints); j++ {
			xDist := allPoints[i].x - allPoints[j].x
			if xDist < 0 {
				xDist = -xDist
			}
			xDist += 1
			yDist := allPoints[i].y - allPoints[j].y
			if yDist < 0 {
				yDist = -yDist
			}
			yDist += 1
			area := xDist * yDist
			if area > maxArea {
				// Check whether there is any point in between for part 2
				if part2 {
					minX := min(allPoints[i].x, allPoints[j].x)
					maxX := max(allPoints[i].x, allPoints[j].x)
					minY := min(allPoints[i].y, allPoints[j].y)
					maxY := max(allPoints[i].y, allPoints[j].y)
					for k := 0; k < len(allPoints); k++ {
						nextPoint := allPoints[(k+1)%len(allPoints)]
						// Check if point k is inside the area
						if allPoints[k].x > minX && allPoints[k].y > minY &&
							allPoints[k].x < maxX && allPoints[k].y < maxY {
							continue inner
						}
						if allPoints[k].x == nextPoint.x && allPoints[k].x > minX && allPoints[k].x < maxX {
							// Might be crossing
							if min(allPoints[k].y, nextPoint.y) < minY && max(allPoints[k].y, allPoints[k].y) > maxY {
								continue inner
							}
						}
						if allPoints[k].y == nextPoint.y && allPoints[k].y > minY && allPoints[k].y < maxY {
							// Might be crossing
							if min(allPoints[k].x, nextPoint.x) < minX && max(allPoints[k].x, allPoints[k].x) > maxX {
								continue inner
							}
						}

						if allPoints[k].x == maxX {
							if allPoints[k].y >= minY && allPoints[k].y <= maxY {
								// On the edge
								// For corner we need to check 1 direction
								if allPoints[k].y == minY {
									// Top right corner
									if !openDirections[k][SW] {
										continue inner
									}
								} else if allPoints[k].y == maxY {
									// Bottom right corner
									if !openDirections[k][NW] {
										continue inner
									}
								} else {
									// Right edge, check both NW and SW
									if !openDirections[k][NW] || !openDirections[k][SW] {
										continue inner
									}
								}
							}
						}
						if allPoints[k].x == minX {
							if allPoints[k].y >= minY && allPoints[k].y <= maxY {
								// On the edge
								// For corner we need to check 1 direction
								if allPoints[k].y == minY {
									// Top left corner
									if !openDirections[k][SE] {
										continue inner
									}
								} else if allPoints[k].y == maxY {
									// Bottom left corner
									if !openDirections[k][NE] {
										continue inner
									}
								} else {
									// Left edge, check both NE and SE
									if !openDirections[k][NE] || !openDirections[k][SE] {
										continue inner
									}
								}
							}
						}
						if allPoints[k].y == maxY {
							if allPoints[k].x >= minX && allPoints[k].x <= maxX {
								// On the edge
								// For corner we need to check 1 direction
								if allPoints[k].x == minX {
									// Bottom left corner
									if !openDirections[k][NE] {
										continue inner
									}
								} else if allPoints[k].x == maxX {
									// Bottom right corner
									if !openDirections[k][NW] {
										continue inner
									}
								} else {
									// Bottom edge, check both NW and NE
									if !openDirections[k][NW] || !openDirections[k][NE] {
										continue inner
									}
								}
							}
						}
						if allPoints[k].y == minY {
							if allPoints[k].x >= minX && allPoints[k].x <= maxX {
								// On the edge
								// For corner we need to check 1 direction
								if allPoints[k].x == minX {
									// Top left corner
									if !openDirections[k][SE] {
										continue inner
									}
								} else if allPoints[k].x == maxX {
									// Top right corner
									if !openDirections[k][SW] {
										continue inner
									}
								} else {
									// Top edge, check both SW and SE
									if !openDirections[k][SW] || !openDirections[k][SE] {
										continue inner
									}
								}
							}
						}
					}
					// We also need to check if any line crosses the area
				}
				maxArea = area
			}
		}
	}

	return maxArea
}
