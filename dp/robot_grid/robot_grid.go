package main

import "fmt"


// recursive approach
// time complexity: O(2^r + c)
//	- there are at most 2 options: go right or down
//	until we get to the end(bottom-right most cell)
// space complexity: O(r * c) - size of the grid
func robotGrid(grid [][]int) bool {
	/*return recursiveHelper(grid, 0, 0)*/
	memo := map[[2]int]bool{}
	return memoized(grid, 0, 0, memo)
}

func recursiveHelper(grid [][]int, r, c int) bool {
	if r > len(grid) - 1 || c > len(grid[0]) - 1 || grid[r][c] == 1 {
		return false
	}
	isAtEnd := r == len(grid) - 1 && c == len(grid[0]) - 1
	if !isAtEnd || recursiveHelper(grid, r + 1, c) || recursiveHelper(grid, r, c + 1) {
		return true
	}
	return false
}

// memoized solution
// time complexity: O(r*c)
//	 - calculations are memoized hence not repeated
// space complexity: O(r*c)
//   - size of the grid
func memoized(grid [][]int, r, c int, memo map[[2]int]bool) bool {
	if r > len(grid) - 1 || c > len(grid[0]) - 1 || grid[r][c] == 1 {
		return false
	}
	if memo[[2]int{r, c}] {
		return true
	}
	isAtEnd := r == len(grid) - 1 && c == len(grid[0]) - 1
	if !isAtEnd || memoized(grid, r + 1, c, memo) || memoized(grid, r, c + 1, memo) {
		memo[[2]int{r, c}] = true
		return true
	}
	return false
}

func main() {
	grid := [][]int{
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 0, 0},
		[]int{0, 0, 1, 0, 0, 1},
		[]int{0, 1, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0},
	}
	fmt.Println(robotGrid(grid))
}
