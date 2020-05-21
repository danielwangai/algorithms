package main

import "fmt"

// recursive solution
// time complexity: O(2^n)
// 2 because of the double recursion i.e. fibRecursive(n - 1) and fibRecursive(n - 2)
// there's lots of repetitive calls
// space complexity: O(n) - from the call stack due to recursion
func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// recursion with memoization
// time complexity: O(n) - there are no repetitions due to memoization
// // space complexity: O(n) - from the call stack due to recursion
func fibMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}

// bottom up approach
// time complexity: O(n) - iterating to n
// space complexity: O(n) - because of the memoization data store
func bottomUp(n int, memo map[int]int) int {
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

// top down approach
// time complexity: O(n) for iterating to n
// space complexity: O(1) we're only storing the previous two numbers which are updated till we get to n
func topDown(n int) int {
	a := 0
	b := 1
	for i := 2; i < n; i++ {
		temp := b
		b = a + b
		a = temp
	}
	return a + b
}

func main() {
	fmt.Println(fibRecursive(6))
	fmt.Println(fibMemo(6, map[int]int{}))
	fmt.Println(bottomUp(6, map[int]int{}))
	fmt.Println(topDown(6))
}
