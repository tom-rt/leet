package main

func twoSum(nums []int, target int) []int {
	var ret []int = make([]int, 2)
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				ret[0] = i
				ret[1] = j
			}
		}
	}
	return ret
}
