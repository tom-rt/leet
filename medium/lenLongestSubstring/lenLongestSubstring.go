package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var max int = 1
	var sum int = 1
	if len(s) <= 1 {
		return len(s)
	}
	var sub string = string(s[0])
	for i := 1; i < len(s); i++ {
		if strings.Contains(sub, string(s[i])) {
			if sum > max {
				max = sum
			}
			sum = 1
			i = getPrevious(s, i-1, string(s[i]))
			sub = string(s[i])
		} else {
			sum++
			sub = sub + string(s[i])
		}
	}
	if sum > max {
		max = sum
	}
	return max
}

func getPrevious(s string, idx int, x string) int {
	for i := idx; i >= 0; i-- {
		if string(s[i]) == x {
			return i + 1
		}
	}
	return 0
}
