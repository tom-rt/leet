package main

import "fmt"

func swapNextValue(nums *[]int, i int) {
	if i+1 == len(*nums) {
		return
	}
	save := (*nums)[i]
	(*nums)[i] = (*nums)[i+1]
	(*nums)[i+1] = save
}

func areTrailingZeroes(nums []int, i int) bool {
	for i < len(nums) {
		if nums[i] != 0 {
			return false
		}
		i++
	}
	return true
}

func dummyMoveZeroes(nums []int) {
	i := 0
	for i < len(nums) {
		if nums[i] == 0 && !areTrailingZeroes(nums, i) {
			j := i
			for j < len(nums) {
				swapNextValue(&nums, j)
				j++
			}
			i = 0
		} else {
			i++
		}
	}
}

func moveZeroes(nums []int) {
	slow := 0
	for fast := range len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}

	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}

	fmt.Println(nums)
}
