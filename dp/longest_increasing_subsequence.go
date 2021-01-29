package dp

import (
	"fmt"
	"github.com/danielwangai/algorithms/dp/utils"
)

// given an array of integers, return the longest increasing subsequence
// of integers
// e.g given array = [10, 9, 55, 7, 98]
// return 3 since {10, 55, 98} is the longest increasing subsequence
func longestIncreasingSubsequence(arr []int) int {
	// array to store the result
	var res []int
	for i := 0; i < len(arr); i++ {
		res = append(res, 1)
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				res[i] = utils.Max(res[j] + 1, res[i])
			}
		}
	}
	max := 0
	for i := 0; i < len(res); i++ {
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}

func LISExamples() {
	fmt.Println(longestIncreasingSubsequence([]int{10, 9, 55, 7, 98}))
}
