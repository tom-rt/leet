package main

import (
	"testing"
)

func TestMaxAscendingSum(t *testing.T) {
	nums := []int{1, 100, 3, 4, 2, 5, 200}
	ret := maxAscendingSum(nums)
	if ret != 207 {
		t.Fatalf("expected 207, got %d", ret)
	}
}
