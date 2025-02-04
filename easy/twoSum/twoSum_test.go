package main

import "testing"

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1]
func TestTwoSum(t *testing.T) {
	input := make([]int, 4)
	input[0] = 2
	input[1] = 7
	input[2] = 11
	input[3] = 15
	ret := twoSum(input, 9)
	if ret[0] != 0 && ret[1] != 1 {
		t.Fatal("expected [0,1]")
	}
}
