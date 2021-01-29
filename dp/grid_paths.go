package dp

import "fmt"

func IsPath(grid [][]int) bool {
	if len(grid) == 0 {
		return false
	}
	if isPathHelper(grid, 0, 0) {
		return true
	}
	return false
}

func isPathHelper(grid [][]int, row int, col int) bool {
	// operate within the grid
	if row > len(grid)-1 || col > len(grid[0])-1 {
		return false
	}
	// if we are at the last cell of the grid
	if row == len(grid)-1 && col == len(grid[0])-1 {
		return true
	}
	// if the path is okay up to the current point
	if isPathHelper(grid, row+1, col) || isPathHelper(grid, row, col+1) {
		return true
	}
	return false
}

// total number of ways to bottom right
func PathCountRecursive(grid [][]int, row, col int) int {
	// base case
	if row == len(grid)-1 && col == len(grid[0])-1 {
		return 1
	}
	// stay within the grid
	if row > len(grid)-1 || col > len(grid[0])-1 {
		return 0
	}
	// recursive case
	return PathCountRecursive(grid, row+1, col) + PathCountRecursive(grid, row, col+1)
}

/*
TODO
func allPathsRecursive(grid [][]int, row, col int, path [][]int) [][]int {
	// base case
	if row == len(grid) - 1 || col == len(grid[0]) - 1 {
		return path
	}
	//return allPathsRecursive(grid, row + 1, col) + allPathsRecursive(grid, row, col + 1)
	return 0
}*/

func GridPathExamples() {
	grid := [][]int{
		[]int{1, 1, 1, 1},
		[]int{1, 0, 0, 0},
	}
	//var path [][]int
	//fmt.Println(path)
	//fmt.Println(IsPath(grid))
	fmt.Println(PathCountRecursive(grid, 0, 0))
}
