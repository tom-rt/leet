package main

func maxAscendingSum(nums []int) int {
	idx := 0
	j := 1
	var asc [][]int = make([][]int, 1)
	for i := range nums {
		if i == 0 {
			asc[idx] = make([]int, 0)
			asc[idx] = append(asc[0], nums[i])
		} else if asc[idx][j-1] < nums[i] {
			asc[idx] = append(asc[idx], nums[i])
			j++
		} else {
			idx++
			j = 1
			new := []int{nums[i]}
			asc = append(asc, new)
		}
	}

	max, curr := 0, 0
	for _, i := range asc {
		curr = 0
		for _, j := range i {
			curr += j
		}
		if curr > max {
			max = curr
		}
	}
	return max
}
