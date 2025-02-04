package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	lenLongestSub := lengthOfLongestSubstring("wlkdubcnnkawelfhbsdhhhhhhhh")
	if lenLongestSub != 11 {
		t.Fatalf("expected 11, got %d", lenLongestSub)
	}
}
