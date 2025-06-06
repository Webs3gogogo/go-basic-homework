package main

import "strconv"

func isPalindrome(x int) bool {
	numberStr := strconv.Itoa(x)
	for p1, p2 := 0, len(numberStr)-1; p1 < p2; p1, p2 = p1+1, p2-1 {
		if numberStr[p1] != numberStr[p2] {
			return false
		}

	}
	return true
}
