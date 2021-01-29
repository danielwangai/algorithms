package dp

import "fmt"

// n:- number of elements in the numbers array
// m:- the sum i.e. the height in the worst case i.e. when there exists number 1 as one of the elements in the
// numbers array
// Time Complexity: O(n^m * m)
// Space Complexity: O(m * m) or O(m^2): because in the worst case if there exists and element 1 in the nums array,
// and every recursive call has a potential to have a length of m
func bestSumRecursive(sum int, numbers []int) []int {
	if sum == 0 {
		return []int{}
	}
	if sum < 0 {
		return nil
	}
	var shortest []int
	for _, n := range numbers {
		rem := sum - n
		remCombo := bestSumRecursive(rem, numbers)
		if remCombo != nil {
			remCombo = append(remCombo, n)
			if shortest == nil || len(remCombo) < len(shortest) {
				shortest = remCombo
			}
		}
	}
	return shortest
}

// n:- number of elements in the numbers array
// m:- the sum i.e. the height in the worst case i.e. when there exists number 1 as one of the elements in the
// numbers array
// Time Complexity: O(m * n * m) or O(m^2 * n)
// Space Complexity: O(m * m) or O(m^2): because in the worst case if there exists and element 1 in the nums array,
// and every recursive call has a potential to have a length of m
func bestSumMemo(sum int, numbers []int, memo map[int][]int) []int {
	if _, ok := memo[sum]; ok {
		return memo[sum]
	}
	if sum == 0 {
		return []int{}
	}
	if sum < 0 {
		return nil
	}
	var shortest []int
	for _, n := range numbers {
		rem := sum - n
		remCombo := bestSumMemo(rem, numbers, memo)
		if remCombo != nil {
			remCombo = append(remCombo, n)
			if shortest == nil || len(remCombo) < len(shortest) {
				shortest = remCombo
			}
		}
	}
	memo[sum] = shortest
	return memo[sum]
}

func BestSumExamples() {
	// recursive
	fmt.Println(bestSumRecursive(7, []int{5, 4, 3, 7}))
	fmt.Println(bestSumRecursive(200, []int{7, 14}))

	// memoized
	fmt.Println(bestSumMemo(7, []int{5, 4, 3, 7}, map[int][]int{}))
	fmt.Println(bestSumMemo(200, []int{7, 14}, map[int][]int{}))
}
