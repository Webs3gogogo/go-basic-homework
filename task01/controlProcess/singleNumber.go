package main

func singleNumber(nums []int) int {
	tempMap := make(map[int]int)
	var result int
	for _, value := range nums {
		tempMap[value]++
	}
	for _, value := range nums {
		if tempMap[value] == 1 {
			result = value
		}
	}
	return result
}
