package dp

import (
	"fmt"
	"math"

	"github.com/danielwangai/algorithms/dp/utils"
)

func EditDistance(str1, str2 string) int {
	memo := utils.Array2D(len(str1)+1, len(str2)+1)
	return EditDistanceBottomUp(str1, str2, memo)
}

func EditDistanceBottomUp(str1, str2 string, memo [][]int) int {
	//fmt.Println(memo)
	for i := 0; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if i == 0 {
				memo[i][j] = j
			}
			if j == 0 {
				memo[i][j] = i
			}
			if string(str1[i-1]) == string(str2[j-1]) {
				memo[i][j] = memo[i-1][j-1]
			} else {
				memo[i][j] = 1 + int(math.Min(
					float64(memo[i][j-1]), // insert
					math.Min(
						float64(memo[i-1][j-1]), // replace
						float64(memo[i-1][j])))) // delete
			}
		}
	}
	fmt.Println(memo)
	return memo[len(str1)][len(str2)]
}

func EditDistanceRecursive(str1, str2 string, pos1, pos2, count int) int {
	//if pos1 == len(str1) || pos2 == len(str2) {
	//	return count + (len(str1) - pos1) + (len(str2) - pos2)
	//}
	if pos1 == len(str1) {
		return count + len(str2) - pos2
	}
	if pos2 == len(str2) {
		return count + len(str1) - pos1
	}
	//if string(str1[pos1]) == string(str2[pos2]) {
	//	return EditDistanceRecursive(str1, str2, pos1 + 1, pos2 + 2, count)
	//}
	oldCount := count
	count = EditDistanceRecursive(str1, str2, pos1, pos2+1, oldCount+1)                                         // add
	count = int(math.Min(float64(count), float64(EditDistanceRecursive(str1, str2, pos1+1, pos2, oldCount+1)))) // remove
	equalCount := 1
	if string(str1[pos1]) == string(str2[pos2]) {
		equalCount = 0
	}
	count = int(math.Min(float64(count), float64(EditDistanceRecursive(str1, str2, pos1+1, pos2+2, oldCount+equalCount)))) // replace or equal
	return count
}

/*
func EditDistanceRecursive(str1, str2 string, pos1, pos2, count int) int {
	if pos1 == len(str1) - 1 {
		return count + len(str2) - pos2
	}
	if pos2 == len(str2) - 1 {
		return count + len(str1) - pos1
	}
	oldCount := count
	count = EditDistanceRecursive(str1, str2, pos1 + 1, pos2 + 1, oldCount + 1)
	count = int(math.Min(float64(count), float64(EditDistanceRecursive(str1, str2, pos1 + 1, pos2, oldCount + 1))))
	equalCount := 1
	if string(str1[pos1]) == string(str2[pos2]) {
		equalCount = 0
	}
	count = int(math.Min(float64(count), float64(EditDistanceRecursive(str1, str2, pos1 + 1, pos2 + 2, oldCount + equalCount))))
	return count
}
*/

/*
func EditDistanceRecursive(str1, str2 string, pos1, pos2, count int) int {
	 fmt.Println(string(str1[pos1]), string(str2[pos2]))
	if pos1 == len(str1)  || pos2 == len(str2) {
		return (len(str1) - pos1) + (len(str2) - pos2)
	}
	//if pos1 >= len(str1) {
	//	return len(str2) - pos2
	//}
	//if pos2 >= len(str2) {
	//	return len(str1) - pos1
	//}
	if string(str1[pos1]) == string(str2[pos2]) {
		// strings are similar - increment counter
		fmt.Println("3: Similar")
		return EditDistanceRecursive(str1, str2, pos1 + 1, pos2 + 1, count)
		//return EditDistanceRecursive(str1, str2, pos1, pos2)
	}
	insert := EditDistanceRecursive(str1, str2, pos1, pos2 + 1, count + 1)
	delete := EditDistanceRecursive(str1, str2, pos1 + 1, pos2, count + 1)
	modify := EditDistanceRecursive(str1, str2, pos1 + 1, pos2 + 1, count + 1)

	//if pos2 > pos1 {
	//	// insert
	//	fmt.Println("HERE 1")
	//	return 1 + EditDistanceRecursive(str1, str2, pos1, pos2 + 1)
	//}
	//if pos1 > pos2 {
	//	// insert
	//	fmt.Println("HERE 2")
	//	return 1 + EditDistanceRecursive(str1, str2, pos1 + 1, pos2)
	//}


	if string(str1[pos1]) == string(str2[pos2]) {
		// strings are similar - increment counter
		fmt.Println("3: Similar")
		return EditDistanceRecursive(str1, str2, pos1 + 1, pos2 + 1)
		//return EditDistanceRecursive(str1, str2, pos1, pos2)
	} else {
		//modify
		fmt.Println("4: Not Similar")
		return 1 + EditDistanceRecursive(str1, str2, pos1 + 1, pos2 + 1)
	}

	return int(
		math.Min(
			float64(insert), math.Min(float64(delete), float64(modify))))
}
*/

func EditDistanceExamples() {
	str1 := "cart"
	str2 := "march"
	//fmt.Println(">> ", EditDistanceRecursive(str1, str2, 6, 6, 0))
	EditDistance(str1, str2)
}
