package main

func plusOne(digits []int) []int {
	len := len(digits)
	var flow int = 0
	for i := len - 1; i >= 0; i-- {
		if i == len-1 {
			digits[i] += 1
		}
		digits[i] += flow
		flow = digits[i] / 10
		digits[i] %= 10
	}
	if flow == 1 {
		result := []int{1}
		return append(result, digits...)
	}
	return digits
}
