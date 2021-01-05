package main

import (
	"fmt"
	"math"
)

func coinChange2(amt int, coins []int) int {
	memo := generate2D(amt, len(coins))
	for i, row := range memo {
		for j, col := range row {
			if j > 0 && i == 0 {
				memo[i][j] = 1
			}
			if j == 0 {
				memo[i][j] = 0
			}
			if j > i {
				memo[i][j] = int(math.Max(
					float64(memo[i][j]),
					float64(memo[i][j - col] + 1)))
			}
		}
	}
	fmt.Println(memo)
	return 1
}

func generate2D(rows, columns int) [][]int {
	memo := make([][]int, 0)
	inner := make([]int, rows + 1)
	for i := 0; i <= columns; i++ {
		memo = append(memo, inner)
	}
	return memo
}

func main() {
	fmt.Println(coinChange2(5, []int{1,2,3}))
}
