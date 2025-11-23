package main

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{2}
	ret := coinChange(coins, 3)
	if ret != -1 {
		t.Fatalf("Test failed with ({2}, 3). Wanted 0, got %d", ret)
		t.Fail()
	}
	coins = []int{1, 2, 3}
	ret = coinChange(coins, 3)
	if ret != 1 {
		t.Fatalf("Test failed with ({1, 2, 3}, 3). Wanted 1, got %d", ret)
		t.Fail()
	}
	ret = coinChange(coins, 6)
	if ret != 2 {
		t.Fatalf("Test failed with ({1, 2, 3}, 6). Wanted 2, got %d", ret)
		t.Fail()
	}
	ret = coinChange(coins, 4)
	if ret != 2 {
		t.Fatalf("Test failed with ({1, 2, 3}, 4). Wanted 4, got %d", ret)
		t.Fail()
	}
}
