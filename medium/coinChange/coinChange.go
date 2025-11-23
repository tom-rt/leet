package main

import "fmt"

func main() {
	coins := []int{2}
	ret := coinChange(coins, 3)
	fmt.Println(ret)
}

func coinChange(coins []int, amount int) int {
	var coinNumber, total, biggest int
	var remain = amount

	for total != amount && remain != 0 && biggest != -1 {
		biggest = getBiggestCoin(coins, remain)
		if biggest != -1 {
			total += biggest
			remain = remain - biggest
			coinNumber++
		}
		fmt.Println("biggest", biggest)
		// fmt.Println("total", total)
		// fmt.Println("remain", remain)
	}
	return coinNumber
}

func getBiggestCoin(coins []int, limit int) int {
	biggest := -1
	for _, coin := range coins {
		if coin > biggest && coin <= limit {
			biggest = coin
		}
	}
	return biggest
}
