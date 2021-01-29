package dp

import "fmt"

/*
Source: Cracking the Coding Interview;

Q 8.1 - Triple Step: A child is running up a staircase with n steps and can hop either 1 step, 2 steps, or 3
steps at a time. Implement a method to count how many possible ways the child can run up the
stairs.
*/

// recursive solution
// time complexity: O(3^n)
// 3 because of the triple recursion i.e. tripleStepRecursive(n - 1), tripleStepRecursive(n - 2), tripleStepRecursive(n - 3
// there's lots of repetitive calls
// space complexity: O(n) - from the call stack due to recursion
func tripleStepRecursive(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return tripleStepRecursive(n-1) + tripleStepRecursive(n-2) + tripleStepRecursive(n-3)
}

// recursion with memoization
// time complexity: O(n) - there are no repetitions due to memoization
// // space complexity: O(n) - from the call stack due to recursion
func tripleStepMemo(n int, memo map[int]int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		memo[0] = 0
		return 1
	}
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	memo[n] = tripleStepRecursive(n-1) + tripleStepRecursive(n-2) + tripleStepRecursive(n-3)
	return memo[n]
}

// bottom up approach
// time complexity: O(n) - iterating to n
// space complexity: O(n) - because of the memoization data store
func bottomUp(n int, memo map[int]int) int {
	memo[0] = 1
	memo[1] = 1
	memo[2] = 2
	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2] + memo[i-3]
	}
	return memo[n]
}

func TripleStepExamples() {
	fmt.Println(tripleStepRecursive(3))
	fmt.Println(tripleStepMemo(3, map[int]int{}))
	fmt.Println(bottomUp(3, map[int]int{}))
}
