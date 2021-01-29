package dp

import (
	"fmt"
	"math"
)

// source Cracking the Coding Interview; Question 8.3
// Magic Index: A magic index in an array A[ 1 ... n-1] is defined to be an index such that A[ i]
// i. Given a sorted array of distinct integers, write a method to find a magic index, if one exists, in
// array A.

func magicIndexIterative(A []int) int {
	if len(A) == 0 {
		return -1
	}
	for k := range A {
		if k == A[k] {
			return k
		}
	}
	return -1
}

func magicIndexRecursive(A []int) int {
	//return distinctValueArray(A, 0, len(A)-1)
	return repeatedValueArray(A, 0, len(A)-1)
}

func distinctValueArray(A []int, start, end int) int {
	// magic index not found
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if A[mid] == mid {
		return mid
	} else if A[mid] < mid {
		// search left
		return distinctValueArray(A, start, mid-1)
	}
	// search right
	return distinctValueArray(A, mid+1, end)
}

func repeatedValueArray(A []int, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if A[mid] == mid {
		return mid
	} else if A[mid] < mid {
		end = int(math.Min(float64(mid - 1), float64(A[mid])))
		return repeatedValueArray(A, start, end)
	}
	start = int(math.Max(float64(mid + 1), float64(A[mid])))
	return repeatedValueArray(A, start, end)
}

func MagicIndexExamples() {
	A := []int{1, 2, 3, 5, 5, 5, 7, 8, 8, 9}
	fmt.Println(magicIndexRecursive(A))
}
