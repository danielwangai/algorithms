package recursion

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	runeStr := []rune(s)
	return palindromeHelper(runeStr, 0, len(s)-1, true)
}

func palindromeHelper(runeStr []rune, l, r int, isPal bool) bool {
	if l >= r {
		return isPal
	}
	if runeStr[l] != runeStr[r] {
		return false
	}
	return palindromeHelper(runeStr, l+1, r-1, true)
}
