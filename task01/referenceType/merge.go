package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	result := [][]int{}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	temp := intervals[0]
	for index, item := range intervals[0:] {
		if item[0] > temp[1] {
			result = append(result, temp)
			temp = item
			if index == len(intervals)-1 {
				result = append(result, temp)
			}
		} else {
			temp[1] = max(temp[1], item[1])
			if index == len(intervals)-1 {
				result = append(result, temp)
			}
		}
	}
	return result
}
