package main

import "fmt"

func isSubsequence(a, b string) bool {
	if len(a) == 0 {
		return false
	}
	if len(b) == 0 {
		return false
	}
	return helper([]rune(a), []rune(b), 0, 0)
}

func helper(a, b []rune, l, r int) bool {
	if l == len(a) {
		return true
	}
	if r == len(b) {
		return false
	}
	if a[l] == b[r] {
		return helper(a, b, l+1, r+1)
	}
	return helper(a, b, l, r+1)
}

func main() {
	fmt.Println(isSubsequence("cc", "bbcacdbac"))
}
