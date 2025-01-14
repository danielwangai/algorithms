package leetcode

// https://leetcode.com/problems/valid-parentheses
func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	braces := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	_ = braces
	stack := []string{string(s[0])}
	_ = stack

	for i := 1; i < len(s); i++ {
		// fmt.Printf("%d: %s\n", i, string(s[i]))
		item := string(s[i])
		if _, ok := braces[item]; ok {
			stack = append(stack, item)
			// continue
		} else {
			// item is not an opening bracket type

			if len(stack) == 0 {
				return false
			}
			// if item closes the item at top of stack, pop stack
			if braces[stack[len(stack)-1]] == item {
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
