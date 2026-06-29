package recursion

func isSubsequence(a, b string) bool {
	if len(a) == 0 {
		return false
	}
	if len(b) == 0 {
		return false
	}
	return subsequenceHelper([]rune(a), []rune(b), 0, 0)
}

func subsequenceHelper(a, b []rune, l, r int) bool {
	if l == len(a) {
		return true
	}
	if r == len(b) {
		return false
	}
	if a[l] == b[r] {
		return subsequenceHelper(a, b, l+1, r+1)
	}
	return subsequenceHelper(a, b, l, r+1)
}
