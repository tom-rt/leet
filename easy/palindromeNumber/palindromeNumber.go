package main

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	j := len(str) - 1
	if j <= 0 {
		return true
	}
	for i := 0; i <= j; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
}
