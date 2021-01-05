package main

import (
	"fmt"
	"math"
)

// given an array of integers, return the longest increasing subsequence
// of integers
// e.g given array = [10, 9, 55, 7, 98]
// return 3 since {10, 55, 98} is the longest increasing subsequence
func longestIncreasingSubsequence(arr []int) int {
	// array to store the result
	var lis []int
	for i := 0; i < len(arr); i++ {
		lis = append(lis, 1)
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				sum := lis[j] + 1
				lis[i] = int(math.Max(float64(sum), float64(lis[i])))
			}
		}
	}
	max := 0
	for i := 0; i < len(lis); i++ {
		if lis[i] > max {
			max = lis[i]
		}
	}
	return max
}

func main() {
	fmt.Println(longestIncreasingSubsequence([]int{10, 9, 55, 7, 98}))
}
