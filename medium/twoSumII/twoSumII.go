package main

func twoSumII(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	// [1,3,5] -> 4
	for {
		if nums[i]+nums[j] == target {
			return []int{i + 1, j + 1}
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
}
