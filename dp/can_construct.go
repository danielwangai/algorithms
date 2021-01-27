package main

import (
	"fmt"
	"strings"
)

// given a target string and an array of strings
// check if it is possible to use an unlimited number of elements in the array
// to form the target string

// let m be the length of the target string
// let n be the length of the wordBank array
// Time Complexity: O(n^m * m):-
// n^m because:- there's a possibility of making the suffix of the target string
// using all string in the word bank array i.e. if all elements are suffixes of target i.e.
// n branches possible
// O(n^m * m) because of the cost of calculating the suffix of target since we'd have to iterate
// to where the prefix in the target word ends so as to get the suffix
// Space Complexity: O(m*m) i.e. O(m^2) because if the wordBank array contains single character elements
// that build up to form the target string, the resultant tree formed will be of height m i.e. the length
// of the string. Also, each of the m stack frames (in the call stack) will have to store the created
// suffix string of possibly length m in the worst hence the space complexity of O(m^2)
func canConstruct(target string, wordBank []string) bool {
	if target == "" {
		return true
	}
	for _, w := range wordBank {
		if strings.HasPrefix(target, w) == true {
			// use a word in the bank if possible
			// i.e. if the word is a prefix of the target word
			remTarget := strings.Split(target, w)
			if len(remTarget) > 1 {
				suffix := strings.Join(remTarget[1:], "")
				if canConstruct(suffix, wordBank) == true {
					return true
				}
			}
		}
	}
	return false
}

// Time Complexity: O(n*m^2)
// Space Complexity: O(m^2)
func canConstructMemo(target string, wordBank []string, memo map[string]bool) bool {
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == "" {
		return true
	}
	for _, w := range wordBank {
		if strings.HasPrefix(target, w) == true {
			remTarget := strings.Split(target, w)
			if len(remTarget) > 1 {
				suffix := strings.Join(remTarget[1:], "")
				if canConstructMemo(suffix, wordBank, memo) == true {
					memo[suffix] = true
					return memo[suffix]
				}
			}
		}
	}
	return false
}

func main() {
	// recursive
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))

	// memoized
	fmt.Println(canConstructMemo("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string]bool{}))
	fmt.Println(canConstructMemo("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]bool{}))
	fmt.Println(canConstructMemo("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, map[string]bool{}))
}
