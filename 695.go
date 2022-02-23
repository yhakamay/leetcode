package main

import "fmt"

func main() {
	//grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	grid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	fmt.Printf("maxAreaOfIsland(): %v\n", maxAreaOfIsland(grid))
}

func maxAreaOfIsland(grid [][]int) (max int) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 1 {
				if area := dfs(grid, x, y); area > max {
					max = area
				}
			}
		}
	}

	return
}

func dfs(grid [][]int, x, y int) (area int) {
	// area >= 1 because already found 1 in maxAreaOfIsland
	area = 1
	// Ignore the starting point to avoid duplicated count
	grid[x][y] = 0
	// Positions which we have to check this time
	nextPositions := [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}}

	for _, nextPosition := range nextPositions {
		nextX, nextY := nextPosition[0], nextPosition[1]

		if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) {
			if grid[nextX][nextY] == 1 {
				area += dfs(grid, nextX, nextY)
			}
		}
	}

	return
}
