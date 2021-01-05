package main

import (
	"fmt"
	"math"

	"github.com/danielwangai/dp/utils"
)

func lcs(str1, str2 string) int {
	// recursive
	// return lcsRecursive(str1, str2, 0, 0)

	// memoized
	s := utils.Strings{
		Str1: str1,
		Str2: str2,
	}
	memo := s.Array2D(len(str1), len(str2))
	return lcsMemo(str1, str2, 0, 0, memo)

	// bottom up approach - TODO
	/*memo := initMemo(len(str1) + 1, len(str2) + 1)
	return lcsBottomUp(str1, str2, memo)*/
}

// recursive approach
func lcsRecursive(str1, str2 string, pos1, pos2 int) int {
	if pos1 >= len(str1)-1 || pos2 >= len(str2)-1 {
		return 0
	}
	if string(str1[pos1]) == string(str2[pos2]) {
		return 1 + lcsRecursive(str1, str2, pos1+1, pos2+1)
	}
	return int(
		math.Max(
			float64(lcsRecursive(str1, str2, pos1+1, pos2)),
			float64(lcsRecursive(str1, str2, pos1+1, pos2))))
}

func lcsMemo(len1, len2 string, pos1, pos2 int, memo [][]int) int {
	if memo[pos1][pos2] != 0 {
		return memo[pos1][pos2]
	}
	if pos1 >= len(len1)-1 || pos2 >= len(len2)-1 {
		return 0
	}
	if string(len1[pos1]) == string(len2[pos2]) {
		memo[pos1][pos2] = 1 + lcsMemo(len1, len2, pos1+1, pos2+1, memo)
		return memo[pos1][pos2]
	}
	memo[pos1][pos2] = int(
		math.Max(
			float64(lcsMemo(len1, len2, pos1+1, pos2, memo)),
			float64(lcsMemo(len1, len2, pos1+1, pos2, memo))))
	return memo[pos1][pos2]
}

/*
TODO
func lcsBottomUp(str1, str2 string, memo [][]int) int {
	//fmt.Println(len(memo), len(memo[0]))
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}
	l1 := len(str1) + 1
	l2 := len(str2) + 1
	for row := 1; row <= l1; row++ {
		fmt.Println("Row: ", row)
		for col := 1; col <= l2; col++ {
			if string(str1[row - 1]) == string(str2[col - 1]) {
				memo[row][col] = memo[row - 1][col - 1] + 1
			} else {
				memo[row][col] = int(
					math.Max(
						float64(memo[row - 1][col]),
						float64(memo[row][col - 1])))
			}
		}
	}
	//fmt.Println(">>>: ", len(str2), len(str1))
	//return 1
	return memo[len(str1)][len(str2)]
}
*/

func main() {
	//fmt.Println(lcs("abaacfdee", "aaabceed"))
	//var memo map[[2]int]int
	fmt.Println(lcs("abaac", "aabaa"))
}
