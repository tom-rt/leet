package main

import (
	"slices"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	ret := twoSumII([]int{2, 7, 11, 15}, 9)
	if !slices.Equal([]int{1, 2}, ret) {
		t.Fatalf("Nope")
	}
}
