package main

func twoSum(nums []int, target int) []int {
	tempMap := make(map[int]int)
	for i, num := range nums {
		currentTarget := target - num
		index, exist := tempMap[currentTarget]
		if exist {
			return []int{index, i}
		}
		tempMap[num] = i
	}
	return []int{}
}
