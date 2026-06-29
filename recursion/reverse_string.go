package recursion

func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := []rune(s)
	reversed := reverseHelper(s, res, 0, len(s)-1)
	return string(reversed)
}

func reverseHelper(s string, res []rune, l, r int) []rune {
	if l >= r {
		return res
	}
	res[l], res[r] = res[r], res[l]
	return reverseHelper(s, res, l+1, r-1)
}
