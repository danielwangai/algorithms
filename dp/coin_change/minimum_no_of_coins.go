package main

import (
	"fmt"
	"math"
)

func minimumCoinsUsed(amt int, coins []int) int {
	//return bottomUp(amt, coins)
	return len(minCoinsRecursive(amt, coins))
}

func minCoinsRecursive(amt int, coins []int) []int {
	if amt == 0 {
		return []int{}
	}
	if amt < 0 {
		return nil
	}
	var count []int
	for _, c := range coins {
		rem := amt - c
		remCombo := minCoinsRecursive(rem, coins)
		if remCombo != nil {
			remCombo = append(remCombo, c)
			if count == nil || len(remCombo) < len(count) {
				count = remCombo
			}
		}
	}
	return count
}

func minCoinsMemoized(amt int, coins []int, memo map[int][]int) []int {
	if _, ok := memo[amt]; ok {
		return memo[amt]
	}
	if amt == 0 {
		return []int{}
	}
	if amt < 0 {
		return nil
	}
	var minCoins []int
	for _, c := range coins {
		rem := amt - c
		remCombo := minCoinsMemoized(rem, coins, memo)
		if remCombo != nil {
			remCombo = append(remCombo, c)
			if minCoins == nil || len(remCombo) < len(minCoins) {
				minCoins = remCombo
			}
		}
	}
	return minCoins
}

func bottomUp(amt int, coins []int) int {
	memo := generateArray(amt)
	for i := 1; i <= amt; i++ {
		for _, coin := range coins {
			if coin > i {
				continue
			}
			memo[i] = int(math.Min(
				float64(memo[i - coin] + 1),
				float64(memo[i])))
		}
	}
	// denomination of coins can't make change
	// e.g coins = [2], amount = 3
	if amt < memo[amt] {
		return -1
	}
	return memo[amt]
}

/*
TODO
func topDown(amt int, coins []int) int {
	if amt == 0 {
		return 0
	}
	memo := generateArray(amt)
	return topDownHelper(amt, coins, memo)
}

func topDownHelper(amt int, coins, memo []int) int {

}
*/

func generateArray(n int) []int {
	var memo []int
	for i := 0; i <= n; i++ {
		if i == 0 {
			memo = append(memo, 0)
			continue
		}
		memo = append(memo, n + 1)
	}
	return memo
}

func main() {
	fmt.Println(minimumCoinsUsed(11, []int{5, 2, 1}))
	fmt.Println(minCoinsRecursive(7, []int{5, 4, 3, 7}))
	fmt.Println(minCoinsMemoized(7, []int{5, 4, 3, 7}, map[int][]int{}))
}
