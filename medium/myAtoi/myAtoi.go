package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	var min32 int = math.MinInt32    // -2147483648
	var max32 uint64 = math.MaxInt32 // 2147483647

	var ret uint64
	if len(s) == 0 {
		return 0
	}

	cleaned, isNeg := cleanString(s)
	if len(cleaned) == 0 {
		return 0
	}

	if len(cleaned) > 20 && !isNeg { // 20 is the max amount of digits for uint64
		return int(max32)
	}

	var fact uint64 = 1
	var x uint64
	for i := 0; i < len(cleaned); i++ {
		fact = fact * 10
	}
	for i := 0; i < len(cleaned); i++ {
		fact = fact / 10
		if int(cleaned[i]-48) == 0 {
			x = 0
		} else {
			x = uint64(cleaned[i]-48) * fact
		}
		ret += x
	}

	if ret > max32 && isNeg {
		return int(min32)
	} else if ret > max32 {
		return int(max32)
	} else if isNeg {
		return int(ret) * -1
	} else {
		return int(ret)
	}
}

func cleanString(s string) (string, bool) {
	var isPos bool
	var isNeg bool
	var begin int
	var idx int
	var checkSign int

	for idx < len(s) && s[idx] == ' ' { // removing spaces
		idx++
	}

	if idx < len(s) && s[idx] != '-' && s[idx] != '+' && !unicode.IsDigit(rune(s[idx])) { // checking validity
		return "", false
	}

	for idx < len(s) && (s[idx] == '-' || s[idx] == '+') { // removing + -
		// handling signs
		if s[idx] == '+' {
			checkSign++
			isPos = true
		} else if s[idx] == '-' {
			checkSign++
			isNeg = true
		}

		if checkSign > 1 {
			return "", false
		}
		idx++
	}
	if isPos && isNeg {
		return "", false
	}

	if idx < len(s) && !unicode.IsDigit(rune(s[idx])) { // checking is digit
		return "", false
	}

	for idx < len(s) && s[idx] == '0' { // removing trailing 0
		if s[idx] == '-' { // handling isNeg
			isNeg = true
		}
		idx++
	}

	if idx < len(s) && !unicode.IsDigit(rune(s[idx])) { // checking is digit
		return "0", false
	}

	begin = idx
	for idx < len(s) && unicode.IsDigit(rune(s[idx])) { // reading all digits
		idx++
	}

	return s[begin:idx], isNeg
}
