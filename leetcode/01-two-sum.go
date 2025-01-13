package leetcode

func twoSum(nums []int, target int) []int {
	return twoSum1(nums, target)
	//return twoSum2(nums, target)
}

// Time: O(n²)
// Space: O(n)
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// use memoization
// Time: O(n)
// Space: O(n)
func twoSum2(nums []int, target int) []int {
	// key: element, value: index
	memo := map[int]int{}
	for i := 0; i < len(nums); i++ {
		// store all numbers in the map
		memo[nums[i]] = i
	}

	for i, num := range nums {
		secondNum := target - num
		if x, ok := memo[secondNum]; ok && x != i {
			return []int{x, i}
		}
	}

	return []int{}
}
