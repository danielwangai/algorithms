package utils

type Strings struct {
	Str1 string
	Str2 string
}

// helper method to initialize a 2D array
func Array2D(r, c int) [][]int {
	outer := make([][]int, 0)
	inner := make([]int, c)
	for i := 0; i < r; i++ {
		outer = append(outer, inner)
	}
	return outer
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
