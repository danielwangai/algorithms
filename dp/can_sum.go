package dp

import "fmt"

// given an integer targetSum and an array nums of integers,
// write a program that returns if it is possible for any combination(s)
// of elements in nums to sum up to targetSum.
// NOTE: an element can be used multiple times

// n is the number of elements in the nums array
// m is the sum
// time complexity: O(n^m):- we can do a maximum difference from the sum (i.e. each node in the resultant tree) n times
// i.e. n branches until m i.e. the sum becomes 0 (or not in some cases)
// There can also ne repeated calculations of the same problem
// space complexity: O(m):- because in the worst case if there exists and element 1 in the nums array,
// in order reduce the sum value progressively to zero, this will be done m times
func CanSum(sum int, nums []int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 {
		return false
	}
	for _, v := range nums {
		rem := sum - v
		if CanSum(rem, nums) == true {
			return true
		}
	}
	return false
}

// n is the number of elements in the nums array
// m is the sum
// time complexity: O(n*m):- results of previously calculated problems are cached and reused to avoid repeat
// calculations
// space complexity: O(m):- because in the worst case if there exists and element 1 in the nums array,
// in order reduce the sum value progressively to zero, this will be done m times
func CanSumMemo(sum int, nums []int, memo map[int]bool) bool {
	if _, ok := memo[sum]; ok {
		return memo[sum]
	}
	if sum == 0 {
		return true
	}
	if sum < 0 {
		return false
	}
	for _, v := range nums {
		rem := sum - v
		if CanSumMemo(rem, nums, memo) == true {
			memo[rem] = true
			return true
		}
	}
	memo[sum] = false
	return false
}

func CanSumExamples() {
	// recursive
	fmt.Println(CanSum(7, []int{5, 4, 3, 7}))
	fmt.Println(CanSum(200, []int{7, 14}))

	// memoized
	fmt.Println(CanSumMemo(7, []int{5, 4, 3, 7}, map[int]bool{}))
	fmt.Println(CanSumMemo(200, []int{7, 14}, map[int]bool{}))
}
