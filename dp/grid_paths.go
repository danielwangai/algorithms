package main

import "fmt"

func isPath(grid [][]int) bool {
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
func pathCountRecursive(grid [][]int, row, col int) int {
	// base case
	if row == len(grid)-1 && col == len(grid[0])-1 {
		return 1
	}
	// stay within the grid
	if row > len(grid)-1 || col > len(grid[0])-1 {
		return 0
	}
	// recursive case
	return pathCountRecursive(grid, row+1, col) + pathCountRecursive(grid, row, col+1)
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

func main() {
	grid := [][]int{
		[]int{1, 1, 1, 1},
		[]int{1, 0, 0, 0},
	}
	//var path [][]int
	//fmt.Println(path)
	//fmt.Println(isPath(grid))
	fmt.Println(pathCountRecursive(grid, 0, 0))
}
