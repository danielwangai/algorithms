package main

import "fmt"

// get total number of moves from top left to bottom right on a 2D grid

func totalMoves(grid [][]int) int {
	return recursive(grid, 0, 0)
	//return totalMovesMemoized(grid, 0, 0, map[[2]int]bool{})
}

func recursive(grid [][]int, r, c int) int {
	if r == len(grid) - 1 || c == len(grid[0]) - 1 {
		return 1
	}
	return recursive(grid, r + 1, c) + recursive(grid, r, c + 1)
}

// TODO - not working as expected
// func totalMovesMemoized(grid [][]int, r, c int, memo map[[2]int]bool) int {
//	if r == len(grid) - 1 || c == len(grid[0]) - 1 {
//		return 1
//	}
//	if _, ok := memo[[2]int{r, c}]; ok {
//		return 1
//	} else {
//		memo[[2]int{r, c}] = true
//	}
//	return totalMovesMemoized(grid, r + 1, c, memo) + totalMovesMemoized(grid, r, c + 1, memo)
//}

func main() {
	grid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}
	fmt.Println(totalMoves(grid))
}
