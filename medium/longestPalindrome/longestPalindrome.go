package main

func longestPalindrome(s string) string {
	var ret string
	var curr string

	if len(s) == 1 {
		return s
	}
	for j := range s {
		i := getPrev(s, j-1, j)
		k := getNext(s, j, j+1)

		if i == 0 && len(s) > 1 && s[0] == s[1] { // Handle first char exception
			curr = s[:k]
		} else if k == len(s)+1 && len(s) > 1 && s[k-1] == s[k] { // Handle first char exception
			curr = s[i+1:]
		} else { // Normal case
			curr = s[i+1 : k]
		}
		if len(curr) > len(ret) {
			ret = curr
		}
		for i >= 0 && k < len(s) {
			if s[i] == s[k] {
				curr = string(s[i]) + curr + string(s[k])
			} else {
				break
			}
			i--
			k++
			if len(curr) > len(ret) {
				ret = curr
			}
		}
	}
	return ret
}

// Returns first next different index
func getNext(s string, j int, k int) int {
	for k < len(s) {
		if s[j] != s[k] {
			if j != k {
				return k
			} else {
				return k + 1
			}
		}
		k++
	}
	return k
}

// Returns first previous different index
func getPrev(s string, i int, j int) int {
	for i > 0 {
		if s[i] != s[j] {
			if i != j {
				return i
			} else {
				return i - 1
			}
		}
		i--
	}
	return 0
}
