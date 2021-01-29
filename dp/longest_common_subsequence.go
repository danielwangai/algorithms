package dp

import (
	"fmt"
	"github.com/danielwangai/algorithms/dp/utils"
	"math"

	//"github.com/danielwangai/dp/utils"
)

func lcs(str1, str2 string) int {
	// recursive
	return lcsRecursive(str1, str2, 0, 0, 0)

	// memoized
	/*
	s := utils.Strings{
		Str1: str1,
		Str2: str2,
	}
	memo := s.Array2D(len(str1), len(str2))
	return lcsMemo(str1, str2, 0, 0, memo)
	 */

	// bottom up approach - TODO
	/*memo := initMemo(len(str1) + 1, len(str2) + 1)
	return lcsBottomUp(str1, str2, memo)*/
}

// recursive approach
func lcsRecursive(str1, str2 string, pos1, pos2, res int) int {
	if pos1 == len(str1) || pos2 == len(str2) {
		return res
	}
	if string(str1[pos1]) == string(str2[pos2]) {
		res = 1 + lcsRecursive(str1, str2, pos1+1, pos2+1, res)
	} else {
		res =  int(
			math.Max(
			float64(lcsRecursive(str1, str2, pos1+1, pos2, res)),
			float64(lcsRecursive(str1, str2, pos1, pos2+1, res))))
	}
	return res
}

func lcsMemo(str1, str2 string, pos1, pos2, res int, memo map[[2]int]int) int {
	if pos1 == len(str1) || pos2 == len(str2) {
		memo[[2]int{pos1, pos2}] = res
		return memo[[2]int{pos1, pos2}]
	}
	if string(str1[pos1]) == string(str2[pos2]) {
		res = 1 + lcsMemo(str1, str2, pos1+1, pos2+1, res, memo)
	} else {
		res =  int(
			math.Max(
				float64(lcsMemo(str1, str2, pos1+1, pos2, res, memo)),
				float64(lcsMemo(str1, str2, pos1, pos2+1, res, memo))))
	}
	memo[[2]int{pos1, pos2}] = res
	return memo[[2]int{pos1, pos2}]
}

func lcsBottomUp(str1, str2 string) int {
	//fmt.Println(len(memo), len(memo[0]))
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}
	memo := utils.Array2D(len(str1) + 1, len(str2) + 1)
	fmt.Println("len: ", memo)
	longest := 0
	for r := 1; r <= len(memo) - 1; r++ {
		for c := 1; c <= len(memo[0]) - 1; c++ {
			fmt.Println("r - c: ", r, c)
			if string(str1[r - 1]) == string(str2[c - 1]) {
				memo[r][c] = memo[r - 1][c - 1] + 1
			} else {
				memo[r][c] = utils.Max(memo[r][c - 1], memo[r - 1][c])
			}

			if memo[r][c] > longest {
				longest = memo[r][c]
			}
		}
	}
	fmt.Println("bottoms up: ", memo)
	return longest
	/*l1 := len(str1) + 1
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
	 */
	//fmt.Println(">>>: ", len(str2), len(str1))
	//return 1
	//return memo[len(str1)][len(str2)]
}


func LCSExample() {
	// recursive
	str1 := "abaac"
	str2 := "aabaa"
	fmt.Println(lcsRecursive(str1, str2, 0, 0, 0))
	// memoized - DP
	//s := utils.Strings{Str1: str1, Str2: str2}
	//var memo map[[2]int]int
	memo := map[[2]int]int{}
	//memo := s.Array2D(-1)
	fmt.Println("Memo: ", lcsMemo(str1, str2, 0, 0, 0, memo))
	fmt.Println("Memo: ", memo)
	//fmt.Println(string(str1[0]))
	//lcsBottomUp(str1, str2)
}
