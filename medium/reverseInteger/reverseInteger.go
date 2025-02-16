package main

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	var revS string
	var isNeg bool
	var min32 int64 = math.MinInt32 // -2147483648
	var max32 int64 = math.MaxInt32 // 2147483647

	if x < 0 {
		isNeg = true
		x = x * -1
	}
	s := strconv.Itoa(x)

	i := len(s) - 1
	for i >= 0 {
		revS += string(s[i])
		i--
	}

	if isNeg {
		revS = "-" + revS
	}
	ret, _ := strconv.ParseInt(revS, 10, 64)
	if ret > max32 || ret < min32 {
		return 0
	}
	return int(ret)
}
