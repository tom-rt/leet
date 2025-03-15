package main

import (
	"fmt"
	"math"
	"strings"
)

func romanToInt(s string) int {
	// Replace str by str
	var total uint64

	// I can be placed before V (5) and X (10) to make 4 and 9.
	newStr := strings.ReplaceAll(s, "IV", "0")
	newStr = strings.ReplaceAll(newStr, "IX", "1")
	// X can be placed before L (50) and C (100) to make 40 and 90.
	newStr = strings.ReplaceAll(newStr, "XL", "2")
	newStr = strings.ReplaceAll(newStr, "XC", "3")
	// C can be placed before D (500) and M (1000) to make 400 and 900.
	newStr = strings.ReplaceAll(newStr, "CD", "4")
	newStr = strings.ReplaceAll(newStr, "CM", "5")

	for _, i := range newStr {
		c := string(i)
		switch c {
		case "I":
			total += 1
		case "V":
			total += 5
		case "X":
			total += 10
		case "L":
			total += 50
		case "C":
			total += 100
		case "D":
			total += 500
		case "M":
			total += 1000
		case "0":
			total += 4
		case "1":
			total += 9
		case "2":
			total += 40
		case "3":
			total += 90
		case "4":
			total += 400
		case "5":
			total += 900
		default:
			fmt.Printf("default\n")
		}
	}
	var ret int
	if total > math.MaxInt {
		ret = 0
	} else {
		ret = int(total)
	}
	return ret
}
