package main

import "sort"

func longestCommonPrefix(strs []string) string {
	slice := []rune{}
	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	minWordLength := len(strs[0])
	r := []rune(strs[0])
outer:
	for i := range minWordLength {
		currIndex := r[i]
		for _, str := range strs {
			runStr := []rune(str)
			if currIndex != runStr[i] {
				break outer
			}
		}
		slice = append(slice, currIndex)
	}
	return string(slice)
}
func longestCommonPrefixOfficialAns(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
