package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	allowedChars := "()[]{}"
	parenthesis := 0
	bracket := 0
	brace := 0

	var prev int
	var curr int
	var isClosing bool
	var prevIsClosing bool
	for _, r := range s {
		if !strings.ContainsRune(allowedChars, r) {
			return false
		}

		if string(r) == "(" {
			isClosing = false
			curr = 1
			parenthesis++
		} else if string(r) == ")" {
			isClosing = true
			curr = 1
			parenthesis--
		} else if string(r) == "[" {
			isClosing = false
			curr = 2
			bracket++
		} else if string(r) == "]" {
			isClosing = true
			curr = 2
			bracket--
		} else if string(r) == "{" {
			isClosing = false
			curr = 3
			brace++
		} else if string(r) == "}" {
			isClosing = true
			curr = 3
			brace--
		}

		// fmt.Printf("%c, %d, %d", r, prev, curr)
		if parenthesis < 0 || bracket < 0 || brace < 0 || (prev != curr && prev != 0 && isClosing && prevIsClosing == false) {
			fmt.Printf("%c %d %d %t %t", r, curr, prev, isClosing, prevIsClosing)
			return false
		}

		prev = curr
		prevIsClosing = isClosing
	}

	if parenthesis != 0 || bracket != 0 || brace != 0 {
		return false
	}

	return true
}
