package main

func main() {
	intToRoman(1234)
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for i, val := range values {
		for num >= val {
			num -= val
			result += symbols[i]
		}
	}
	return result
}
