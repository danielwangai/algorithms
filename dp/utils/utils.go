package utils

type Strings struct {
	Str1 string
	Str2 string
}

// helper method to initialize a 2D array
func (s *Strings) Array2D(len1, len2 int) [][]int {
	outer := make([][]int, 0)
	inner := make([]int, len2)
	for i := 0; i < len1; i++ {
		outer = append(outer, inner)
	}
	return outer
}
